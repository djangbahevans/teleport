/*
 * Teleport
 * Copyright (C) 2024  Gravitational, Inc.
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

package plugin

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/alecthomas/kingpin/v2"
	"github.com/gravitational/trace"
	"golang.org/x/crypto/bcrypt"

	pluginsv1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/plugins/v1"
	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/lib/utils"
)

func (p *PluginsCommand) initInstallOkta(parent *kingpin.CmdClause) {
	p.install.okta.cmd = parent.Command("okta", "Install an okta integration")
	p.install.okta.cmd.
		Flag("name", "Name of the plugin resource to create").
		Default("okta").
		StringVar(&p.install.name)
	p.install.okta.cmd.
		Flag("org", "URL of Okta organization").
		Required().
		URLVar(&p.install.okta.org)
	p.install.okta.cmd.
		Flag("api-token", "Okta API token for the plugin to use").
		StringVar(&p.install.okta.apiToken)
	p.install.okta.cmd.
		Flag("saml-connector", "SAML connector used for Okta SSO login.").
		Required().
		StringVar(&p.install.okta.samlConnector)
	p.install.okta.cmd.
		Flag("app-id", "Okta ID of the APP used for SSO via SAML").
		StringVar(&p.install.okta.appID)
	p.install.okta.cmd.
		Flag("scim", "Enable SCIM Okta integration").
		BoolVar(&p.install.okta.scimEnabled)
	p.install.okta.cmd.
		Flag("scim-token", "Okta SCIM auth token for the plugin to use").
		StringVar(&p.install.okta.scimToken)
	p.install.okta.cmd.
		Flag("users-sync", "Enable user synchronization").
		Default("true").
		BoolVar(&p.install.okta.userSync)
	p.install.okta.cmd.
		Flag("owner", "Add default owners for synced Access Lists").
		Short('o').
		StringsVar(&p.install.okta.defaultOwners)
	p.install.okta.cmd.
		Flag("accesslist-sync", "Enable group to Access List synchronization").
		Default("true").
		BoolVar(&p.install.okta.accessListSync)
	p.install.okta.cmd.
		Flag("appgroup-sync", "Enable Okta Applications and Groups sync").
		Default("true").
		BoolVar(&p.install.okta.appGroupSync)
	p.install.okta.cmd.
		Flag("group-filter", "Add a group filter. Supports globbing by default. Enclose in `^pattern$` for full regex support.").
		Short('g').
		StringsVar(&p.install.okta.groupFilters)
	p.install.okta.cmd.
		Flag("app-filter", "Add an app filter. Supports globbing by default. Enclose in `^pattern$` for full regex support.").
		Short('a').
		StringsVar(&p.install.okta.appFilters)
}

type oktaArgs struct {
	cmd            *kingpin.CmdClause
	org            *url.URL
	appID          string
	samlConnector  string
	apiToken       string
	scimEnabled    bool
	scimToken      string
	userSync       bool
	accessListSync bool
	defaultOwners  []string
	appFilters     []string
	groupFilters   []string
	appGroupSync   bool

	autoGeneratedSCIMToken bool
}

func (s *oktaArgs) validateAndCheckDefaults(ctx context.Context, args *installPluginArgs) error {
	if s.apiToken == "" {
		if !s.scimEnabled {
			return trace.BadParameter("API token is required")
		}
		if s.userSync {
			return trace.BadParameter("User sync requires API token to be set")
		}
		if s.accessListSync {
			return trace.BadParameter("AccessList sync requires API token to be set")
		}
		if s.appGroupSync {
			return trace.BadParameter("AppGroup sync requires API token to be set")
		}
	}
	if s.accessListSync {
		if len(s.defaultOwners) == 0 {
			return trace.BadParameter("AccessList sync requires at least one default owner to be set")
		}
		if !s.appGroupSync {
			return trace.BadParameter("AppGroup sync is required for AccessList sync")
		}
		if !s.userSync {
			return trace.BadParameter("User sync is required for AccessList sync")
		}
	}
	if s.scimEnabled {
		if s.scimToken == "" {
			var err error
			s.scimToken, err = utils.CryptoRandomHex(32)
			if err != nil {
				return trace.Wrap(err)
			}
			s.autoGeneratedSCIMToken = true
		}
	}
	if s.scimToken != "" {
		s.scimEnabled = true
	}
	connector, err := args.authClient.GetSAMLConnector(ctx, s.samlConnector, false)
	if err != nil {
		return trace.Wrap(err)
	}
	if s.appID == "" {
		appID, ok := connector.GetMetadata().Labels[types.OktaAppIDLabel]
		if ok {
			s.appID = appID
		}
	}
	if s.scimToken != "" && s.appID == "" && s.userSync {
		msg := []string{
			"SCIM support requires App ID, which was not supplied and couldn't be deduced from the SAML connector",
			"Specify the App ID explicitly with --app-id",
			"SCIM support requires app-id to be set",
		}
		return trace.BadParameter(strings.Join(msg, "\n"))
	}
	return nil
}

func (p *PluginsCommand) InstallOkta(ctx context.Context, args installPluginArgs) error {
	oktaSettings := p.install.okta
	if err := oktaSettings.validateAndCheckDefaults(ctx, &args); err != nil {
		return trace.Wrap(err)
	}
	creds, err := generateCredentials(p.install.name, oktaSettings)
	if err != nil {
		return trace.Wrap(err)
	}
	settings := &types.PluginOktaSettings{
		OrgUrl: oktaSettings.org.String(),
		SyncSettings: &types.PluginOktaSyncSettings{
			SsoConnectorId:  oktaSettings.samlConnector,
			AppId:           oktaSettings.appID,
			SyncUsers:       oktaSettings.userSync,
			SyncAccessLists: oktaSettings.accessListSync,
			DefaultOwners:   oktaSettings.defaultOwners,
			GroupFilters:    oktaSettings.groupFilters,
			AppFilters:      oktaSettings.appFilters,
		},
	}
	req := &pluginsv1.CreatePluginRequest{
		Plugin: &types.PluginV1{
			SubKind: types.PluginSubkindAccess,
			Metadata: types.Metadata{
				Labels: map[string]string{
					types.HostedPluginLabel: "true",
				},
				Name: p.install.name,
			},
			Spec: types.PluginSpecV1{Settings: &types.PluginSpecV1_Okta{Okta: settings}},
		},
		StaticCredentialsList: creds,
		CredentialLabels: map[string]string{
			types.OktaOrgURLLabel: oktaSettings.org.String(),
		},
	}

	if _, err := args.plugins.CreatePlugin(ctx, req); err != nil {
		return trace.Wrap(err)
	}

	fmt.Printf("Successfully created Okta plugin %q\n\n", p.install.name)
	if oktaSettings.scimEnabled {
		pingResp, err := args.authClient.Ping(ctx)
		if err != nil {
			return trace.Wrap(err, "failed fetching cluster info")
		}
		scimBaseURL := fmt.Sprintf("https://%s/v1/webapi/scim/%s", pingResp.GetProxyPublicAddr(), p.install.name)
		fmt.Printf("SCIM Base URL: %s\n", scimBaseURL)
		fmt.Printf("SCIM Identifier field for users: %s\n", "userName")
		if oktaSettings.autoGeneratedSCIMToken {
			fmt.Printf("SCIM Bearer Token: %s\n", oktaSettings.scimToken)
		}
	}

	fmt.Println("\nSee https://goteleport.com/docs/application-access/okta/hosted-guide for help configuring provisioning in Okta")
	return nil
}

func generateCredentials(pluginName string, oktaSettings oktaArgs) ([]*types.PluginStaticCredentialsV1, error) {
	var creds []*types.PluginStaticCredentialsV1
	if oktaSettings.apiToken != "" {
		label := types.OktaCredPurposeAuth
		if !oktaSettings.appGroupSync {
			label = types.CredPurposeOKTAAPITokenWithSCIMOnlyIntegration
		}

		oktaAPICreds := &types.PluginStaticCredentialsV1{
			ResourceHeader: types.ResourceHeader{
				Metadata: types.Metadata{
					Name: pluginName,
					Labels: map[string]string{
						types.OktaCredPurposeLabel: label,
					},
				},
			},
			Spec: &types.PluginStaticCredentialsSpecV1{
				Credentials: &types.PluginStaticCredentialsSpecV1_APIToken{
					APIToken: oktaSettings.apiToken,
				},
			},
		}
		creds = append(creds, oktaAPICreds)
	}

	if oktaSettings.scimToken != "" {
		scimTokenHash, err := bcrypt.GenerateFromPassword([]byte(oktaSettings.scimToken), bcrypt.DefaultCost)
		if err != nil {
			return nil, trace.Wrap(err)
		}
		oktaSCIMCreds := &types.PluginStaticCredentialsV1{
			ResourceHeader: types.ResourceHeader{
				Metadata: types.Metadata{
					Name: pluginName + "-scim-token",
					Labels: map[string]string{
						types.OktaCredPurposeLabel: types.OktaCredPurposeSCIMToken,
					},
				},
			},
			Spec: &types.PluginStaticCredentialsSpecV1{
				Credentials: &types.PluginStaticCredentialsSpecV1_APIToken{
					APIToken: string(scimTokenHash),
				},
			},
		}
		creds = append(creds, oktaSCIMCreds)
	}
	return creds, nil
}
