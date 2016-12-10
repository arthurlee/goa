package route

import (
	"github.com/arthurlee/goa/middleware"
	"github.com/arthurlee/goa/server"
	"net/http"
)

type tSrvHandler struct{}

var SrvHandler tSrvHandler

func (tSrvHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	context := server.CreateHttpContext(w, r)

	for e := MiddlewareList.Front(); e != nil; e = e.Next() {
		m := e.Value.(*middleware.Entry)

		context.Log.Debug("Do middleware [%s] ...", m.Name)
		ret, err := m.Handler(context)
		if ret == server.HR_ERROR {
			if err != nil {
				context.SendError("-1", err.Error())
			}
			break
		}
	}
}
