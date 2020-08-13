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
	"github.com/codeBehindMe/gosom/utilx"
	"math"
)

// Contains the behaviour specific to the neighbourhood radius commonly denoted
// by sigma.
// It is most appropriate to define the starting sigma denoted as sigma zero
// which then decays after some iteration number t and scaled by some time
// constant lambda.

type SigmaZero float64

// SigmaAfterDecay returns the value of sigma after some decay has taken place.
// The value t is the iteration number and lambda is the time constant.
func (s *SigmaZero) SigmaAfterDecay(t int, lambda Lambda64) float64 {
	return float64(*s) * math.Exp(-float64(t)/float64(lambda))
}

// NewSigmaZero initialises sigma zero based on a commonly used mechanism by
// taking the half the size of the largest dimension of the map.
func NewSigmaZero(width, height int) SigmaZero {
	return SigmaZero(float64(utilx.Max(width, height)) / 2)
}
