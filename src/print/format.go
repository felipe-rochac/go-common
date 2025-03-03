package formating

import (
	"fmt"
	"time"
)

const formatStrToStr string = "%s: %s"

// FormatDecimal returns the formatted string for a decimal number.
func FormatDecimal(message string, number int) string {
	return fmt.Sprintf("%s: %d", message, number)
}

// FormatBinary returns the formatted string for a binary number.
func FormatBinary(message string, number int) string {
	return fmt.Sprintf("%s: %b", message, number)
}

// FormatHexadecimal returns the formatted string for a hexadecimal number.
func FormatHexadecimal(message string, number int) string {
	return fmt.Sprintf("%s: %x", message, number)
}

// FormatOctal returns the formatted string for an octal number.
func FormatOctal(message string, number int) string {
	return fmt.Sprintf("%s: %o", message, number)
}

// FormatFloat returns the formatted string for a float number.
func FormatFloat(message string, number float64) string {
	return fmt.Sprintf("%s: %f", message, number)
}

// FormatScientific returns the formatted string for a scientific notation number.
func FormatScientific(message string, number float64) string {
	return fmt.Sprintf("%s: %e", message, number)
}

// FormatCharacter returns the formatted string for a character.
func FormatCharacter(message string, character rune) string {
	return fmt.Sprintf("%s: %c", message, character)
}

// FormatText returns the formatted string for a text.
func FormatText(message string, text string) string {
	return fmt.Sprintf(formatStrToStr, message, text)
}

// FormatLargeExponent returns the formatted string for a large exponent number.
func FormatLargeExponent(message string, number float64) string {
	return fmt.Sprintf("%s: %g", message, number)
}

// FormatPointerAddress returns the formatted string for a pointer address.
func FormatPointerAddress(message string, pointer *int) string {
	return fmt.Sprintf("%s: %p", message, pointer)
}

// FormatBoolean returns the formatted string for a boolean value.
func FormatBoolean(message string, value bool) string {
	return fmt.Sprintf("%s: %t", message, value)
}

// FormatStruct returns the formatted string for a struct value.
func FormatStruct(message string, value any) string {
	return fmt.Sprintf("%s: %+v", message, value)
}

// FormatType returns the formatted string for a type.
func FormatType(message string, value any) string {
	return fmt.Sprintf("%s: %T", message, value)
}

// FormatQuoted returns the formatted string for a quoted value.
func FormatQuoted(message string, value string) string {
	return fmt.Sprintf("%s: %q", message, value)
}

// FormatUnicode returns the formatted string for a unicode value.
func FormatUnicode(message string, value rune) string {
	return fmt.Sprintf("%s: %U", message, value)
}

// FormatValue returns the formatted string for a value.
func FormatValue(message string, value any) string {
	return fmt.Sprintf("%s: %#v", message, value)
}

// FormatPercentage returns the formatted string for a percentage value.
func FormatPercentage(message string, value float64) string {
	return fmt.Sprintf("%s: %f%%", message, value)
}

// FormatTimeStampRFC3339 returns the formatted string for a time in RFC3339 format.
func FormatTimeStampRFC3339(message string, value time.Time) string {
	return fmt.Sprintf(formatStrToStr, message, value.Format(time.RFC3339))
}

// FormatTimeStampUTC returns the formatted string for a time in UTC format.
func FormatTimeStampUTC(message string, value time.Time) string {
	return fmt.Sprintf(formatStrToStr, message, value.Format(time.RFC822Z))
}

// FormatTimeStampRFC822 returns the formatted string for a time in RFC822 format.
func FormatTimeStampRFC822(message string, value time.Time) string {
	return fmt.Sprintf(formatStrToStr, message, value.Format(time.RFC822))
}

// FormatWithPrecision returns the formatted string for a value with the given precision.
func FormatWithPrecision(message string, value float64, precision int) string {
	return fmt.Sprintf("%s: %.*f", message, precision, value)
}
