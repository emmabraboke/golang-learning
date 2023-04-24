package oddNumber

import (
	"basic-number-filtering/helper"
	"log"
	"testing"
)

func TestOddNumbers(t *testing.T) {
	sample := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := OddNumbers(sample)

	
	expect := []int{1,3,5,7,9}

	if !helper.CompareSlice(result, expect) {
		t.Errorf("OddNumbers(%d) returned %d, expected %d", sample, result, expect)
	}


	log.Println(result)

}
