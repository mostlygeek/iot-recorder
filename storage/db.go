package storage

import (
	"database/sql"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

type DB struct {
	sync.RWMutex
	db   *sql.DB
	path string
}

// New creates or opens a new SQLite Database
func New(path string) (*DB, error) {
	db := &DB{path: path}
	if err := db.open(); err != nil {
		return nil, errors.Wrap(err, "Could not initialize DB")
	}

	return db, nil
}

func (d *DB) Close() error {
	return d.db.Close()
}

func (d *DB) open() error {

	var err error
	d.db, err = sql.Open("sqlite3", d.path)
	if err != nil {
		return errors.Wrap(err, "Unable to open database")
	}

	var userVersion int
	if err := d.db.QueryRow("PRAGMA user_version;").Scan(&userVersion); err != nil {
		return errors.Wrap(err, "Unable to read PRAGMA user_version")
	}

	if userVersion == 0 {
		if _, err := d.db.Exec(SCHEMA_0); err != nil {
			return errors.Wrap(err, "Failed applying SCHEMA_0")
		}
	}

	return nil
}

func (d *DB) RecordDemand(ts time.Time, demand int) error {
	d.Lock()
	defer d.Unlock()

	dml := "INSERT INTO InstantDemand (Timestamp, Demand) VALUES (?,?)"

	_, err := d.db.Exec(dml, ts.Unix(), demand)
	if err != nil {
		return errors.Wrap(err, "Unable to INSERT Instantaneous Demand Record")
	}

	return nil
}

// RecordSummation records the current KWh recieved and delivered as measured by the meter
func (d *DB) RecordSummation(ts time.Time, delivered, received int) error {
	d.Lock()
	defer d.Unlock()

	dml := `INSERT INTO Summations (Timestamp, Delivered, Received)
			VALUES (?,?,?)`

	_, err := d.db.Exec(dml, ts.Unix(), delivered, received)

	if err != nil {
		return errors.Wrap(err, "Unable to INSERT Summation Record")
	}

	return nil

	return nil
}
