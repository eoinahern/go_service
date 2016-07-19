package tests

import (
	"testing"
)

func Testtest(t *testing.T) {

	shit := -1

	if shit > 0 {
		t.Log("test failed!!!!!")
		t.Fail()
	}

	t.Fail()

}
