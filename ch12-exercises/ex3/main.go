package main

import (
	"fmt"
	"math"
	"sync"
)

/*
Write a function that builds a map[int]float64 where the keys are the numbers from
0 (inclusive) to 100,000 (exclusive) and the values are the square roots of those
numbers (use the math.Sqrt function to calculate square roots).
Use sync.OnceValue to generate a function that caches the map returned by this
function and use the cached value to look up square roots for every 1,000th number
from 0 to 100,000.
*/
func generateMap() map[int]float64{
	values := make(map[int]float64, 100_000)
	for n := range 100_000{
		values[n] = math.Sqrt(float64(n))
	}
	return values
}

var savedVals = sync.OnceValue(generateMap)

func main(){
	for i := 0; i < 100_000; i += 1_000{
		fmt.Println(i, savedVals()[i])
	}
}
