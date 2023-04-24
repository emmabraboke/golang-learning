package evenMultipleOfFive

import (
	"basic-number-filtering/helper"
	"log"
	"testing"
)

func TestEvenMultipleOfFive(t *testing.T) {
	sample := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	result := EvenMultipleOfFive(sample)

	expect := []int{10,20}

	if !helper.CompareSlice(result, expect) {
		t.Errorf("EvenMultipleOfFive(%d) returned %d, expected %d", sample, result, expect)
	}


	log.Println(result)
}
