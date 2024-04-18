/*
 * Teleport
 * Copyright (C) 2023  Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package slack

import (
	"context"
	"fmt"
	"os/user"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gravitational/trace"
	"github.com/jonboulle/clockwork"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/gravitational/teleport/api/client/proto"
	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/api/types/accesslist"
	"github.com/gravitational/teleport/api/types/header"
	"github.com/gravitational/teleport/integrations/access/accessrequest"
	"github.com/gravitational/teleport/integrations/access/common"
	"github.com/gravitational/teleport/integrations/access/common/auth"
	"github.com/gravitational/teleport/integrations/lib"
	"github.com/gravitational/teleport/integrations/lib/logger"
	"github.com/gravitational/teleport/integrations/lib/testing/integration"
)

var msgFieldRegexp = regexp.MustCompile(`(?im)^\*([a-zA-Z ]+)\*: (.+)$`)
var requestReasonRegexp = regexp.MustCompile("(?im)^\\*Reason\\*:\\ ```\\n(.*?)```(.*?)$")

type SlackSuite struct {
	integration.Suite
	appConfig *Config
	userNames struct {
		ruler     string
		requestor string
		reviewer1 string
		reviewer2 string
		plugin    string
	}
	requestorUser  User
	raceNumber     int
	fakeSlack      *FakeSlack
	fakeStatusSink *fakeStatusSink

	clients          map[string]*integration.Client
	teleportFeatures *proto.Features
	teleportConfig   lib.TeleportConfig
}

func TestSlackbot(t *testing.T) { suite.Run(t, &SlackSuite{}) }

func (s *SlackSuite) SetupSuite() {
	var err error
	t := s.T()

	logger.Init()
	err = logger.Setup(logger.Config{Severity: "debug"})
	require.NoError(t, err)
	s.raceNumber = runtime.GOMAXPROCS(0)
	me, err := user.Current()
	require.NoError(t, err)

	// We set such a big timeout because integration.NewFromEnv could start
	// downloading a Teleport *-bin.tar.gz file which can take a long time.
	ctx := s.SetContextTimeout(2 * time.Minute)

	teleport, err := integration.NewFromEnv(ctx)
	require.NoError(t, err)
	t.Cleanup(teleport.Close)

	auth, err := teleport.NewAuthService()
	require.NoError(t, err)
	s.StartApp(auth)

	s.clients = make(map[string]*integration.Client)

	// Set up the user who has an access to all kinds of resources.

	s.userNames.ruler = me.Username + "-ruler@example.com"
	client, err := teleport.MakeAdmin(ctx, auth, s.userNames.ruler)
	require.NoError(t, err)
	s.clients[s.userNames.ruler] = client

	// Get the server features.

	pong, err := client.Ping(ctx)
	require.NoError(t, err)
	teleportFeatures := pong.GetServerFeatures()

	var bootstrap integration.Bootstrap

	// Set up user who can request the access to role "editor".

	conditions := types.RoleConditions{Request: &types.AccessRequestConditions{Roles: []string{"editor"}}}
	if teleportFeatures.AdvancedAccessWorkflows {
		conditions.Request.Thresholds = []types.AccessReviewThreshold{{Approve: 2, Deny: 2}}
	}
	role, err := bootstrap.AddRole("foo", types.RoleSpecV6{Allow: conditions})
	require.NoError(t, err)

	user, err := bootstrap.AddUserWithRoles(me.Username+"@example.com", role.GetName())
	require.NoError(t, err)
	s.userNames.requestor = user.GetName()

	// Set up TWO users who can review access requests to role "editor".

	conditions = types.RoleConditions{}
	if teleportFeatures.AdvancedAccessWorkflows {
		conditions.ReviewRequests = &types.AccessReviewConditions{Roles: []string{"editor"}}
	}
	role, err = bootstrap.AddRole("foo-reviewer", types.RoleSpecV6{Allow: conditions})
	require.NoError(t, err)

	user, err = bootstrap.AddUserWithRoles(me.Username+"-reviewer1@example.com", role.GetName())
	require.NoError(t, err)
	s.userNames.reviewer1 = user.GetName()

	user, err = bootstrap.AddUserWithRoles(me.Username+"-reviewer2@example.com", role.GetName())
	require.NoError(t, err)
	s.userNames.reviewer2 = user.GetName()

	// Set up plugin user.

	role, err = bootstrap.AddRole("access-slack", types.RoleSpecV6{
		Allow: types.RoleConditions{
			Rules: []types.Rule{
				types.NewRule(types.KindAccessList, []string{"list", "read"}),
				types.NewRule(types.KindAccessRequest, []string{"list", "read"}),
				types.NewRule(types.KindAccessPluginData, []string{"update"}),
			},
		},
	})
	require.NoError(t, err)

	user, err = bootstrap.AddUserWithRoles("access-slack", role.GetName())
	require.NoError(t, err)
	s.userNames.plugin = user.GetName()

	// Bake all the resources.

	err = teleport.Bootstrap(ctx, auth, bootstrap.Resources())
	require.NoError(t, err)

	// Initialize the clients.

	client, err = teleport.NewClient(ctx, auth, s.userNames.requestor)
	require.NoError(t, err)
	s.clients[s.userNames.requestor] = client

	if teleportFeatures.AdvancedAccessWorkflows {
		client, err = teleport.NewClient(ctx, auth, s.userNames.reviewer1)
		require.NoError(t, err)
		s.clients[s.userNames.reviewer1] = client

		client, err = teleport.NewClient(ctx, auth, s.userNames.reviewer2)
		require.NoError(t, err)
		s.clients[s.userNames.reviewer2] = client
	}

	identityPath, err := teleport.Sign(ctx, auth, s.userNames.plugin)
	require.NoError(t, err)

	s.teleportConfig.Addr = auth.AuthAddr().String()
	s.teleportConfig.Identity = identityPath
	s.teleportFeatures = teleportFeatures
}

func (s *SlackSuite) SetupTest() {
	t := s.T()

	err := logger.Setup(logger.Config{Severity: "debug"})
	require.NoError(t, err)

	s.fakeSlack = NewFakeSlack(User{Name: "slackbot"}, s.raceNumber)
	t.Cleanup(s.fakeSlack.Close)

	s.requestorUser = s.fakeSlack.StoreUser(User{Name: "Vladimir", Profile: UserProfile{Email: s.userNames.requestor}})

	s.fakeStatusSink = &fakeStatusSink{}

	var conf Config
	conf.Teleport = s.teleportConfig
	conf.Slack.Token = "000000"
	conf.Slack.APIURL = s.fakeSlack.URL() + "/"
	conf.AccessTokenProvider = auth.NewStaticAccessTokenProvider(conf.Slack.Token)
	conf.StatusSink = s.fakeStatusSink

	s.appConfig = &conf
	s.SetContextTimeout(5 * time.Second)
}

func (s *SlackSuite) startApp() {
	t := s.T()
	t.Helper()

	app := NewSlackApp(s.appConfig)
	s.StartApp(app)
}

func (s *SlackSuite) ruler() *integration.Client {
	return s.clients[s.userNames.ruler]
}

func (s *SlackSuite) requestor() *integration.Client {
	return s.clients[s.userNames.requestor]
}

func (s *SlackSuite) reviewer1() *integration.Client {
	return s.clients[s.userNames.reviewer1]
}

func (s *SlackSuite) reviewer2() *integration.Client {
	return s.clients[s.userNames.reviewer2]
}

func (s *SlackSuite) newAccessRequest(reviewers []User) types.AccessRequest {
	t := s.T()
	t.Helper()

	req, err := types.NewAccessRequest(uuid.New().String(), s.userNames.requestor, "editor")
	require.NoError(t, err)
	// max size of request was decreased here: https://github.com/gravitational/teleport/pull/13298
	req.SetRequestReason("because of " + strings.Repeat("A", 4000))
	var suggestedReviewers []string
	for _, user := range reviewers {
		suggestedReviewers = append(suggestedReviewers, user.Profile.Email)
	}
	req.SetSuggestedReviewers(suggestedReviewers)
	return req
}

func (s *SlackSuite) createAccessRequest(reviewers []User) types.AccessRequest {
	t := s.T()
	t.Helper()

	req := s.newAccessRequest(reviewers)
	out, err := s.requestor().CreateAccessRequestV2(s.Context(), req)
	require.NoError(t, err)
	return out
}

func (s *SlackSuite) checkPluginData(reqID string, cond func(accessrequest.PluginData) bool) accessrequest.PluginData {
	t := s.T()
	t.Helper()

	for {
		rawData, err := s.ruler().PollAccessRequestPluginData(s.Context(), "slack", reqID)
		require.NoError(t, err)
		data, err := accessrequest.DecodePluginData(rawData)
		require.NoError(t, err)
		if cond(data) {
			return data
		}
	}
}

func (s *SlackSuite) TestMessagePosting() {
	t := s.T()

	reviewer1 := s.fakeSlack.StoreUser(User{Profile: UserProfile{Email: s.userNames.reviewer1}})
	reviewer2 := s.fakeSlack.StoreUser(User{Profile: UserProfile{Email: s.userNames.reviewer2}})

	s.startApp()
	const numMessages = 3
	request := s.createAccessRequest([]User{reviewer2, reviewer1})

	pluginData := s.checkPluginData(request.GetName(), func(data accessrequest.PluginData) bool {
		return len(data.SentMessages) > 0
	})
	assert.Len(t, pluginData.SentMessages, numMessages)

	var messages []Message
	messageSet := make(SlackDataMessageSet)
	for i := 0; i < numMessages; i++ {
		msg, err := s.fakeSlack.CheckNewMessage(s.Context())
		require.NoError(t, err)
		messageSet.Add(accessrequest.MessageData{ChannelID: msg.Channel, MessageID: msg.Timestamp})
		messages = append(messages, msg)
	}

	assert.Len(t, messageSet, numMessages)
	for i := 0; i < numMessages; i++ {
		assert.Contains(t, messageSet, pluginData.SentMessages[i])
	}

	sort.Sort(SlackMessageSlice(messages))

	assert.Equal(t, s.requestorUser.ID, messages[0].Channel)
	assert.Equal(t, reviewer1.ID, messages[1].Channel)
	assert.Equal(t, reviewer2.ID, messages[2].Channel)

	msgUser, err := parseMessageField(messages[0], "User")
	require.NoError(t, err)
	assert.Equal(t, s.userNames.requestor, msgUser)

	block, ok := messages[0].BlockItems[1].Block.(SectionBlock)
	require.True(t, ok)
	t.Logf("%q", block.Text.GetText())
	matches := requestReasonRegexp.FindAllStringSubmatch(block.Text.GetText(), -1)
	require.Len(t, matches, 1)
	require.Len(t, matches[0], 3)
	assert.Equal(t, "because of "+strings.Repeat("A", 489), matches[0][1])
	assert.Equal(t, " (truncated)", matches[0][2])

	statusLine, err := getStatusLine(messages[0])
	require.NoError(t, err)
	assert.Equal(t, "*Status*: ⏳ PENDING", statusLine)

	assert.Equal(t, types.PluginStatusCode_RUNNING, s.fakeStatusSink.Get().GetCode())
}

func (s *SlackSuite) TestRecipientsConfig() {
	t := s.T()

	reviewer1 := s.fakeSlack.StoreUser(User{Profile: UserProfile{Email: s.userNames.reviewer1}})
	reviewer2 := s.fakeSlack.StoreUser(User{Profile: UserProfile{Email: s.userNames.reviewer2}})
	s.appConfig.Recipients = common.RawRecipientsMap{
		types.Wildcard: []string{reviewer2.Profile.Email, reviewer1.ID},
	}

	s.startApp()
	const numMessages = 3

	request := s.createAccessRequest(nil)
	pluginData := s.checkPluginData(request.GetName(), func(data accessrequest.PluginData) bool {
		return len(data.SentMessages) > 0
	})
	assert.Len(t, pluginData.SentMessages, numMessages)

	var messages []Message

	messageSet := make(SlackDataMessageSet)

	for i := 0; i < numMessages; i++ {
		msg, err := s.fakeSlack.CheckNewMessage(s.Context())
		require.NoError(t, err)
		messageSet.Add(accessrequest.MessageData{ChannelID: msg.Channel, MessageID: msg.Timestamp})
		messages = append(messages, msg)
	}

	assert.Len(t, messageSet, numMessages)
	for i := 0; i < numMessages; i++ {
		assert.Contains(t, messageSet, pluginData.SentMessages[i])
	}

	sort.Sort(SlackMessageSlice(messages))

	assert.Equal(t, s.requestorUser.ID, messages[0].Channel)
	assert.Equal(t, reviewer1.ID, messages[1].Channel)
	assert.Equal(t, reviewer2.ID, messages[2].Channel)
}

func (s *SlackSuite) TestApproval() {
	t := s.T()

	reviewer := s.fakeSlack.StoreUser(User{Profile: UserProfile{Email: s.userNames.reviewer1}})

	s.startApp()

	req := s.createAccessRequest([]User{reviewer})
	msgs := s.checkNewMessages(t, channelsToMessages(s.requestorUser.ID, reviewer.ID), matchOnlyOnChannel)

	err := s.ruler().ApproveAccessRequest(s.Context(), req.GetName(), "okay")
	require.NoError(t, err)

	s.checkNewMessageUpdateByAPI(t, msgs, matchByTimestamp, func(t *testing.T, msgUpdate Message) {
		statusLine, err := getStatusLine(msgUpdate)
		require.NoError(t, err)
		assert.Equal(t, "*Status*: ✅ APPROVED\n*Resolution reason*: ```\nokay```", statusLine)
	})

	s.checkNewMessages(t, channelsToMessages(s.requestorUser.ID), matchOnlyOnChannel, func(t *testing.T, m Message) {
		line := fmt.Sprintf("Request with ID %q has been updated: *%s*", req.GetName(), types.RequestState_APPROVED.String())
		assert.Equal(t, line, m.BlockItems[0].Block.(SectionBlock).Text.GetText())
	})
}

func (s *SlackSuite) TestDenial() {
	t := s.T()

	reviewer := s.fakeSlack.StoreUser(User{Profile: UserProfile{Email: s.userNames.reviewer1}})

	s.startApp()

	req := s.createAccessRequest([]User{reviewer})
	msgs := s.checkNewMessages(t, channelsToMessages(s.requestorUser.ID, reviewer.ID), matchOnlyOnChannel)

	// max size of request was decreased here: https://github.com/gravitational/teleport/pull/13298
	err := s.ruler().DenyAccessRequest(s.Context(), req.GetName(), "not okay "+strings.Repeat("A", 4000))
	require.NoError(t, err)

	s.checkNewMessageUpdateByAPI(t, msgs, matchByTimestamp, func(t *testing.T, msgUpdate Message) {
		statusLine, err := getStatusLine(msgUpdate)
		require.NoError(t, err)
		assert.Equal(t, "*Status*: ❌ DENIED\n*Resolution reason*: ```\nnot okay "+strings.Repeat("A", 491)+"``` (truncated)", statusLine)
	})
}

func (s *SlackSuite) TestReviewReplies() {
	t := s.T()

	if !s.teleportFeatures.AdvancedAccessWorkflows {
		t.Skip("Doesn't work in OSS version")
	}

	reviewer := s.fakeSlack.StoreUser(User{Profile: UserProfile{Email: s.userNames.reviewer1}})

	s.startApp()

	req := s.createAccessRequest([]User{reviewer})
	s.checkPluginData(req.GetName(), func(data accessrequest.PluginData) bool {
		return len(data.SentMessages) > 0
	})

	msgs := s.checkNewMessages(t, channelsToMessages(s.requestorUser.ID, reviewer.ID), matchOnlyOnChannel)

	err := s.reviewer1().SubmitAccessRequestReview(s.Context(), req.GetName(), types.AccessReview{
		Author:        s.userNames.reviewer1,
		ProposedState: types.RequestState_APPROVED,
		Created:       time.Now(),
		Reason:        "okay",
	})
	require.NoError(t, err)

	s.checkNewMessages(t, msgs, matchByThreadTs, func(t *testing.T, reply Message) {
		assert.Contains(t, reply.Text, s.userNames.reviewer1+" reviewed the request", "reply must contain a review author")
		assert.Contains(t, reply.Text, "Resolution: ✅ APPROVED", "reply must contain a proposed state")
		assert.Contains(t, reply.Text, "Reason: ```\nokay```", "reply must contain a reason")
	})

	err = s.reviewer2().SubmitAccessRequestReview(s.Context(), req.GetName(), types.AccessReview{
		Author:        s.userNames.reviewer2,
		ProposedState: types.RequestState_DENIED,
		Created:       time.Now(),
		Reason:        "not okay",
	})
	require.NoError(t, err)

	s.checkNewMessages(t, msgs, matchByThreadTs, func(t *testing.T, reply Message) {
		assert.Contains(t, reply.Text, s.userNames.reviewer2+" reviewed the request", "reply must contain a review author")
		assert.Contains(t, reply.Text, "Resolution: ❌ DENIED", "reply must contain a proposed state")
		assert.Contains(t, reply.Text, "Reason: ```\nnot okay```", "reply must contain a reason")
	})
}

func (s *SlackSuite) TestApprovalByReview() {
	t := s.T()

	if !s.teleportFeatures.AdvancedAccessWorkflows {
		t.Skip("Doesn't work in OSS version")
	}

	reviewer := s.fakeSlack.StoreUser(User{Profile: UserProfile{Email: s.userNames.reviewer1}})

	s.startApp()

	req := s.createAccessRequest([]User{reviewer})
	msgs := s.checkNewMessages(t, channelsToMessages(s.requestorUser.ID, reviewer.ID), matchOnlyOnChannel)

	err := s.reviewer1().SubmitAccessRequestReview(s.Context(), req.GetName(), types.AccessReview{
		Author:        s.userNames.reviewer1,
		ProposedState: types.RequestState_APPROVED,
		Created:       time.Now(),
		Reason:        "okay",
	})
	require.NoError(t, err)

	s.checkNewMessages(t, msgs, matchByThreadTs, func(t *testing.T, reply Message) {
		assert.Contains(t, reply.Text, s.userNames.reviewer1+" reviewed the request", "reply must contain a review author")
	})

	s.checkNewMessageUpdateByAPI(t, msgs, matchByTimestamp, func(t *testing.T, msgUpdate Message) {
		statusLine, err := getStatusLine(msgUpdate)
		require.NoError(t, err)
		assert.Equal(t, "*Status*: ⏳ PENDING", statusLine)
	})

	err = s.reviewer2().SubmitAccessRequestReview(s.Context(), req.GetName(), types.AccessReview{
		Author:        s.userNames.reviewer2,
		ProposedState: types.RequestState_APPROVED,
		Created:       time.Now(),
		Reason:        "finally okay",
	})
	require.NoError(t, err)

	s.checkNewMessages(t, msgs, matchByThreadTs, func(t *testing.T, reply Message) {
		assert.Contains(t, reply.Text, s.userNames.reviewer2+" reviewed the request", "reply must contain a review author")
	})

	s.checkNewMessageUpdateByAPI(t, msgs, matchByTimestamp, func(t *testing.T, msgUpdate Message) {
		statusLine, err := getStatusLine(msgUpdate)
		require.NoError(t, err)
		assert.Equal(t, "*Status*: ✅ APPROVED\n*Resolution reason*: ```\nfinally okay```", statusLine)
	})
}

func (s *SlackSuite) TestDenialByReview() {
	t := s.T()

	if !s.teleportFeatures.AdvancedAccessWorkflows {
		t.Skip("Doesn't work in OSS version")
	}

	reviewer := s.fakeSlack.StoreUser(User{Profile: UserProfile{Email: s.userNames.reviewer1}})

	s.startApp()

	req := s.createAccessRequest([]User{reviewer})
	msgs := s.checkNewMessages(t, channelsToMessages(s.requestorUser.ID, reviewer.ID), matchOnlyOnChannel)

	err := s.reviewer1().SubmitAccessRequestReview(s.Context(), req.GetName(), types.AccessReview{
		Author:        s.userNames.reviewer1,
		ProposedState: types.RequestState_DENIED,
		Created:       time.Now(),
		Reason:        "not okay",
	})
	require.NoError(t, err)

	s.checkNewMessages(t, msgs, matchByThreadTs, func(t *testing.T, reply Message) {
		assert.Contains(t, reply.Text, s.userNames.reviewer1+" reviewed the request", "reply must contain a review author")
	})

	s.checkNewMessageUpdateByAPI(t, msgs, matchByTimestamp, func(t *testing.T, msgUpdate Message) {
		statusLine, err := getStatusLine(msgUpdate)
		require.NoError(t, err)
		assert.Equal(t, "*Status*: ⏳ PENDING", statusLine)
	})

	err = s.reviewer2().SubmitAccessRequestReview(s.Context(), req.GetName(), types.AccessReview{
		Author:        s.userNames.reviewer2,
		ProposedState: types.RequestState_DENIED,
		Created:       time.Now(),
		Reason:        "finally not okay",
	})
	require.NoError(t, err)

	s.checkNewMessages(t, msgs, matchByThreadTs, func(t *testing.T, reply Message) {
		assert.Contains(t, reply.Text, s.userNames.reviewer2+" reviewed the request", "reply must contain a review author")
	})

	s.checkNewMessageUpdateByAPI(t, msgs, matchByTimestamp, func(t *testing.T, msgUpdate Message) {
		statusLine, err := getStatusLine(msgUpdate)
		require.NoError(t, err)
		assert.Equal(t, "*Status*: ❌ DENIED\n*Resolution reason*: ```\nfinally not okay```", statusLine)
	})
}

func (s *SlackSuite) TestExpiration() {
	t := s.T()

	reviewer := s.fakeSlack.StoreUser(User{Profile: UserProfile{Email: s.userNames.reviewer1}})

	s.startApp()

	request := s.createAccessRequest([]User{reviewer})
	msgs := s.checkNewMessages(t, channelsToMessages(s.requestorUser.ID, reviewer.ID), matchOnlyOnChannel)

	s.checkPluginData(request.GetName(), func(data accessrequest.PluginData) bool {
		return len(data.SentMessages) > 0
	})

	err := s.ruler().DeleteAccessRequest(s.Context(), request.GetName()) // simulate expiration
	require.NoError(t, err)

	s.checkNewMessageUpdateByAPI(t, msgs, matchByTimestamp, func(t *testing.T, msgUpdate Message) {
		statusLine, err := getStatusLine(msgUpdate)
		require.NoError(t, err)
		assert.Equal(t, "*Status*: ⌛ EXPIRED", statusLine)
	})
}

func (s *SlackSuite) TestAccessListReminder() {
	t := s.T()

	if !s.teleportFeatures.AdvancedAccessWorkflows {
		t.Skip("Doesn't work in OSS version")
	}

	reviewer := s.fakeSlack.StoreUser(User{Profile: UserProfile{Email: s.userNames.reviewer1}})

	clock := clockwork.NewFakeClockAt(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC))
	s.appConfig.clock = clock
	s.startApp()

	accessList, err := accesslist.NewAccessList(header.Metadata{
		Name: "access-list",
	}, accesslist.Spec{
		Title: "simple title",
		Grants: accesslist.Grants{
			Roles: []string{"grant"},
		},
		Owners: []accesslist.Owner{
			{Name: s.userNames.reviewer1},
		},
		Audit: accesslist.Audit{
			NextAuditDate: time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC),
		},
	})
	require.NoError(t, err)

	ctx := context.Background()
	_, err = s.ruler().AccessListClient().UpsertAccessList(ctx, accessList)
	require.NoError(t, err)

	// 2 weeks before date, should trigger reminder.
	clock.Advance(45 * 24 * time.Hour)
	s.requireReminderMsgEqual(reviewer.ID, "Access List *simple title* is due for a review by 2023-03-01. Please review it soon!")

	// 1 weeks before date, should trigger reminder.
	clock.Advance(7 * 24 * time.Hour)
	s.requireReminderMsgEqual(reviewer.ID, "Access List *simple title* is due for a review by 2023-03-01. Please review it soon!")

	// On the date, should trigger reminder.
	clock.Advance(7 * 24 * time.Hour)
	s.requireReminderMsgEqual(reviewer.ID, "Access List *simple title* is due for a review by 2023-03-01. Please review it soon!")

	// Past the date, should trigger reminder.
	clock.Advance(7 * 24 * time.Hour)
	s.requireReminderMsgEqual(reviewer.ID, "Access List *simple title* is 7 day(s) past due for a review! Please review it.")
}

func (s *SlackSuite) requireReminderMsgEqual(id, text string) {
	t := s.T()

	msg, err := s.fakeSlack.CheckNewMessage(s.Context())
	require.NoError(t, err)
	require.Equal(t, id, msg.Channel)
	require.IsType(t, SectionBlock{}, msg.BlockItems[0].Block)
	require.Equal(t, text, (msg.BlockItems[0].Block).(SectionBlock).Text.GetText())
}

func (s *SlackSuite) TestRace() {
	t := s.T()

	if !s.teleportFeatures.AdvancedAccessWorkflows {
		t.Skip("Doesn't work in OSS version")
	}

	err := logger.Setup(logger.Config{Severity: "info"}) // Turn off noisy debug logging
	require.NoError(t, err)

	reviewer1 := s.fakeSlack.StoreUser(User{Profile: UserProfile{Email: s.userNames.reviewer1}})
	reviewer2 := s.fakeSlack.StoreUser(User{Profile: UserProfile{Email: s.userNames.reviewer2}})

	s.SetContextTimeout(20 * time.Second)
	s.startApp()

	var (
		raceErr             error
		raceErrOnce         sync.Once
		threadMsgIDs        sync.Map
		threadMsgsCount     int32
		msgUpdateCounters   sync.Map
		reviewReplyCounters sync.Map
	)
	setRaceErr := func(err error) error {
		raceErrOnce.Do(func() {
			raceErr = err
		})
		return err
	}

	process := lib.NewProcess(s.Context())
	for i := 0; i < s.raceNumber; i++ {
		process.SpawnCritical(func(ctx context.Context) error {
			req, err := types.NewAccessRequest(uuid.New().String(), s.userNames.requestor, "editor")
			if err != nil {
				return setRaceErr(trace.Wrap(err))
			}
			req.SetSuggestedReviewers([]string{reviewer1.Profile.Email, reviewer2.Profile.Email})
			if _, err := s.requestor().CreateAccessRequestV2(ctx, req); err != nil {
				return setRaceErr(trace.Wrap(err))
			}
			return nil
		})
	}

	// Having TWO suggested reviewers will post THREE messages for each request (including the requestor).
	// We also have approval threshold of TWO set in the role properties
	// so lets simply submit the approval from each of the suggested reviewers.
	//
	// Multiplier NINE means that we handle THREE messages for each request and also
	// TWO comments for each message: 2 * (1 message + 2 comments).
	for i := 0; i < 9*s.raceNumber; i++ {
		process.SpawnCritical(func(ctx context.Context) error {
			msg, err := s.fakeSlack.CheckNewMessage(ctx)
			if err != nil {
				return setRaceErr(trace.Wrap(err))
			}

			if msg.ThreadTs == "" {
				// Handle "root" notifications.

				threadMsgKey := accessrequest.MessageData{ChannelID: msg.Channel, MessageID: msg.Timestamp}
				if _, loaded := threadMsgIDs.LoadOrStore(threadMsgKey, struct{}{}); loaded {
					return setRaceErr(trace.Errorf("thread %v already stored", threadMsgKey))
				}
				atomic.AddInt32(&threadMsgsCount, 1)

				user, ok := s.fakeSlack.GetUser(msg.Channel)
				if !ok {
					return setRaceErr(trace.Errorf("user %s is not found", msg.Channel))
				}

				reqID, err := parseMessageField(msg, "ID")
				if err != nil {
					return setRaceErr(trace.Wrap(err))
				}

				// The requestor can't submit reviews.
				if user.ID == s.requestorUser.ID {
					return nil
				}

				if err = s.clients[user.Profile.Email].SubmitAccessRequestReview(ctx, reqID, types.AccessReview{
					Author:        user.Profile.Email,
					ProposedState: types.RequestState_APPROVED,
					Created:       time.Now(),
					Reason:        "okay",
				}); err != nil {
					return setRaceErr(trace.Wrap(err))
				}
			} else {
				// Handle review comments.

				threadMsgKey := accessrequest.MessageData{ChannelID: msg.Channel, MessageID: msg.ThreadTs}
				var newCounter int32
				val, _ := reviewReplyCounters.LoadOrStore(threadMsgKey, &newCounter)
				counterPtr := val.(*int32)
				atomic.AddInt32(counterPtr, 1)
			}

			return nil
		})
	}

	// Multiplier THREE means that we handle the 2 updates for each of the two messages posted to reviewers.
	for i := 0; i < 3*2*s.raceNumber; i++ {
		process.SpawnCritical(func(ctx context.Context) error {
			msg, err := s.fakeSlack.CheckMessageUpdateByAPI(ctx)
			if err != nil {
				return setRaceErr(trace.Wrap(err))
			}

			threadMsgKey := accessrequest.MessageData{ChannelID: msg.Channel, MessageID: msg.Timestamp}
			var newCounter int32
			val, _ := msgUpdateCounters.LoadOrStore(threadMsgKey, &newCounter)
			counterPtr := val.(*int32)
			atomic.AddInt32(counterPtr, 1)

			return nil
		})
	}

	process.Terminate()
	<-process.Done()
	require.NoError(t, raceErr)

	assert.Equal(t, int32(3*s.raceNumber), threadMsgsCount)
	threadMsgIDs.Range(func(key, value interface{}) bool {
		next := true

		val, loaded := reviewReplyCounters.LoadAndDelete(key)
		next = next && assert.True(t, loaded)
		counterPtr := val.(*int32)
		next = next && assert.Equal(t, int32(2), *counterPtr)

		val, loaded = msgUpdateCounters.LoadAndDelete(key)
		next = next && assert.True(t, loaded)
		counterPtr = val.(*int32)
		// Each message should be updated 2 times
		next = next && assert.Equal(t, int32(2), *counterPtr)

		return next
	})
}

func parseMessageField(msg Message, field string) (string, error) {
	block := msg.BlockItems[1].Block
	sectionBlock, ok := block.(SectionBlock)
	if !ok {
		return "", trace.Errorf("invalid block type %T", block)
	}

	if sectionBlock.Text.TextObject == nil {
		return "", trace.Errorf("section block does not contain text")
	}

	text := sectionBlock.Text.GetText()
	matches := msgFieldRegexp.FindAllStringSubmatch(text, -1)
	if matches == nil {
		return "", trace.Errorf("cannot parse fields from text %s", text)
	}
	var fields []string
	for _, match := range matches {
		if match[1] == field {
			return match[2], nil
		}
		fields = append(fields, match[1])
	}
	return "", trace.Errorf("cannot find field %s in %v", field, fields)
}

func getStatusLine(msg Message) (string, error) {
	block := msg.BlockItems[2].Block
	contextBlock, ok := block.(ContextBlock)
	if !ok {
		return "", trace.Errorf("invalid block type %T", block)
	}

	elementItems := contextBlock.ElementItems
	if n := len(elementItems); n != 1 {
		return "", trace.Errorf("expected only one context element, got %v", n)
	}

	element := elementItems[0].ContextElement
	textBlock, ok := element.(TextObject)
	if !ok {
		return "", trace.Errorf("invalid element type %T", element)
	}

	return textBlock.GetText(), nil
}

// matchFns are functions that tell how to match two messages together after matching on the channel ID.
type matchFn func(matchAgainst Message, newMsg Message) bool

func matchOnlyOnChannel(_, _ Message) bool {
	return true
}

func matchByTimestamp(matchAgainst, newMsg Message) bool {
	return matchAgainst.Timestamp == newMsg.Timestamp
}

func matchByThreadTs(matchAgainst, newMsg Message) bool {
	return matchAgainst.Timestamp == newMsg.ThreadTs
}

// checkMsgTestFn is a test function to run on a new message after it has been matched.
type checkMsgTestFn func(*testing.T, Message)

func (s *SlackSuite) checkNewMessages(t *testing.T, matchMessages []Message, matchBy matchFn, testFns ...checkMsgTestFn) []Message {
	t.Helper()
	return s.matchAndCallFn(t, matchMessages, matchBy, testFns, s.fakeSlack.CheckNewMessage)
}

func (s *SlackSuite) checkNewMessageUpdateByAPI(t *testing.T, matchMessages []Message, matchBy matchFn, testFns ...checkMsgTestFn) []Message {
	t.Helper()
	return s.matchAndCallFn(t, matchMessages, matchBy, testFns, s.fakeSlack.CheckMessageUpdateByAPI)
}

func channelsToMessages(channels ...string) (messages []Message) {
	for _, channel := range channels {
		messages = append(messages, Message{BaseMessage: BaseMessage{Channel: channel}})
	}

	return messages
}

type slackCheckMessage func(context.Context) (Message, error)

func (s *SlackSuite) matchAndCallFn(t *testing.T, matchMessages []Message, matchBy matchFn, testFns []checkMsgTestFn, slackCall slackCheckMessage) []Message {
	matchingTimestamps := map[string]Message{}

	for _, matchMessage := range matchMessages {
		matchingTimestamps[matchMessage.Channel] = matchMessage
	}

	var messages []Message
	var notMatchingMessages []Message

	// Try for 5 seconds to get the expected messages
	require.Eventually(t, func() bool {
		msg, err := slackCall(s.Context())
		if err != nil {
			return false
		}

		if matchMsg, ok := matchingTimestamps[msg.Channel]; ok {
			if matchBy(matchMsg, msg) {
				messages = append(messages, msg)
			}
		} else {
			notMatchingMessages = append(notMatchingMessages, msg)
		}

		return len(messages) == len(matchMessages)
	}, 2*time.Second, 100*time.Millisecond)

	require.Len(t, messages, len(matchMessages), "missing required messages, found %v", notMatchingMessages)

	for _, testFn := range testFns {
		for _, message := range messages {
			testFn(t, message)
		}
	}

	return messages
}
