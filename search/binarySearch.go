package search

import "math"

func binarySearch(values []int, target int) bool {
	low := float64(0)
	high := float64(len(values))

	for low < high {
		i := math.Floor(low + (high-low)/2)
		v := values[int(i)]

		if v == target {
			return true
		} else if v > target {
			high = i
		} else {
			low = i + 1
		}

	}

	return false
}

func TestBinarySearch() string {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	test1 := binarySearch(values, 6)
	test2 := binarySearch(values, 12)

	if test1 && !test2 {
		return "Pass"
	}

	return "Fail"
}
