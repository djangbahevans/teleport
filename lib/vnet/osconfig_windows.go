// Teleport
// Copyright (C) 2024 Gravitational, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package vnet

import (
	"context"
	"log/slog"
	"os/exec"
	"strings"

	"github.com/gravitational/trace"
)

func configureOS(ctx context.Context, cfg *osConfig) error {
	// There is no need to remove IP addresses or routes, they will automatically be cleaned up when the
	// process exits and the TUN is deleted.

	if cfg.tunIPv4 != "" {
		log.InfoContext(ctx, "Setting IPv4 address for the TUN device.", "device", cfg.tunName, "address", cfg.tunIPv4)
		// TODO(nklaassen) handle proper CIDR ranges
		if err := runCommand(ctx,
			"netsh", "interface", "ip", "set", "address", cfg.tunName, "static", cfg.tunIPv4,
		); err != nil {
			return trace.Wrap(err)
		}
		if err := runCommand(ctx,
			"route", "add", "100.64.0.0", "mask", "255.192.0.0", cfg.tunIPv4,
		); err != nil {
			return trace.Wrap(err)
		}
	}

	if cfg.tunIPv6 != "" {
		log.InfoContext(ctx, "Setting IPv6 address for the TUN device.", "device", cfg.tunName, "address", cfg.tunIPv6)
		if err := runCommand(ctx,
			"netsh", "interface", "ipv6", "set", "address", cfg.tunName, cfg.tunIPv6,
		); err != nil {
			return trace.Wrap(err)
		}

		log.InfoContext(ctx, "Setting an IPv6 route for the VNet.")
		if err := runCommand(ctx,
			"netsh", "interface", "ipv6", "set", "route", cfg.tunIPv6+"/64", cfg.tunName, cfg.tunIPv6,
		); err != nil {
			return trace.Wrap(err)
		}
	}

	if err := configureDNS(ctx, cfg.tunName, cfg.dnsAddr); err != nil {
		return trace.Wrap(err, "configuring DNS")
	}

	return nil
}

func configureDNS(ctx context.Context, tunName, nameserver string) error {
	log.InfoContext(ctx, "Setting up DNS for the tun",
		"tunName", tunName,
		"addr", nameserver)
	if err := runCommand(ctx,
		"netsh", "interface", "ipv6", "set", "dns", "name="+tunName, "source=static", "addr="+nameserver, "validate=no",
	); err != nil {
		return trace.Wrap(err)
	}
	return nil
}

func (c *osConfigurator) doWithDroppedRootPrivileges(ctx context.Context, fn func() error) (err error) {
	// TODO(nklaassen): actually do with dropped privileges.
	return trace.Wrap(fn())
}

func runCommand(ctx context.Context, path string, args ...string) error {
	cmdString := strings.Join(append([]string{path}, args...), " ")
	log.InfoContext(ctx, "Running command", "cmd", cmdString)
	cmd := exec.CommandContext(ctx, path, args...)
	var output strings.Builder
	cmd.Stderr = &output
	cmd.Stdout = &output
	if err := cmd.Run(); err != nil {
		slog.WarnContext(ctx, "Failed to run osconfig command",
			"cmd", strings.Join(append([]string{path}, args...), " "),
			"output", output.String())
		//return trace.Wrap(err, `running "%s" output: %s`, strings.Join(append([]string{path}, args...), " "), output.String())
	}
	return nil
}
