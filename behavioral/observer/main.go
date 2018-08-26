package main

import (
	"fmt"
)

type News struct {
	Title   string
	Content string
}

type IObserver interface {
	Update(*News)
}

type NewsAgency struct {
	news      *News
	observers []IObserver
}

func (n *NewsAgency) AddObserver(observer IObserver) {
	n.observers = append(n.observers, observer)
}

func (n *NewsAgency) SetNews(news *News) {
	n.news = news
	for _, observer := range n.observers {
		observer.Update(n.news)
	}
}

type NewsChannel struct {
	news *News
}

func (n *NewsChannel) Update(news *News) {
	n.news = news
}

func (n *NewsChannel) GetNews() *News {
	return n.news
}

func main() {
	observeable := new(NewsAgency)
	observer := new(NewsChannel)

	observeable.AddObserver(observer)
	observeable.SetNews(&News{
		Title:   "First News",
		Content: "Hello World",
	})

	fmt.Println(observer.GetNews())

	observeable.SetNews(&News{
		Title:   "Second News",
		Content: "Hello World 2",
	})

	fmt.Println(observer.GetNews())
}
