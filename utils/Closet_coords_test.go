package utils

import (
	"fmt"
	"testing"

	"github.com/eoinahern/go_service/domain/entities"
)

func Test_Nearest(t *testing.T) {

	list := map[string]float64{
		"cork":   5.00,
		"dublin": 6.00,
	}

	if ShortestDist(list) != "cork" {
		t.Error("wrong shortest found!!")
	}
}

func Test_Closest_Place(t *testing.T) {

	citlist := getList()
	distmap := GetDist(citlist, 40.9048, 2.15003)

	fmt.Printf("cork : %f and barca : %f", distmap["cork"], distmap["barcelona"])

	if distmap["cork"] < distmap["barcelona"] {
		t.Error("innacurate distance calculation")
	}

}

//helper small list

func getList() []*entities.City {

	cities := []*entities.City{
		{Name: "cork", Latitude: 51.8969, Longitude: 8.4863},
		{Name: "barcelona", Latitude: 41.390205, Longitude: 2.154007},
	}

	return cities

}
