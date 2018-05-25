package server

import (
	"time"

	"github.com/go-kit/kit/log"
	"github.com/moov-io/ach"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(Service) Service

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   Service
	logger log.Logger
}

func (mw loggingMiddleware) CreateFile(f ach.File) (id string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "CreateFile", "id", f.ID, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.CreateFile(f)
}

func (mw loggingMiddleware) GetFile(id string) (f ach.File, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetFile", "id", id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.GetFile(id)
}

func (mw loggingMiddleware) GetFiles() []ach.File {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetFiles", "took", time.Since(begin))
	}(time.Now())
	return mw.next.GetFiles()
}

func (mw loggingMiddleware) DeleteFile(id string) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "DeleteFile", "id", id, "took", time.Since(begin))
	}(time.Now())
	return mw.next.DeleteFile(id)
}
