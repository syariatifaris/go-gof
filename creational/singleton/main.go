package main

import (
	"sync"
	"fmt"
)

type developer struct {
	state string
}

var singleton *developer
var once sync.Once

func GetDeveloper() *developer {
	once.Do(func() {
		singleton = &developer{state: "off"}
	})
	return singleton
}

func (sm *developer) GetState() string {
	return sm.state
}

func (sm *developer) SetState(s string) {
	sm.state = s
}

func main(){
	dev := GetDeveloper()
	dev.SetState("resigned")
	fmt.Println(dev.GetState())

	dev2 := GetDeveloper()
	fmt.Println(dev2.GetState())
}
