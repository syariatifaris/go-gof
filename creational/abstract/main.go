package main

import (
	"fmt"
)

type IAnimalFactory interface {
	GetBird(string) IBird
	GetFish(string) IFish
}

type tameAnimalFactory struct{}

func (a *tameAnimalFactory) GetBird(continent string) IBird {
	switch continent {
	case "asia":
		return &peafowl{}
	case "america":
		return &toco{}
	}
	return nil
}

func (a *tameAnimalFactory) GetFish(name string) IFish {
	switch name {
	case "asia":
		return &burmundi{}
	case "america":
		return &pirarucu{}
	}
	return nil
}

type fierceAnimalFactory struct{}

func (a *fierceAnimalFactory) GetBird(continent string) IBird {
	switch continent {
	case "asia":
		return &falcon{}
	case "america":
		return &condor{}
	}
	return nil
}

func (a *fierceAnimalFactory) GetFish(continent string) IFish {
	switch continent {
	case "asia":
		return &tigerFish{}
	case "america":
		return &piranha{}
	}
	return nil
}

type IBird interface {
	Name() string
	Fly(to string) error
}

type peafowl struct{}

func (p *peafowl) Name() string {
	return "peafowl"
}

func (p *peafowl) Fly(to string) error {
	fmt.Println(p.Name(), "flies to", to)
	return nil
}

type toco struct{}

func (p *toco) Name() string {
	return "toco"
}

func (p *toco) Fly(to string) error {
	fmt.Println(p.Name(), "flies to", to)
	return nil
}

type condor struct{}

func (p *condor) Name() string {
	return "condor"
}

func (p *condor) Fly(to string) error {
	fmt.Println(p.Name(), "flies to", to)
	return nil
}

type falcon struct{}

func (p *falcon) Name() string {
	return "falcon"
}

func (p *falcon) Fly(to string) error {
	fmt.Println(p.Name(), "flies to", to)
	return nil
}

type IFish interface {
	Name() string
	Swim(to string) error
}

type tigerFish struct{}

func (f *tigerFish) Name() string {
	return "tiger fish"
}

func (f *tigerFish) Swim(to string) error {
	fmt.Println(f.Name(), "swim to", to)
	return nil
}

type piranha struct{}

func (f *piranha) Name() string {
	return "piranha"
}

func (f *piranha) Swim(to string) error {
	fmt.Println(f.Name(), "swim to", to)
	return nil
}

type burmundi struct{}

func (f *burmundi) Name() string {
	return "burmundi"
}

func (f *burmundi) Swim(to string) error {
	fmt.Println(f.Name(), "swim to", to)
	return nil
}

type pirarucu struct{}

func (f *pirarucu) Name() string {
	return "pirarucu"
}

func (f *pirarucu) Swim(to string) error {
	fmt.Println(f.Name(), "swim to", to)
	return nil
}

func NewContinentalAnimalClient(factory IAnimalFactory, continent string) IAnimalClient {
	return &animalClient{
		bird: factory.GetBird(continent),
		fish: factory.GetFish(continent),
	}
}

type IAnimalClient interface {
	GetFishName() string
	DoSwim(string) error
	GetBirdName() string
	DoFly(string) error
}

type animalClient struct {
	bird IBird
	fish IFish
}

func (v *animalClient) GetFishName() string {
	return v.fish.Name()
}

func (v *animalClient) DoSwim(to string) error {
	return v.fish.Swim(to)
}

func (v *animalClient) GetBirdName() string {
	return v.bird.Name()
}

func (v *animalClient) DoFly(to string) error {
	return v.bird.Fly(to)
}

func main() {
	asianTameAnimalClt := NewContinentalAnimalClient(new(tameAnimalFactory), "asia")
	fmt.Println(asianTameAnimalClt.GetBirdName())
	asianTameAnimalClt.DoFly("to indonesia")
	fmt.Println(asianTameAnimalClt.GetFishName())
	asianTameAnimalClt.DoSwim("to singapore")

	americanTameAnimalClt := NewContinentalAnimalClient(new(tameAnimalFactory), "america")
	fmt.Println(americanTameAnimalClt.GetBirdName())
	americanTameAnimalClt.DoFly("to brazilia")
	fmt.Println(americanTameAnimalClt.GetFishName())
	americanTameAnimalClt.DoSwim("to mexico")
}
