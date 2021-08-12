package revcal

import (
	"testing"
	"time"
)

func TestCoupOf18Brumaire(t *testing.T) {
	date := NewDate(time.Date(1799, time.November, 9, 0, 0, 0, 0, time.UTC))
	if date.Format() != "18 Brumaire 8" {
		t.Errorf(`18 Brumaire 8 != %s`, date.Format())
	}
}

func TestJulyRevolution(t *testing.T) {
	date := NewDate(time.Date(1830, time.July, 26, 0, 0, 0, 0, time.UTC))
	if date.Format() != "7 Thermidor 38" {
		t.Errorf(`7 Thermidor 38 != %s`, date.Format())
	}
}

func TestElectionOfNapoleonIII(t *testing.T) {
	date := NewDate(time.Date(1851, time.December, 2, 0, 0, 0, 0, time.UTC))
	if date.Format() != "11 Frimaire 60" {
		t.Errorf(`11 Frimaire 60 != %s`, date.Format())
	}
}
