package cassandra

import (
	"log"

	"github.com/gocql/gocql"
	"github.com/ni/checkin"
)

type checkinRepository struct {
	cql *gocql.Session
}

func NewCassandraCheckinRepository(cql *gocql.Session) checkin.Repository {
	return &checkinRepository{
		cql,
	}
}

func (r *checkinRepository) AddCheckIn(chkin checkin.CheckIn) error {

	if err := r.cql.Query("INSERT INTO checkin (checkinid, userid, placeid, name, lat,lng, category, checkintimestamp) VALUES (?,?,?,?,?,?,?,?)", gocql.TimeUUID(), chkin.UserID, chkin.Place.PlaceID, chkin.Place.Name, chkin.Place.Latitude, chkin.Place.Longitude, chkin.Place.Category, chkin.CheckinTimestamp).Exec(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
