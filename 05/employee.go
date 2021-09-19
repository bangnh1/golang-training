package main

import (
	"fmt"
)

type Employee interface {
	Clone() Employee
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId    int
	basicpay int
}

func (c *Permanent) Clone() Employee {
	newId := c.empId + 1
	return &Permanent{
		empId:    newId,
		basicpay: c.basicpay,
		pf:       c.pf,
	}
}

func (c *Contract) Clone() Employee {
	newId := c.empId + 1
	return &Contract{
		empId:    newId,
		basicpay: c.basicpay,
	}
}

func (p *Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

func (c *Contract) CalculateSalary() int {
	return c.basicpay
}

func totalExpense(s []Employee) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d \n", expense)
}
