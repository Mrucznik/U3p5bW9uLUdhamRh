package in_memory

import "github.com/Mrucznik/U3p5bW9uLUdhamRh/grpc/proto/urls"

// This type provides in-memory saving mechanism.
type Saver struct {
	results []*urls.Response
}

func NewSaver() *Saver {
	return &Saver{results: make([]*urls.Response, 0)}
}

func (s *Saver) Save(data *urls.Response) error {
	s.results = append(s.results, data)
	return nil
}

func (s *Saver) GetResults() ([]*urls.Response, error) {
	return s.results, nil
}
