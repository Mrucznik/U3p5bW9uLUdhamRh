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
	Url      string
	Interval time.Duration
	saver    ISaver
	cancel   context.CancelFunc
}

func NewWorker(url string, interval int32, dataContainer ISaver) *Worker {
	return &Worker{
		Url:      url,
		Interval: time.Duration(interval),
		saver:    dataContainer,
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
			time.Sleep(w.Interval * time.Millisecond)
		}
	}()
}

func (w *Worker) Stop() {
	w.cancel()
}

func (w *Worker) GetResults() ([]*urls.Response, error) {
	return w.saver.GetResults()
}

func (w *Worker) fetch() {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	startTime := time.Now()
	resp, err := client.Get(w.Url)
	duration := time.Since(startTime).Seconds()

	if err == nil && resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		err = w.saver.Save(&urls.Response{
			Response:  bodyString,
			Duration:  duration,
			CreatedAt: time.Now().Unix(),
		})
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = w.saver.Save(&urls.Response{
			Response:  "null",
			Duration:  duration,
			CreatedAt: time.Now().Unix(),
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}
