package tz

import (
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	tz, err := Decode("Europe/London")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if tz.TzIdentifier() != "Europe/London" {
		t.Errorf("Expected Europe/London, got %s", tz.TzIdentifier())
	}

	if tz.CountryCode() != "GB" {
		t.Errorf("Expected GB, got %s", tz.CountryCode())
	}

	if tz.UtcOffset() != 0 {
		t.Errorf("Expected 0, got %f", tz.UtcOffset())
	}

	_, err = Decode("Invalid/Timezone")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestTimezoneMethods(t *testing.T) {
	tz := Timezone{
		tzIdentifier:       "America/New_York",
		defaultCountryCode: "US",
		otherCountryCodes:  []string{"CA"},
		utcOffset:          -5,
	}

	if tz.TzIdentifier() != "America/New_York" {
		t.Errorf("Expected America/New_York, got %s", tz.TzIdentifier())
	}

	if tz.CountryCode() != "US" {
		t.Errorf("Expected US, got %s", tz.CountryCode())
	}

	expectedCountryCodes := []string{"CA", "US"}
	if !reflect.DeepEqual(tz.CountryCodes(), expectedCountryCodes) {
		t.Errorf("Expected %v, got %v", expectedCountryCodes, tz.CountryCodes())
	}

	if tz.UtcOffset() != -5 {
		t.Errorf("Expected -5, got %f", tz.UtcOffset())
	}
}
