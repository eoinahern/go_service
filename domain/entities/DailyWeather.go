package entities

func NewDailyWeather() *DailyWeather {
    m := new(DailyWeather)
    return m
}

type DailyWeather struct {

  
  Name string `json: "name"`
  Time int      `json:"time"`
  Summary string `json:"summary"`
  Icon string	`json:"icon"`
  SunriseTime int `json:"sunriseTime"`
  SunsetTime int  `json:"sunsetTime"`
  PrecipProbability float32 `json:"precipProbability"`
  TemperatureMin float32 `json:"temperatureMin"`
  TemperatureMinTime int `json:"temperatureMinTime"`
  TemperatureMax float32	`json:"temperatureMax"`
  TemperatureMaxTime int	`json:"temperatureMaxTime"`
  ApparentTemperatureMaxTime int	`json:"apparentTemperatureMaxTime"`
  DewPoint float32	`json:"dewPoint"`
  WindSpeed float32	`json:"windSpeed"`
  Humidity float32	`json:"humidity"`
  Pressure float32	`json:"pressure"`
  CloudCover float32	`json:"cloudCover"`
}

func (w *DailyWeather) getTime() int {
    return w.Time
}


