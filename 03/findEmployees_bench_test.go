package main

import (
	"testing"
)

func BenchmarkNameSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
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
		newEmployees.nameSort()
	}
}

func BenchmarkSalaryDescendingSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
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
		newEmployees.salaryDescendingSort()
	}
}

func BenchmarkGet2ndMaxSalary(b *testing.B) {
	for i := 0; i < b.N; i++ {
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
		newEmployees.get2ndMaxSalary()
	}
}
