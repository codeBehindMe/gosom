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

import "math"

// GetMinInFloat64Slice returns the smallest value and its index.
func GetMinInFloat64Slice(s []float64) (float64, int) {
	var minValue float64 = math.MaxFloat64
	var minIndex int = 0

	for i, v := range s {
		if v < minValue {
			minValue = v
			minIndex = i
		}
	}
	return minValue, minIndex
}

// FIXME: This function is duplicated.
func EuclidianDistance2D(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
