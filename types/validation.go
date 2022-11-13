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
	"fmt"
	"time"

	v "github.com/go-ozzo/ozzo-validation/v4"
)

var (
	_ v.Validatable = &AggregatedMeasurementRequest{}
	_ v.Validatable = &RawMeasurementRequest{}
)

// Returns error if "begin" is after "after".
func isBefore(end time.Time) v.RuleFunc {
	return func(value interface{}) error {
		begin, _ := value.(time.Time)
		if begin.After(end) {
			return fmt.Errorf("Begin must be before End")
		}
		return nil
	}
}

// Returns error if "end" is before "begin".
func isAfter(begin time.Time) v.RuleFunc {
	return func(value interface{}) error {
		end, _ := value.(time.Time)
		if end.Before(begin) {
			return fmt.Errorf("End must be after Begin")
		}
		return nil
	}
}

// Validates the measurement request.
func (amr AggregatedMeasurementRequest) Validate() error {
	return v.ValidateStruct(&amr,
		v.Field(&amr.Begin, v.Required, v.By(isBefore(amr.End))),
		v.Field(&amr.End, v.Required, v.By(isAfter(amr.Begin))),
		v.Field(&amr.FreqMin, v.Required, v.Min(0), v.Max(amr.FreqMax)),
		v.Field(&amr.FreqMax, v.Required, v.Min(amr.FreqMin)),
		v.Field(&amr.FreqRes, v.Required, v.Max(amr.FreqMax-amr.FreqMin)),
		v.Field(&amr.TimeRes, v.Required, v.Min(0)),
	)
}

// Validates the measurement request.
func (rmr RawMeasurementRequest) Validate() error {
	return v.ValidateStruct(&rmr,
		v.Field(&rmr.Begin, v.Required, v.Min(0), v.Max(rmr.End)),
		v.Field(&rmr.End, v.Required, v.Min(rmr.Begin)),
		v.Field(&rmr.FreqCenter, v.Required, v.Min(0)),
	)
}
