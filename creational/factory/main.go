package main

import (
	"fmt"
)

func NewVehicleFactory() IVehicleFactory {
	return &vehicleFactory{}
}

type IVehicleFactory interface {
	GetVehicle(string) IVehicle
}

type vehicleFactory struct{}

func (*vehicleFactory) GetVehicle(name string) IVehicle {
	switch name {
	case "scooter":
		return &scooter{name: name}
	case "bike":
		return &bike{name: name}
	}
	return nil
}

type IVehicle interface {
	Drive(miles int) error
}

type scooter struct {
	name string
}

func (s *scooter) Drive(miles int) error {
	fmt.Println("vehicle", s.name, "drives for", miles, "miles")
	return nil
}

type bike struct {
	name string
}

func (s *bike) Drive(miles int) error {
	fmt.Println("vehicle", s.name, "drives for", miles, "miles")
	return nil
}

func main() {
	vehicleFactory := NewVehicleFactory()
	scooter := vehicleFactory.GetVehicle("scooter")
	bike := vehicleFactory.GetVehicle("bike")
	scooter.Drive(10)
	bike.Drive(20)
}
