package storage

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	timestamp time.Time
)

func init() {
	timestamp, _ = time.Parse("Jan 2 2006 15:04:05", "Apr 3 2014 11:22:21")
}

// tmpDB creates a memory backed sqlite instance, great for testing
func tmpDB() *DB {
	db, err := New(":memory:")
	if err != nil {
		panic(err)
	}
	return db
}

func TestRecordDemand(t *testing.T) {
	assert := assert.New(t)
	db := tmpDB()
	defer db.Close()

	err := db.RecordDemand(timestamp, 1992)
	if !assert.Nil(err) {
		return
	}

	row := db.db.QueryRow(`SELECT Timestamp,Demand FROM InstantDemand`)

	// place holder for data
	var (
		ts     int64
		demand int
	)
	if err := row.Scan(&ts, &demand); assert.NoError(err) {
		assert.Equal(timestamp.Unix(), ts)
		assert.Equal(1992, demand)
	}
}

func TestRecordSummation(t *testing.T) {
	assert := assert.New(t)
	db := tmpDB()
	defer db.Close()

	if !assert.NoError(db.RecordSummation(timestamp, 392, 5)) {
		return
	}

	row := db.db.QueryRow(`SELECT Timestamp,Delivered,Received FROM Summations`)

	// place holder for data
	var (
		ts                  int64
		delivered, received int
	)
	if err := row.Scan(&ts, &delivered, &received); assert.NoError(err) {
		assert.Equal(timestamp.Unix(), ts)
		assert.Equal(392, delivered)
		assert.Equal(5, received)
	}

	// insert 2 more records to test shortcutting
	for i := 0; i < 2; i++ {
		// makes a row based on diff delivered
		// skips a row due to same values
		assert.NoError(db.RecordSummation(timestamp, 400, 5))
	}
	for i := 0; i < 2; i++ {
		// makes a row based on diff recieved
		// skips a row due to same values
		assert.NoError(db.RecordSummation(timestamp, 400, 6))
	}

	// there should be three new records in total
	var count int
	if !assert.NoError(db.db.QueryRow("SELECT COUNT(1) c FROM Summations").Scan(&count)) {
		assert.Equal(3, count)
	}
}
