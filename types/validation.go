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

import (
	"time"

	v "github.com/go-ozzo/ozzo-validation/v4"
)

var _ v.Validatable = &AggregatedMeasurementRequest{}

// Validates the measurement request
func (amr AggregatedMeasurementRequest) Validate() error {
	return v.ValidateStruct(&amr,
		v.Field(&amr.Begin, v.Required, v.Min(0), v.Max(amr.End)),
		v.Field(&amr.End, v.Required, v.Min(amr.Begin), v.Max(time.Now().UnixMilli())),
		v.Field(&amr.FreqMin, v.Required, v.Min(0), v.Max(amr.FreqMax)),
		v.Field(&amr.FreqMax, v.Required, v.Min(amr.FreqMin)),
		v.Field(&amr.FreqRes, v.Required, v.Max(amr.FreqMax-amr.FreqMin)),
		v.Field(&amr.TimeRes, v.Required, v.Min(0)),
	)
}
