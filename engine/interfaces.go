package engine

import "github.com/Mrucznik/U3p5bW9uLUdhamRh/grpc/proto/urls"

type IURLsService interface {
	// Create a new worker and start fetching data from Url.
	Create(url string, interval int32) (int32, error)
	// Delete an existing worker, it's history and stop fetching data from Url.
	Delete(id int32) error
	// Get all existing urls.
	Get() ([]*urls.Url, error)
	// Get fetching data from URL history.
	History(id int32) ([]*urls.Response, error)
}

type ISaver interface {
	Save(data *urls.Response) error
	GetResults() []*urls.Response
}
