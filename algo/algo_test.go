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
	"fmt"
	"github.com/codeBehindMe/gosom/feed"
	"github.com/codeBehindMe/gosom/mapx"
	"testing"
)

// FIXME: Convert to test right now its just printing a value.
func TestSigma_Decay(t *testing.T) {
	var s Sigma = 10.0
	s.Decay(1, 10)
	fmt.Printf("%v", s)
}

func TestLearningRate_Decay(t *testing.T) {
	var lr LearningRate = 10.0
	lr.Decay(1, 10)
	fmt.Printf("%v", lr)
}

func TestBestMatchingUnit(t *testing.T) {
	testNeurons := []mapx.NeuronDouble{
		{0.1, 0.1, 0.1},
		{0.9, 0.9, 0.9},
		{0.1, 0.1, 0.1},
	}
	testInput := []float64{0.8, 0.8, 0.8}
	bmu := bestMatchingUnit(testInput, testNeurons)
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
	s := Sigma(5)

	distances := GetDistanceOfNeighboursOfBMU(bmuIndex, *mpx)
	influence := GetInfluenceOfBMU(distances, s)

	_ = influence
}

func TestUpdateWeights(t *testing.T) {
	mpx := mapx.New(10, 10, 3)
	_ = mpx.Initialise(mapx.OnesInitialiser)
	trInstance := []float64{0.1, 0.1, 0.1}
	bmuIndex := bestMatchingUnit(trInstance, mpx.Data)
	lr := LearningRate(0.1)
	sigma := Sigma(5) // max(10,10)/2

	distances := GetDistanceOfNeighboursOfBMU(bmuIndex, *mpx)

	influence := GetInfluenceOfBMU(distances, sigma)

	updateWeights(influence, &mpx.Data, lr, trInstance)
}

func TestSOM_Train(t *testing.T) {
	som := NewSOM(feed.CSVFileFeeder{
		Filename:    "test.csv",
		FeatureSize: 3,
	}, 10, 10, 3, mapx.PseudoZerosOnesInitialiser, 50, 0.1)

	_ = som
	som.Train()

	som.DumpWeightsToFile("weights.json")
}
