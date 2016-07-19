package tests

import (
	"testing"
)

//intentionally failing tests!!!

func Test_stuff(t *testing.T) {
	t.Error("an error!!!!")
}

func Test_tyrfail(t *testing.T) {
	t.Error("another error!!!")
}
