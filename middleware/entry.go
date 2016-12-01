package middleware

import (
	"github.com/arthurlee/goa/server"
)

type Entry struct {
	Name    string
	Version string
	Handler server.HttpHandler
}
