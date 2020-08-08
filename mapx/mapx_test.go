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

package mapx

import "testing"

func TestNew(t *testing.T) {
	mpx := New(10, 10, 3)

	if mpx.NRows != 10 {
		t.Errorf("Row count mismatch - got: %v, want %v", mpx.NRows, 10)
		t.Fail()
	}
	if mpx.NCols != 10 {
		t.Errorf("Col count mismatc - got: %v, want %v", mpx.NCols, 10)
	}

	if len(mpx.data) != 10*10 {
		t.Errorf("Data container size mismatch - got: %v , want %v", len(mpx.data), 10*10)
	}
}

// FIXME: Do more test cases with the other initialisation schemes.
func TestMapx_Initialise(t *testing.T) {
	// FIXME: This is just testing ones.
	mpx := New(10, 10, 3)
	err := mpx.Initialise(OnesInitialiser)
	if err != nil {
		t.Errorf("initialisation failed with error: %v", err)
	}
	for i := range mpx.data {
		neuron := mpx.data[i]
		for _, weight := range neuron {
			if weight != 1 {
				t.Errorf("incorrect initialisation- got %v, want %v", weight, 1)
			}
		}
	}
}
