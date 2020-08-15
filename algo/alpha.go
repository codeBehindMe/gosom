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

// Contains behaviour specific to the learning rate commonly denoted by alpha.

type Alpha64 struct {
	AlphaZero float64
	Alpha     float64
}

func NewAlpha64(alphaZero float64) *Alpha64 {
	return &Alpha64{
		AlphaZero: alphaZero,
		Alpha:     alphaZero,
	}
}

func (a *Alpha64) Decay(t int, lambda Lambda64) {
	a.Alpha = a.AlphaZero * math.Exp(-float64(t)/float64(lambda))
}

func (a *Alpha64) GetCurrentValue() float64 {
	return a.Alpha
}

func (a *Alpha64) DecayAndGetValue(t int, lambda64 Lambda64) float64 {
	a.Decay(t, lambda64)
	return a.GetCurrentValue()
}

