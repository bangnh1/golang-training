package main

func removeDuplicates(numbers []int) []int {

	uniqueArray := []int{}

	check := make(map[int]int)
	for _, num := range numbers {
		check[num] = 1
	}

	for i := range check {
		uniqueArray = append(uniqueArray, i)
	}

	return uniqueArray
}
