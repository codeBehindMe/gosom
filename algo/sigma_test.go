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

func TestSigma64_SigmaAfterDecay(t *testing.T) {
	sigmaZero := SigmaZero(math.E)
	iter, lambda := 10, NewIterationBasedLambda(10, sigmaZero)

	want := 1.0 // 0.1 * e^(- 10/10)
	got := sigmaZero.SigmaAfterDecay(iter, lambda)

	if got != want {
		t.Errorf("Incorrect sigma after decay: got %v, want %v", got, want)
	}
}

func TestNewSigmaZero(t *testing.T) {
	width, height := 12, 10
	want := 6.
	got := NewSigmaZero(width, height)
	if float64(got) != want {
		t.Errorf("Incorrect sigma zero init: got %v, want %v", got, want)
	}
}
