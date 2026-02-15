# tz

[![CI](https://github.com/infobits-io/tz/actions/workflows/ci.yml/badge.svg)](https://github.com/infobits-io/tz/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/infobits-io/tz.svg)](https://pkg.go.dev/github.com/infobits-io/tz)

A lightweight Go package that maps IANA timezone identifiers to country codes and UTC offsets. Zero dependencies, 400+ timezones.

## Installation

```bash
go get github.com/infobits-io/tz
```

## Usage

```go
package main

import (
    "fmt"
    "log"

    "github.com/infobits-io/tz"
)

func main() {
    timezone, err := tz.Decode("Europe/Berlin")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(timezone.Identifier())  // Europe/Berlin
    fmt.Println(timezone.CountryCode()) // DE
    fmt.Println(timezone.UtcOffset())   // 1
}
```

### Reverse lookups

```go
// Find all timezones in a country.
zones := tz.ByCountryCode("US")
for _, z := range zones {
    fmt.Printf("%s (UTC%+g)\n", z.Identifier(), z.UtcOffset())
}

// Find all timezones at a given UTC offset.
zones = tz.ByUtcOffset(5.5)
for _, z := range zones {
    fmt.Printf("%s (%s)\n", z.Identifier(), z.CountryCode())
}
```

### Validation and enumeration

```go
// Check if a timezone identifier is recognized.
if tz.IsValid("Asia/Tokyo") {
    fmt.Println("valid")
}

// Get all supported timezone identifiers (sorted).
all := tz.All()
fmt.Printf("%d timezones supported\n", len(all))

// Detect the system's current timezone.
current, err := tz.Current()
if err == nil {
    fmt.Printf("System timezone: %s\n", current.Identifier())
}
```

## API

### `Decode(identifier string) (Timezone, error)`

Looks up a timezone by its IANA identifier and returns a `Timezone` value, or an error wrapping `ErrNotFound` if the identifier is not recognized.

### `IsValid(identifier string) bool`

Reports whether the given identifier is a recognized timezone.

### `All() []string`

Returns a sorted slice of all supported IANA timezone identifiers.

### `ByCountryCode(code string) []Timezone`

Returns all timezones for the given ISO 3166-1 alpha-2 country code, sorted by identifier.

### `ByUtcOffset(offset float32) []Timezone`

Returns all timezones with the given standard UTC offset in hours, sorted by identifier.

### `Current() (Timezone, error)`

Returns the timezone for the system's current location.

### `Timezone` methods

| Method | Return type | Description |
|---|---|---|
| `Identifier()` | `string` | IANA timezone identifier |
| `CountryCode()` | `string` | ISO 3166-1 alpha-2 country code |
| `CountryCodes()` | `[]string` | All associated country codes |
| `UtcOffset()` | `float32` | Standard UTC offset in hours |

### Sentinel error

```go
var ErrNotFound = errors.New("timezone not found")
```

Use `errors.Is(err, tz.ErrNotFound)` to check for unknown timezone identifiers.

> **Note:** UTC offsets represent standard time only. Daylight saving time adjustments are not reflected.

## Supported Timezones

Covers 400+ IANA timezones across all regions: Africa, Americas, Antarctica, Asia, Atlantic, Australia, Europe, Indian Ocean, and Pacific. See [`data.go`](data.go) for the full list.

## License

[BSD 3-Clause](LICENSE)
