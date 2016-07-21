package cron

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/eoinahern/go_service/src/domain/entities"
	"github.com/eoinahern/go_service/src/domain/model"
)

//1.call external service for each city in my db!!
//2. data for each call returns json obj

type Fullapi struct {
	Dailydata entities.DailyWeatherContainer `json:"daily"`
}

func LoadServiceDataPerCity() {

	var apikey string = "63f0914cdd082e76d25b40161cbe70c4"
	dbconn := model.NewDatabase("eoin", "pass", "weather_app")
	citydao := model.NewCityDAO(dbconn)
	dailyweatherdao := model.NewDailyWeatherDAO(dbconn)
	cities := citydao.GetAllCities()

	for _, cityval := range cities {

		lat := strconv.FormatFloat(cityval.Latitude, 'E', 8, 64)
		longit := strconv.FormatFloat(cityval.Longitude, 'E', 8, 64)

		println(lat)
		println(longit)

		cal := fmt.Sprintf("https://api.forecast.io/forecast/%s/%s,%s", apikey, lat, longit)
		println(cal)
		resp, err := http.Get(cal)

		if err != nil {
			println("error on api call")
		}

		defer resp.Body.Close()
		dailyweather := unmarshallData(resp)
		weatherdatawname := appendName(cityval.Name, dailyweather.Dailydata.Dw)
		go InsertData(weatherdatawname, dailyweatherdao)
		resp.Body.Close()
	}

	//	InsertData(dailyweatherslice, dailyweatherdao)
}

func unmarshallData(resp *http.Response) *Fullapi {
	dailyweather := new(Fullapi)
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &dailyweather)
	return dailyweather
}

func appendName(citname string, dailywlist []*entities.DailyWeather) []*entities.DailyWeather {

	for ind, item := range dailywlist {
		item.Name = citname
		dailywlist[ind] = item
	}
	return dailywlist
}

//4. add each dailyweather element to the dailweather database

func InsertData(dailyweather []*entities.DailyWeather, weatherdao *model.DailyWeatherDAO) {
	weatherdao.Insert(dailyweather)

	println(dailyweather[0].Summary)
	println(dailyweather[0].Icon)
	println("")

}
