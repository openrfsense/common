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
	"testing"
	"time"
)

func TestValidateAggregatedMeasurementRequest(t *testing.T) {
	t.Run("valid measurement request", func(t *testing.T) {
		now := time.Now()
		amr := AggregatedMeasurementRequest{
			Begin:   now.Add(-time.Minute),
			End:     now,
			FreqMin: 10e8,   // 100MHz
			FreqMax: 16e8,   // 160Mhz
			FreqRes: 100000, // 100kHz
			TimeRes: 30,     // 30 seconds
		}

		err := amr.Validate()
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("begin > end", func(t *testing.T) {
		now := time.Now()
		amr := AggregatedMeasurementRequest{
			Begin:   now,
			End:     now.Add(-time.Minute),
			FreqMin: 18e8,   // 180MHz
			FreqMax: 16e8,   // 160Mhz
			FreqRes: 100000, // 100kHz
			TimeRes: 30,     // 30 seconds
		}

		err := amr.Validate()
		if err == nil {
			t.Fatalf("begin (%v) must be later than end (%v)", amr.Begin, amr.End)
		}
	})

	t.Run("freqMin > freqMax", func(t *testing.T) {
		now := time.Now()
		amr := AggregatedMeasurementRequest{
			Begin:   now.Add(-time.Minute),
			End:     now,
			FreqMin: 18e8,   // 180MHz
			FreqMax: 16e8,   // 160Mhz
			FreqRes: 100000, // 100kHz
			TimeRes: 30,     // 30 seconds
		}

		err := amr.Validate()
		if err == nil {
			t.Fatalf("freqMin (%d) must be higher than freqMax (%d)", amr.FreqMin, amr.FreqMax)
		}
	})

	t.Run("freqRes too high (freqMax - freqMin)", func(t *testing.T) {
		now := time.Now()
		amr := AggregatedMeasurementRequest{
			Begin:   now.Add(-time.Minute),
			End:     now,
			FreqMin: 10e8, // 100MHz
			FreqMax: 16e8, // 160Mhz
			FreqRes: 10e8, // 100MHz > 60MHz
			TimeRes: 30,   // 30 seconds
		}

		err := amr.Validate()
		if err == nil {
			t.Fatalf("freqRes (%d) must be higher than freqMax (%d) - freqMin (%d)", amr.FreqRes, amr.FreqMax, amr.FreqMin)
		}
	})
}

func TestValidateRawMeasurementRequest(t *testing.T) {
	t.Run("valid measurement request", func(t *testing.T) {
		now := time.Now()
		rmr := RawMeasurementRequest{
			Begin:      now.Add(-time.Minute),
			End:        now,
			FreqCenter: -1,
		}

		err := rmr.Validate()
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("begin > end", func(t *testing.T) {
		now := time.Now()
		rmr := RawMeasurementRequest{
			Begin:      now,
			End:        now.Add(-time.Minute),
			FreqCenter: -1,
		}

		err := rmr.Validate()
		if err == nil {
			t.Fatalf("begin (%v) must be later than end (%v)", rmr.Begin, rmr.End)
		}
	})

	t.Run("invalid freqCenter", func(t *testing.T) {
		now := time.Now()
		rmr := RawMeasurementRequest{
			Begin:      now,
			End:        now.Add(-time.Minute),
			FreqCenter: -1,
		}

		err := rmr.Validate()
		if err == nil {
			t.Fatalf("begin (%v) must be later than end (%v)", rmr.Begin, rmr.End)
		}
	})
}
