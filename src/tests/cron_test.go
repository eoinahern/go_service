package tests

import (
	"testing"

	"github.com/eoinahern/go_service/src/cron"
)

func Test_cron(t *testing.T) {

	count := dailywdao.CountRows()
	cron.LoadServiceDataPerCity()

	if count == dailywdao.CountRows() {
		t.Error("no new data inserted!")
	}
}
