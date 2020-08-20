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

package gosom

import (
	"encoding/json"
	"github.com/codeBehindMe/gosom/algo"
	"github.com/codeBehindMe/gosom/feed"
	"github.com/codeBehindMe/gosom/mapx"
	"io"
	"io/ioutil"
)

type SOM64 struct {
	Feed                  feed.Feeder
	Mapx                  *mapx.Mapx
	Radius                *algo.Sigma64
	LearningRate          *algo.Alpha64
	TimeConstant          *algo.Lambda64
	MaxTrainingIterations int
}

func NewSOM64(feeder feed.Feeder, nRows, nCols, maxIterations int, initialLearningRate float64, initialisationScheme mapx.Scheme) *SOM64 {
	sigma := algo.NewSigma64(nCols, nRows)
	alpha := algo.NewAlpha64(initialLearningRate)
	lambda := algo.NewIterationBasedLambda(maxIterations, sigma.SigmaZero)
	somMap := mapx.New(nRows, nCols, feeder.GetFeatureSize())
	_ = somMap.Initialise(initialisationScheme)
	return &SOM64{
		Feed:                  feeder,
		Mapx:                  somMap,
		Radius:                sigma,
		LearningRate:          alpha,
		TimeConstant:          lambda,
		MaxTrainingIterations: maxIterations,
	}
}

func (s *SOM64) Train() {
trainingLoop:
	for t := 0; t < s.MaxTrainingIterations; {
		pipe := s.Feed.MakeFeaturePipe()
		go s.Feed.Start(pipe)
		for feature := range pipe {
			if t >= s.MaxTrainingIterations {
				close(pipe)
				break trainingLoop
			}
			algo.StepForward(feature, &s.Mapx.Data, s.Mapx, s.Radius.DecayAndGetValue(t, s.TimeConstant.GetValue()), s.LearningRate.DecayAndGetValue(t, s.TimeConstant.GetValue()))
			t++
		}
	}
}

func (s *SOM64) DumpMapToFile(path string) error {
	data := s.Mapx.Data

	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (s *SOM64) WriteMap(w io.Writer) (int, error) {

	data := s.Mapx.Data

	b, err := json.Marshal(data)
	if err != nil {
		return -1, err
	}

	n, err := w.Write(b)
	if err != nil {
		return n, err
	}
	return n, nil
}
