package search

func linearSearch(values []int, target int) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}

	return false
}

func TestLinearSearch() string {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	test1 := linearSearch(values, 3)
	test2 := linearSearch(values, 12)

	if test1 && !test2 {
		return "Pass"
	}

	return "Fail"
}
