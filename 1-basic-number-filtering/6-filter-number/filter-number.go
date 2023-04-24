package filterNumber

import (
	"basic-number-filtering/helper"

	"strconv"
	"strings"
)

func FilterNumber(numbers []int, conditions ...string) []int {

	var filter []int
	isValid := true

	condition1, condition2 := helper.MapCondition()

	for _, num := range numbers {
		isValid = true
		for _, condition := range conditions {
			if strings.Contains(condition, "multiple of") {
				numString := strings.Split(condition, "multiple of")[1]

				value, err := strconv.Atoi(strings.TrimSpace(numString))

				if err != nil {
					isValid = false
					break
				}

				istrue := condition2["multiple"](num, value)

				if !istrue {
					isValid = false
					break
				}
			}

			if strings.Contains(condition, "greater than") {

				numString := strings.Split(condition, "greater than")[1]

				value, err := strconv.Atoi(strings.TrimSpace(numString))

				if err != nil {
					isValid = false
					break
				}

				istrue := condition2["greater than"](num, value)

				if !istrue {
					isValid = false
					break
				}
			}

			if strings.Contains(condition, "less than") {

				numString := strings.Split(condition, "less than")[1]

				value, err := strconv.Atoi(strings.TrimSpace(numString))

				if err != nil {
					isValid = false
					break
				}

				istrue := condition2["less than"](num, value)
				if !istrue {
					isValid = false
					break
				}

			}

			if condition == "prime" || condition == "even" || condition == "odd" {
				if !condition1[condition](num) {
					isValid = false
					break
				}

			}

		}

		if isValid {
			filter = append(filter, num)
		}

	}

	return filter
}
