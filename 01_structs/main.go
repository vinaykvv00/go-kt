package main

import (
	"encoding/json"
	"fmt"
)

func (e *Employee) updateName(newName string) {
	e.Name = newName
}

func (e *Employee) printEmployee() {
	fmt.Println(e.Name, e.Age, e.Isremote, e.Address.City, e.Address.Street)
}

type Employee struct {
	Name     string `json:"employeeName"`
	Age      int    `json:"employeeAge"`
	Isremote bool   `json:"isRemote"`
	Address  `json:"address"`
}

type Address struct {
	City   string `json:"MyCity"`
	Street string `json:"Mystreet"`
}

func main() {

	a := Address{
		City:   "Bengaluru",
		Street: "5th cross",
	}
	e := Employee{
		Name:     "artur",
		Age:      30,
		Isremote: true,
		Address:  a,
	}
	e.updateName("Vinay")
	e.printEmployee()

	job := struct {
		Title  string
		Salary int
	}{
		Title:  "developer",
		Salary: 50000,
	}
	fmt.Println(job.Title, job.Salary)

	jsonData, _ := json.Marshal(e)
	fmt.Println(string(jsonData))
}
