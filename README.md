# tz

[![CI](https://github.com/infobits-io/tz/actions/workflows/ci.yml/badge.svg)](https://github.com/infobits-io/tz/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/infobits-io/tz.svg)](https://pkg.go.dev/github.com/infobits-io/tz)

A lightweight Go package that maps IANA timezone identifiers to country codes and UTC offsets. Zero dependencies.

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

    fmt.Println(timezone.TzIdentifier()) // Europe/Berlin
    fmt.Println(timezone.CountryCode())   // DE
    fmt.Println(timezone.CountryCodes())  // [DE]
    fmt.Println(timezone.UtcOffset())     // 1
}
```

## API

### `Decode(timezoneString string) (*Timezone, error)`

Looks up a timezone by its IANA identifier and returns a `Timezone` struct, or an error if the identifier is not found.

### `Timezone` methods

| Method | Return type | Description |
|---|---|---|
| `TzIdentifier()` | `string` | IANA timezone identifier |
| `CountryCode()` | `string` | Default ISO 3166-1 alpha-2 country code |
| `CountryCodes()` | `[]string` | All associated country codes |
| `UtcOffset()` | `float32` | UTC offset in hours |

## Supported Timezones

Covers major timezones across Europe, the Americas, Asia, and Australia. See [`tz.go`](tz.go) for the full list.

## License

[BSD 3-Clause](LICENSE)
