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

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Implements feeding to the algorithm.

type Feeder interface {
	GetFeed() chan []float64
}

type FeaturePipe chan []float64

type CSVFileFeeder struct {
	Filename    string
	FeatureSize int
}

func (c *CSVFileFeeder) MakeFeaturePipe() FeaturePipe {
	return make(FeaturePipe)
}

// FIXME: Remove
func (c *CSVFileFeeder) GetFeed() (chan []float64, error) {
	file, err := os.Open(c.Filename)
	if err != nil {
		return nil, fmt.Errorf("Could not open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return make(chan []float64), nil
}

func splitCSVToFloat64Slice(s string) ([]float64, error) {
	strSlice := strings.Split(s, ",")
	floatSlice := make([]float64, len(strSlice))
	for i, v := range strSlice {

		parsedFloat, err := strconv.ParseFloat(strings.TrimSpace(v), 64)
		if err != nil {
			return nil, err
		}
		floatSlice[i] = parsedFloat
	}
	return floatSlice, nil
}
func (c *CSVFileFeeder) StartFeeder(dp FeaturePipe) {
	file, err := os.Open(c.Filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		floatArray, err := splitCSVToFloat64Slice(scanner.Text())
		if err != nil {
			panic(err)
		}
		dp <- floatArray
	}
}

func (c *CSVFileFeeder) Start(sc FeaturePipe ) {
	file, _ := os.Open(c.Filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		floatArray, err := splitCSVToFloat64Slice(scanner.Text())
		if err != nil {
			panic(err)
		}
		sc <- floatArray
	}
	close(sc)
}
