package engine

import (
	"errors"
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/proto/urls"
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

func (i *InMemoryUrls) Create(url string, interval int32) int32 {
	id := i.ids
	i.workers[id] = NewWorker(url, interval)
	i.workers[id].Start()
	i.ids++
	return id
}

func (i *InMemoryUrls) Delete(id int32) error {
	if worker, ok := i.workers[id]; ok {
		worker.Stop()
		delete(i.workers, id)
		return nil
	}
	return errors.New("not found")
}

func (i *InMemoryUrls) Get() []*urls.Url {
	result := make([]*urls.Url, 0, len(i.workers))
	for id, worker := range i.workers {
		result = append(result, &urls.Url{
			Id:       id,
			Url:      worker.url,
			Interval: int32(worker.interval.Milliseconds()),
		})
	}
	return result
}

func (i *InMemoryUrls) History(id int32) ([]*urls.Response, error) {
	if worker, ok := i.workers[id]; ok {
		return worker.GetResults(), nil
	}
	return nil, errors.New("not found")
}
