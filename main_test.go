package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ni/checkin"
	"github.com/ni/storage/cassandra"
	"github.com/stretchr/testify/assert"
)

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("An error occurred. %v", err)
	}
}

func TestCheckInHandler(t *testing.T) {

	jsonStr := []byte(`{
		"userId": "2342342424",
		"place": {
		"placeId":"37298472",
		"name": "Bistro 65",
		"lng": 31.112232,
		"lat": 20.123221,
		"category": "restaurant"
		},
		"checkinTimestamp": "2018-09-22T12:42:31Z"
		}`)
	req, err := http.NewRequest("POST", "/addCheckIn", bytes.NewBuffer(jsonStr))

	checkError(err, t)

	rr := httptest.NewRecorder()

	var checkinRepo checkin.Repository

	checkinRepo = cassandra.NewCassandraCheckinRepository(cassandraConnection())

	checkinService := checkin.NewService(checkinRepo)
	checkinHandler := checkin.NewHandler(checkinService)
	http.HandlerFunc(checkinHandler.AddCheckIn).
		ServeHTTP(rr, req)

	//Check for 200
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
	}

	//The expected json we want
	expected := string(`{
		"userId": "2342342424",
		"place": {
		"placeId":"37298472",
		"name": "Bistro 65",
		"lng": 31.112232,
		"lat": 20.123221,
		"category": "restaurant"
		},
		"checkinTimestamp": "2018-09-22T12:42:31Z"
		}`)

	//Assert Validates the expected json with the body string
	assert.JSONEq(t, expected, rr.Body.String(), "Response not same")
}

func BenchmarkCheckInHandler(b *testing.B) {

	jsonStr := []byte(`{
		"userId": "2342342424",
		"place": {
		"placeId":"37298472",
		"name": "Bistro 65",
		"lng": 31.112232,
		"lat": 20.123221,
		"category": "restaurant"
		},
		"checkinTimestamp": "2018-09-22T12:42:31Z"
		}`)

	var checkinRepo checkin.Repository

	checkinRepo = cassandra.NewCassandraCheckinRepository(cassandraConnection())

	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest("POST", "/addCheckIn", bytes.NewBuffer(jsonStr))

		rr := httptest.NewRecorder()

		checkinService := checkin.NewService(checkinRepo)
		checkinHandler := checkin.NewHandler(checkinService)
		http.HandlerFunc(checkinHandler.AddCheckIn).
			ServeHTTP(rr, req)
	}
}
