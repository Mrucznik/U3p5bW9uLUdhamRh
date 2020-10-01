package database

import (
	"database/sql"
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/engine"
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/grpc/proto/urls"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type URLsService struct {
	db      *sql.DB
	workers map[int32]*engine.Worker
}

func NewURLsService(db *sql.DB) *URLsService {
	ret := &URLsService{db: db}
	if err := ret.loadWorkers(); err != nil {
		logrus.Errorln("cannot load workers from database")
	}
	return ret
}

// Fetch all existing URL's from database and start worker for each one.
func (d *URLsService) loadWorkers() error {
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

		// Create & start worker
		d.workers[row.Id] = engine.NewWorker(row.Url, row.Interval, NewSaver(d.db, row.Id))
		d.workers[row.Id].Start()
	}

	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}

// Create a new worker in database and start fetching data from Url.
func (d *URLsService) Create(url string, interval int) (int, error) {
	result, err := d.db.Exec("INSERT INTO urls(address, `interval`) VALUE (?, ?)",
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

	// Create & start worker
	id32 := int32(id)
	d.workers[id32] = engine.NewWorker(url, int32(interval), NewSaver(d.db, id32))
	d.workers[id32].Start()

	return int(id), nil
}

// Delete an existing worker, it's history and stop fetching data from Url.
func (d *URLsService) Delete(id int) error {
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
	d.workers[int32(id)].Stop()

	return nil
}

// Get all existing urls.
func (d *URLsService) Get() ([]*urls.Url, error) {
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
func (d *URLsService) History(id int) ([]*urls.Response, error) {
	return d.workers[int32(id)].GetResults()
}
