package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNameSort(t *testing.T) {
	assert := assert.New(t)

	newEmployees := employees{
		{"Hoang", 5.6, 1000000},
		{"Le", 7.5, 2000000},
		{"Anh", 8.5, 2000000},
		{"Vu", 5.5, 3000000},
		{"Chi", 7.5, 2000000},
		{"Long", 5.5, 5000000},
		{"Thuc", 7.5, 2000000},
		{"Khoan", 8.5, 2000000},
		{"Anh", 5.6, 1000000},
	}
	nameSortedArray := []employee{
		{"Anh", 8.5, 2000000},
		{"Anh", 5.6, 1000000},
		{"Chi", 7.5, 2000000},
		{"Hoang", 5.6, 1000000},
		{"Khoan", 8.5, 2000000},
		{"Le", 7.5, 2000000},
		{"Long", 5.5, 5000000},
		{"Thuc", 7.5, 2000000},
		{"Vu", 5.5, 3000000},
	}
	assert.Equal(newEmployees.nameSort(), nameSortedArray)
}

func TestSalaryDescendingSort(t *testing.T) {
	assert := assert.New(t)

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
	salaryDescendingArray := []employee{
		{"Anh", 8.5, 2000000},
		{"Khoan", 8.5, 2000000},
		{"Le", 7.5, 2000000},
		{"Chi", 7.5, 2000000},
		{"Long", 5.5, 5000000},
		{"Thuc", 7.5, 2000000},
		{"Vu", 5.5, 3000000},
		{"Hoang", 5.6, 1000000},
		{"Thao", 5.6, 1000000},
	}
	assert.Equal(newEmployees.salaryDescendingSort(), salaryDescendingArray)
}

func TestGet2ndMaxSalary(t *testing.T) {
	assert := assert.New(t)

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
	max2ndSalaryArray := []employee{
		{"Le", 7.5, 2000000},
		{"Chi", 7.5, 2000000},
		{"Long", 5.5, 5000000},
		{"Thuc", 7.5, 2000000},
	}
	assert.Equal(newEmployees.get2ndMaxSalary(), max2ndSalaryArray)
}
