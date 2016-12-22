package checkparam

import (
	"github.com/arthurlee/goa/middleware"
	"github.com/arthurlee/goa/server"
	"strconv"
)

// middleware entry

var RM_CheckParameter = middleware.Entry{"CheckParameter", "0.0.1", checkParameter}

func checkParameter(ctx *server.HttpContext) (server.HResult, error) {
	ctx.ParseParam()

	checkItems, ok := checkItemsMap[ctx.R.URL.Path]
	ctx.Log.Debug("checkParameter %s, %s", ctx.R.URL.Path, strconv.FormatBool(ok))
	if ok {
		for e := checkItems.Front(); e != nil; e = e.Next() {
			item := e.Value.(CheckBase)
			val, err := item.GetHandler()(item, ctx)
			if err != nil {
				return server.HR_ERROR, err
			}

			// store value to parameters with specific type
			ctx.Params.Set(item.GetName(), val)
		}
	}

	return server.HR_OK, nil
}
