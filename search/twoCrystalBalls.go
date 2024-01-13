package search

import (
	"math"
)

func twoCrystalBalls(values []bool) int {
	vlen := float64(len(values))
	jumpAmt := math.Floor(math.Sqrt(vlen))
	minBreak := -1

	i := jumpAmt
	for ; i < vlen; i += jumpAmt {
		if values[int(i)] {
			minBreak = int(i)
			break
		}
	}

	i -= jumpAmt

	for j := 0; j < int(jumpAmt) && i < vlen; j, i = j+1, i+1 {
		if values[int(i)] {
			minBreak = int(i)
			break
		}
	}

	return minBreak
}

func TestTwoCrystalBalls() string {
	values := []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true}

	test1 := twoCrystalBalls(values)
	test2 := twoCrystalBalls([]bool{false, false, false, false, false})

	if test1 == 13 && test2 == -1 {
		return "Pass"
	}

	return "Fail"
}
