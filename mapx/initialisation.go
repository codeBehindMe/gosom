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

import (
	"fmt"
	"math/rand"
)

type Scheme int

const (
	ZeroesInitialiser = iota
	OnesInitialiser
	RandomNormalInitialiser
	PseudoZerosOnesInitialiser
)

// TODO: Initialise should take options which could describe additional
// behaviour to initialisation methods.

func (p *Mapx) Initialise(s Scheme) error {
	switch s {
	case ZeroesInitialiser:
		return fmt.Errorf("unsupported initalisation")
	case OnesInitialiser:
		initialiseOnes(&p.Data, p.NWeights)
		return nil
	case RandomNormalInitialiser:
		initialiseRandomNormal(&p.Data, p.NWeights)
		return fmt.Errorf("unsupported initialisation")
	case PseudoZerosOnesInitialiser:
		initialisePseudoZeroOnes(&p.Data, p.NWeights)
	}
	return fmt.Errorf("undefined initialisation scheme: %v", s)
}

// initialiseOnes initialises the weights in the map with ones.
func initialiseOnes(d *[]NeuronDouble, numWeights int) {
	data := *d
	ones := make(NeuronDouble, numWeights)
	for i := 0; i < numWeights; i++ {
		ones[i] = float64(1)
	}
	for i := 0; i < len(data); i++ {
		data[i] = ones
	}
}

// FIXME: Requires test
// initialiseRandomNormal initialises the weights in the map by sampling from a
//random normal distribution with a mean of 0 and standard deviation of 1.
func initialiseRandomNormal(d *[]NeuronDouble, numWeights int) {
	data := *d
	for i := 0; i < len(data); i++ {
		neuron := make(NeuronDouble, numWeights)
		for j := 0; j < numWeights; j++ {
			neuron[j] = rand.NormFloat64()
		}
		data[i] = neuron
	}
}

// FIXME: Requires test
// initialisePseudoZeroOnes initialises the weights with values from [0.0,1.0)
func initialisePseudoZeroOnes(d *[]NeuronDouble, numWeights int) {
	data := *d
	for i := 0; i < len(data); i++ {
		neuron := make(NeuronDouble, numWeights)
		for j := 0; j < numWeights; j++ {
			neuron[j] = rand.Float64()
		}
		data[i] = neuron
	}
}
