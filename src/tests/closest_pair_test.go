package tests

import "testing"

func Test_Nearest(t *testing.T) {

	list := map[string]float64{
		"cork":   5.00,
		"dublin": 6.00,
	}

	if Closest_coords.ShortestDist(list) != "cork" {
		t.Error("wrong shortest found!!")
	}

}
