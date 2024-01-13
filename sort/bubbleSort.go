package sort

import "reflect"

func bubbleSort(values []int) []int {
	for i := range values {
		for j := 0; j < len(values)-1-i; j++ {
			if values[j] > values[j+1] {
				tmp := values[j]
				values[j] = values[j+1]
				values[j+1] = tmp
			}
		}
	}
	return values
}

func TestBubbleSort() string {
	values := []int{1, 4, 3, 6, 2, 5}

	test1 := bubbleSort(values)

	if reflect.DeepEqual(test1, []int{1, 2, 3, 4, 5, 6}) {
		return "Pass"
	}

	return "False"
}
