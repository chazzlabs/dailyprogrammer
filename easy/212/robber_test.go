package main

import "testing"

func TestEncode(t *testing.T) {
	testCases := []struct {
		inputString, expectedString string
	} {
		{"Jag talar Rövarspråket!", "Jojagog totalolaror Rorövovarorsospoproråkoketot!"},
		{"I'm speaking Robber's language!", "I'mom sospopeakokinongog Rorobobboberor'sos lolanongoguagoge!"},
		{"", ""},
	}
	for _, testCase := range testCases {
		result := encode(testCase.inputString)
		if result != testCase.expectedString {
			t.Errorf("encode(%q) == %q, want %q", testCase.inputString, result, testCase.expectedString)
		}
	}
}

func TestDecode(t *testing.T) {
	testCases := []struct {
		inputString, expectedString string
	} {
        {"Jojagog totalolaror Rorövovarorsospoproråkoketot!", "Jag talar Rövarspråket!"},
        {"I'mom sospopeakokinongog Rorobobboberor'sos lolanongoguagoge!", "I'm speaking Robber's language!"},
		{"", ""},
	}
	for _, testCase := range testCases {
		result := decode(testCase.inputString)
		if result != testCase.expectedString {
			t.Errorf("encode(%q) == %q, want %q", testCase.inputString, result, testCase.expectedString)
		}
	}
}
