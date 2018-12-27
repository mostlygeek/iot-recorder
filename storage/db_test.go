package storage

import (
	"testing"
	"time"

	"github.com/mostlygeek/iot-recorder/storage/models"
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

	err := db.RecordDemand(timestamp, int32(1992))
	if !assert.Nil(err) {
		return
	}

	demands, err := models.Demands().All(db)
	if !assert.NoError(err) {
		return
	}

	if assert.Len(demands, 1) {
		assert.Equal(timestamp.Unix(), demands[0].Timestamp)
		assert.Equal(int32(1992), demands[0].Demand)
	}
}

func TestRecordSummation(t *testing.T) {
	assert := assert.New(t)
	db := tmpDB()
	defer db.Close()

	delivered := float32(5)
	received := float32(7)

	for i := 1; i <= 5; i++ {
		if !assert.NoError(db.RecordSummation(timestamp, delivered, received)) {
			return
		}
	}

	sums, err := models.Summations().All(db)
	if !assert.NoError(err) {
		return
	}

	assert.Len(sums, 5)
	for _, sum := range sums {
		assert.Equal(timestamp.Unix(), sum.Timestamp)
		assert.Equal(delivered, sum.Delivered)
		assert.Equal(received, sum.Received)
	}
}
