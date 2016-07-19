package cron

import (
	"domain/model"
	"domain/entities"
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
	"io/ioutil"
)

//1.call external service for each city in my db!!
//2. data for each call returns json obj

type Fullapi struct {
		Dailydata entities.DailyWeatherContainer `json:"daily"`
}

func LoadServiceDataPerCity(){
	
	var apikey string = "63f0914cdd082e76d25b40161cbe70c4"
	dbconn := model.NewDatabase("eoin", "pass","weather_app");
	citydao := model.NewCityDAO(dbconn)
	//dailyweatherslice = [][] entities.DailyWeather
	dailyweatherdao := model.NewDailyWeatherDAO(dbconn);
	cities := citydao.GetAllCities()
	
	for _, cityval := range cities {
		
		lat := strconv.FormatFloat(cityval.Latitude,'E',8,64)
		longit := strconv.FormatFloat(cityval.Longitude,'E',8,64)
		
		cal := fmt.Sprintf("https://api.forecast.io/forecast/?/?,?" ,apikey, lat, longit);
		resp, err := http.Get(cal)
		
		if(err != nil){
			println("error on api call");
		}
		 
		 defer resp.Body.Close()
		 dailyweather := unmarshallData(resp)
		 dailyweatherslice = append(dailyweatherslice, dailyweather.Dailydata.Dw)   // is a slice. iterate over these
		 resp.Body.Close()
	}
	
	//InsertData(,dailyweatherdao)
}

func unmarshallData(resp *http.Response)  *Fullapi {
	
		 dailyweather := new(Fullapi)
		 body, _ := ioutil.ReadAll(resp.Body)
		 json.Unmarshal(body, &dailyweather)
		 return dailyweather
}


//4. add each dailyweather element to the dailweather database

func InsertData(dailyweather []*entities.DailyWeather, weatherdao *model.DailyWeatherDAO){
	weatherdao.Insert(dailyweather)
}










