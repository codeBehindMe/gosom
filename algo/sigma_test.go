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

import "testing"

func TestNewSigma64(t *testing.T) {
	sigma := NewSigma64(6, 5)

	got := sigma.SigmaZero
	want := 3.
	if got != want {
		t.Errorf("Incorrect sigma: got %v, want %v", got, want)
	}
}

func TestSigma64_Decay(t *testing.T) {
	sigma := NewSigma64(6, 4)

	sigma.Decay(0, Lambda64(0.1))

	got := sigma.SigmaZero
	want := 3.

	if got != want {
		t.Errorf("Incorrect sigma: got %v, want %v", got, want)
	}
}
