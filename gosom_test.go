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
	"github.com/codeBehindMe/gosom/feed"
	"github.com/codeBehindMe/gosom/mapx"
	"testing"
)

func TestNewSOM64(t *testing.T) {
	feeder := feed.RandomFeeder{
		FeatureSize:  3,
		NumInstances: 10,
	}
	som := NewSOM64(feeder, 10, 10, 10, 0.1, mapx.RandomNormalInitialiser)

	if som.TimeConstant.GetValue() != 6.213349345596119 {
		t.Errorf("Incorrect lambda: got %v, want %v", som.TimeConstant.GetValue(), 6.213349345596119)
	}
}

func TestSOM64_Train(t *testing.T) {

	feeder := feed.RandomFeeder{
		FeatureSize:  3,
		NumInstances: 10,
	}
	som := NewSOM64(feeder, 10, 10, 10, 0.1, mapx.RandomNormalInitialiser)

	if som.TimeConstant.GetValue() != 6.213349345596119 {
		t.Errorf("Incorrect lambda: got %v, want %v", som.TimeConstant.GetValue(), 6.213349345596119)
	}

	som.Train()
}
