package business

import (
	"github.com/syariatifaris/go-gof/sampleapp/module"
	"net/http"
)

type IOrderBusinessFactory interface {
	NewCreateOrderBusiness(*http.Request) module.IBusiness
}
