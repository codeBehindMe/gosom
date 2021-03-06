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

func TestNewIterationBasedLambda(t *testing.T) {
	maxIter := 10

	want := 10.
	got := float64(*NewIterationBasedLambda(maxIter, math.E))
	if got != want {
		t.Errorf("Incorrect lambda, got %v, want %v", got, want)
	}
}

func TestLambda64_GetValue(t *testing.T) {

	lambda := Lambda64(10.)
	want := 10.
	got := lambda.GetValue()
	if got != want {
		t.Errorf("Incorrect lambda: got %v, want %v", got, want)
	}
}
