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
	"fmt"
	"testing"
)

func TestCSVFileFeeder_GetFeed(t *testing.T) {
	filePath := "test.csv"
	feeder := CSVFileFeeder{
		Filename:    filePath,
		FeatureSize: 0,
	}
	feed, err := feeder.GetFeed()
	if err != nil {
		t.Errorf("Failed to execute test: %v", err)
	}
	_ = feed
}

func SendValues(s chan string) {
	for i := 0; i < 20; i++ {
		s <- "hello"
	}
}

func TestSendValues(t *testing.T) {
	c := make(chan string)
	go SendValues(c)

	y := <-c
	fmt.Println(y)
}
// FIXME: This needs to be converted to test
func TestCSVFileFeeder_Start(t *testing.T) {

	filePath := "test.csv"
	feeder := CSVFileFeeder{
		Filename:    filePath,
		FeatureSize: 0,
	}
	sc := make(FeaturePipe)

	go feeder.Start(sc)

	for i := range sc{
		fmt.Println(i)
	}
}
