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

package logging

import (
	"io"
	"testing"
)

var testLogger = New().
	WithOutput(io.Discard).
	WithPrefix("benchmark")

func BenchmarkLog(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testLogger.Info(b.N)
	}
}

func BenchmarkLog100(b *testing.B) {
	for i := 0; i < 100; i++ {
		for n := 0; n < b.N; n++ {
			testLogger.Info(b.N)
		}
	}
}

func BenchmarkLog1000(b *testing.B) {
	for i := 0; i < 10000; i++ {
		for n := 0; n < b.N; n++ {
			testLogger.Info(b.N)
		}
	}
}
