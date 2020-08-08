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
	data         []NeuronDouble
	NRows, NCols int
	NWeights int
}

type NeuronDouble []float64

func (p *Mapx) GetNode(row int, col int) NeuronDouble {
	return p.data[p.NCols*row+col]
}

func New(nrows, ncols , nweights int) *Mapx {
	nDouble := NeuronDouble{1,2,3,4}
	mpx:=  &Mapx{
		data:  make([]NeuronDouble, nrows*ncols),
		NRows: nrows,
		NCols: ncols,
		NWeights: nweights,
	}

	mpx.data[0] = nDouble
	return mpx
}
