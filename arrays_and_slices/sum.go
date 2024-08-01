package main

func sum(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, sum(numbers))
	}
	return sums

}

func SumAllTails(numbersToSum ...[]int) []int {
	var tailsSum []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			tailsSum = append(tailsSum, 0)
			continue
		}
		tailsSum = append(tailsSum, sum(numbers[1:]))
	}

	return tailsSum
}
