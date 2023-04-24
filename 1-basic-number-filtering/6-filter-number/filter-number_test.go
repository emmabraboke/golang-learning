package filterNumber

import (
	"basic-number-filtering/helper"
	"log"
	"testing"
)

func TestFilterNumber(t *testing.T) {

	sample := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	result := FilterNumber(sample, "even", "less than 15", "multiple of 3")
	expect := []int{6, 12}

	if !helper.CompareSlice(result, expect) {
		t.Errorf("FilterNumber(%d) returned %d, expected %d", sample, result, expect)
	}

	log.Println(result)

	result = FilterNumber(sample, "odd", "greater than 5", "multiple of 3")
	expect = []int{9, 15}

	if !helper.CompareSlice(result, expect) {
		t.Errorf("FilterNumber(%d) returned %d, expected %d", sample, result, expect)
	}

	log.Println(result)
}
