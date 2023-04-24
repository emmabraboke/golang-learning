package primeNumber

import (
	"basic-number-filtering/helper"
	"log"
	"testing"
)

func TestPrimeNumber(t *testing.T){
	sample := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result := PrimeNumbers(sample)

	
	expect := []int{2, 3, 5, 7}

	if !helper.CompareSlice(result, expect) {
		t.Errorf("PrimeNumbers(%d) returned %d, expected %d", sample, result, expect)
	}

	log.Println(result)
}