package cron

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/eoinahern/go_service/domain/entities"
	"github.com/eoinahern/go_service/domain/model"
)

//1.call external service for each city in my db!!
//2. data for each call returns json obj

type Fullapi struct {
	Dailydata entities.DailyWeatherContainer `json:"daily"`
}

func main() {
	LoadServiceDataPerCity()
}

func LoadServiceDataPerCity() {

	var apikey string = "63f0914cdd082e76d25b40161cbe70c4"
	dbconn := model.NewDatabase("eoin", "pass", "weather_app")
	citydao := model.NewCityDAO(dbconn)
	dailyweatherdao := model.NewDailyWeatherDAO(dbconn)
	cities := citydao.GetAllCities()

	for _, cityval := range cities {

		//helper in utils possibly
		lat := strconv.FormatFloat(cityval.Latitude, 'f', 5, 64)
		longit := strconv.FormatFloat(cityval.Longitude, 'f', 5, 64)

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
		fmt.Println(dailyweather.Dailydata.Dw)
		go InsertData(weatherdatawname, dailyweatherdao)
		resp.Body.Close()
	}

}

func unmarshallData(resp *http.Response) *Fullapi {
	dailyweather := new(Fullapi)
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &dailyweather)
	return dailyweather
}

func appendName(citname string, dailywlist []*entities.DailyWeather) []*entities.DailyWeather {

	name := fmt.Sprintf("citname is %s", citname)
	println(name)

	for ind, item := range dailywlist {
		item.Name = citname
		dailywlist[ind] = item
	}
	return dailywlist
}

//4. add each dailyweather element to the dailweather database

func InsertData(dailyweather []*entities.DailyWeather, weatherdao *model.DailyWeatherDAO) {
	weatherdao.Insert(dailyweather)
}
