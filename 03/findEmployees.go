package main

import (
	"sort"
)

type employee struct {
	name        string
	salaryRatio float64
	bonus       int
}

type employees []employee

func (e employee) salary() int {
	return int(e.salaryRatio*1500000) + e.bonus
}

func (e employees) nameSort() []employee {

	sort.SliceStable(e, func(i, j int) bool {
		return e[i].name < e[j].name
	})

	return e
}

func (e employees) salaryDescendingSort() []employee {
	sort.SliceStable(e, func(i, j int) bool {
		return e[i].salary() > e[j].salary()
	})

	return e
}

func (e employees) get2ndMaxSalary() []employee {
	var sortedSalary []int
	for _, employee := range e {
		sortedSalary = append(sortedSalary, employee.salary())
	}
	uniqueSalaryList := removeDuplicates(sortedSalary)

	sort.Slice(uniqueSalaryList, func(i, j int) bool {
		return uniqueSalaryList[i] > uniqueSalaryList[j]
	})

	sorted2ndMaxSalary := employees{}
	for _, employee := range e {
		salary := int(employee.salaryRatio*1500000) + employee.bonus
		if salary == uniqueSalaryList[1] {
			sorted2ndMaxSalary = append(sorted2ndMaxSalary, employee)
		}
	}
	return sorted2ndMaxSalary
}
