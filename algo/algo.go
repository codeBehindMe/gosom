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
	"encoding/json"
	"github.com/codeBehindMe/gosom/feed"
	"github.com/codeBehindMe/gosom/mapx"
	"github.com/codeBehindMe/gosom/utilx"
	"io"
	"io/ioutil"
	"math"
)

type Trainer interface {
	Train()
}

type Learner interface {
	SetLearningRate(r float64)
	GetLearningRate() float64
}

type SOM struct {
	Feed       feed.Feeder
	Mapx       *mapx.Mapx
	Radius     Sigma
	LR         LearningRate
	Iterations int
	Lambda     float64
}

type JSONWeight struct {
	Row     int       `json:"row"`
	Col     int       `json:"col"`
	Weights []float64 `json:"weights"`
}

// Train the SOM
func (s *SOM) Train() {

	for t := 0; t < s.Iterations; t++ {
		pipe := s.Feed.MakeFeaturePipe()
		go s.Feed.Start(pipe)

		for feature := range pipe {
			trainingStep(feature, &s.Mapx.Data, s.Mapx, s.Radius.DecayedForIteration(float64(t), s.Lambda), s.LR)
			s.LR.Decay(float64(t), s.Lambda)
		}
	}
}

func NewSOM(feeder feed.Feeder, height int, width int, weights int, initialisationScheme mapx.Scheme, maxIter int, initialLearningRate LearningRate) *SOM {

	mpx := mapx.New(height, width, weights)
	err := mpx.Initialise(initialisationScheme)
	if err != nil {
		panic(err)
	}
	// FIXME: Make this max(width,height)/2
	sigmaZero := float64(width) / 2
	lambda := float64(maxIter) / math.Log(sigmaZero)
	return &SOM{
		Feed:       feeder,
		Mapx:       mpx,
		Radius:     Sigma(sigmaZero),
		LR:         initialLearningRate,
		Iterations: maxIter,
		Lambda:     lambda,
	}
}

func (s *SOM) dumpWeights(r *io.WriteCloser) {
	// FIXME: Implement this much cleaner and reusable version.
}

// FIXME: Docstring
// Dump the weights of the map to a json file.
func (s *SOM) DumpWeightsToFile(path string) {
	data := s.Mapx.Data

	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(path, b, 0644)
	if err != nil {
		panic(err)
	}
}

func trainingStep(featureInstance []float64, m *[]mapx.NeuronDouble, mapx *mapx.Mapx, radius float64, learningRate LearningRate) {
	bmu := bestMatchingUnit(featureInstance, *m)
	distances := GetDistanceOfNeighboursOfBMU(bmu, *mapx)
	influence := GetInfluenceOfBMU(distances, radius)
	updateWeights(influence, m, learningRate, featureInstance)
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

func (s *Sigma) DecayedForIteration(t float64, lambda float64) float64 {
	return float64(*s) * math.Exp(-t/lambda)
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

func GetInfluenceOfBMU(distances []float64, s float64) []float64 {
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
