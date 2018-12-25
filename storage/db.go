package storage

import (
	"database/sql"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mostlygeek/iot-recorder/storage/models"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
)

type DB struct {
	*sql.DB
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
	return d.DB.Close()
}

func (d *DB) open() error {
	var err error

	query := []string{
		"cache=shared",       // helps with concurrency
		"_busy_timeout=1500", // wait up to 1500ms when db is busy, set to ~max transaction time
		"_foreign_keys=true", // db has fk relationships, want these enabled
		"_journal_mode=wal",  // WAL + SetMaxOpenConns(1) is better than sync.Mutex everywhere in go code
	}
	d.DB, err = sql.Open("sqlite3", d.path+"?"+strings.Join(query, "&"))

	if err != nil {
		return errors.Wrap(err, "Unable to open database")
	}

	d.DB.SetMaxOpenConns(1)

	var userVersion int
	if err := d.DB.QueryRow("PRAGMA user_version;").Scan(&userVersion); err != nil {
		return errors.Wrap(err, "Unable to read PRAGMA user_version")
	}

	if userVersion == 0 {
		if _, err := d.DB.Exec(schema); err != nil {
			return errors.Wrap(err, "Failed applying schema")
		}
	}

	return nil
}

func (d *DB) RecordDemand(ts time.Time, demand int32) error {

	rec := models.Demand{
		Timestamp: ts.Unix(),
		Demand:    demand,
	}

	return rec.Insert(d.DB, boil.Infer())
}

// RecordSummation records the current Wh (Watt hour) recieved and delivered as measured by the meter
func (d *DB) RecordSummation(ts time.Time, delivered, received float32) error {
	sum := models.Summation{
		Timestamp: ts.Unix(),
		Delivered: delivered,
		Received:  received,
	}
	return sum.Insert(d.DB, boil.Infer())
}
