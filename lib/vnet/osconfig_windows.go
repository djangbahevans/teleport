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
	"os/exec"

	"github.com/gravitational/trace"
)

func configureOS(ctx context.Context, cfg *osConfig) error {
	// There is no need to remove IP addresses or routes, they will automatically be cleaned up when the
	// process exits and the TUN is deleted.

	if cfg.tunIPv4 != "" {
		log.InfoContext(ctx, "Setting IPv4 address for the TUN device.", "device", cfg.tunName, "address", cfg.tunIPv4)
		// TODO(nklaassen) handle proper CIDR ranges
		cmd := exec.CommandContext(ctx,
			"netsh", "interface", "ip", "set", "address", cfg.tunName, "static", cfg.tunIPv4, "255.192.0.0", cfg.tunIPv4)
		if err := cmd.Run(); err != nil {
			return trace.Wrap(err, "running %v", cmd.Args)
		}
	}

	if cfg.tunIPv6 != "" {
		log.InfoContext(ctx, "Setting IPv6 address for the TUN device.", "device", cfg.tunName, "address", cfg.tunIPv6)
		cmd := exec.CommandContext(ctx,
			"netsh", "interface", "ipv6", "set", "address", cfg.tunName, cfg.tunIPv6)
		if err := cmd.Run(); err != nil {
			return trace.Wrap(err, "running %v", cmd.Args)
		}

		log.InfoContext(ctx, "Setting an IPv6 route for the VNet.")
		cmd = exec.CommandContext(ctx,
			"netsh", "interface", "ipv6", "set", "route", cfg.tunIPv6+"/64", cfg.tunName, cfg.tunIPv6)
		if err := cmd.Run(); err != nil {
			return trace.Wrap(err, "running %v", cmd.Args)
		}
	}

	if err := configureDNS(ctx, cfg.dnsAddr, cfg.dnsZones); err != nil {
		return trace.Wrap(err, "configuring DNS")
	}

	return nil
}

func configureDNS(ctx context.Context, nameserver string, zones []string) error {
	// TODO(nklaassen): actually configure DNS.
	return nil
}

func (c *osConfigurator) doWithDroppedRootPrivileges(ctx context.Context, fn func() error) (err error) {
	// TODO(nklaassen): actually do with dropped privileges.
	return trace.Wrap(fn())
}
