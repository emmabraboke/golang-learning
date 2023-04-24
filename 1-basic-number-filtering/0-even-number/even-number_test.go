package evenNumber

import (
	"basic-number-filtering/helper"
	"log"
	"testing"
)

func TestEvenNumbers(t *testing.T) {
	sample := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := EvenNumbers(sample)

	expect := []int{2, 4, 6, 8, 10}

	if !helper.CompareSlice(result, expect) {
		t.Errorf("EvenNumbers(%d) returned %d, expected %d", sample, result, expect)
	}

	log.Println(result)

}
