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
	"errors"
	"math"
)

// Contains behaviour specific to the learning rate commonly denoted by alpha.

type Alpha64 struct {
	AlphaZero float64
	Alpha     float64
}

func notZero64(v float64) error {
	if v == 0 {
		return errors.New("alpha cannot be zero")
	}
	return nil
}

func notNegative64(v float64) error {
	if v < 0 {
		return errors.New("alpha cannot be negative")
	}
	return nil
}

// NewAlpha64 creates a new Alpha parameter to be used with 64 bit
// wide SOMs.
func NewAlpha64(alphaZero float64) (*Alpha64, error) {
	err := notZero64(alphaZero)
	if err != nil {
		return nil, err
	}

	err = notNegative64(alphaZero)
	if err != nil {
		return nil, err
	}

	return &Alpha64{
		AlphaZero: alphaZero,
		Alpha:     alphaZero,
	}, nil
}

// Decay reduces the value of Alpha based on the value t.
// t is the value of the current iteration which decays alpha
// exponentially to that value t.
// Note: The parameter t takes how many iterations have passed
// and uses that value to caluclate the decay. Calling this function
// in a loop with t = 1 will set the same result as if only one
// iteration has passed.  FIXME: This is not intuitive at all.
func (a *Alpha64) Decay(t int, lambda float64) {
	a.Alpha = a.AlphaZero * math.Exp(-float64(t)/lambda)
}

// GetCurrentValue returns the current value of Alpha.
func (a *Alpha64) GetCurrentValue() float64 {
	return a.Alpha
}

// DecayAndGetValue is a helper function to the common case where
// the algorithm wants the value for Alpha after t iterations. This
// conveniently sets alpha to num iterations and then returns the
// value.
func (a *Alpha64) DecayAndGetValue(t int, lambda float64) float64 {
	a.Decay(t, lambda)
	return a.GetCurrentValue()
}
