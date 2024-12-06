package main

import (
	"fmt"
	"math"
)

func groupTemperatures(temps []float64) map[int][]float64 {
	result := make(map[int][]float64)
	for _, temp := range temps {
		group := int(math.Floor(temp/10) * 10)
		result[group] = append(result[group], temp)
	}
	return result
}

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := groupTemperatures(temps)

	for k, v := range groups {
		fmt.Printf("%d: %v\n", k, v)
	}
}
