package helper

func MapCondition() (map[string]func(number int) bool, map[string]func(number int, value int) bool){
	conditions1 := make(map[string]func(number int) bool)
	conditions2 := make(map[string]func(number, value int) bool)

	conditions1["even"] = func(number int) bool { return IsEven(number) }
	conditions1["odd"] = func(number int) bool { return IsOdd(number)}
	conditions1["prime"] = func(number int) bool {return IsPrime(number)}
	conditions2["multiple"] = func(number int, value int) bool {return IsMultiple(number,value)}
	conditions2["greater than"] = func(number, value int) bool {return number > value}
	conditions2["less than"] = func(number, value int) bool {return number < value}

	return conditions1, conditions2
}
