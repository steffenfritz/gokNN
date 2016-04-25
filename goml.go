/*

a machine learning library

Copyright (C) 2016  Steffen Fritz

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package goml

import (
	"errors"
	"math"
)

// a basic type with a lable and a feature slice
type LabelWithFeatures struct {
	Label   string
	Feature []float64
}

// give the distance functions a common name
type Distance func([]float64, []float64) float64

// calculates Euclidian distance
func Euclidian(inX []float64, setX []float64) (float64, error) {
	if len(inX) != len(setX) {
		err := errors.New("Input slices are not of same length")
		return 0.0, err
	}
	var dist float64

	for i, valElement := range inX {
		temp := math.Pow(valElement-setX[i], 2)
		dist = dist + temp
	}

	return math.Sqrt(dist), nil
}

// calculates Manhattan distance
func Manhatten(inX []float64, dataSet []float64) (float64, error) {
	var dist float64

	return dist, nil
}

// calculates k-Nearest Neighbors, returns label
func KNN(k int, toClassify float64, tData []LabelWithFeatures, distf Distance) string {
	nLabel := "test"
	return nLabel
}
