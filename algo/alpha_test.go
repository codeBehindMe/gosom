/*
   gosom
   Copyright (C) 2020  aarontillekeratne

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

/*
  Author: aarontillekeratne
  Contact: github.com/codeBehindMe
*/

package algo

import (
	"math"
	"testing"
)

func TestNewAlpha64(t *testing.T) {
	alpha, err := NewAlpha64(0.1)
	if err != nil {
		t.Error(err)
	}

	want := 0.1
	if alpha.AlphaZero != want {
		t.Errorf("Incorrect alpha zero: got %v, want %v", alpha.AlphaZero, want)
	}

	if alpha.Alpha != want {
		t.Errorf("Incorrect alpha: got %v, want %v", alpha.Alpha, want)
	}
}

func TestAlpha64_Decay(t *testing.T) {
	alpha, err := NewAlpha64(0.1)
	if err != nil {
		t.Error(t)
	}

	alpha.Decay(0, math.E)
	got := alpha.Alpha
	want := 0.1

	if got != want {
		t.Errorf("Incorrect alpha: got %v, want %v", got, want)
	}
}

func TestAlpha64_DecayAndGetValue(t *testing.T) {
	alpha, err := NewAlpha64(0.1)
	if err != nil {
		t.Error(err)
	}

	got := alpha.DecayAndGetValue(0, math.E)
	want := 0.1

	if got != want {
		t.Errorf("Incorrect alpha: got %v, want %v", got, want)
	}
}

func TestAlpha64_GetCurrentValue(t *testing.T) {
	alpha, err := NewAlpha64(0.1)
	if err != nil {
		t.Error(err)
	}

	got := alpha.GetCurrentValue()
	want := 0.1
	if got != want {
		t.Errorf("Incorrect  alpha: got %v, want %v", got, want)
	}
}

// Alpha cannot be zero
func TestAlpha64_Zero(t *testing.T) {
	_, err := NewAlpha64(0)
	if err == nil {
		t.Fail()
		t.Logf("cannot allow alpha to take zero value")
	}
}
