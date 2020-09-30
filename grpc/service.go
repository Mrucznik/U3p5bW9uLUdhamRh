package grpc

import (
	"context"
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/engine"
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/grpc/proto/urls"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/url"
)

type Server struct {
	urls engine.IUrls
}

func NewServer() *Server {
	return &Server{urls: engine.NewInMemoryUrls()}
}

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func (s Server) CreateUrl(ctx context.Context, request *urls.CreateUrlRequest) (*urls.CreateUrlResponse, error) {
	if !IsUrl(request.Url) {
		return nil, status.Error(codes.InvalidArgument, "Invalid URL.")
	}

	if request.Interval < 0 {
		return nil, status.Error(codes.InvalidArgument, "Invalid interval - must be positive.")
	}

	return &urls.CreateUrlResponse{Id: s.urls.Create(request.Url, request.Interval)}, nil
}

func (s Server) DeleteUrl(ctx context.Context, request *urls.DeleteUrlRequest) (*urls.DeleteUrlResponse, error) {
	err := s.urls.Delete(request.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return &urls.DeleteUrlResponse{Id: request.Id}, nil
}

func (s Server) GetUrls(ctx context.Context, request *urls.GetUrlsRequest) (*urls.GetUrlsResponse, error) {
	return &urls.GetUrlsResponse{Urls: s.urls.Get()}, nil
}

func (s Server) GetUrlHistory(ctx context.Context, request *urls.GetUrlHistoryRequest) (*urls.GetUrlHistoryResponse, error) {
	history, err := s.urls.History(request.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return &urls.GetUrlHistoryResponse{History: history}, nil
}
