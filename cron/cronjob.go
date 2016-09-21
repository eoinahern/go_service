package cron

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/eoinahern/go_service/domain/entities"
	"github.com/eoinahern/go_service/domain/model"
)

//1.call external service for each city in my db!!
//2. data for each call returns json obj
var apikey = "63f0914cdd082e76d25b40161cbe70c4"
var dbconn *model.Database
var citydao *model.CityDAO
var dailyweatherdao *model.DailyWeatherDAO
var ch1 chan []*entities.DailyWeather
var wg sync.WaitGroup

func initvars() {

	//dbconn = model.NewDatabase("bd145d3b601f2e", "532d35c9", "heroku_1587748f259385b")

	dbconn = model.NewDatabase("eoin", "pass", "weather_app")
	citydao = model.NewCityDAO(dbconn)
	dailyweatherdao = model.NewDailyWeatherDAO(dbconn)
	ch1 = make(chan []*entities.DailyWeather)
}

type Fullapi struct {
	Dailydata entities.DailyWeatherContainer `json:"daily"`
}

func LoadServiceDataPerCity() {
	initvars()
	InsertRows()
}

func InsertRows() {

	cities := citydao.GetAllCities()
	wg.Add(len(cities) + 1)

	go func() {

		for _, cityval := range cities {
			lat, longit := getCoordsString(cityval.Latitude, cityval.Longitude)

			dailyweatherdao.DeleteAll(cityval.Name)
			resp := callSevice(lat, longit)
			ch1 <- buildObj(resp, cityval.Name)
			wg.Done()
		}
		close(ch1)

	}()

	go InsertData()
	wg.Wait()
}

func getCoordsString(lat float64, long float64) (slong string, slat string) {

	latit := strconv.FormatFloat(lat, 'f', 5, 64)
	longit := strconv.FormatFloat(long, 'f', 5, 64)
	return latit, longit

}

func callSevice(lat string, long string) (resp *http.Response) {

	cal := fmt.Sprintf("https://api.forecast.io/forecast/%s/%s,%s", apikey, lat, long)
	response, err := http.Get(cal)

	if err != nil {
		println("error on api call")
		log.Fatal(err)
	}

	defer resp.Body.Close()
	return response
}

func buildObj(resp *http.Response, name string) []*entities.DailyWeather {
	dailyweather := unmarshallData(resp)
	return appendName(name, dailyweather.Dailydata.Dw)
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

func InsertData() {

	for val := range ch1 {
		dailyweatherdao.Insert(val)
	}
	wg.Done()
}
