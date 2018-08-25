package main

import (
	"fmt"
)

type IEmployee interface {
	AddEmployee(employee IEmployee)
	GetSubordinateAt(i int) IEmployee
	GetName() string
	GetSalary() float32
	GetInfo() string
}

type manager struct{
	name string
	salary float32
	subordinates []IEmployee
}

func (m *manager) AddEmployee(employee IEmployee) {
	m.subordinates = append(m.subordinates, employee)
}

func (m *manager) GetSubordinateAt(i int) IEmployee {
	return m.subordinates[i]
}

func (m *manager) GetName() string {
	return m.name
}

func (m *manager) GetSalary() float32 {
	return m.salary
}

func (m *manager) GetInfo() string{
	msg := fmt.Sprint("name:", m.name, ". salary:", m.salary, "\n")
	for _, sub := range m.subordinates{
		msg = fmt.Sprint(msg, "---dev: name=", sub.GetName(), ", salary", sub.GetSalary(), "\n")
	}
	return msg
}

type developer struct{
	name string
	salary float32
	programmingLangs []string
}

func (*developer) AddEmployee(employee IEmployee) {
	//child do not need to implement this method
}

func (*developer) GetSubordinateAt(i int) IEmployee {
	//child do not need to implement this method
	return nil
}

func (d *developer) GetName() string {
	return d.name
}

func (d *developer) GetSalary() float32 {
	return d.salary
}

func (d *developer) GetInfo() string {
	return fmt.Sprint("name:", d.name, ". salary:", d.salary, "\n")
}

func main() {
	manager := &manager{
		name: "Panjul",
		salary: 3000,
		subordinates: []IEmployee{},
	}

	dev := &developer{
		name: "Faris",
		salary: 800,
		programmingLangs: []string{"go", "C#"},
	}

	dev2 := &developer{
		name: "Dean",
		salary: 600,
		programmingLangs: []string {"JAVA"},
	}

	manager.AddEmployee(dev)
	manager.AddEmployee(dev2)
	fmt.Println(manager.GetInfo())
}
