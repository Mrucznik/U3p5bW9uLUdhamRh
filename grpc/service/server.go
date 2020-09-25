package service

import (
	"context"
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/engine"
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/proto/urls"
)

type Server struct {
	urls engine.IUrls
}

func NewServer() *Server {
	return &Server{urls: engine.NewInMemoryUrls()}
}

func (s Server) CreateUrl(ctx context.Context, request *urls.CreateUrlRequest) (*urls.CreateUrlResponse, error) {
	return &urls.CreateUrlResponse{Id: s.urls.Create(request.Url, request.Interval)}, nil
}

func (s Server) DeleteUrl(ctx context.Context, request *urls.DeleteUrlRequest) (*urls.DeleteUrlResponse, error) {
	s.urls.Delete(request.Id)
	return &urls.DeleteUrlResponse{Id: request.Id}, nil
}

func (s Server) GetUrls(ctx context.Context, request *urls.GetUrlsRequest) (*urls.GetUrlsResponse, error) {
	return &urls.GetUrlsResponse{Urls: s.urls.Get()}, nil
}

func (s Server) GetUrlHistory(ctx context.Context, request *urls.GetUrlHistoryRequest) (*urls.GetUrlHistoryResponse, error) {
	return &urls.GetUrlHistoryResponse{History: s.urls.History(request.Id)}, nil
}
