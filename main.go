package main

import (
	"fmt"

	"github.com/ClaraSmyth/go-algos/list"
	"github.com/ClaraSmyth/go-algos/search"
	"github.com/ClaraSmyth/go-algos/sort"
)

func main() {
	// Search Tests
	fmt.Printf("LinearSearch: %v \n", search.TestLinearSearch())
	fmt.Printf("BinarySearch: %v \n", search.TestBinarySearch())
	fmt.Printf("TwoCrystalBalls: %v \n", search.TestTwoCrystalBalls())

	// Sort Tests
	fmt.Printf("BubbleSort: %v \n", sort.TestBubbleSort())
	fmt.Printf("QuickSort %v \n", sort.TestQuickSort())

	// List Tests
	fmt.Printf("Queue: %v \n", list.TestQueue())
	fmt.Printf("Stack: %v \n", list.TestStack())
}
