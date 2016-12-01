package route

import (
	"github.com/arthurlee/goa/server"
)

// resources, like user
type GoaResource interface {
	GetPath() string
}

// small resource unit, like user/create
type GoaResourceAction interface {
	GetResource() GoaResource
	GetPath() string
	GetHandler() server.HttpContext
}
