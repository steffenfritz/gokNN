/*
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

//a machine learning library
package goml

import (
	"errors"
	"math"
	"sort"
)

// a basic type with a label and a feature slice
type LabelWithFeatures struct {
	Label   string
	Feature []float64
}

// a slice of type LabelWithFeatures for the Sort methods
type SLabelWithFeatures []LabelWithFeatures

// give the distance functions a common name
type Distance func([]float64, []float64) (float64, error)

// implements the sort.Sort interface for type LabelWithFeatures
func (slice SLabelWithFeatures) Len() int {
	return len(slice)
}

func (slice SLabelWithFeatures) Less(i, j int) bool {
	return slice[i].Feature[0] < slice[j].Feature[0]
}

func (slice SLabelWithFeatures) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// calculates Euclidian distance
func Euclidian(inX []float64, setX []float64) (float64, error) {
	if len(inX) != len(setX) {
		err := errors.New("Input slices are not of same length.\n")
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
func Manhatten(inX []float64, setX []float64) (float64, error) {
	if len(inX) != len(setX) {
		err := errors.New("Input slices are not of same length.\n")
		return 0.0, err
	}
	var dist float64

	for i, valElement := range inX {
		temp := math.Abs(valElement - setX[i])
		dist = dist + temp
	}

	return dist, nil
}

// calculates k-Nearest Neighbors, returns k nearest LabelWithFeatures
func KNN(k int, toClassify []float64, tData []LabelWithFeatures, distf Distance) ([]LabelWithFeatures, error) {
	var tempE LabelWithFeatures
	var unsortedSlice SLabelWithFeatures
	var sortedSlice []LabelWithFeatures

	if k > len(tData) {
		err := errors.New("k larger than training data.\n")
		return sortedSlice, err
	}

	for _, tEntry := range tData {
		dist, err := distf(toClassify, tEntry.Feature)
		tempE.Feature = []float64{dist}
		tempE.Label = tEntry.Label
		unsortedSlice = append(unsortedSlice, tempE)

		if err != nil {
			return unsortedSlice, err
		}
		sort.Sort(unsortedSlice)
	}

	for i := 0; i < k; i++ {
		sortedSlice = append(sortedSlice, unsortedSlice[i])
	}

	return sortedSlice, nil
}
