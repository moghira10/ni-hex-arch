package mysql

import (
	"database/sql"

	"github.com/ni/checkin"
)

type checkinRepository struct {
	db *sql.DB
}

func NewMysqlCheckinRepository(db *sql.DB) checkin.Repository {
	return &checkinRepository{
		db,
	}
}

func (r *checkinRepository) AddCheckIn(chkin checkin.CheckIn) error {
	ts := chkin.CheckinTimestamp.String()
	_, err := r.db.Query("INSERT INTO checkin VALUES ("+chkin.UserID, chkin.Place.PlaceID, chkin.Place.Name, chkin.Place.Latitude, chkin.Place.Longitude, chkin.Place.Category, ts+")")
	if err != nil {
		return err
	}
	return nil
}
