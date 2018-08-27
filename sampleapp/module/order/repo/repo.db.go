package repo

import (
	"github.com/syariatifaris/go-gof/sampleapp/module/order/entity"
)

func NewOrderDBRepo() OrderDBRepo {
	return &orderDBRepo{}
}

type OrderDBRepo interface {
	InsertOrder(*entity.Order) (int64, error)
}

type orderDBRepo struct{}

func (o *orderDBRepo) InsertOrder(order *entity.Order) (int64, error) {
	return order.ID + 1, nil
}
