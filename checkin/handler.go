package checkin

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type CheckinHandler interface {
	AddCheckIn(w http.ResponseWriter, r *http.Request)
}

type checkinHandler struct {
	checkinService Service
}

func NewHandler(checkinService Service) CheckinHandler {
	return &checkinHandler{
		checkinService,
	}
}

func WriteResponse(w http.ResponseWriter, code int, response interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}

func (c *CheckIn) IsValid() (bool, error) {

	if len(c.UserID) == 0 {
		return false, errors.New("userId is required")
	}

	if len(c.Place.PlaceID) == 0 {
		return false, errors.New("placeId is required")
	}

	if &c.Place.Longitude == nil || c.Place.Longitude < 0 || c.Place.Longitude > 180 {
		return false, errors.New("longitude value is missing or invalid")
	}

	if &c.Place.Longitude == nil || c.Place.Latitude < 0 || c.Place.Latitude > 90 {
		return false, errors.New("Latitude value is missing or invalid")
	}

	if len(c.Place.Category) == 0 {
		return false, errors.New("Category is required")
	}

	if c.CheckinTimestamp == nil {
		return false, errors.New("Checkin Timestamp is missing")
	}

	return true, nil
}

func (c *checkinHandler) AddCheckIn(w http.ResponseWriter, r *http.Request) {
	var chkin CheckIn

	data, ioErr := ioutil.ReadAll(r.Body)

	if ioErr != nil {
		// Handling IO Error
		err := map[string]interface{}{"message": "Invalid Request", "validationError": ioErr.Error()}
		WriteResponse(w, http.StatusBadRequest, err)
		return
	}
	if jsonErr := json.Unmarshal([]byte(data), &chkin); jsonErr != nil {
		// Handling JSON parsing error
		err := map[string]interface{}{"message": "Invalid Request", "validationError": jsonErr.Error()}
		WriteResponse(w, http.StatusBadRequest, err)
		return
	}

	isValid, validationError := chkin.IsValid()
	if !isValid {
		err := map[string]interface{}{"message": "Invalid Request", "validationError": validationError.Error()}
		WriteResponse(w, http.StatusBadRequest, err)
		return
	}

	err := c.checkinService.AddCheckIn(chkin)

	if err != nil {
		WriteResponse(w, http.StatusInternalServerError, err)
		return
	}

	WriteResponse(w, http.StatusOK, chkin)
	return
}
