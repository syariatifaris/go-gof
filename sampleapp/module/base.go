package module

import (
	"context"
)

type IBusiness interface {
	GetModelCtx(context.Context) (interface{}, error)
	Validate() error
}
