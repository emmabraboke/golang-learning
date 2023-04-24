package oddPrimeNumber

import (
	"basic-number-filtering/helper"
	"log"
	"testing"
)

func TestOddPrimeNumber(t *testing.T) {
	sample := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := oddPrimeNumbers(sample)

	expect := []int{3, 5, 7}

	if !helper.CompareSlice(result, expect) {
		t.Errorf("OddPrimeNumbers(%d) returned %d, expected %d", sample, result, expect)
	}

	
	log.Println(result)
}
