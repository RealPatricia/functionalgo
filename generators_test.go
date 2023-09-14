package functionalgo

import (
	"reflect"
	"testing"
)

func TestGenerator(t *testing.T) {
	source := func(in int) int {
		return in + 1
	}

	quitter := func(in int) bool {
		return in >= 10
	}

	nums := Generator(0, source, quitter)
	expected := Collection[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	end := FinalValue(StreamToCollection(nums))
	if !reflect.DeepEqual(end, expected) {
		t.Errorf("Generator did not return the correct result, expected: %v, got: %v", expected, end)
	}
}

func TestValidateRange(t *testing.T) {
	tests := []struct {
		name             string
		start, end, step float64
		valid            bool
	}{
		{"Go up step up", 1, 10, 1, true},
		{"Go up step down", 10, 1, 1, false},
		{"Go down step up", 1, 10, -1, false},
		{"Go down step down", 10, 1, -1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateRange(tt.start, tt.end, tt.step)
			if result != tt.valid {
				t.Errorf("ValidateRange produced an incorrect result [%v] for range [%f] to [%f] with step [%f]",
					result, tt.start, tt.end, tt.step)
			}
		})
	}
}

func TestGenerateRangeUpDown(t *testing.T) {
	_, err := GenerateRange(10, 1, 1)
	if err == nil {
		t.Errorf("GenerateRange did not produce an error for an invalid range")
	}
}

func TestGenerateRangeUp(t *testing.T) {
	oneToTen, err := GenerateRange(1, 10, 1)
	if err != nil {
		t.Errorf("Generate range produced an error: %v", err)
	}
	expected := Collection[int]{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := FinalValue(StreamToCollection(oneToTen))

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GenerateRange failed, expected: %v, got %v", expected, result)
	}
}

func TestGenerateRangeDown(t *testing.T) {
	oneToTen, err := GenerateRange(10, 1, -1)
	if err != nil {
		t.Errorf("Generate range produced an error: %v", err)
	}
	expected := Collection[int]{10, 9, 8, 7, 6, 5, 4, 3, 2}
	result := FinalValue(StreamToCollection(oneToTen))

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GenerateRange failed, expected: %v, got %v", expected, result)
	}
}
