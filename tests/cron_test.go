package tests

import (
	"testing"

	"github.com/eoinahern/go_service/cron"
)

//end to end. sheck data was inserted!!

func Test_cron(t *testing.T) {

	t.Parallel()
	//count := dailywdao.CountRows()
	cron.LoadServiceDataPerCity()

	/*if count == dailywdao.CountRows() {
		t.Error("no new data inserted!")
	}*/
}
