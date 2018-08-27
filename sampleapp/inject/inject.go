package inject

import (
	"github.com/gorilla/mux"
	"github.com/syariatifaris/go-gof/sampleapp/handler"
	orderbusiness "github.com/syariatifaris/go-gof/sampleapp/module/order/business"
	orderepo "github.com/syariatifaris/go-gof/sampleapp/module/order/repo"
)

type Injection struct {
	Router   *mux.Router
	Handlers []handler.IHandler
}

func NewDependencyInjection() *Injection {
	var (
		router   *mux.Router
		handlers []handler.IHandler
	)
	//region: router
	router = mux.NewRouter()

	//region repository
	orderDBRepo := orderepo.NewOrderDBRepo()

	//region business factory
	orderBusinessFactory := orderbusiness.NewOrderBusinessFactory(orderDBRepo)

	//region: handler
	handlerFactory := handler.NewHandlerFactory()
	orderHandler := handlerFactory.NewOrderHandler(orderBusinessFactory)
	handlers = append(handlers, orderHandler)

	return &Injection{
		Router:   router,
		Handlers: handlers,
	}
}
