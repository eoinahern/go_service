package entities


func NewCity() (*City) {
	city := new(City)
	return city
}


func NewCityVars(name string, latitude float64, longitude float64) (*City) {
	city := new(City)
	city.Latitude = latitude
	city.Longitude = longitude
	city.Name = name
	return city
}



type City struct {
	Name string 	`json:"name"`
	Latitude float64 	`json :"latitude"`
	Longitude float64	`json: "longitude"`
}









