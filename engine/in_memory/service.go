package in_memory

import (
	"errors"
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/engine"
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/grpc/proto/urls"
)

type InMemoryURLsService struct {
	workers map[int32]*engine.Worker
	ids     int32
}

func NewInMemoryURLsService() *InMemoryURLsService {
	return &InMemoryURLsService{
		workers: make(map[int32]*engine.Worker, 0),
	}
}

// Create a new worker and start fetching data from url.
func (i *InMemoryURLsService) Create(url string, interval int32) (int32, error) {
	id := i.ids
	i.workers[id] = engine.NewWorker(url, interval, NewSaver())
	i.workers[id].Start()
	i.ids++
	return id, nil
}

// Delete an existing worker, it's history and stop fetching data from url.
func (i *InMemoryURLsService) Delete(id int32) error {
	if worker, ok := i.workers[id]; ok {
		worker.Stop()
		delete(i.workers, id)
		return nil
	}
	return errors.New("not found")
}

// Get all existing urls.
func (i *InMemoryURLsService) Get() ([]*urls.Url, error) {
	result := make([]*urls.Url, 0, len(i.workers))
	for id, worker := range i.workers {
		result = append(result, &urls.Url{
			Id:       id,
			Url:      worker.Url,
			Interval: int32(worker.Interval.Milliseconds()),
		})
	}
	return result, nil
}

// Get fetching data from URL history.
func (i *InMemoryURLsService) History(id int32) ([]*urls.Response, error) {
	if worker, ok := i.workers[id]; ok {
		return worker.GetResults(), nil
	}
	return nil, errors.New("not found")
}
