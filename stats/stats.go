// Copyright (C) 2022 OpenRFSense
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package stats

import (
	"time"
)

// Interface Provider describes a generic stats provider which can add more
// information to the Stats type.
type Provider interface {
	// Returns a unique name for the provider, so it can store stats at Stats.Providers[name]
	Name() string

	// The actual stats collected by the provider.
	Stats() interface{}
}

// Type Stats contains in-depth information about a node's hardware and identity.
type Stats struct {
	// A unique identifier for the node (a hardware-bound ID is recommended)
	ID string `json:"id"`

	// Hostname of the system
	Hostname string `json:"hostname"`

	// The model/vendor of the system's hardware, useful for identification
	Model string `json:"model"`

	// Uptime of the system
	Uptime time.Duration `json:"uptime"`

	// Extra, more in-depth information about the system as dynamically returned by providers.
	Providers map[string]interface{} `json:"providers,omitempty"`
}

// Executes the given providers and stores the returned stats in Stats.Providers.
func (s *Stats) Provide(providers ...Provider) {
	if len(providers) == 0 {
		return
	}

	if s.Providers == nil {
		s.Providers = make(map[string]interface{})
	}

	for _, p := range providers {
		s.Providers[p.Name()] = p.Stats()
	}
}
