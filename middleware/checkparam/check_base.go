package checkparam

import (
	"github.com/arthurlee/goa/server"
)

type CheckHandler func(CheckBase, *server.HttpContext) (interface{}, error)

type CheckBase interface {
	GetName() string
	GetErrorCode() string
	IsRequired() bool
	GetHandler() CheckHandler
}
