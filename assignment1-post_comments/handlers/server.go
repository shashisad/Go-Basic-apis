package handlers

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Server struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func NewServer(address string, readTimeout, writeTimeout time.Duration) *Server {
	return &Server{Addr: address, ReadTimeout: readTimeout, WriteTimeout: writeTimeout}
}

// server launches a http server

func (s *Server) Start(ctx context.Context) error {

	srv := http.Server{
		Addr:         s.Addr,
		ReadTimeout:  s.ReadTimeout,
		WriteTimeout: s.WriteTimeout,
		Handler:      s.initHandlers(),
	}

	go func() {
		<-ctx.Done()
		log.Println("attempting graceful shutdown of server")
		srv.SetKeepAlivesEnabled(false)
		closeCtx, closeFn := context.WithTimeout(context.Background(), 30*time.Second)
		defer closeFn()
		_ = srv.Shutdown(closeCtx)
	}()

	return srv.ListenAndServe()
}

func (s *Server) initHandlers() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/getPosts", s.FetchCommentsForAllPosts)
	r.HandleFunc("/getComments/{id}", s.FetchCommentsForPostById)
	return r
}
