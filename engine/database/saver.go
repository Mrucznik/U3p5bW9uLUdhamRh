package database

import (
	"database/sql"
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/grpc/proto/urls"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Saver struct {
	db *sql.DB
}

func NewSaver(db *sql.DB) *Saver {
	return &Saver{db: db}
}

func (s Saver) Save(data *urls.Response) error {
	_, err := s.db.Exec("INSERT INTO url_history(response, duration, created_at) VALUE (?, ?, ?)",
		data.Response, data.Duration, data.CreatedAt)

	if err != nil {
		logrus.Error(err)
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (s Saver) GetResults() []*urls.Response {
	panic("implement me")
}
