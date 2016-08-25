package tests

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/eoinahern/go_service/domain/entities"
	"github.com/eoinahern/go_service/domain/model"
)

var name string
var pass string
var db string
var database model.Database
var dailywdao *model.DailyWeatherDAO

func init() {
	var name = "eoin"
	var pass = "pass"
	var db = "weather_app_test"
	var database = model.NewDatabase(name, pass, db)
	dailywdao = model.NewDailyWeatherDAO(database)
}

//need get test

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
	dailywdao.Delete("cork", 1469287025)

	if count > dailywdao.CountRows() {
		t.Error("delete failed!!!")
	}

}

func Test_Get(t *testing.T) {

	dailylist := dailywdao.Get("cork")
	if dailywdao.CountRows() == 0 {
		t.Error("counted rows returned 0")
	}

	fmt.Printf("num rows = %d", len(dailylist))

	for _, item := range dailylist {
		if item == nil || len(item.Name) > 0 {
			t.Error("null item found")
		}

		if reflect.TypeOf(item.Name).String() != "string" {
			t.Error("name is not in string format")
		}

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