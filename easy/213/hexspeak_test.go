package main

import "testing"

func TestPronounce(t *testing.T) {
    createDictionary()
    
	testCases := []struct {
		inputString, expectedString string
	} {
		{"0xF5", "fleventy-five"},
		{"0xB3", "bibbity-three"},
        {"0xE4", "ebbity-four"},
        {"0xBBBB", "bibbity-bee bitey bibbity-bee"},
        {"0xA0C9", "atta bitey city-nine"},
	}
	for _, testCase := range testCases {
		result := pronounce(testCase.inputString)
		if result != testCase.expectedString {
			t.Errorf("pronounce(%q) == %q, want %q", testCase.inputString, result, testCase.expectedString)
		}
	}
}
