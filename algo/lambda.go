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

import "math"

// Contains the behaviour specific to the time constant commonly denoted by
// lambda.

type Lambda64 float64

func NewIterationBasedLambda(maxIterations int, sigmaZero float64) *Lambda64 {
	var a = new(Lambda64)
	*a = Lambda64(float64(maxIterations) / math.Log(sigmaZero))
	return a
}

func (l *Lambda64) GetValue()float64{
	return float64(*l)
}