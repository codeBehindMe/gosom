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

type Trainer interface {
	Train()
}

type Learner interface {
	SetLearningRate(r float64)
	GetLearningRate() float64
}

// FIXME: Missing test.
// bestMatchingUnit returns the index of the neuron which has the closest
// distance to the input vector.
func bestMatchingUnit(input []float64, m []mapx.NeuronDouble) int {
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

type Sigma float64

func (s *Sigma) Decay(t float64, lambda float64) {
	*s = Sigma(float64(*s) * math.Exp(-t/lambda))
}

type LearningRate float64

func (l *LearningRate) Decay(t, lambda float64) {
	*l = LearningRate(float64(*l) * math.Exp(-t/lambda))
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

func GetInfluenceOfBMU(distances []float64, s Sigma) []float64 {
	influence := make([]float64, len(distances))
	for i := 0; i < len(distances); i++ {
		influence[i] = math.Exp(-1 * math.Pow(distances[i], 2) / (2 * math.Pow(float64(s), 2)))
	}
	return influence
}

func updateWeights(influence []float64, m *[]mapx.NeuronDouble, lr LearningRate, trainingInstance []float64) {
	mpx := *m
	for i := 0; i < len(influence); i++ {
		for j := 0; j < len(mpx[i]); j++ {
			mpx[i][j] = mpx[i][j] + float64(lr)*influence[i]*(trainingInstance[j]-mpx[i][j])
		}
	}
}
