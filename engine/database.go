package engine

import (
	"database/sql"
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/grpc/proto/urls"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DatabaseUrls struct {
	db *sql.DB
}

func NewDatabaseUrls(db *sql.DB) *DatabaseUrls {
	//TODO: get data
	return &DatabaseUrls{db: db}
}

func (d *DatabaseUrls) initWorkers() error {
	rows, err := d.db.Query("SELECT id, address, `interval` FROM urls")
	if err != nil {
		logrus.Error(err)
		if err != sql.ErrNoRows {
			return err
		}
	}
	defer rows.Close()

	for rows.Next() {
		row := urls.Url{}
		err = rows.Scan(&row.Id, &row.Url, &row.Interval)
		if err != nil {
			return err
		}

		//TODO: start worker
	}

	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}

// Create a new worker in database and start fetching data from url.
func (d *DatabaseUrls) Create(url string, interval int) (int, error) {
	result, err := d.db.Exec("INSERT INTO urls(address, interval) VALUE (?, ?)",
		url, interval)
	if err != nil {
		logrus.Error(err)
		return 0, status.Error(codes.Internal, err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		logrus.Error(err)
		return 0, status.Error(codes.Internal, err.Error())
	}
	//TODO: start worker

	return int(id), nil
}

// Delete an existing worker, it's history and stop fetching data from url.
func (d *DatabaseUrls) Delete(id int) error {
	result, err := d.db.Exec("DELETE FROM urls WHERE id = ?", id)
	if err != nil {
		logrus.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	deletedRows, err := result.RowsAffected()
	if err != nil {
		logrus.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	if deletedRows == 0 {
		return status.Errorf(codes.NotFound, "Record of id %v not found.", id)
	}
	//TODO: stop worker

	return nil
}

// Get all existing urls.
func (d *DatabaseUrls) Get() ([]*urls.Url, error) {
	rows, err := d.db.Query("SELECT id, address, `interval` FROM urls")
	if err != nil {
		logrus.Error(err)
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		} else {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	defer rows.Close()

	result := make([]*urls.Url, 0)

	for rows.Next() {
		row := urls.Url{}
		err = rows.Scan(&row.Id, &row.Url, &row.Interval)
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

// Get fetching data from URL history.
func (d *DatabaseUrls) History(id int) ([]*urls.Response, error) {
	rows, err := d.db.Query("SELECT response, duration, created_at FROM url_history")
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
