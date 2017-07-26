package calculation

func findOdd(numbers int) []int {
	var newArr []int
	for index := 1; index < numbers; index++ {
		if index%3 == 0 || index%5 == 0 {
			newArr = append(newArr, index)
		}
	}
	return newArr
}

func SumOddNumber(numbers int) int {
  array := findOdd(numbers)
	sum := 0
	for _, value := range array {
		sum += value
	}
	return sum
}
