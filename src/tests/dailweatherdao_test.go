package tests

import (
	"testing"

	"github.com/eoinahern/go_service/src/domain/entities"
	"github.com/eoinahern/go_service/src/domain/model"
)

//could do with another test db here.
//at present testing prod.

var name string
var pass string
var db string
var database model.Database
var dailywdao *model.DailyWeatherDAO

func init() {
	var name = "eoin"
	var pass = "pass"
	var db = "weather_app"
	var database = model.NewDatabase(name, pass, db)
	dailywdao = model.NewDailyWeatherDAO(database)
}

func Test_Insert(t *testing.T) {

	t.Parallel()
	dailwslice := create_dailyweather()

	inserted := dailywdao.Insert(dailwslice)
	if inserted == false {
		t.Error("insert failed!!! oops!!")
	}

}

//count rows arter item deleted
func Test_Delete(t *testing.T) {

	testdw := create_dailyweather()
	inserted := dailywdao.Insert(testdw)
	if inserted == false {
		t.Error("insert failed!!! oops!! delete test")
	}

	count := dailywdao.CountRows()

	//deleted := dailywdao.Delete("cork", 1469287025)

	//delete
	if count <= dailywdao.CountRows() {
		t.Error("delete failed!!!")
	}

}

func create_dailyweather() []*entities.DailyWeather {

	//fake obj
	dailwslice := make([]*entities.DailyWeather, 0)
	dailyweather := entities.NewDailyWeather()

	dailyweather.Name = "cork"
	dailyweather.Summary = "lovely weather"
	dailyweather.Time = 1469287025
	dailyweather.PrecipProbability = 0.82
	dailyweather.Pressure = 0.01
	dailyweather.Icon = "no icon"
	dailyweather.SunriseTime = 1469287025
	dailyweather.DewPoint = 0.49
	dailyweather.WindSpeed = 0.50
	dailyweather.Humidity = 0.50
	dailyweather.CloudCover = 0.70
	dailyweather.SunsetTime = 1469287025
	dailyweather.TemperatureMin = 21.5
	dailyweather.TemperatureMinTime = 1469287025
	dailyweather.TemperatureMax = 24.5
	dailyweather.TemperatureMaxTime = 1469287025
	dailyweather.ApparentTemperatureMaxTime = 1469287025

	dailwslice = append(dailwslice, dailyweather)

	return dailwslice
}
