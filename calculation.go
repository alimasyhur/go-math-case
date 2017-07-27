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
