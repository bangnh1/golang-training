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

func (e employees) nameSort() []employee {
	var nameArray []string
	for _, employee := range e {
		nameArray = append(nameArray, employee.name)
	}
	sort.Strings(nameArray)
	sortedAlphabetEmployees := employees{}
	for _, name := range nameArray {
		for _, employee := range e {
			if name == employee.name {
				sortedAlphabetEmployees = append(sortedAlphabetEmployees, employee)
			}
		}
	}

	return sortedAlphabetEmployees
}

func (e employees) salaryDescending() []employee {
	var sortedSalary []int
	for _, employee := range e {
		salary := int(employee.salaryRatio*1500000) + employee.bonus
		sortedSalary = append(sortedSalary, salary)
	}

	sort.Slice(sortedSalary, func(i, j int) bool {
		return sortedSalary[i] > sortedSalary[j]
	})

	sortedSalaryDescending := employees{}
	for _, s := range sortedSalary {
		for _, employee := range e {
			salary := int(employee.salaryRatio*1500000) + employee.bonus
			if s == salary {
				sortedSalaryDescending = append(sortedSalaryDescending, employee)
			}
		}
	}
	return sortedSalaryDescending
}

func (e employees) get2ndMaxSalary() []employee {
	var sortedSalary []int
	for _, employee := range e {
		salary := int(employee.salaryRatio*1500000) + employee.bonus
		sortedSalary = append(sortedSalary, salary)
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
