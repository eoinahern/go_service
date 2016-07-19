package entities

import (

)



func NewDailyWeatherContainer() *DailyWeatherContainer{
	cont := new(DailyWeatherContainer)
	return cont
}


type DailyWeatherContainer struct {
		Summary string `json:"summary"`
		Icon    string `json:"icon"`
		Dw  []DailyWeather	`json:"data"`
}

