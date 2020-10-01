package engine

import "github.com/Mrucznik/U3p5bW9uLUdhamRh/grpc/proto/urls"

type IUrls interface {
	// Create a new worker and start fetching data from url.
	Create(url string, interval int32) (int32, error)
	// Delete an existing worker, it's history and stop fetching data from url.
	Delete(id int32) error
	// Get all existing urls.
	Get() ([]*urls.Url, error)
	// Get fetching data from URL history.
	History(id int32) ([]*urls.Response, error)
}
