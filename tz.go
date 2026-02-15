package tz

import "errors"

// Decode returns the timezone for a given timezone string.
func Decode(timezoneString string) (*Timezone, error) {
	// Check if the timezoneString is in the map
	if timezone, ok := tzIdentifierToTimezone[timezoneString]; ok {
		return &timezone, nil
	}

	// If not, return an error
	return nil, errors.New("could not find country code for timezone")
}

type Timezone struct {
	tzIdentifier       string
	defaultCountryCode string
	otherCountryCodes  []string
	utcOffset          float32
}

// Returns the tzIdentifier for the timezone.
func (t Timezone) TzIdentifier() string {
	return t.tzIdentifier
}

// Returns the default country code for the timezone.
func (t Timezone) CountryCode() string {
	return t.defaultCountryCode
}

// Return all the country codes for the timezone.
func (t Timezone) CountryCodes() []string {
	return append(t.otherCountryCodes, t.defaultCountryCode)
}

// Returns the UTC offset for the timezone.
func (t Timezone) UtcOffset() float32 {
	return t.utcOffset
}

// List of tzIdentifiers to timezones.
var tzIdentifierToTimezone = map[string]Timezone{
	"Europe/London": {
		tzIdentifier:       "Europe/London",
		defaultCountryCode: "GB",
		otherCountryCodes:  []string{},
		utcOffset:          0,
	},
	"Europe/Berlin": {
		tzIdentifier:       "Europe/Berlin",
		defaultCountryCode: "DE",
		otherCountryCodes:  []string{},
		utcOffset:          1,
	},
	"Europe/Paris": {
		tzIdentifier:       "Europe/Paris",
		defaultCountryCode: "FR",
		otherCountryCodes:  []string{},
		utcOffset:          1,
	},
	"Europe/Moscow": {
		tzIdentifier:       "Europe/Moscow",
		defaultCountryCode: "RU",
		otherCountryCodes:  []string{},
		utcOffset:          3,
	},
	"Europe/Madrid": {
		tzIdentifier:       "Europe/Madrid",
		defaultCountryCode: "ES",
		otherCountryCodes:  []string{},
		utcOffset:          1,
	},
	"Europe/Rome": {
		tzIdentifier:       "Europe/Rome",
		defaultCountryCode: "IT",
		otherCountryCodes:  []string{},
		utcOffset:          1,
	},
	"Europe/Kiev": {
		tzIdentifier:       "Europe/Kiev",
		defaultCountryCode: "UA",
		otherCountryCodes:  []string{},
		utcOffset:          2,
	},
	"Europe/Istanbul": {
		tzIdentifier:       "Europe/Istanbul",
		defaultCountryCode: "TR",
		otherCountryCodes:  []string{},
		utcOffset:          3,
	},
	"Europe/Bucharest": {
		tzIdentifier:       "Europe/Bucharest",
		defaultCountryCode: "RO",
		otherCountryCodes:  []string{},
		utcOffset:          2,
	},
	"Europe/Budapest": {
		tzIdentifier:       "Europe/Budapest",
		defaultCountryCode: "HU",
		otherCountryCodes:  []string{},
		utcOffset:          1,
	},
	"Europe/Warsaw": {
		tzIdentifier:       "Europe/Warsaw",
		defaultCountryCode: "PL",
		otherCountryCodes:  []string{},
		utcOffset:          1,
	},
	"Europe/Minsk": {
		tzIdentifier:       "Europe/Minsk",
		defaultCountryCode: "BY",
		otherCountryCodes:  []string{},
		utcOffset:          3,
	},
	"Europe/Prague": {
		tzIdentifier:       "Europe/Prague",
		defaultCountryCode: "CZ",
		otherCountryCodes:  []string{},
		utcOffset:          1,
	},
	"Europe/Stockholm": {
		tzIdentifier:       "Europe/Stockholm",
		defaultCountryCode: "SE",
		otherCountryCodes:  []string{},
		utcOffset:          1,
	},
	"Europe/Vienna": {
		tzIdentifier:       "Europe/Vienna",
		defaultCountryCode: "AT",
		otherCountryCodes:  []string{},
		utcOffset:          1,
	},
	"Europe/Sofia": {
		tzIdentifier:       "Europe/Sofia",
		defaultCountryCode: "BG",
		otherCountryCodes:  []string{},
		utcOffset:          2,
	},
	"Europe/Helsinki": {
		tzIdentifier:       "Europe/Helsinki",
		defaultCountryCode: "FI",
		otherCountryCodes:  []string{},
		utcOffset:          2,
	},
	"Europe/Copenhagen": {
		tzIdentifier:       "Europe/Copenhagen",
		defaultCountryCode: "DK",
		otherCountryCodes:  []string{},
		utcOffset:          1,
	},
	"Europe/Athens": {
		tzIdentifier:       "Europe/Athens",
		defaultCountryCode: "GR",
		otherCountryCodes:  []string{},
		utcOffset:          2,
	},
	"Europe/Brussels": {
		tzIdentifier:       "Europe/Brussels",
		defaultCountryCode: "BE",
		otherCountryCodes:  []string{},
		utcOffset:          1,
	},
	"Europe/Dublin": {
		tzIdentifier:       "Europe/Dublin",
		defaultCountryCode: "IE",
		otherCountryCodes:  []string{},
		utcOffset:          0,
	},
	"Europe/Amsterdam": {
		tzIdentifier:       "Europe/Amsterdam",
		defaultCountryCode: "NL",
		otherCountryCodes:  []string{},
		utcOffset:          1,
	},
	"Europe/Oslo": {
		tzIdentifier:       "Europe/Oslo",
		defaultCountryCode: "NO",
		otherCountryCodes:  []string{},
		utcOffset:          1,
	},
	"Europe/Zurich": {
		tzIdentifier:       "Europe/Zurich",
		defaultCountryCode: "CH",
		otherCountryCodes:  []string{},
		utcOffset:          1,
	},
	"Europe/Lisbon": {
		tzIdentifier:       "Europe/Lisbon",
		defaultCountryCode: "PT",
		otherCountryCodes:  []string{},
		utcOffset:          0,
	},
	"Europe/Bratislava": {
		tzIdentifier:       "Europe/Bratislava",
		defaultCountryCode: "SK",
		otherCountryCodes:  []string{},
		utcOffset:          1,
	},
	"Europe/Ljubljana": {
		tzIdentifier:       "Europe/Ljubljana",
		defaultCountryCode: "SI",
		otherCountryCodes:  []string{},
		utcOffset:          1,
	},
	"Europe/Riga": {
		tzIdentifier:       "Europe/Riga",
		defaultCountryCode: "LV",
		otherCountryCodes:  []string{},
		utcOffset:          2,
	},
	"Europe/Tallinn": {
		tzIdentifier:       "Europe/Tallinn",
		defaultCountryCode: "EE",
		otherCountryCodes:  []string{},
		utcOffset:          2,
	},
	"Europe/Vilnius": {
		tzIdentifier:       "Europe/Vilnius",
		defaultCountryCode: "LT",
		otherCountryCodes:  []string{},
		utcOffset:          2,
	},
	"Europe/Reykjavik": {
		tzIdentifier:       "Europe/Reykjavik",
		defaultCountryCode: "IS",
		otherCountryCodes:  []string{},
		utcOffset:          0,
	},
	"America/New_York": {
		tzIdentifier:       "America/New_York",
		defaultCountryCode: "US",
		otherCountryCodes:  []string{},
		utcOffset:          -5,
	},
	"America/Chicago": {
		tzIdentifier:       "America/Chicago",
		defaultCountryCode: "US",
		otherCountryCodes:  []string{},
		utcOffset:          -6,
	},
	"America/Denver": {
		tzIdentifier:       "America/Denver",
		defaultCountryCode: "US",
		otherCountryCodes:  []string{},
		utcOffset:          -7,
	},
	"America/Los_Angeles": {
		tzIdentifier:       "America/Los_Angeles",
		defaultCountryCode: "US",
		otherCountryCodes:  []string{},
		utcOffset:          -8,
	},
	"America/Anchorage": {
		tzIdentifier:       "America/Anchorage",
		defaultCountryCode: "US",
		otherCountryCodes:  []string{},
		utcOffset:          -9,
	},
	"America/Sao_Paulo": {
		tzIdentifier:       "America/Sao_Paulo",
		defaultCountryCode: "BR",
		otherCountryCodes:  []string{},
		utcOffset:          -3,
	},
	"America/Manaus": {
		tzIdentifier:       "America/Manaus",
		defaultCountryCode: "BR",
		otherCountryCodes:  []string{},
		utcOffset:          -4,
	},
	"America/Mexico_City": {
		tzIdentifier:       "America/Mexico_City",
		defaultCountryCode: "MX",
		otherCountryCodes:  []string{},
		utcOffset:          -6,
	},
	"America/Bogota": {
		tzIdentifier:       "America/Bogota",
		defaultCountryCode: "CO",
		otherCountryCodes:  []string{},
		utcOffset:          -5,
	},
	"America/Caracas": {
		tzIdentifier:       "America/Caracas",
		defaultCountryCode: "VE",
		otherCountryCodes:  []string{},
		utcOffset:          -4,
	},
	"America/Santiago": {
		tzIdentifier:       "America/Santiago",
		defaultCountryCode: "CL",
		otherCountryCodes:  []string{},
		utcOffset:          -3,
	},
	"America/Argentina/Buenos_Aires": {
		tzIdentifier:       "America/Argentina/Buenos_Aires",
		defaultCountryCode: "AR",
		otherCountryCodes:  []string{},
		utcOffset:          -3,
	},
	"America/Lima": {
		tzIdentifier:       "America/Lima",
		defaultCountryCode: "PE",
		otherCountryCodes:  []string{},
		utcOffset:          -5,
	},
	"America/Halifax": {
		tzIdentifier:       "America/Halifax",
		defaultCountryCode: "CA",
		otherCountryCodes:  []string{},
		utcOffset:          -4,
	},
	"America/Regina": {
		tzIdentifier:       "America/Regina",
		defaultCountryCode: "CA",
		otherCountryCodes:  []string{},
		utcOffset:          -6,
	},
	"America/St_Johns": {
		tzIdentifier:       "America/St_Johns",
		defaultCountryCode: "CA",
		otherCountryCodes:  []string{},
		utcOffset:          -3.5,
	},
	"America/Vancouver": {
		tzIdentifier:       "America/Vancouver",
		defaultCountryCode: "CA",
		otherCountryCodes:  []string{},
		utcOffset:          -8,
	},
	"America/Edmonton": {
		tzIdentifier:       "America/Edmonton",
		defaultCountryCode: "CA",
		otherCountryCodes:  []string{},
		utcOffset:          -7,
	},
	"America/Toronto": {
		tzIdentifier:       "America/Toronto",
		defaultCountryCode: "CA",
		otherCountryCodes:  []string{},
		utcOffset:          -5,
	},
	"America/Whitehorse": {
		tzIdentifier:       "America/Whitehorse",
		defaultCountryCode: "CA",
		otherCountryCodes:  []string{},
		utcOffset:          -8,
	},
	"America/Winnipeg": {
		tzIdentifier:       "America/Winnipeg",
		defaultCountryCode: "CA",
		otherCountryCodes:  []string{},
		utcOffset:          -6,
	},
	"America/Havana": {
		tzIdentifier:       "America/Havana",
		defaultCountryCode: "CU",
		otherCountryCodes:  []string{},
		utcOffset:          -5,
	},
	"America/La_Paz": {
		tzIdentifier:       "America/La_Paz",
		defaultCountryCode: "BO",
		otherCountryCodes:  []string{},
		utcOffset:          -4,
	},
	"America/Asuncion": {
		tzIdentifier:       "America/Asuncion",
		defaultCountryCode: "PY",
		otherCountryCodes:  []string{},
		utcOffset:          -3,
	},
	"America/Montevideo": {
		tzIdentifier:       "America/Montevideo",
		defaultCountryCode: "UY",
		otherCountryCodes:  []string{},
		utcOffset:          -3,
	},
	"Asia/Tokyo": {
		tzIdentifier:       "Asia/Tokyo",
		defaultCountryCode: "JP",
		otherCountryCodes:  []string{},
		utcOffset:          9,
	},
	"Asia/Shanghai": {
		tzIdentifier:       "Asia/Shanghai",
		defaultCountryCode: "CN",
		otherCountryCodes:  []string{},
		utcOffset:          8,
	},
	"Asia/Kolkata": {
		tzIdentifier:       "Asia/Kolkata",
		defaultCountryCode: "IN",
		otherCountryCodes:  []string{},
		utcOffset:          5.5,
	},
	"Asia/Seoul": {
		tzIdentifier:       "Asia/Seoul",
		defaultCountryCode: "KR",
		otherCountryCodes:  []string{},
		utcOffset:          9,
	},
	"Asia/Jakarta": {
		tzIdentifier:       "Asia/Jakarta",
		defaultCountryCode: "ID",
		otherCountryCodes:  []string{},
		utcOffset:          7,
	},
	"Asia/Tehran": {
		tzIdentifier:       "Asia/Tehran",
		defaultCountryCode: "IR",
		otherCountryCodes:  []string{},
		utcOffset:          3.5,
	},
	"Asia/Taipei": {
		tzIdentifier:       "Asia/Taipei",
		defaultCountryCode: "TW",
		otherCountryCodes:  []string{},
		utcOffset:          8,
	},
	"Asia/Hong_Kong": {
		tzIdentifier:       "Asia/Hong_Kong",
		defaultCountryCode: "HK",
		otherCountryCodes:  []string{},
		utcOffset:          8,
	},
	"Asia/Bangkok": {
		tzIdentifier:       "Asia/Bangkok",
		defaultCountryCode: "TH",
		otherCountryCodes:  []string{},
		utcOffset:          7,
	},
	"Asia/Singapore": {
		tzIdentifier:       "Asia/Singapore",
		defaultCountryCode: "SG",
		otherCountryCodes:  []string{},
		utcOffset:          8,
	},
	"Asia/Kuala_Lumpur": {
		tzIdentifier:       "Asia/Kuala_Lumpur",
		defaultCountryCode: "MY",
		otherCountryCodes:  []string{},
		utcOffset:          8,
	},
	"Australia/Sydney": {
		tzIdentifier:       "Australia/Sydney",
		defaultCountryCode: "AU",
		otherCountryCodes:  []string{},
		utcOffset:          11,
	},
	"Australia/Perth": {
		tzIdentifier:       "Australia/Perth",
		defaultCountryCode: "AU",
		otherCountryCodes:  []string{},
		utcOffset:          8,
	},
	"Australia/Brisbane": {
		tzIdentifier:       "Australia/Brisbane",
		defaultCountryCode: "AU",
		otherCountryCodes:  []string{},
		utcOffset:          10,
	},
	"Australia/Adelaide": {
		tzIdentifier:       "Australia/Adelaide",
		defaultCountryCode: "AU",
		otherCountryCodes:  []string{},
		utcOffset:          10.5,
	},
	"Australia/Darwin": {
		tzIdentifier:       "Australia/Darwin",
		defaultCountryCode: "AU",
		otherCountryCodes:  []string{},
		utcOffset:          9.5,
	},
	"Australia/Melbourne": {
		tzIdentifier:       "Australia/Melbourne",
		defaultCountryCode: "AU",
		otherCountryCodes:  []string{},
		utcOffset:          11,
	},
	"Australia/Hobart": {
		tzIdentifier:       "Australia/Hobart",
		defaultCountryCode: "AU",
		otherCountryCodes:  []string{},
		utcOffset:          11,
	},
	"Australia/Canberra": {
		tzIdentifier:       "Australia/Canberra",
		defaultCountryCode: "AU",
		otherCountryCodes:  []string{},
		utcOffset:          11,
	},
}
