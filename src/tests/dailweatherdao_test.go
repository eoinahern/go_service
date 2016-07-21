package tests

import (
	"testing"

	"github.com/eoinahern/go_service/src/domain/entities"
	"github.com/eoinahern/go_service/src/domain/model"
)

var name = "eoin"
var pass = "pass"
var db = "weather_app"
var database = model.NewDatabase(name, pass, db)
var dailywdao = model.NewDailyWeatherDAO(database)
var dailyweather = entities.NewDailyWeather()

func Test_Insert(t *testing.T) {

	testdw = create_dailyweather()

	if "" != "eoin" {
		t.Error("incorrect name param!!")
	}

	t.Error("error because!!")
}

func create_dailyweather() entities.DailyWeather {
	dailyweather.Name = "ea"
	dailyweather.Summary = "ea"
	dailyweather.Time = "ea"
	dailyweather.PrecipProbability = "ea"
	dailyweather.Pressure = "ea"
	dailyweather.Icon = "ea"
	dailyweather.SunriseTime = "ea"
	dailyweather.Name = "ea"

}
