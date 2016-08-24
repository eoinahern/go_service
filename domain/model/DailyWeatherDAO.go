package model

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/eoinahern/go_service/src/domain/entities"
)

func NewDailyWeatherDAO(dbconnin *Database) *DailyWeatherDAO {
	dweather := new(DailyWeatherDAO)
	dweather.base = NewBaseDao(dbconnin)
	//dweather.dbconn = dbconnin
	return dweather

}

type DailyWeatherDAO struct {
	base   *baseDao
	dbconn *Database
}

func (dw *DailyWeatherDAO) Insert(weatheritems []*entities.DailyWeather) bool {

	keyStrings := make([]string, 0)
	values := make([]interface{}, 0)

	for _, weatherval := range weatheritems {

		keyStrings = append(keyStrings, `(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)

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

	stmt := fmt.Sprintf(`INSERT into dailyweather (name, time, summary,
		icon, sunriseTime, sunsetTime, precipProbability, temperatureMin, temperatureMinTime,
		temperatureMax, temperatureMaxTime, apparentTemperatureMaxTime, dewPoint,
		windSpeed, humidity, pressure, cloudCover) VALUES %s`, strings.Join(keyStrings, ","))

	fmt.Println(stmt)
	fmt.Println(strings.Join(keyStrings, ","))
	fmt.Println(values)

	//_, err := dw.dbconn.mydbconn.Exec(stmt, values...)
	_, err := dw.base.db.mydbconn.Exec(stmt, values...)

	if err != nil {
		println("data not inserted")
		return false
	}

	return true

}

//pretty verbose to just count rows in a DB table lol.

func (dw *DailyWeatherDAO) CountRows() int {
	return dw.base.CountRows("dailyweather")
}

func (dw *DailyWeatherDAO) Delete(city string, time int) int {

	stmt := fmt.Sprintf("Delete FROM dailyweather WHERE name = '%s' AND time = %d", city, time)
	println(stmt)

	rows, err := dw.base.db.mydbconn.Query(stmt)
	if err != nil {
		println("delete failed")
		return -1
	}

	defer rows.Close()
	return checkcount(rows)
}

func (dw *DailyWeatherDAO) Get(city string) []entities.DailyWeather {

	rows, err := dw.dbconn.mydbconn.Query("SELECT * FROM dailyweather WHERE name = ?;", city)

	if err != nil {
		println("error calling query")
		log.Fatal(err)
	}

	defer rows.Close()
	newrows := createJsonWeather(rows)
	rows.Close()
	return newrows
}

//wasnt sure how to make this more generic just yet.
//params differ etc need to be scanned into slice

func createJsonWeather(rows *sql.Rows) []entities.DailyWeather {

	data := make([]entities.DailyWeather, 0)
	for rows.Next() {

		dailyweather := entities.NewDailyWeather()
		rows.Scan(&dailyweather.Name, &dailyweather.Icon, &dailyweather.Time,
			dailyweather.Summary)
		data = append(data, *dailyweather)
	}

	if len(data) < 1 {
		log.Output(1, "no data returned")
		println("error reading from db")
	}

	return data
}

func (dw *DailyWeatherDAO) Update() {

}
