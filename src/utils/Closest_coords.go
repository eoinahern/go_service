package utils

import (
	"math"

	"github.com/eoinahern/go_service/src/domain/entities"
)

//brute force algorithm for now

func findClosest(cities []*entities.City, latitude float64, longitude float64) *entities.City {

	return entities.NewCity()
}

func ShortestDist(citiesmap map[string]float64) string {
	var shortest = math.MaxFloat64
	var smallestplacename string

	for name, val := range citiesmap {
		if val < shortest {
			smallestplacename = name
			shortest = val
		}
	}

	return smallestplacename
}

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

//Haversine Formula

func GetDist(cities []*entities.City, latitude float64, longitude float64) map[string]float64 {

	distmap := make(map[string]float64)

	for _, place := range cities {

		/*dlat := (place.Latitude * (math.Pi / 180)) - (latitude * (math.Pi / 180))
		dlon := (place.Longitude * (math.Pi / 180)) - (longitude * (math.Pi / 180))
		sum := (math.Sin(dlat/2)*math.Sin(dlat/2) + math.Cos((latitude*math.Pi/180))*math.Cos((place.Latitude*math.Pi/180)) + math.Sin(dlon/2)*math.Sin(dlon/2))
		c := 2 * math.Atan2(math.Sqrt(sum), math.Sqrt(1-sum))
		var r float64 = 6378100
		d := r * c*/

		var r float64 = 6378100

		h := hsin((place.Latitude*(math.Pi/180))-(latitude*(math.Pi/180))) + math.Cos((latitude*(math.Pi/180)))*math.Cos((place.Latitude*(math.Pi/180)))*hsin((place.Longitude*(math.Pi/180))-(longitude*(math.Pi/180)))
		distmap[place.Name] = 2 * r * math.Asin(math.Sqrt(h))

		//distmap[place.Name] = d
	}

	return distmap
}
