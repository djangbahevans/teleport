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

package git

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"sync"

	"github.com/go-git/go-git/v5/plumbing/format/pktline"
	"github.com/go-git/go-git/v5/plumbing/protocol/packp"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/gravitational/trace"

	"github.com/gravitational/teleport"
	apievents "github.com/gravitational/teleport/api/types/events"
	"github.com/gravitational/teleport/lib/events"
	"github.com/gravitational/teleport/lib/utils/log"
)

type gitCommandEmitter struct {
	events.StreamEmitter
	discard apievents.Emitter
}

// NewEmitter returns an emitter for Git proxy usage.
func NewEmitter(emitter events.StreamEmitter) events.StreamEmitter {
	return &gitCommandEmitter{
		StreamEmitter: emitter,
		discard:       events.NewDiscardEmitter(),
	}
}

// EmitAuditEvent overloads EmitAuditEvent to only emit Git command events.
func (e *gitCommandEmitter) EmitAuditEvent(ctx context.Context, event apievents.AuditEvent) error {
	switch event.GetType() {
	case events.GitCommandEvent:
		return trace.Wrap(e.StreamEmitter.EmitAuditEvent(ctx, event))
	default:
		return trace.Wrap(e.discard.EmitAuditEvent(ctx, event))
	}
}

// CommandRecorder records Git commands.
type CommandRecorder interface {
	// WriteCloser is the basic interface for the recorder to receive payload.
	io.WriteCloser

	GetService() string
	GetPath() string
	GetActions() []*apievents.GitCommandAction
}

func NewCommandRecorder(command string) (CommandRecorder, error) {
	sshCommand, err := parseSSHCommand(command)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	// For now, only record details on the push. Fetch is not very interesting.
	if sshCommand.gitService == transport.ReceivePackServiceName {
		return newPushCommandRecorder(sshCommand), nil
	}
	return newBaseRecorder(sshCommand), nil
}

type baseRecorder struct {
	sshCommand *sshCommand
}

func newBaseRecorder(sshCommand *sshCommand) *baseRecorder {
	return &baseRecorder{
		sshCommand: sshCommand,
	}
}

func (r *baseRecorder) GetService() string {
	return r.sshCommand.gitService
}
func (r *baseRecorder) GetPath() string {
	return r.sshCommand.path
}
func (r *baseRecorder) GetActions() []*apievents.GitCommandAction {
	return nil
}
func (r *baseRecorder) Write(p []byte) (int, error) {
	return len(p), nil
}
func (r *baseRecorder) Close() error {
	return nil
}

type pushCommandRecorder struct {
	*baseRecorder

	logger    *slog.Logger
	payload   []byte
	close     chan struct{}
	closeOnce sync.Once
}

func newPushCommandRecorder(sshCommand *sshCommand) *pushCommandRecorder {
	return &pushCommandRecorder{
		baseRecorder: newBaseRecorder(sshCommand),
		logger:       slog.With(teleport.ComponentKey, "git:packp"),
		close:        make(chan struct{}),
	}
}

func (r *pushCommandRecorder) Close() error {
	r.closeOnce.Do(func() {
		close(r.close)
	})
	return nil
}

func (r *pushCommandRecorder) Write(p []byte) (int, error) {
	select {
	case <-r.close:
		if len(p) > 0 {
			r.logger.Log(context.Background(), log.TraceLevel, "Discarding packet protocol", "packet_length", len(p))
		}
		return len(p), nil
	default:
		r.logger.Log(context.Background(), log.TraceLevel, "Recording Git command in packet protocol", "packet", string(p))
		r.payload = append(r.payload, p...)
		// Only record the header and avoid caching the packfile. Close early.
		//
		// https://git-scm.com/docs/pack-protocol#_reference_update_request_and_packfile_transfer
		if bytes.HasSuffix(r.payload, pktline.FlushPkt) {
			r.Close()
		}
		return len(p), nil
	}
}

func (r *pushCommandRecorder) GetActions() (actions []*apievents.GitCommandAction) {
	<-r.close

	// Noop push (e.g. "Everything up-to-date")
	if bytes.Equal(r.payload, pktline.FlushPkt) {
		return nil
	}
	request := packp.NewReferenceUpdateRequest()
	if err := request.Decode(bytes.NewReader(r.payload)); err != nil {
		r.logger.WarnContext(context.Background(), "failed to decode push command", "error", err)
	}
	for _, command := range request.Commands {
		actions = append(actions, &apievents.GitCommandAction{
			Action:    string(command.Action()),
			Reference: string(command.Name),
			Old:       command.Old.String(),
			New:       command.New.String(),
		})
	}
	return
}
