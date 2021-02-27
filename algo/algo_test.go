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
	"testing"

	"github.com/codeBehindMe/gosom/mapx"
)

func TestBestMatchingUnit(t *testing.T) {
	testNeurons := []mapx.NeuronDouble{
		{0.1, 0.1, 0.1},
		{0.9, 0.9, 0.9},
		{0.1, 0.1, 0.1},
	}
	testInput := []float64{0.8, 0.8, 0.8}
	bmu := BestMatchingUnit(testInput, testNeurons)
	if bmu != 1 {
		t.Errorf("Failed to get best matching unit - got %v, want %v", bmu, 1)
	}
}

func TestGetDistanceOfNeighboursOfBMU(t *testing.T) {
	mpx := mapx.New(10, 10, 3)
	_ = mpx.Initialise(mapx.OnesInitialiser)
	bmuIndex := 0

	distances := GetDistanceOfNeighboursOfBMU(bmuIndex, *mpx)
	_ = distances
}

func TestGetInfluenceOfBMU(t *testing.T) {

	mpx := mapx.New(10, 10, 3)
	_ = mpx.Initialise(mapx.OnesInitialiser)
	bmuIndex := 0
	s := NewSigma64(10, 10)

	distances := GetDistanceOfNeighboursOfBMU(bmuIndex, *mpx)
	influence := GetInfluenceOfBMU(distances, s.GetCurrentValue())

	_ = influence
}

func TestUpdateWeights(t *testing.T) {
	mpx := mapx.New(10, 10, 3)
	_ = mpx.Initialise(mapx.OnesInitialiser)
	trInstance := []float64{0.1, 0.1, 0.1}
	bmuIndex := BestMatchingUnit(trInstance, mpx.Data)
	lr, err := NewAlpha64(0.1)
	if err != nil {
		t.Error(err)
	}
	sigma := NewSigma64(10, 10)

	distances := GetDistanceOfNeighboursOfBMU(bmuIndex, *mpx)

	influence := GetInfluenceOfBMU(distances, sigma.GetCurrentValue())

	UpdateWeights(influence, &mpx.Data, lr.DecayAndGetValue(0, 1.), trInstance)
}
