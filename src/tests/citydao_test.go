package tests

import (
	"testing"

	"github.com/eoinahern/go_service/src/domain/model"
)

//intentionally failing tests!!!

func Test_getbycity(t *testing.T) {

	dbconn := model.NewDatabase("eoin", "pass", "weather_app")
	citydao := model.NewCityDAO(dbconn)
	cities := citydao.GetAllCities()

	for _, city := range cities {

		if city.Name != "cork" && city.Name != "dublin" {
			t.Error("unknown city")
		}

		if city.Latitude > 55.00 || city.Latitude < 51.00 {
			t.Error("lat incorrect for Ireland!")
		}

	}

	if len(cities) != 2 {
		t.Error("incorrect cities number!!")
	}
}
