package main

import (
	"fmt"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/syariatifaris/go-gof/sampleapp/inject"
	"net/http"
)

func main() {
	injection := inject.NewDependencyInjection()
	router := injection.Router
	handlers := injection.Handlers

	for _, handler := range handlers {
		fmt.Println("register handler: ", handler.GetName())
		handler.RegisterHandler(router)
	}

	fmt.Println("starting service")
	gracehttp.Serve(&http.Server{
		Addr:    "0.0.0.0:9098",
		Handler: router,
	})
}
