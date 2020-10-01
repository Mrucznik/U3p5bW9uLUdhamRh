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
	id int32
}

func NewSaver(db *sql.DB, id int32) *Saver {
	return &Saver{db: db, id: id}
}

func (s *Saver) Save(data *urls.Response) error {
	_, err := s.db.Exec("INSERT INTO url_history(url_id, response, duration, created_at) VALUE (?, ?, ?, ?)",
		s.id, data.Response, data.Duration, data.CreatedAt)

	if err != nil {
		logrus.Error(err)
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (s *Saver) GetResults() ([]*urls.Response, error) {
	rows, err := s.db.Query("SELECT response, duration, created_at FROM url_history WHERE url_id=?", s.id)
	if err != nil {
		logrus.Error(err)
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		} else {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	defer rows.Close()

	result := make([]*urls.Response, 0)

	for rows.Next() {
		row := urls.Response{}
		err = rows.Scan(&row.Response, &row.Duration, &row.CreatedAt)
		if err != nil {
			logrus.Error(err)
			return result, status.Error(codes.Internal, err.Error())
		}
		result = append(result, &row)
	}

	if err = rows.Err(); err != nil {
		logrus.Error(err)
		return result, status.Error(codes.Internal, err.Error())
	}

	return result, nil
}
