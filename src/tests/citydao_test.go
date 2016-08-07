package tests

import (
	"testing"

	"github.com/eoinahern/go_service/src/domain/model"
)

//failing tests now as more cities added to db.

func Test_getbycity(t *testing.T) {

	t.Parallel()
	dbconn := model.NewDatabase("eoin", "pass", "weather_app_test")
	citydao := model.NewCityDAO(dbconn)
	cities := citydao.GetAllCities()

	for _, city := range cities {

		//check cit within bounds

		if city.Latitude > 55.00 || city.Latitude < 40.00 {
			t.Error("lat out of bounded regoin")
		}

		if city.Longitude > 3.00 || city.Longitude < -9.00 {
			t.Error("long out of bounded region")
		}

	}

	if len(cities) != 4 {
		t.Error("incorrect cities number!!")
	}
}
