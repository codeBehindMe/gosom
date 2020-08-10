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

package mapx

// Mapx will contain the actual SOM Map itself.
type Mapx struct {
	Data         []NeuronDouble
	NRows, NCols int
	NWeights     int
}

type NeuronDouble []float64

func (n *NeuronDouble) ToSliceOfFloat64() []float64 {
	sliceF64 := make([]float64, len(*n))
	for i, v := range *n {
		sliceF64[i] = v
	}
	return sliceF64
}

func (p *Mapx) GetNode(row int, col int) NeuronDouble {
	return p.Data[p.NCols*row+col]
}

// FIXME: Is this the best position for this?
func (p *Mapx) IndexToRowCol(index int) (row, col int) {
	col = index % p.NCols
	row = index / p.NCols
	return row, col
}

func New(nrows, ncols, nweights int) *Mapx {
	mpx := &Mapx{
		Data:     make([]NeuronDouble, nrows*ncols),
		NRows:    nrows,
		NCols:    ncols,
		NWeights: nweights,
	}

	return mpx
}
