package engine

import (
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/proto/urls"
)

type InMemoryUrls struct {
	urls    []*urls.Url
	history map[int32][]*urls.Response
}

func NewInMemoryUrls() *InMemoryUrls {
	return &InMemoryUrls{
		urls:    make([]*urls.Url, 0),
		history: make(map[int32][]*urls.Response, 0),
	}
}

func (i *InMemoryUrls) Create(url string, interval int32) int32 {
	id := int32(len(i.urls))
	i.urls = append(i.urls,
		&urls.Url{
			Id:       id,
			Url:      url,
			Interval: interval,
		})
	i.history[id] = make([]*urls.Response, 0)
	return id
}

func (i *InMemoryUrls) Delete(id int32) {
	// TODO: O(n) for delete :/ maybe do smth else here
	i.urls = append(i.urls[:id], i.urls[id+1:]...)
	delete(i.history, id)
}

func (i *InMemoryUrls) Get() []*urls.Url {
	return i.urls
}

func (i *InMemoryUrls) History(id int32) []*urls.Response {
	return i.history[id]
}
