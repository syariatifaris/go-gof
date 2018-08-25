package main

import (
	"fmt"
)

type Laptop struct {
	model         string
	processorCore int
}

func (l *Laptop) ShowInfo() {
	fmt.Println("-model:", l.model)
	fmt.Println("-processor core:", l.processorCore)
}

type ILaptopBuilder interface {
	Init()
	SetModel()
	SetProcessor()
	GetLaptop() *Laptop
}

type asusBuilder struct{
	asusLaptop *Laptop
}

func (a *asusBuilder) Init() {
	a.asusLaptop = new(Laptop)
}

func (a *asusBuilder) SetModel() {
	a.asusLaptop.model = "asus"
}

func (a *asusBuilder) SetProcessor() {
	a.asusLaptop.processorCore = 4
}

func (a *asusBuilder) GetLaptop() *Laptop {
	return a.asusLaptop
}

type lenovoBuilder struct{
	lenovoLaptop *Laptop
}

func (l *lenovoBuilder) Init() {
	l.lenovoLaptop = new(Laptop)
}

func (l *lenovoBuilder) SetModel() {
	l.lenovoLaptop.model = "lenovo"
}

func (l *lenovoBuilder) SetProcessor() {
	l.lenovoLaptop.processorCore = 8
}

func (l *lenovoBuilder) GetLaptop() *Laptop {
	return l.lenovoLaptop
}

func NewLaptopCreator(builder ILaptopBuilder) ILaptopCreator{
	return &laptopCreator{
		builder: builder,
	}
}

type ILaptopCreator interface{
	CreateLaptop()
	GetLaptop() *Laptop
}

type laptopCreator struct{
	builder ILaptopBuilder
}

func (l *laptopCreator) CreateLaptop(){
	l.builder.Init()
	l.builder.SetModel()
	l.builder.SetProcessor()
}

func (l *laptopCreator) GetLaptop() *Laptop{
	return l.builder.GetLaptop()
}

func main() {
	lenovoCreator := NewLaptopCreator(new(lenovoBuilder))
	lenovoCreator.CreateLaptop()
	lenovoLaptop := lenovoCreator.GetLaptop()
	lenovoLaptop.ShowInfo()

	asusCreator := NewLaptopCreator(new(asusBuilder))
	asusCreator.CreateLaptop()
	asusLaptop := asusCreator.GetLaptop()
	asusLaptop.ShowInfo()
}
