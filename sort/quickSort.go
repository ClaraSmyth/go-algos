package sort

import (
	"reflect"
)

func partition(arr []int, low, high int) int {
	pivot := arr[high]

	index := low - 1

	for i := low; i < high; i++ {
		if arr[i] <= pivot {
			index++

			curr := arr[i]
			arr[i] = arr[index]
			arr[index] = curr
		}
	}

	index++
	arr[high] = arr[index]
	arr[index] = pivot

	return index
}

func quickSort(arr []int, low, high int) {
	if low >= high {
		return
	}

	pivot := partition(arr, low, high)

	quickSort(arr, low, pivot-1)
	quickSort(arr, pivot+1, high)
}

func TestQuickSort() string {
	arr := []int{8, 3, 9, 4, 6, 1, 7, 2, 5}

	quickSort(arr, 0, len(arr)-1)

	if reflect.DeepEqual(arr, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		return "Pass"
	}

	return "Fail"
}
