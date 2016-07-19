package model

import (
	"domain/entities"
	"database/sql"
	"log"
	"strings"
	"fmt"
	
)

func NewDailyWeatherDAO(dbconnin *Database) *DailyWeatherDAO {
	d_weather := new(DailyWeatherDAO)
	d_weather.dbconn = dbconnin
	return d_weather
	
}

type DailyWeatherDAO struct {
	dbconn *Database
	
}

func (dw *DailyWeatherDAO) Insert(weatheritems []*entities.DailyWeather){
	
	keyStrings := make([]string,  17)
	values := make([]interface{}, 0)
	
	for _, weatherval := range weatheritems {
		
		keyStrings = append(keyStrings, `(?, ?, ?,
		 ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
		
		values = append(values, weatherval.Name)
		values = append(values, weatherval.Time)
		values = append(values, weatherval.Summary)
		values = append(values, weatherval.Icon)
		values = append(values, weatherval.SunriseTime)
		values = append(values, weatherval.SunsetTime)
		values = append(values, weatherval.PrecipProbability)
		values = append(values, weatherval.TemperatureMin)
		values = append(values, weatherval.TemperatureMinTime)
		values = append(values, weatherval.TemperatureMax)
		values = append(values, weatherval.TemperatureMaxTime)
		values = append(values, weatherval.ApparentTemperatureMaxTime)
		values = append(values, weatherval.DewPoint)
		values = append(values, weatherval.WindSpeed)
		values = append(values, weatherval.Humidity)
		values = append(values, weatherval.Pressure)
		values = append(values, weatherval.CloudCover)
		
	}
		
		stmt := fmt.Sprintf(`INSERT into dailyweather () VALUES %s`, strings.Join(keyStrings, ","))
		 _, err := dw.dbconn.mydbconn.Exec(stmt, values...)
		 
		 if(err != nil){
		 	println("data not inserted")
		 }
}


func (dw *DailyWeatherDAO) Delete(){
	
}


func (dw *DailyWeatherDAO) Get(city string) ([]entities.DailyWeather){
	
	rows, err := dw.dbconn.mydbconn.Query("SELECT * FROM dailyweather WHERE name = ?;", city)
	 
	 if err != nil {
		 	println( "error calling query")
			log.Fatal(err)
	 }
	 
	 defer rows.Close()
	 newrows := createJsonWeather(rows)
	 rows.Close()
	 return newrows
}

//wasnt sure how to make this more generic just yet.
//params differ etc need to be scanned into slice

func  createJsonWeather(rows *sql.Rows) ([]entities.DailyWeather){
	
	data :=  make([]entities.DailyWeather,0) 
	for rows.Next(){
		
		dailyweather := entities.NewDailyWeather()
		rows.Scan(&dailyweather.Name, &dailyweather.Icon, &dailyweather.Time,
		dailyweather.Summary)			
		data = append(data, *dailyweather)
	}
	
	if(len(data) < 1){
		log.Output(1, "no data returned")
		println("error reading from db")
	}
	  
	return data
}

func (dw *DailyWeatherDAO) Update(){
	
}



