package GoServe

import (
	"context"
	"fmt"
	"net/http"
)

type Server struct {
	status bool
	Server *http.Server
	Router *Router
}

func CreateServer(Port int) *Server {
	s := new(Server)

	s.status = false
	s.Server = &http.Server{
		Addr:              fmt.Sprintf(":%v", Port),
		Handler:           nil,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	return s
}

func (S *Server) Stop() error {
	if err := S.Server.Shutdown(context.TODO()); err != nil {
		return err
	}
	S.status = false
	return nil
}

func (S *Server) Start() error {
	if err := S.Server.ListenAndServe(); err != nil {
		return err
	}

	S.status = true
	return nil
}
