package checkin

import "time"

type CheckIn struct {
	UserID string `json:"userId"`
	Place  struct {
		PlaceID   string  `json:"placeId"`
		Name      string  `json:"name"`
		Longitude float64 `json:"lng"`
		Latitude  float64 `json:"lat"`
		Category  string  `json:"category"`
	} `json:"place"`
	CheckinTimestamp *time.Time `json:"checkinTimestamp"`
}
