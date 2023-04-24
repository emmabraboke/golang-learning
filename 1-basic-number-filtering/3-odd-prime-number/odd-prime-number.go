package oddPrimeNumber

import "basic-number-filtering/helper"

func oddPrimeNumbers(numbers []int) []int {
	var filter []int
	for _, num := range numbers {
		isOdd := helper.IsOdd(num)
		isPrime := helper.IsPrime(num)
		if isOdd && isPrime {
			filter = append(filter, num)
		}

	}

	return filter
}
