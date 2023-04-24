package oddNumber

import "basic-number-filtering/helper"

func OddNumbers(numbers []int) []int {
	var filter []int
	for _, num := range numbers {
		if isTrue := helper.IsOdd(num); isTrue {
			filter = append(filter, num)
		}
	}

	return filter
}
