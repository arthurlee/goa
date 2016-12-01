package server

import (
	"encoding/json"
	"github.com/arthurlee/goa/logger"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	//"strconv"
	"fmt"
	"time"
)

type HttpContext struct {
	W         *http.ResponseWriter
	R         *http.Request
	Form      url.Values
	SessionId string
	Log       *logger.Logger
	items     map[string]interface{}
}

// randStr(1) => 0 ~ 9
// randStr(2) => 0 ~ 99
func randStr(maxLen int) string {
	n := int64(math.Pow10(maxLen))
	rand.Seed(time.Now().UnixNano())
	return string(rand.Int63n(n))
}

func createSessionId() string {
	t := time.Now()
	d := t.Format("20060102150405000")
	return fmt.Sprintf("%s%s", d, randStr(7))
}

func CreateHttpContext(w *http.ResponseWriter, r *http.Request) *HttpContext {
	context := HttpContext{W: w, R: r, Form: r.Form}

	context.SessionId = createSessionId()
	context.Log = logger.GetLogger(context.SessionId)
	context.items = make(map[string]interface{}, 20)

	return &context
}

func (me *HttpContext) Set(key string, value interface{}) {
	me.items[key] = value
}

func (me *HttpContext) Get(key string) (interface{}, bool) {
	v, ok := me.items[key]
	return v, ok
}

func (me *HttpContext) SendJson(data interface{}) {
	err := json.NewEncoder(*me.W).Encode(data)
	if err != nil {
		me.Log.WarnError(err)
	}
}

type GoaBaseRes struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (me *HttpContext) SendError(code string, message string) {
	res := GoaBaseRes{code, message}
	me.SendJson(res)
}

func (me *HttpContext) SendOK() {
	res := GoaBaseRes{"0", "ok"}
	me.SendJson(res)
}

// ------------------------------------
// convenient functions

func GetSendOK() func(*HttpContext) (HResult, error) {
	return func(res *HttpContext) (HResult, error) {
		res.SendOK()
		return HR_OK, nil
	}
}

func GetSendError(code string, message string) func(*HttpContext) (HResult, error) {
	return func(res *HttpContext) (HResult, error) {
		res.SendError(code, message)
		return HR_OK, nil
	}
}
