package handler

import (
	orderbusiness "github.com/syariatifaris/go-gof/sampleapp/module/order/business"
)

func NewHandlerFactory() IHandlerFactory {
	return &handlerFactory{}
}

type handlerFactory struct {
}

func (h *handlerFactory) NewOrderHandler(businessFactory orderbusiness.IOrderBusinessFactory) IHandler {
	return &orderHandler{
		businessFactory: businessFactory,
	}
}
