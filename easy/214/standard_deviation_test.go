package main

import "testing"

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
