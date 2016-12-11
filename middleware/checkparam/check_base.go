package checkparam

import (
	"github.com/arthurlee/goa/server"
)

type CheckHandler func(CheckBase, *server.HttpContext) error

type CheckBase interface {
	GetName() string
	GetErrorCode() string
	GetHandler() CheckHandler
}
