package evenNumber

import "basic-number-filtering/helper"

func EvenNumbers(numbers []int) []int {
	var filter []int
	for _, num := range numbers {
		if isTrue := helper.IsEven(num); isTrue {
			filter = append(filter, num)
		}
	}

	return filter
}
