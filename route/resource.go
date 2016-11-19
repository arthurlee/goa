package route

import (
	"github.com/arthurlee/goa/server"
)

// resources, like user
type GoaResource interface {
	GetPath() string
}

type GoaHandler func(*server.GoaResponse)

// small resource unit, like user/create
type GoaResourceAction interface {
	GetResource() GoaResource
	GetPath() string
	GetHandler() GoaHandler
}
