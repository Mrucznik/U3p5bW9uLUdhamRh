package service

import (
	"context"
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/proto/urls"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s Server) CreateUrl(ctx context.Context, request *urls.CreateUrlRequest) (*urls.CreateUrlResponse, error) {
	panic("implement me")
}

func (s Server) DeleteUrl(ctx context.Context, request *urls.DeleteUrlRequest) (*urls.DeleteUrlResponse, error) {
	panic("implement me")
}

func (s Server) GetUrls(ctx context.Context, request *urls.GetUrlsRequest) (*urls.GetUrlsResponse, error) {
	panic("implement me")
}

func (s Server) GetUrlHistory(ctx context.Context, request *urls.GetUrlHistoryRequest) (*urls.GetUrlHistoryResponse, error) {
	panic("implement me")
}
