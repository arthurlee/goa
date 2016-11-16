package route

import (
//"github.com/arthurlee/goa/server"
)

// middleware support

func Use() {

}

func Get(path string, handler GoaHandler) {
	getHandlerMap[path] = handler
}

func Post(path string, handler GoaHandler) {
	postHandlerMap[path] = handler
}
