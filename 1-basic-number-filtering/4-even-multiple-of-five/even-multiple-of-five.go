package evenMultipleOfFive

import "basic-number-filtering/helper"


func EvenMultipleOfFive(numbers []int) []int {

	var filter []int
	val := 5

	for _, num := range numbers {
		isEven := helper.IsEven(num)
		istrue := helper.IsMultiple(num, val)

		if isEven && istrue {
			filter = append(filter, num)
		}
	}
	return filter
}
