package main

import "fmt"

type IEvent interface {
	Execute() error
}

type EventStore struct {
	events []IEvent
}

func (e *EventStore) StoreAndExecute(event IEvent){
	e.events = append(e.events, event)
	event.Execute()
}

type ShippingEvent struct {}

func (*ShippingEvent) Execute() error {
	fmt.Println("shipping in process")
	return nil
}

type InvalidAWBEvent struct{}

func (*InvalidAWBEvent) Execute() error {
	fmt.Println("invalid AWB handled")
	return nil
}

func main(){
	eventStore := new(EventStore)
	eventStore.StoreAndExecute(new(ShippingEvent))
	eventStore.StoreAndExecute(new(InvalidAWBEvent))
}


