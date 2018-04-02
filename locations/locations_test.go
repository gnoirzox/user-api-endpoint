package locations

import (
	"strconv"
	"testing"
)

func TestIsValidLongitude(t *testing.T) {
	var tests = []struct {
		input float64
		want  bool
	}{
		{0.00, true},
		{-0.00, true},
		{180.00, true},
		{125.463, true},
		{-125.463, true},
		{-223.407, false},
		{223.407, false},
	}

	for _, test := range tests {
		var l Location
		l.Longitude = test.input

		if got := l.IsValidLongitude(); got != test.want {
			t.Errorf("l.Longitude = %q l.IsValidLongitude() = %v", strconv.FormatFloat(test.input, 'f', -1, 64), got)
		}
	}
}

func TestIsValidLatitude(t *testing.T) {
	var tests = []struct {
		input float64
		want  bool
	}{
		{0.00, true},
		{-0.00, true},
		{90.00, true},
		{-90, true},
		{-23.07, true},
		{23.47, true},
		{125.463, false},
		{-125.463, false},
	}

	for _, test := range tests {
		var l Location
		l.Latitude = test.input

		if got := l.IsValidLatitude(); got != test.want {
			t.Errorf("l.Latitude = %q l.IsValidLatitude() = %v", strconv.FormatFloat(test.input, 'f', -1, 64), got)
		}
	}
}
