// Copyright 2024 Gravitational, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package net

import (
	"github.com/gravitational/trace"
)

// ValidatePortRange checks if the given port range is within bounds. If endPort is not zero, it
// also checks if it's bigger than port. A port range with zero as endPort is assumed to describe a
// single port.
func ValidatePortRange(port, endPort int) error {
	const minPort = 1
	const maxPort = 65535

	if port < minPort || port > maxPort {
		return trace.BadParameter("port must be between %d and %d, but got %d", minPort, maxPort, port)
	}

	if endPort != 0 {
		if endPort < minPort+1 || endPort > maxPort {
			return trace.BadParameter("end port must be between %d and %d, but got %d", minPort+1, maxPort, endPort)
		}

		if endPort <= port {
			return trace.BadParameter("end port must be greater than port (%d vs %d)", endPort, port)
		}
	}

	return nil
}
