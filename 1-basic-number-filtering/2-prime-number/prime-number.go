package primeNumber

import "basic-number-filtering/helper"

func PrimeNumbers(numbers []int) []int {
	var filter []int
	for _, num := range numbers {

		if isTrue := helper.IsPrime(num); isTrue {
			filter = append(filter, num)
		}
	}
	return filter
}
