package tz

import (
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"
)

// ErrNotFound is returned when a timezone identifier is not in the dataset.
var ErrNotFound = errors.New("timezone not found")

// Timezone holds the decoded information for an IANA timezone identifier.
// UTC offsets represent standard time only (DST adjustments are not reflected).
type Timezone struct {
	identifier  string
	countryCode string
	utcOffset   float32
}

// Decode looks up a timezone by its IANA identifier.
// Returns ErrNotFound (wrapped) if the identifier is not recognized.
func Decode(identifier string) (Timezone, error) {
	data, ok := timezones[identifier]
	if !ok {
		return Timezone{}, fmt.Errorf("timezone %q: %w", identifier, ErrNotFound)
	}

	return Timezone{
		identifier:  identifier,
		countryCode: data.countryCode,
		utcOffset:   data.utcOffset,
	}, nil
}

// Identifier returns the IANA timezone identifier.
func (t Timezone) Identifier() string {
	return t.identifier
}

// TzIdentifier returns the IANA timezone identifier.
//
// Deprecated: Use Identifier instead.
func (t Timezone) TzIdentifier() string {
	return t.identifier
}

// CountryCode returns the ISO 3166-1 alpha-2 country code.
func (t Timezone) CountryCode() string {
	return t.countryCode
}

// CountryCodes returns all associated country codes as a slice.
func (t Timezone) CountryCodes() []string {
	return []string{t.countryCode}
}

// UtcOffset returns the standard UTC offset in hours.
func (t Timezone) UtcOffset() float32 {
	return t.utcOffset
}

// IsValid reports whether the given identifier is a recognized timezone.
func IsValid(identifier string) bool {
	_, ok := timezones[identifier]

	return ok
}

// All returns a sorted slice of all supported IANA timezone identifiers.
func All() []string {
	result := make([]string, 0, len(timezones))

	for id := range timezones {
		result = append(result, id)
	}

	sort.Strings(result)

	return result
}

// Lazy-built reverse lookup indices.
var (
	countryIndex     map[string][]Timezone
	countryIndexOnce sync.Once

	offsetIndex     map[float32][]Timezone
	offsetIndexOnce sync.Once
)

func buildCountryIndex() {
	countryIndex = make(map[string][]Timezone, len(timezones))

	for id, data := range timezones {
		tz := Timezone{identifier: id, countryCode: data.countryCode, utcOffset: data.utcOffset}
		countryIndex[data.countryCode] = append(countryIndex[data.countryCode], tz)
	}

	// Sort each slice by identifier for deterministic output.
	for code := range countryIndex {
		slice := countryIndex[code]

		sort.Slice(slice, func(i, j int) bool {
			return slice[i].identifier < slice[j].identifier
		})
	}
}

func buildOffsetIndex() {
	offsetIndex = make(map[float32][]Timezone, len(timezones))

	for id, data := range timezones {
		tz := Timezone{identifier: id, countryCode: data.countryCode, utcOffset: data.utcOffset}
		offsetIndex[data.utcOffset] = append(offsetIndex[data.utcOffset], tz)
	}

	// Sort each slice by identifier for deterministic output.
	for offset := range offsetIndex {
		slice := offsetIndex[offset]

		sort.Slice(slice, func(i, j int) bool {
			return slice[i].identifier < slice[j].identifier
		})
	}
}

// ByCountryCode returns all timezones for the given ISO 3166-1 alpha-2 country code.
// Results are sorted by identifier. Returns nil if no timezones match.
func ByCountryCode(code string) []Timezone {
	countryIndexOnce.Do(buildCountryIndex)

	return countryIndex[code]
}

// ByUtcOffset returns all timezones with the given standard UTC offset in hours.
// Results are sorted by identifier. Returns nil if no timezones match.
func ByUtcOffset(offset float32) []Timezone {
	offsetIndexOnce.Do(buildOffsetIndex)

	return offsetIndex[offset]
}

// Current returns the timezone for the system's current location.
// Returns ErrNotFound if the system timezone is not in the dataset.
func Current() (Timezone, error) {
	return Decode(time.Now().Location().String())
}
