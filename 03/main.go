package main

import (
	"fmt"
)

func main() {
	// numbers := []int{2, 1, 3, 4}
	// max2Numbers(numbers)
	// fmt.Println("Second largest element is: ", max2Numbers(numbers))

	// strings := []string{"aba", "aa", "ad", "c", "vcd"}
	// findMaxLengthElement(strings)

	// uniqueNumbers := []int{1, 2, 5, 2, 6, 2, 5}
	// fmt.Println(removeDuplicates(uniqueNumbers))

	newEmployees := employees{
		{"Hoang", 5.6, 1000000},
		{"Le", 7.5, 2000000},
		{"Anh", 8.5, 2000000},
		{"Vu", 5.5, 3000000},
		{"Chi", 7.5, 2000000},
		{"Long", 5.5, 5000000},
		{"Thuc", 7.5, 2000000},
		{"Khoan", 8.5, 2000000},
		{"Thao", 5.6, 1000000},
	}
	// sortedAlphabetEmployees := newEmployees.nameSort()
	sortedSalaryDescending := newEmployees.salaryDescendingSort()
	// sorted2ndMaxSalary := newEmployees.get2ndMaxSalary()
	// fmt.Println(sortedAlphabetEmployees)
	fmt.Println(sortedSalaryDescending)
	// fmt.Println(sorted2ndMaxSalary)
}
