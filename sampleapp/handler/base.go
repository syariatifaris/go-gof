package handler

import (
	orderbusiness "github.com/syariatifaris/go-gof/sampleapp/module/order/business"
	"github.com/gorilla/mux"
)

type IHandler interface {
	GetName() string
	RegisterHandler(router *mux.Router)
}

type IHandlerFactory interface {
	NewOrderHandler(orderbusiness.IOrderBusinessFactory) IHandler
}
