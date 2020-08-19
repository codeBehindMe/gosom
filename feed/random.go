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

package feed

import "math/rand"

// RandomFeeder generates features drawing from a random distribution. This is
// useful when you just want to just run and test the algorithm. The features
// are drawn from a uniform distribution of real numbers element of [0,1)
type RandomFeeder struct {
	FeatureSize  int
	NumInstances int
}

func (r RandomFeeder) MakeFeaturePipe() FeaturePipe {
	return make(FeaturePipe)
}

func (r RandomFeeder) Start(sc FeaturePipe) {
	arr := make([]float64, r.FeatureSize)
	for i := 0; i < r.NumInstances; i++ {
		for j := 0; j < r.FeatureSize; j++ {
			arr[j] = rand.Float64()
		}
		sc <- arr
	}
	close(sc)
}

func (r RandomFeeder) GetFeatureSize() int {
	return r.FeatureSize
}
