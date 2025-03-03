package formating

import (
	"testing"
	"time"
)

func TestFormatDecimal(t *testing.T) {
	got := FormatDecimal("Decimal", 123)
	want := "Decimal: 123"
	if got != want {
		t.Errorf("FormatDecimal() = %v, want %v", got, want)
	}
}

func TestFormatBinary(t *testing.T) {
	got := FormatBinary("Binary", 5)
	want := "Binary: 101"
	if got != want {
		t.Errorf("FormatBinary() = %v, want %v", got, want)
	}
}

func TestFormatHexadecimal(t *testing.T) {
	got := FormatHexadecimal("Hexadecimal", 255)
	want := "Hexadecimal: ff"
	if got != want {
		t.Errorf("FormatHexadecimal() = %v, want %v", got, want)
	}
}

func TestFormatOctal(t *testing.T) {
	got := FormatOctal("Octal", 8)
	want := "Octal: 10"
	if got != want {
		t.Errorf("FormatOctal() = %v, want %v", got, want)
	}
}

func TestFormatFloat(t *testing.T) {
	got := FormatFloat("Float", 3.14)
	want := "Float: 3.140000"
	if got != want {
		t.Errorf("FormatFloat() = %v, want %v", got, want)
	}
}

func TestFormatScientific(t *testing.T) {
	got := FormatScientific("Scientific", 123456.789)
	want := "Scientific: 1.234568e+05"
	if got != want {
		t.Errorf("FormatScientific() = %v, want %v", got, want)
	}
}

func TestFormatCharacter(t *testing.T) {
	got := FormatCharacter("Character", 'A')
	want := "Character: A"
	if got != want {
		t.Errorf("FormatCharacter() = %v, want %v", got, want)
	}
}

func TestFormatText(t *testing.T) {
	got := FormatText("Text", "Hello")
	want := "Text: Hello"
	if got != want {
		t.Errorf("FormatText() = %v, want %v", got, want)
	}
}

func TestFormatLargeExponent(t *testing.T) {
	got := FormatLargeExponent("LargeExponent", 123456.789)
	want := "LargeExponent: 123456.789"
	if got != want {
		t.Errorf("FormatLargeExponent() = %v, want %v", got, want)
	}
}

func TestFormatPointerAddress(t *testing.T) {
	num := 123
	got := FormatPointerAddress("Pointer", &num)
	want := "Pointer: 0x"
	if got[:len(want)] != want {
		t.Errorf("FormatPointerAddress() = %v, want prefix %v", got, want)
	}
}

func TestFormatBoolean(t *testing.T) {
	got := FormatBoolean("Boolean", true)
	want := "Boolean: true"
	if got != want {
		t.Errorf("FormatBoolean() = %v, want %v", got, want)
	}
}

func TestFormatStruct(t *testing.T) {
	type Sample struct {
		Field1 string
		Field2 int
	}
	sample := Sample{"test", 123}
	got := FormatStruct("Struct", sample)
	want := "Struct: {Field1:test Field2:123}"
	if got != want {
		t.Errorf("FormatStruct() = %v, want %v", got, want)
	}
}

func TestFormatType(t *testing.T) {
	got := FormatType("Type", 123)
	want := "Type: int"
	if got != want {
		t.Errorf("FormatType() = %v, want %v", got, want)
	}
}

func TestFormatQuoted(t *testing.T) {
	got := FormatQuoted("Quoted", "Hello")
	want := `Quoted: "Hello"`
	if got != want {
		t.Errorf("FormatQuoted() = %v, want %v", got, want)
	}
}

func TestFormatUnicode(t *testing.T) {
	got := FormatUnicode("Unicode", 'A')
	want := "Unicode: U+0041"
	if got != want {
		t.Errorf("FormatUnicode() = %v, want %v", got, want)
	}
}

func TestFormatValue(t *testing.T) {
	got := FormatValue("Value", 123)
	want := "Value: 123"
	if got != want {
		t.Errorf("FormatValue() = %v, want %v", got, want)
	}
}

func TestFormatPercentage(t *testing.T) {
	got := FormatPercentage("Percentage", 0.123)
	want := "Percentage: 0.123000%"
	if got != want {
		t.Errorf("FormatPercentage() = %v, want %v", got, want)
	}
}

func TestFormatTimeStampRFC3339(t *testing.T) {
	timestamp := time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC)
	got := FormatTimeStampRFC3339("RFC3339", timestamp)
	want := "RFC3339: 2023-10-01T12:00:00Z"
	if got != want {
		t.Errorf("FormatTimeStampRFC3339() = %v, want %v", got, want)
	}
}

func TestFormatTimeStampUTC(t *testing.T) {
	timestamp := time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC)
	got := FormatTimeStampUTC("UTC", timestamp)
	want := "UTC: 01 Oct 23 12:00 +0000"
	if got != want {
		t.Errorf("FormatTimeStampUTC() = %v, want %v", got, want)
	}
}

func TestFormatTimeStampRFC822(t *testing.T) {
	timestamp := time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC)
	got := FormatTimeStampRFC822("RFC822", timestamp)
	want := "RFC822: 01 Oct 23 12:00 UTC"
	if got != want {
		t.Errorf("FormatTimeStampRFC822() = %v, want %v", got, want)
	}
}

func TestFormatWithPrecision(t *testing.T) {
	got := FormatWithPrecision("Precision", 3.14159, 2)
	want := "Precision: 3.14"
	if got != want {
		t.Errorf("FormatWithPrecision() = %v, want %v", got, want)
	}
}
