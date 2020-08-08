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

package utilx

import "testing"

func TestGetMinInFloat64Slice(t *testing.T) {
	s := []float64{0.0, 100.1, 20.2, -3.24}
	minVal, minIndex := GetMinInFloat64Slice(s)
	if minVal != -3.24 {
		t.Errorf("Incorrect min value - got %v want %v", minVal, -3.24)
	}
	if minIndex != 3 {
		if minIndex != 3 {
			t.Errorf("Incorrect min index - got %v want %v", minIndex, 3)
		}
	}
}
