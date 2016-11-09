package internal

import (
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	if location == nil {
		t.Fatal("location is nil.")
	}
}

func TestDate(t *testing.T) {
	s := Date(time.Now())
	if len(s) != len(dateLayout) {
		t.Errorf("date: %s", s)
	}
}

func TestDatetime(t *testing.T) {
	s := Datetime(time.Now())
	if len(s) != len(datetimeLayout) {
		t.Errorf("datetime: %s", s)
	}
}
