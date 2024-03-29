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

package types

import "time"

// Type AggregatedMeasurementRequest describes a HTTP request for a measurement
// campaign on multiple sensors, where the results are averaged over the specified
// time resolution.
type AggregatedMeasurementRequest struct {
	// List of sensor hardware IDs to run the measurement campaign on
	Sensors []string `json:"sensors"`

	// Start time in ISO 8601
	Begin time.Time `json:"begin"`

	// End time in ISO 8601
	End time.Time `json:"end"`

	// Lower bound for frequency in Hz
	FreqMin int64 `json:"freqMin"`

	// Upper bound for frequency in Hz
	FreqMax int64 `json:"freqMax"`

	// Frequency resolution in Hz
	FreqRes int64 `json:"freqRes"`

	// Time resolution in seconds
	TimeRes int64 `json:"timeRes"`

	// Campaign ID. For internal use only, will be ignored if not null
	CampaignId string `json:"campaignId"`

	// AggregationFunc? (defaults to AVG/average)
}

// Type RawMeasurementRequest describes a HTTP request for a measurement
// campaign on multiple sensors, where the results are not sampled and
// returned just as the sensor fetched them.
type RawMeasurementRequest struct {
	// List of sensor hardware IDs to run the measurement campaign on
	Sensors []string `json:"sensors"`

	// Center frequency for measurement
	FreqCenter int64

	// Start time in ISO 8601
	Begin time.Time `json:"begin"`

	// End time in ISO 8601
	End time.Time `json:"end"`

	// Campaign ID. For internal use only, will be ignored if not null
	CampaignId string `json:"campaignId"`
}
