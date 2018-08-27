package business

import (
	"github.com/syariatifaris/go-gof/sampleapp/module"
	"github.com/syariatifaris/go-gof/sampleapp/module/order/repo"
	"net/http"
)

func NewOrderBusinessFactory(dbRepo repo.OrderDBRepo) IOrderBusinessFactory {
	return &orderBusinessFactory{
		dbRepo: dbRepo,
	}
}

type orderBusinessFactory struct {
	dbRepo repo.OrderDBRepo
}

func (o *orderBusinessFactory) NewCreateOrderBusiness(request *http.Request) module.IBusiness {
	return &createBusiness{
		request: request,
		dbRepo:  o.dbRepo,
	}
}
