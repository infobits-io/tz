package tz

import (
	"errors"
	"fmt"
	"sort"
	"sync"
	"testing"
)

func TestDecode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		identifier string
		wantCode   string
		wantOffset float32
		wantErr    bool
	}{
		// Positive offsets.
		{name: "Europe/Berlin CET", identifier: "Europe/Berlin", wantCode: "DE", wantOffset: 1},
		{name: "Asia/Tokyo JST", identifier: "Asia/Tokyo", wantCode: "JP", wantOffset: 9},
		{name: "Pacific/Auckland NZST", identifier: "Pacific/Auckland", wantCode: "NZ", wantOffset: 12},
		{name: "Pacific/Kiritimati UTC+14", identifier: "Pacific/Kiritimati", wantCode: "KI", wantOffset: 14},

		// Negative offsets.
		{name: "America/New_York EST", identifier: "America/New_York", wantCode: "US", wantOffset: -5},
		{name: "America/Los_Angeles PST", identifier: "America/Los_Angeles", wantCode: "US", wantOffset: -8},
		{name: "Pacific/Pago_Pago SST", identifier: "Pacific/Pago_Pago", wantCode: "AS", wantOffset: -11},

		// Half-hour offsets.
		{name: "Asia/Kolkata IST +5.5", identifier: "Asia/Kolkata", wantCode: "IN", wantOffset: 5.5},
		{name: "Asia/Tehran IRST +3.5", identifier: "Asia/Tehran", wantCode: "IR", wantOffset: 3.5},
		{name: "America/St_Johns NST -3.5", identifier: "America/St_Johns", wantCode: "CA", wantOffset: -3.5},
		{name: "Australia/Adelaide ACST +9.5", identifier: "Australia/Adelaide", wantCode: "AU", wantOffset: 9.5},
		{name: "Australia/Lord_Howe +10.5", identifier: "Australia/Lord_Howe", wantCode: "AU", wantOffset: 10.5},

		// Quarter-hour offsets.
		{name: "Asia/Kathmandu NPT +5.75", identifier: "Asia/Kathmandu", wantCode: "NP", wantOffset: 5.75},
		{name: "Australia/Eucla ACWST +8.75", identifier: "Australia/Eucla", wantCode: "AU", wantOffset: 8.75},
		{name: "Pacific/Chatham CHAST +12.75", identifier: "Pacific/Chatham", wantCode: "NZ", wantOffset: 12.75},
		{name: "Pacific/Marquesas -9.5", identifier: "Pacific/Marquesas", wantCode: "PF", wantOffset: -9.5},

		// Zero offset.
		{name: "Europe/London GMT", identifier: "Europe/London", wantCode: "GB", wantOffset: 0},
		{name: "Etc/UTC", identifier: "Etc/UTC", wantCode: "", wantOffset: 0},
		{name: "UTC alias", identifier: "UTC", wantCode: "", wantOffset: 0},

		// Three-part identifiers.
		{name: "America/Argentina/Buenos_Aires", identifier: "America/Argentina/Buenos_Aires", wantCode: "AR", wantOffset: -3},
		{name: "America/Indiana/Indianapolis", identifier: "America/Indiana/Indianapolis", wantCode: "US", wantOffset: -5},
		{name: "America/Kentucky/Louisville", identifier: "America/Kentucky/Louisville", wantCode: "US", wantOffset: -5},
		{name: "America/North_Dakota/Center", identifier: "America/North_Dakota/Center", wantCode: "US", wantOffset: -6},

		// Africa.
		{name: "Africa/Cairo EET", identifier: "Africa/Cairo", wantCode: "EG", wantOffset: 2},
		{name: "Africa/Lagos WAT", identifier: "Africa/Lagos", wantCode: "NG", wantOffset: 1},
		{name: "Africa/Nairobi EAT", identifier: "Africa/Nairobi", wantCode: "KE", wantOffset: 3},

		// Indian Ocean.
		{name: "Indian/Maldives MVT", identifier: "Indian/Maldives", wantCode: "MV", wantOffset: 5},

		// Caribbean.
		{name: "America/Barbados AST", identifier: "America/Barbados", wantCode: "BB", wantOffset: -4},

		// Middle East.
		{name: "Asia/Dubai GST", identifier: "Asia/Dubai", wantCode: "AE", wantOffset: 4},
		{name: "Asia/Baghdad AST", identifier: "Asia/Baghdad", wantCode: "IQ", wantOffset: 3},

		// Error cases.
		{name: "empty string", identifier: "", wantErr: true},
		{name: "invalid timezone", identifier: "Invalid/Timezone", wantErr: true},
		{name: "wrong case", identifier: "europe/london", wantErr: true},
		{name: "partial match", identifier: "Europe", wantErr: true},
		{name: "trailing slash", identifier: "Europe/London/", wantErr: true},
		{name: "whitespace", identifier: " Europe/London ", wantErr: true},
		{name: "just a slash", identifier: "/", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tz, err := Decode(tt.identifier)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}

				if !errors.Is(err, ErrNotFound) {
					t.Errorf("expected errors.Is(err, ErrNotFound), got: %v", err)
				}

				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if tz.Identifier() != tt.identifier {
				t.Errorf("Identifier() = %q, want %q", tz.Identifier(), tt.identifier)
			}

			if tz.CountryCode() != tt.wantCode {
				t.Errorf("CountryCode() = %q, want %q", tz.CountryCode(), tt.wantCode)
			}

			if tz.UtcOffset() != tt.wantOffset {
				t.Errorf("UtcOffset() = %v, want %v", tz.UtcOffset(), tt.wantOffset)
			}
		})
	}
}

func TestDecodeCountryCodes(t *testing.T) {
	t.Parallel()

	tz, err := Decode("Europe/Berlin")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	codes := tz.CountryCodes()

	if len(codes) != 1 || codes[0] != "DE" {
		t.Errorf("CountryCodes() = %v, want [DE]", codes)
	}
}

func TestDecodeTzIdentifierDeprecated(t *testing.T) {
	t.Parallel()

	tz, err := Decode("Europe/London")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if tz.TzIdentifier() != "Europe/London" {
		t.Errorf("TzIdentifier() = %q, want %q", tz.TzIdentifier(), "Europe/London")
	}

	if tz.TzIdentifier() != tz.Identifier() {
		t.Error("TzIdentifier() and Identifier() should return the same value")
	}
}

func TestDataIntegrity(t *testing.T) {
	t.Parallel()

	for id, data := range timezones {
		// Country code must be empty (Etc/UTC) or exactly 2 uppercase ASCII letters.
		cc := data.countryCode
		if cc != "" {
			if len(cc) != 2 {
				t.Errorf("%s: country code %q is not 2 characters", id, cc)
			}

			for _, c := range cc {
				if c < 'A' || c > 'Z' {
					t.Errorf("%s: country code %q contains non-uppercase ASCII letter", id, cc)
				}
			}
		}

		// UTC offset must be in [-12, +14].
		if data.utcOffset < -12 || data.utcOffset > 14 {
			t.Errorf("%s: UTC offset %v is out of range [-12, 14]", id, data.utcOffset)
		}

		// Offset must be a quarter-hour multiple (0, 0.25, 0.5, 0.75).
		remainder := data.utcOffset - float32(int(data.utcOffset))
		if remainder < 0 {
			remainder = -remainder
		}

		validFractions := remainder == 0 || remainder == 0.25 || remainder == 0.5 || remainder == 0.75
		if !validFractions {
			t.Errorf("%s: UTC offset %v is not a quarter-hour multiple", id, data.utcOffset)
		}

		// Decode round-trip must succeed.
		tz, err := Decode(id)
		if err != nil {
			t.Errorf("%s: Decode() returned unexpected error: %v", id, err)

			continue
		}

		if tz.Identifier() != id {
			t.Errorf("%s: Decode() Identifier() = %q", id, tz.Identifier())
		}

		if tz.CountryCode() != data.countryCode {
			t.Errorf("%s: Decode() CountryCode() = %q, want %q", id, tz.CountryCode(), data.countryCode)
		}

		if tz.UtcOffset() != data.utcOffset {
			t.Errorf("%s: Decode() UtcOffset() = %v, want %v", id, tz.UtcOffset(), data.utcOffset)
		}
	}
}

func TestIsValid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		identifier string
		want       bool
	}{
		{"Europe/London", true},
		{"America/New_York", true},
		{"Asia/Tokyo", true},
		{"Etc/UTC", true},
		{"UTC", true},
		{"Invalid/Timezone", false},
		{"", false},
		{"europe/london", false},
	}

	for _, tt := range tests {
		t.Run(tt.identifier, func(t *testing.T) {
			t.Parallel()

			if got := IsValid(tt.identifier); got != tt.want {
				t.Errorf("IsValid(%q) = %v, want %v", tt.identifier, got, tt.want)
			}
		})
	}
}

func TestAll(t *testing.T) {
	t.Parallel()

	all := All()

	// Must match the map length.
	if len(all) != len(timezones) {
		t.Errorf("All() returned %d entries, want %d", len(all), len(timezones))
	}

	// Must be sorted.
	if !sort.StringsAreSorted(all) {
		t.Error("All() result is not sorted")
	}

	// Must contain no duplicates.
	seen := make(map[string]bool, len(all))

	for _, id := range all {
		if seen[id] {
			t.Errorf("All() contains duplicate: %q", id)
		}

		seen[id] = true
	}
}

func TestByCountryCode(t *testing.T) {
	t.Parallel()

	// US should have multiple timezones.
	us := ByCountryCode("US")
	if len(us) == 0 {
		t.Fatal("ByCountryCode(\"US\") returned empty slice")
	}

	// Results must be sorted by identifier.
	for i := 1; i < len(us); i++ {
		if us[i].Identifier() <= us[i-1].Identifier() {
			t.Errorf("ByCountryCode(\"US\") not sorted: %q >= %q", us[i-1].Identifier(), us[i].Identifier())
		}
	}

	// All results must have country code US.
	for _, tz := range us {
		if tz.CountryCode() != "US" {
			t.Errorf("ByCountryCode(\"US\") contains timezone with country code %q", tz.CountryCode())
		}
	}

	// Nonexistent country code should return nil.
	if got := ByCountryCode("ZZ"); got != nil {
		t.Errorf("ByCountryCode(\"ZZ\") = %v, want nil", got)
	}

	// Single-timezone country.
	jp := ByCountryCode("JP")
	if len(jp) != 1 || jp[0].Identifier() != "Asia/Tokyo" {
		t.Errorf("ByCountryCode(\"JP\") = %v, want [Asia/Tokyo]", jp)
	}
}

func TestByUtcOffset(t *testing.T) {
	t.Parallel()

	// UTC+0 should include Europe/London.
	utcZero := ByUtcOffset(0)
	if len(utcZero) == 0 {
		t.Fatal("ByUtcOffset(0) returned empty slice")
	}

	found := false

	for _, tz := range utcZero {
		if tz.Identifier() == "Europe/London" {
			found = true

			break
		}
	}

	if !found {
		t.Error("ByUtcOffset(0) does not contain Europe/London")
	}

	// Results must be sorted by identifier.
	for i := 1; i < len(utcZero); i++ {
		if utcZero[i].Identifier() <= utcZero[i-1].Identifier() {
			t.Errorf("ByUtcOffset(0) not sorted: %q >= %q", utcZero[i-1].Identifier(), utcZero[i].Identifier())
		}
	}

	// Half-hour offset.
	halfHour := ByUtcOffset(5.5)
	if len(halfHour) == 0 {
		t.Fatal("ByUtcOffset(5.5) returned empty slice")
	}

	// Nonexistent offset should return nil.
	if got := ByUtcOffset(99); got != nil {
		t.Errorf("ByUtcOffset(99) = %v, want nil", got)
	}
}

func TestCurrent(t *testing.T) {
	t.Parallel()

	// Current should succeed on most systems with our expanded dataset.
	tz, err := Current()
	if err != nil {
		t.Skipf("system timezone not in dataset: %v", err)
	}

	if tz.Identifier() == "" {
		t.Error("Current() returned timezone with empty identifier")
	}
}

func TestErrNotFound(t *testing.T) {
	t.Parallel()

	_, err := Decode("Nonexistent/Zone")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !errors.Is(err, ErrNotFound) {
		t.Errorf("errors.Is(err, ErrNotFound) = false, err = %v", err)
	}

	// The error message should contain the identifier.
	if expected := `timezone "Nonexistent/Zone": timezone not found`; err.Error() != expected {
		t.Errorf("err.Error() = %q, want %q", err.Error(), expected)
	}
}

func TestConcurrentDecode(t *testing.T) {
	t.Parallel()

	identifiers := []string{
		"Europe/London", "America/New_York", "Asia/Tokyo",
		"Australia/Sydney", "Africa/Cairo", "Pacific/Auckland",
	}

	var wg sync.WaitGroup

	for i := range 100 {
		wg.Go(func() {
			id := identifiers[i%len(identifiers)]

			tz, err := Decode(id)
			if err != nil {
				t.Errorf("Decode(%q) error: %v", id, err)

				return
			}

			if tz.Identifier() != id {
				t.Errorf("Decode(%q) Identifier() = %q", id, tz.Identifier())
			}
		})
	}

	wg.Wait()
}

func TestConcurrentByCountryCode(t *testing.T) {
	t.Parallel()

	var wg sync.WaitGroup

	for range 100 {
		wg.Go(func() {
			result := ByCountryCode("US")
			if len(result) == 0 {
				t.Error("ByCountryCode(\"US\") returned empty in concurrent access")
			}
		})
	}

	wg.Wait()
}

func TestConcurrentByUtcOffset(t *testing.T) {
	t.Parallel()

	var wg sync.WaitGroup

	for range 100 {
		wg.Go(func() {
			result := ByUtcOffset(1)
			if len(result) == 0 {
				t.Error("ByUtcOffset(1) returned empty in concurrent access")
			}
		})
	}

	wg.Wait()
}

// Benchmarks.

func BenchmarkDecode(b *testing.B) {
	for b.Loop() {
		_, _ = Decode("America/New_York")
	}
}

func BenchmarkDecodeNotFound(b *testing.B) {
	for b.Loop() {
		_, _ = Decode("Invalid/Timezone")
	}
}

func BenchmarkAll(b *testing.B) {
	for b.Loop() {
		_ = All()
	}
}

func BenchmarkByCountryCode(b *testing.B) {
	// Warm up the index.
	_ = ByCountryCode("US")

	b.ResetTimer()

	for b.Loop() {
		_ = ByCountryCode("US")
	}
}

func BenchmarkByUtcOffset(b *testing.B) {
	// Warm up the index.
	_ = ByUtcOffset(1)

	b.ResetTimer()

	for b.Loop() {
		_ = ByUtcOffset(1)
	}
}

func BenchmarkIsValid(b *testing.B) {
	for b.Loop() {
		_ = IsValid("Europe/Berlin")
	}
}

// Examples.

func ExampleDecode() {
	tz, err := Decode("Europe/Berlin")
	if err != nil {
		panic(err)
	}

	fmt.Println(tz.Identifier())
	fmt.Println(tz.CountryCode())
	fmt.Println(tz.UtcOffset())
	// Output:
	// Europe/Berlin
	// DE
	// 1
}

func ExampleDecode_notFound() {
	_, err := Decode("Invalid/Timezone")
	fmt.Println(err)
	// Output:
	// timezone "Invalid/Timezone": timezone not found
}

func ExampleIsValid() {
	fmt.Println(IsValid("Europe/London"))
	fmt.Println(IsValid("Invalid/Zone"))
	// Output:
	// true
	// false
}

func ExampleAll() {
	all := All()
	fmt.Printf("Total timezones: %d\n", len(all))
	fmt.Printf("First: %s\n", all[0])
	// Output:
	// Total timezones: 418
	// First: Africa/Abidjan
}

func ExampleByCountryCode() {
	zones := ByCountryCode("JP")
	for _, tz := range zones {
		fmt.Printf("%s (UTC%+g)\n", tz.Identifier(), tz.UtcOffset())
	}
	// Output:
	// Asia/Tokyo (UTC+9)
}

func ExampleCurrent() {
	tz, err := Current()
	if err != nil {
		fmt.Println("system timezone not recognized")

		return
	}

	fmt.Printf("Current timezone: %s\n", tz.Identifier())
}
