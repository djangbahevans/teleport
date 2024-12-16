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

package agent

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gravitational/trace"
)

// IsActive returns true if the local Teleport binary is managed by teleport-update.
// Note that true may be returned even if auto-updates is disabled or the version is pinned.
func IsActive() (bool, error) {
	teleportPath, err := os.Readlink("/proc/self/exe")
	if err != nil {
		return false, trace.Wrap(err, "cannot find Teleport binary")
	}
	updaterBasePath := filepath.Clean(teleportOptDir) + "/"
	absPath, err := filepath.Abs(teleportPath)
	if err != nil {
		return false, trace.Wrap(err, "cannot get absolute path for Teleport binary")
	}
	if !strings.HasPrefix(absPath, updaterBasePath) {
		return false, nil
	}
	systemDir := filepath.Join(teleportOptDir, systemNamespace)
	return !strings.HasPrefix(absPath, systemDir), nil
}
