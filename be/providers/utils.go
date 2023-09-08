package providers

import "math"

func findLargest(numbers ...float64) float64 {
	largest := math.Inf(-1)
	for _, num := range numbers {
		if num > largest {
			largest = num
		}
	}
	return largest
}

func findSmallest(numbers ...float64) float64 {
	smallest := math.Inf(1)
	for _, num := range numbers {
		if num < smallest {
			smallest = num
		}
	}
	return smallest
}
