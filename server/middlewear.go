package server

import (
	"fmt"
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

// timer represents a one-time stoppable time tracker
// it's assumed startTimer() is called to initialize a timer
type timer struct {
	start, end time.Time
}

func startTimer() *timer {
	return &timer{
		start: time.Now(),
	}
}

// stop returns the number of milliseconds for the given timer
// and stops the timer.
func (t *timer) stop() int64 {
	when := time.Now()
	if t.end.IsZero() {
		t.end = when
	}
	return t.end.Sub(t.start).Nanoseconds() / 1e6
}
func (t *timer) String() string {
	return fmt.Sprintf("%dms", t.stop())
}

// Middleware endpoints

func (mw loggingMiddleware) CreateFile(f *ach.FileHeader) (id string, err error) {
	t := startTimer()
	defer func() {
		// f.ID can be empty here if the request does not specify it, do we care?
		// The ID is properly generated and stored, so there's no "issue", just flakey logs.
		// TODO(adam): figure out if we care to fix this
		mw.logger.Log("method", "CreateFile", "id", f.ID, "took", t, "err", err)
	}()
	return mw.next.CreateFile(f)
}

func (mw loggingMiddleware) GetFile(id string) (f *ach.File, err error) {
	t := startTimer()
	defer func() {
		mw.logger.Log("method", "GetFile", "id", id, "took", t, "err", err)
	}()
	return mw.next.GetFile(id)
}

func (mw loggingMiddleware) GetFiles() []*ach.File {
	t := startTimer()
	defer func() {
		mw.logger.Log("method", "GetFiles", "took", t)
	}()
	return mw.next.GetFiles()
}

func (mw loggingMiddleware) DeleteFile(id string) (err error) {
	t := startTimer()
	defer func() {
		mw.logger.Log("method", "DeleteFile", "id", id, "took", t)
	}()
	return mw.next.DeleteFile(id)
}

//** BATCHES ** //

func (mw loggingMiddleware) CreateBatch(fileID string, bh *ach.BatchHeader) (id string, err error) {
	t := startTimer()
	defer func() {
		mw.logger.Log("method", "CreateBatch", "FileID", fileID, "batchID", bh.ID, "took", t, "err", err)
	}()
	return mw.next.CreateBatch(fileID, bh)
}

func (mw loggingMiddleware) GetBatch(fileID string, batchID string) (b ach.Batcher, err error) {
	t := startTimer()
	defer func() {
		mw.logger.Log("method", "GetBatch", "fileID", fileID, "batchID", batchID, "took", t, "err", err)
	}()
	return mw.next.GetBatch(fileID, batchID)
}

func (mw loggingMiddleware) GetBatches(fileID string) []ach.Batcher {
	t := startTimer()
	defer func() {
		mw.logger.Log("method", "GetBatches", "fileID", fileID, "took", t)
	}()
	return mw.next.GetBatches(fileID)
}

func (mw loggingMiddleware) DeleteBatch(fileID string, batchID string) (err error) {
	t := startTimer()
	defer func() {
		mw.logger.Log("method", "DeleteBatch", "fileID", fileID, "batchID", batchID, "took", t, "err", err)
	}()
	return mw.next.DeleteBatch(fileID, batchID)
}
