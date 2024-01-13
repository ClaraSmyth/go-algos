package main

import (
	"fmt"

	"github.com/ClaraSmyth/go-algos/search"
)

func main() {
	// Search Tests
	fmt.Printf("LinearSearch: %v \n", search.TestLinearSearch())
	fmt.Printf("BinarySearch: %v \n", search.TestBinarySearch())
	fmt.Printf("TwoCrystalBalls: %v \n", search.TestTwoCrystalBalls())
}
