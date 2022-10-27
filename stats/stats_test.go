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
	"encoding/json"
	"errors"
	"testing"
	"time"
)

type statsFs struct {
	Device string
}

type fsProvider struct{}

func (fsProvider) Stats() (interface{}, error) {
	return []statsFs{
		{Device: "device"},
	}, nil
}

func (fsProvider) Name() string {
	return "fs"
}

type errProvider struct{}

func (errProvider) Stats() (interface{}, error) {
	return nil, errors.New("error")
}

func (errProvider) Name() string {
	return "err"
}

type staticDataProvider struct {
	Data string
}

func (sp staticDataProvider) Stats() (interface{}, error) {
	return sp.Data, nil
}

func (staticDataProvider) Name() string {
	return "staticData"
}

func TestProviders(t *testing.T) {
	t.Run("standard mock provider", func(t *testing.T) {
		s := Stats{
			ID:       "id",
			Hostname: "hostname",
			Model:    "model",
			Uptime:   time.Hour,
		}

		err := s.Provide(fsProvider{})
		if err != nil {
			t.Fatal(err)
		}

		raw, err := json.Marshal(&s)
		if err != nil {
			t.Fatal(err)
		}

		out := Stats{}
		err = json.Unmarshal(raw, &out)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("error provider", func(t *testing.T) {
		s := Stats{
			ID:       "id",
			Hostname: "hostname",
			Model:    "model",
			Uptime:   time.Hour,
		}

		err := s.Provide(errProvider{}, errProvider{})
		if err == nil {
			t.Log(err)
			t.Fatal("err should not be nil")
		}
	})

	t.Run("static data provider", func(t *testing.T) {
		s := Stats{
			ID:       "id",
			Hostname: "hostname",
			Model:    "model",
			Uptime:   time.Hour,
		}

		exp := "data"

		err := s.Provide(staticDataProvider{
			Data: exp,
		})
		if err != nil {
			t.Log(err)
			t.Fatal("err should be nil")
		}

		got := s.Providers["staticData"]
		if got != exp {
			t.Logf("expected staticData provider to be '%s', got '%s'", exp, got)
			t.Fail()
		}
	})
}
