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

// FIXME: GoDocs
// Contains the behaviour specific to the neighbourhood radius commonly denoted
// by sigma.
// It is most appropriate to define the starting sigma denoted as sigma zero
// which then decays after some iteration number t and scaled by some time
// constant lambda.

type Sigma64 struct {
	SigmaZero float64
	Sigma     float64
}

func NewSigma64(width, height int) *Sigma64 {
	sigmaZero := float64(utilx.Max(width, height)) / 2
	return &Sigma64{
		SigmaZero: sigmaZero,
		Sigma:     sigmaZero,
	}
}

// FIXME: Potential confusion with decay
// This needs some serious documentation since decay indicates that the state
// contained in this variable is important. However, that is *NOT* the case.
// All variables that impact the decayed value are passed in and only the
// initial value of sigma is used.
func (s *Sigma64) Decay(t int, lambda float64) {
	s.Sigma = s.SigmaZero * math.Exp(-float64(t)/lambda)
}

func (s *Sigma64) GetCurrentValue() float64 {
	return s.Sigma
}
func (s *Sigma64) DecayAndGetValue(t int, lambda float64) float64 {
	s.Decay(t, lambda)
	return s.GetCurrentValue()
}
