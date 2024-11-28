package arraysandslices

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}


func SumAll(numberToSum ...[]int) []int {
	lengthOfNumbers := len(numberToSum)
	sums := make([]int, lengthOfNumbers)

	for _, numbers := range numberToSum{
		sums = append(sums, Sum(numbers))
	}
	return sums
}