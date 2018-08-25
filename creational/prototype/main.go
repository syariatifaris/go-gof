package main

import (
	"encoding/json"
	"fmt"
)

type IEmployee interface{
	Clone() IEmployee
	GetDetails() string
}

type Developer struct{
	Name             string
	Age              int
	ProgrammingLangs []string
}

//Clone performs shallow copy
func (d *Developer) Clone() IEmployee {
	return &Developer{
		Name:             d.Name,
		Age:              d.Age,
		ProgrammingLangs: d.ProgrammingLangs,
	}
}

func (d *Developer) GetDetails() string {
	byteData, _ := json.Marshal(d)
	return fmt.Sprintf("details of developer: %s", string(byteData))
}

func main(){
	dev := new(Developer)
	dev.Name = "Faris"
	dev.Age = 27
	dev.ProgrammingLangs = []string{"go", "C#", "JAVA"}

	devCpy := dev.Clone().(*Developer)
	devCpy.Name = "Dean"
	devCpy.Age = 1
	devCpy.ProgrammingLangs = []string{"baba", "papa"}

	fmt.Println(dev.GetDetails())
	fmt.Println(devCpy.GetDetails())
}
