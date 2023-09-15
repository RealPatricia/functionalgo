package functionalgo

import (
	"github.com/RealPatricia/pth/assert"
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
	expected := MakeSlice(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	end := FinalValue(ChanToSlice(nums))
	if assert.Equal(end, expected, false) {
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
	_, err := Range(10, 1, 1)
	if err == nil {
		t.Errorf("GenerateRange did not produce an error for an invalid range")
	}
}

func TestGenerateRangeUp(t *testing.T) {
	oneToTen, err := Range(1, 10, 1)
	if err != nil {
		t.Errorf("Generate range produced an error: %v", err)
	}
	expected := MakeSlice(1, 2, 3, 4, 5, 6, 7, 8, 9)
	result := FinalValue(ChanToSlice(oneToTen))

	if assert.Equal(result, expected, false) {
		t.Errorf("GenerateRange failed, expected: %v, got %v", expected, result)
	}
}

func TestGenerateRangeDown(t *testing.T) {
	oneToTen, err := Range(10, 0, -1)
	if err != nil {
		t.Errorf("Generate range produced an error: %v", err)
	}
	expected := MakeSlice(10, 9, 8, 7, 6, 5, 4, 3, 2, 1)
	result := FinalValue(ChanToSlice(oneToTen))

	if assert.Equal(result, expected, false) {
		t.Errorf("GenerateRange failed, expected: %v, got %v", expected, result)
	}
}
