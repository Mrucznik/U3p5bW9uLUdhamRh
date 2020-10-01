package engine

import (
	"context"
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/grpc/proto/urls"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

// Worker provides a mechanism for retrieving and saving data from a URL at specified intervals
type Worker struct {
	Url      string
	Interval time.Duration
	saver    ISaver
	cancel   context.CancelFunc
}

func NewWorker(url string, interval int32, saver ISaver) *Worker {
	return &Worker{
		Url:      url,
		Interval: time.Duration(interval) * time.Second,
		saver:    saver,
	}
}

// Start begins fetching data process. It run a new goroutine, which can be stopped by Stop method.
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
			time.Sleep(w.Interval)
		}
	}()
}

// Stop ends fetching data process.
func (w *Worker) Stop() {
	if w.cancel != nil {
		w.cancel()
	}
}

// GetResults returns fetching results.
func (w *Worker) GetResults() ([]*urls.Response, error) {
	return w.saver.GetResults()
}

// fetch method fetch data from URL and saves it.
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
			logrus.Errorln(err)
		}
		bodyString := string(bodyBytes)
		err = w.saver.Save(&urls.Response{
			Response:  bodyString,
			Duration:  duration,
			CreatedAt: time.Now().Unix(),
		})
		if err != nil {
			logrus.Errorln(err)
		}
	} else {
		err = w.saver.Save(&urls.Response{
			Response:  "null",
			Duration:  duration,
			CreatedAt: time.Now().Unix(),
		})
		if err != nil {
			logrus.Errorln(err)
		}
	}
}
