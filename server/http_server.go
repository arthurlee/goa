package server

import (
	"fmt"
	"github.com/arthurlee/goa/logger"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func HttpListenAndServe(addr string, handler http.Handler) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	// exit gracefully
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		s := <-c
		fmt.Println("Got signal", s)
		logger.Warn("Got signal %s", s)
		listener.Close()
	}()

	return http.Serve(listener, handler)
}
