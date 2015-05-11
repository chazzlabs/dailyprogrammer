package main

import (
    "testing"
)

func sliceEqual(a, b[]float64) bool {
    if len(a) != len(b) {
        return false
    }

    for index, value := range a {
        if value != b[index] {
            return false
        }
    }

    return true
}

func TestCalculateMean(t *testing.T) {
    testCases := []struct {
		inputValues []float64
        expectedMean float64
	} {
		{[]float64{1, 2, 6}, 3},
        {[]float64{5, 10, 15}, 10},
        {[]float64{1, 2}, 1.5},
	}
	for _, testCase := range testCases {
		result := calculateMean(testCase.inputValues)
		if result != testCase.expectedMean {
			t.Errorf("calculateMean(%q) == %f, want %f", testCase.inputValues, result, testCase.expectedMean)
		}
	}
}

func TestCalculateVariance(t *testing.T) {
    testCases := []struct {
		inputValues []float64
        mean float64
        expectedVariance float64
	} {
		{[]float64{5, 6, 11, 13, 19, 20, 25, 26, 28, 37}, 19.0, 95.6},
        {[]float64{1, 2}, 1.5, 0.25},
	}
	for _, testCase := range testCases {
		result := calculateVariance(testCase.inputValues, testCase.mean)
		if result != testCase.expectedVariance {
			t.Errorf("calculateVariance(%f) == %f, want %f", testCase.inputValues, result, testCase.expectedVariance)
		}
	}
}

func TestCalculateStdDevidation(t *testing.T) {
    testCases := []struct {
		inputValues []float64
        expectedStdDeviation float64
	} {
		{[]float64{5, 6, 11, 13, 19, 20, 25, 26, 28, 37}, 9.777525249264253},
	}
	for _, testCase := range testCases {
		result := calculateStdDeviation(testCase.inputValues)
		if result != testCase.expectedStdDeviation {
			t.Errorf("calculateStdDeviation(%q) == %f, want %f", testCase.inputValues, result, testCase.expectedStdDeviation)
		}
	}
}

func TestRound(t *testing.T) {
    testCases := []struct {
		inputValue float64
        places int
        expectedRoundedValue float64
	} {
        {10.12, 1, 10.1},
        {10.12345, 4, 10.1235},
		{10.0, 1, 10.0},
	}
	for _, testCase := range testCases {
		result := round(testCase.inputValue, testCase.places)
		if result != testCase.expectedRoundedValue {
			t.Errorf("round(%f, %d) == %f, want %f", testCase.inputValue, testCase.places, result, testCase.expectedRoundedValue)
		}
	}
}

func TestConvertToFloatArray(t *testing.T) {
    testCases := []struct {
		inputArray []string
        expectedArray []float64
	} {
        {[]string{"10", "11", "12"}, []float64{10.000000, 11.000000, 12.000000}},

	}
	for _, testCase := range testCases {
		result := convertToFloatArray(testCase.inputArray)
		if !sliceEqual(result, testCase.expectedArray) {
			t.Errorf("convertToFloatArray(%q) == %f, want %f", testCase.inputArray, result, testCase.expectedArray)
		}
	}
}
