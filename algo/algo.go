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
	"github.com/codeBehindMe/gosom/mapx"
	"github.com/codeBehindMe/gosom/utilx"
	"math"
)

func StepForward(featureInstance []float64, m *[]mapx.NeuronDouble, mapx *mapx.Mapx, radius float64, learningRate float64) {
	bmu := BestMatchingUnit(featureInstance, *m)
	distances := GetDistanceOfNeighboursOfBMU(bmu, *mapx)
	influence := GetInfluenceOfBMU(distances, radius)
	UpdateWeights(influence, m, learningRate, featureInstance)
}

// FIXME: Missing test.
// BestMatchingUnit returns the index of the neuron which has the closest
// distance to the input vector.
func BestMatchingUnit(input []float64, m []mapx.NeuronDouble) int {
	distMatrix := make([]float64, len(m))

	// FIXME: We need to do input dimension safety check somewhere.
	// Not worth doing inside distance function since it will get called over
	// and over again.

	for i := 0; i < len(distMatrix); i++ {
		distMatrix[i] = CalcEuclidianDistance(input, m[i])
	}
	_, minIndex := utilx.GetMinInFloat64Slice(distMatrix)
	return minIndex
}

type LearningRate float64

func (l *LearningRate) DecayForIteration(t, lambda float64) float64 {
	return float64(*l) * math.Exp(-t/lambda)
}

func GetDistanceOfNeighboursOfBMU(bmuIndex int, m mapx.Mapx) []float64 {
	// FIXME: Recalculating expression is unidiomatic
	mask := make([]float64, m.NRows*m.NCols)
	bmuRow, bmuCol := m.IndexToRowCol(bmuIndex)
	for i := 0; i < len(mask); i++ {
		row, col := m.IndexToRowCol(i)
		mask[i] = utilx.EuclidianDistance2D(float64(bmuRow), float64(bmuCol), float64(row), float64(col))
	}
	return mask
}

func GetInfluenceOfBMU(distances []float64, s float64) []float64 {
	influence := make([]float64, len(distances))
	for i := 0; i < len(distances); i++ {
		influence[i] = math.Exp(-1 * math.Pow(distances[i], 2) / (2 * math.Pow(s, 2)))
	}
	return influence
}

func UpdateWeights(influence []float64, m *[]mapx.NeuronDouble, lr float64, trainingInstance []float64) {
	mpx := *m
	for i := 0; i < len(influence); i++ {
		for j := 0; j < len(mpx[i]); j++ {
			mpx[i][j] = mpx[i][j] + lr*influence[i]*(trainingInstance[j]-mpx[i][j])
		}
	}
}
