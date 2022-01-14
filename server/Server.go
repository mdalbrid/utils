package server

import (
	"context"
	"errors"
	"github.com/mdalbrid/utils/contexts"
	"github.com/mdalbrid/utils/logger"
	"net/http"
	"strconv"
	"time"
)

type Config struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Server struct {
	ctx        contexts.InternalModuleMainContext
	httpServer http.Server
	logger     logger.Logger
}

func New(ctx contexts.InternalModuleMainContext, cfg Config, handler http.Handler) *Server {
	return &Server{
		ctx:    ctx,
		logger: logger.NewLogger("HttpServer"),
		httpServer: http.Server{
			Addr:         cfg.Addr,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
			Handler:      handler,
		},
	}
}

func (srv *Server) Listen() {
	go func() {
		srv.logger.Info("Start server on ", srv.httpServer.Addr)
		err := srv.httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			srv.logger.Fatal("[Listen] httpServer.ListenAndServe() error: ", err)
		}
	}()

	<-srv.ctx.Done() // wait main context interrupt

	shutdownTimeout := srv.httpServer.WriteTimeout + srv.httpServer.ReadTimeout
	srv.logger.Info("Shutdown server with timeout ", strconv.Itoa(int(shutdownTimeout/time.Second)), " seconds")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	err := srv.httpServer.Shutdown(shutdownCtx)
	if err != nil {
		srv.logger.Error("[Listen] httpServer.Shutdown() error: ", err)
	}
	cancel()
}
