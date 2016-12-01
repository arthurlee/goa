package server

type HResult int

const (
	HR_OK   HResult = iota
	HR_WARN         // can continue
	HR_ERROR
)

type HttpHandler func(*HttpContext) (HResult, error)
