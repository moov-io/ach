package test

import (
	"net/http"
	"time"

	"github.com/moov-io/ach/server"
	"github.com/moov-io/base/log"

	kitlog "github.com/go-kit/log"
)

type Server struct {
	Repository server.Repository
	Service    server.Service
	Handler    http.Handler
}

func NewServer() *Server {
	repo := server.NewRepositoryInMemory(0*time.Second, log.NewNopLogger())
	svc := server.NewService(repo)

	handler := server.MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

	return &Server{
		Repository: repo,
		Service:    svc,
		Handler:    handler,
	}
}
