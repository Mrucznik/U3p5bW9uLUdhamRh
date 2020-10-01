package engine

import (
	"errors"
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/grpc/proto/urls"
)

type InMemoryUrls struct {
	workers map[int32]*Worker
	ids     int32
}

func NewInMemoryUrls() *InMemoryUrls {
	return &InMemoryUrls{
		workers: make(map[int32]*Worker, 0),
	}
}

// Create a new worker and start fetching data from url.
func (i *InMemoryUrls) Create(url string, interval int32) (int32, error) {
	id := i.ids
	i.workers[id] = NewWorker(url, interval)
	i.workers[id].Start()
	i.ids++
	return id, nil
}

// Delete an existing worker, it's history and stop fetching data from url.
func (i *InMemoryUrls) Delete(id int32) error {
	if worker, ok := i.workers[id]; ok {
		worker.Stop()
		delete(i.workers, id)
		return nil
	}
	return errors.New("not found")
}

// Get all existing urls.
func (i *InMemoryUrls) Get() ([]*urls.Url, error) {
	result := make([]*urls.Url, 0, len(i.workers))
	for id, worker := range i.workers {
		result = append(result, &urls.Url{
			Id:       id,
			Url:      worker.url,
			Interval: int32(worker.interval.Milliseconds()),
		})
	}
	return result, nil
}

// Get fetching data from URL history.
func (i *InMemoryUrls) History(id int32) ([]*urls.Response, error) {
	if worker, ok := i.workers[id]; ok {
		return worker.GetResults(), nil
	}
	return nil, errors.New("not found")
}
