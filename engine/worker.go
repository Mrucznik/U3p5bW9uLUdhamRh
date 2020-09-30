package engine

import (
	"context"
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/grpc/proto/urls"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Worker struct {
	url      string
	interval time.Duration
	results  []*urls.Response
	cancel   context.CancelFunc
}

func NewWorker(url string, interval int32) *Worker {
	return &Worker{
		url:      url,
		interval: time.Duration(interval),
		results:  make([]*urls.Response, 0),
	}
}

func (w *Worker) Start() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	w.cancel = cancelFunc
	go func() {
		for {
			select {
			case <-ctx.Done():
			default:
				w.fetch()
			}
			time.Sleep(w.interval * time.Millisecond)
		}
	}()
}

func (w *Worker) Stop() {
	w.cancel()
}

func (w *Worker) GetResults() []*urls.Response {
	return w.results
}

func (w *Worker) fetch() {
	startTime := time.Now()
	resp, err := http.Get(w.url)
	duration := time.Since(startTime).Seconds()

	if err == nil && resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		w.results = append(w.results, &urls.Response{
			Response:  bodyString,
			Duration:  duration,
			CreatedAt: time.Now().Unix(),
		})
	} else {
		w.results = append(w.results, &urls.Response{
			Response:  "null",
			Duration:  duration,
			CreatedAt: time.Now().Unix(),
		})
	}
}
