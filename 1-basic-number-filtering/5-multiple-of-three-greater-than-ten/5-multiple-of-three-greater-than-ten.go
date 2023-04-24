package oddMultipleOfThreeGreaterThanTen

import "basic-number-filtering/helper"

func OddMultipleOfThreeGreaterThanTen(numbers []int) []int {
	var filter []int
	val := 3

	for _, num := range numbers {
		isMultipleOfThree := helper.IsMultiple(num, val)
		isgreaterThanTen := num > 10
		isOdd := helper.IsOdd(num)

		if isOdd && isMultipleOfThree && isgreaterThanTen {
			filter = append(filter, num)
		}
	}

	return filter
}
