package internal_test

import (
	"testing"

	"github.com/gi4nks/cesar/internal"
)

// TestNewConverter_ValidRange tests the NewConverter function with a valid range.
// It checks if NewConverter returns nil error for a valid range.
func TestNewConverter_ValidRange(t *testing.T) {
	start, end := 1, 10
	_, err := internal.NewConverter(start, end)

	if err != nil {
		t.Errorf("NewConverter() returned unexpected error: %v", err)
	}
}

// TestNewConverter_InvalidStart tests the NewConverter function with an invalid start value.
// It checks if NewConverter returns the expected error when the start value is invalid.
func TestNewConverter_InvalidStart(t *testing.T) {
	start, end := 0, 10
	expectedError := "start and end values must be between 1 and 3999"

	cnv, err := internal.NewConverter(start, end)

	if err.Error() != expectedError {
		t.Errorf("NewConverter() error = %v, expectedError %v", err.Error(), expectedError)
	}

	if cnv != nil {
		t.Errorf("NewConverter() cnv %v, expected nil", cnv)
	}
}

// TestNewConverter_InvalidEnd tests the NewConverter function with an invalid end value.
// It checks if NewConverter returns the expected error when the end value is invalid.
func TestNewConverter_InvalidEnd(t *testing.T) {
	start, end := 1, 4000
	expectedError := "start and end values must be between 1 and 3999"

	cnv, err := internal.NewConverter(start, end)

	if err.Error() != expectedError {
		t.Errorf("NewConverter() error = %v, expectedError %v", err.Error(), expectedError)
	}

	if cnv != nil {
		t.Errorf("NewConverter() cnv %v, expected nil", cnv)
	}
}

// TestNewConverter_StartGreaterThanEnd tests the NewConverter function
// when the start value is greater than the end value.
// It checks if NewConverter returns the expected error.
func TestNewConverter_StartGreaterThanEnd(t *testing.T) {
	start, end := 10, 1
	expectedError := "start value must be smaller than or equal to end value"

	cnv, err := internal.NewConverter(start, end)

	if err.Error() != expectedError {
		t.Errorf("NewConverter() error = %v, expectedError %v", err.Error(), expectedError)
	}

	if cnv != nil {
		t.Errorf("NewConverter() cnv %v, expected nil", cnv)
	}
}

// TestConverter_Convert_1 tests the Convert method with a single value range.
// It checks if Convert returns the correct Roman numeral for the input value 1.
func TestConverter_Convert_1(t *testing.T) {
	converter, err := internal.NewConverter(1, 1)

	if err != nil {
		t.Errorf("NewConverter() returned unexpected error: %v", err)
	}

	expected := "I"

	cnv := converter.Convert()

	if cnv[1] != expected {
		t.Errorf("Convert() for 1 cnv %s; expected %s", cnv[1], expected)
	}
}

// TestConverter_Convert_5 tests the Convert method with a single value range.
// It checks if Convert returns the correct Roman numeral for the input value 5.
func TestConverter_Convert_5(t *testing.T) {
	converter, err := internal.NewConverter(5, 5)

	if err != nil {
		t.Errorf("NewConverter() returned unexpected error: %v", err)
	}

	expected := "V"

	cnv := converter.Convert()

	if cnv[5] != expected {
		t.Errorf("Convert() for 5 cnv %s; expected %s", cnv[5], expected)
	}
}

// TestConverter_Convert_10 tests the Convert method with a single value range.
// It checks if Convert returns the correct Roman numeral for the input value 10.
func TestConverter_Convert_10(t *testing.T) {
	converter, err := internal.NewConverter(10, 10)

	if err != nil {
		t.Errorf("NewConverter() returned unexpected error: %v", err)
	}

	expected := "X"

	cnv := converter.Convert()

	if cnv[10] != expected {
		t.Errorf("Convert() for 10 cnv %s; expected %s", cnv[10], expected)
	}
}

// TestConverter_Convert_3999 tests the Convert method with a single value range.
// It checks if Convert returns the correct Roman numeral for the input value 3999.
func TestConverter_Convert_3999(t *testing.T) {
	converter, err := internal.NewConverter(3999, 3999)

	if err != nil {
		t.Errorf("NewConverter() returned unexpected error: %v", err)
	}

	expected := "MMMCMXCIX"

	cnv := converter.Convert()

	if cnv[3999] != expected {
		t.Errorf("Convert() for 3999 cnv %s; expected %s", cnv[3999], expected)
	}
}

// TestConverter_Convert_1234 tests the Convert method with a single value range.
// It checks if Convert returns the correct Roman numeral for the input value 1234.
func TestConverter_Convert_1234(t *testing.T) {
	converter, err := internal.NewConverter(1234, 1234)

	if err != nil {
		t.Errorf("NewConverter() returned unexpected error: %v", err)
	}

	expected := "MCCXXXIV"

	cnv := converter.Convert()

	if cnv[1234] != expected {
		t.Errorf("Convert() for 1234 cnv %s; expected %s", cnv[1234], expected)
	}
}
