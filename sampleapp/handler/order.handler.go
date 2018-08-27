package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	orderbusiness "github.com/syariatifaris/go-gof/sampleapp/module/order/business"
	"net/http"
)

type orderHandler struct {
	businessFactory orderbusiness.IOrderBusinessFactory
}

func (o *orderHandler) GetName() string {
	return "OrderHandler"
}

func (o *orderHandler) RegisterHandler(router *mux.Router) {
	orderRouter := router.PathPrefix("/order").Subrouter()
	orderRouter.HandleFunc("/create", o.create).Methods(http.MethodPost)
}

func (o *orderHandler) create(w http.ResponseWriter, r *http.Request) {
	bm := o.businessFactory.NewCreateOrderBusiness(r)
	model, err := bm.GetModelCtx(r.Context())
	if err != nil {
		w.Write([]byte(fmt.Sprint("err:", err.Error())))
	}
	w.Write([]byte(fmt.Sprint(model)))
}
