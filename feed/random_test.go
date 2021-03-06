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

package feed_test

import (
	"github.com/codeBehindMe/gosom"
	"github.com/codeBehindMe/gosom/feed"
	"github.com/codeBehindMe/gosom/mapx"
	"testing"
)

func TestRandomFeeder_Start(t *testing.T) {
	randFeeder := feed.RandomFeeder{
		FeatureSize:  3,
		NumInstances: 1000,
	}
	_ = randFeeder
	som := gosom.NewSOM64(randFeeder, 10, 10, 10, 0.1, mapx.PseudoZerosOnesInitialiser)
	som.Train()
	som.DumpMapToFile("weights.json")
}
