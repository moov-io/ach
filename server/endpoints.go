package server

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/moov-io/ach"
)

type Endpoints struct {
	CreateFileEndpoint endpoint.Endpoint
	GetFileEndpoint    endpoint.Endpoint
	GetFilesEndpoint   endpoint.Endpoint
	DeleteFileEndpoint endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateFileEndpoint: MakeCreateFileEndpoint(s),
		GetFileEndpoint:    MakeGetFileEndpoint(s),
		GetFilesEndpoint:   MakeGetFilesEndpoint(s),
		DeleteFileEndpoint: MakeDeleteFileEndpoint(s),
	}
}

// MakeCreateFileEndpoint returns an endpoint via the passed service.
func MakeCreateFileEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(createFileRequest)
		id, e := s.CreateFile(req.File)
		return createFileResponse{ID: id, Err: e}, nil
	}
}

type createFileRequest struct {
	File ach.File
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
	Files []ach.File `json:"files,omitempty"`
	Err   error      `json:"error,omitempty"`
}

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
	File ach.File `json:"file,omitempty"`
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
