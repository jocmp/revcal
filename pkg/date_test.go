package revcal

import "testing"

func TestCoupOf18Brumaire(t *testing.T) {
	date := FormatDate(1799, 11, 9)
	if date != "18 Brumaire 8" {
		t.Errorf(`FormatDate(1799, 11, 9) = %s`, date)
	}
}

func TestJulyRevolution(t *testing.T) {
	date := FormatDate(1830, 7, 26)
	if date != "7 Thermidor 38" {
		t.Errorf(`FormatDate(1830, 7, 26) = %s`, date)
	}
}

func TestElectionOfNapoleonIII(t *testing.T) {
	date := FormatDate(1851, 12, 2)
	if date != "11 Frimaire 60" {
		t.Errorf(`FormatDate(1851, 12, 2) = %s`, date)
	}
}
