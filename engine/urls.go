package engine

import "github.com/Mrucznik/U3p5bW9uLUdhamRh/grpc/proto/urls"

type IUrls interface {
	Create(url string, interval int32) int32
	Delete(id int32) error
	Get() []*urls.Url
	History(id int32) ([]*urls.Response, error)
}
