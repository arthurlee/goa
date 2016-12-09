package route

import (
	"container/list"
	// "encoding/json"
	// "errors"
	"github.com/arthurlee/goa/logger"
	"github.com/arthurlee/goa/middleware"
	// "github.com/arthurlee/goa/server"
	// "net/http"
)

var MiddlewareList = list.New()

func Use(rm *middleware.Entry, after string) {
	var element *list.Element = nil
	if len(after) > 0 {
		for e := MiddlewareList.Front(); e != nil; e = e.Next() {
			if e.Value.(*middleware.Entry).Name == rm.Name {
				element = e
				break
			}
		}
	}

	if element != nil {
		MiddlewareList.InsertAfter(rm, element)
	} else {
		MiddlewareList.PushBack(rm)
	}
}

func DumpRoutes() {
	logger.Info("------------------- route list -------------------")
	for e := MiddlewareList.Front(); e != nil; e = e.Next() {
		m := e.Value.(*middleware.Entry)
		logger.Info("    %-20s    %s", m.Name, m.Version)
	}
	logger.Info("--------------------------------------------------")
}
