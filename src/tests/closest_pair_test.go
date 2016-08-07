package tests

import (
	"testing"

	"github.com/eoinahern/go_service/src/utils"
)

func Test_Nearest(t *testing.T) {

	list := map[string]float64{
		"cork":   5.00,
		"dublin": 6.00,
	}

	if utils.ShortestDist(list) != "cork" {
		t.Error("wrong shortest found!!")
	}
}
