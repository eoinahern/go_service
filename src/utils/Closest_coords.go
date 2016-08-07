package utils

import (
	"math"

	"github.com/eoinahern/go_service/src/domain/entities"
)

//brute force algorithm for now

func findClosest(cities []*entities.City, latitude float64, longitude float64) *entities.City {

	return entities.NewCity()

}

//Haversine Formula

func getDist(cities []*entities.City, latitude float64, longitude float64) map[string]float64 {

	distmap := make(map[string]float64)

	for _, place := range cities {

		dlat := ((place.Latitude * math.Pi / 180) - (latitude * math.Pi / 180))
		dlon := ((place.Longitude * math.Pi / 180) - (longitude * math.Pi / 180))

		sum := (math.Sin(dlat/2)*math.Sin(dlat/2) + math.Cos((latitude*math.Pi/180))*math.Cos((place.Latitude*math.Pi/180)) + math.Sin(dlon/2)*math.Sin(dlon/2))
		c := 2 * math.Atan2(math.Sqrt(sum), math.Sqrt(1-sum))
		var r float64 = 6378100
		d := r * c

		distmap[place.Name] = d
	}

	return distmap
}
