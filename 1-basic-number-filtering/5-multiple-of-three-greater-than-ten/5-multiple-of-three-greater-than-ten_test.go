package oddMultipleOfThreeGreaterThanTen

import (
	"log"
	"testing"
	"basic-number-filtering/helper"
)

func TestOddMultipleOfThreeGreaterThanTen(t *testing.T) {
	sample := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}

	result := OddMultipleOfThreeGreaterThanTen(sample)
	expect := []int{15}


	if !helper.CompareSlice(result, expect){
		t.Errorf("OddMultipleOfThreeGreaterThanTen(%d) returned %d, expected %d",sample, result, expect)
	}

	log.Println(result)
}

