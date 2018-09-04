package server

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/moov-io/ach"
)

type Endpoints struct {
	CreateFileEndpoint  endpoint.Endpoint
	GetFileEndpoint     endpoint.Endpoint
	GetFilesEndpoint    endpoint.Endpoint
	DeleteFileEndpoint  endpoint.Endpoint
	CreateBatchEndpoint endpoint.Endpoint
	GetBatchesEndpoint  endpoint.Endpoint
	GetBatchEndpoint    endpoint.Endpoint
	DeleteBatchEndpoint endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateFileEndpoint:  MakeCreateFileEndpoint(s),
		GetFileEndpoint:     MakeGetFileEndpoint(s),
		GetFilesEndpoint:    MakeGetFilesEndpoint(s),
		DeleteFileEndpoint:  MakeDeleteFileEndpoint(s),
		CreateBatchEndpoint: MakeCreateBatchEndpoint(s),
		GetBatchesEndpoint:  MakeGetBatchesEndpoint(s),
		GetBatchEndpoint:    MakeGetBatchEndpoint(s),
		DeleteBatchEndpoint: MakeDeleteBatchEndpoint(s),
	}
}

// MakeCreateFileEndpoint returns an endpoint via the passed service.
func MakeCreateFileEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(createFileRequest)
		id, e := s.CreateFile(&req.FileHeader)
		return createFileResponse{ID: id, Err: e}, nil
	}
}

type createFileRequest struct {
	FileHeader ach.FileHeader
}

type createFileResponse struct {
	ID  string `json:"id,omitempty"`
	Err error  `json:"err,omitempty"`
}

func (r createFileResponse) error() error { return r.Err }

func MakeGetFilesEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(getFilesRequest)
		return getFilesResponse{Files: s.GetFiles(), Err: nil}, nil
	}
}

type getFilesRequest struct{}

type getFilesResponse struct {
	Files []*ach.File `json:"files,omitempty"`
	Err   error      `json:"error,omitempty"`
}

func (r getFilesResponse) error() error { return r.Err }

// MakeGetFileEndpoint returns an endpoint via the passed service.
// Primarily useful in a server.
func MakeGetFileEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getFileRequest)
		f, e := s.GetFile(req.ID)
		return getFileResponse{File: f, Err: e}, nil
	}
}

type getFileRequest struct {
	ID string `json:"id,omitempty"`
}

type getFileResponse struct {
	File *ach.File `json:"file,omitempty"`
	Err  error    `json:"err,omitempty"`
}

func (r getFileResponse) error() error { return r.Err }

func MakeDeleteFileEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(deleteFileRequest)
		e := s.DeleteFile(req.ID)
		return deleteFileResponse{Err: e}, nil
	}
}

type deleteFileRequest struct {
	ID string `json:"id,omitempty"`
}

type deleteFileResponse struct {
	Err error `json:"err,omitempty"`
}

func (r deleteFileResponse) error() error { return r.Err }

//** Batches ** //

// MakeCreateFileEndpoint returns an endpoint via the passed service.
func MakeCreateBatchEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(createBatchRequest)
		id, e := s.CreateBatch(req.FileID, &req.BatchHeader)
		return createBatchResponse{ID: id, Err: e}, nil
	}
}

type createBatchRequest struct {
	FileID      string
	BatchHeader ach.BatchHeader
}

type createBatchResponse struct {
	ID  string `json:"id,omitempty"`
	Err error  `json:"err,omitempty"`
}

func (r createBatchResponse) error() error { return r.Err }

func MakeGetBatchesEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getBatchesRequest)
		return getBatchesResponse{Batches: s.GetBatches(req.fileID), Err: nil}, nil
	}
}

type getBatchesRequest struct {
	fileID string
}

type getBatchesResponse struct {
	Batches []ach.Batcher `json:"batches,omitempty"`
	Err     error         `json:"error,omitempty"`
}

func (r getBatchesResponse) error() error { return r.Err }

func MakeGetBatchEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getBatchRequest)
		batch, e := s.GetBatch(req.fileID, req.batchID)
		return getBatchResponse{Batch: batch, Err: e}, nil
	}
}

type getBatchRequest struct {
	fileID  string
	batchID string
}

type getBatchResponse struct {
	Batch ach.Batcher `json:"batch,omitempty"`
	Err   error       `json:"err,omitempty"`
}

func (r getBatchResponse) error() error { return r.Err }

func MakeDeleteBatchEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(deleteBatchRequest)
		e := s.DeleteBatch(req.fileID, req.batchID)
		return deleteBatchResponse{Err: e}, nil
	}
}

type deleteBatchRequest struct {
	fileID  string
	batchID string
}

type deleteBatchResponse struct {
	Err error `json:"err,omitempty"`
}

func (r deleteBatchResponse) error() error { return r.Err }
