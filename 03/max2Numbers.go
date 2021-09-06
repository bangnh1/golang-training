package main

func max2Numbers(numbers []int) int {

	var large1st int = 0
	var large2nd int = 0

	for i := 1; i < len(numbers); i++ {
		if large1st < numbers[i] {
			large2nd = large1st
			large1st = numbers[i]
		} else if large2nd < numbers[i] {
			large2nd = numbers[i]
		}
	}
	return large2nd
}
