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

//SumOddNumber sum int: case One
func SumOddNumber(numbers int) int {
	array := findOdd(numbers)
	sum := 0
	for _, value := range array {
		sum += value
	}
	return sum
}

//Fibonacci n int
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

//CalculateOddNumberFibonacci n: Case Two
func CalculateOddNumberFibonacci(n int) int {
	sum := 0
	for index := 1; index <= n; index++ {
		fibonacci := Fibonacci(index)
		if fibonacci%2 != 0 {
			sum += fibonacci
		}
	}
	return sum
}

//PrimeFactor number
func PrimeFactor(number int) []int {
	var arrPrimeFactor []int
	for index := 2; index < number; index++ {
		if (index%2 != 0) || (index%3 != 0) {
			if number%index == 0 {
				sum := 1
				for _, value := range arrPrimeFactor {
					sum = sum * value
				}
				if sum == number {
					break
				}
				arrPrimeFactor = append(arrPrimeFactor, index)
			}
		}
	}

	return arrPrimeFactor
}

//HighestPrimeFctor number int: Case Three
func HighestPrimeFctor(number int) int {
	arrayPrimeFactor := PrimeFactor(number)
	highestValue := 0
	for _, value := range arrayPrimeFactor {
		if value > highestValue {
			highestValue = value
		}
	}

	return highestValue
}
