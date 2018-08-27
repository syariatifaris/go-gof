package business

import (
	"context"
	"errors"
	"fmt"
	"github.com/syariatifaris/go-gof/sampleapp/module/order/entity"
	"github.com/syariatifaris/go-gof/sampleapp/module/order/repo"
	"net/http"
)

type createBusiness struct {
	dbRepo  repo.OrderDBRepo
	request *http.Request
}

func (c *createBusiness) GetModelCtx(context.Context) (interface{}, error) {
	if err := c.Validate(); err != nil {
		return nil, err
	}
	id, err := c.dbRepo.InsertOrder(&entity.Order{})
	if err != nil {
		return nil, err
	}
	//logic business disini
	return fmt.Sprint("create order success, order_id:", id), nil
}

func (c *createBusiness) Validate() error {
	if c.request == nil {
		return errors.New("request cannot be nil")
	}
	return nil
}
