package main

import "testing"

func TestUnpack(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"ok", "a4bc2d5e", "aaaabccddddde"},
		{"withoutChanges", "abcd", "abcd"},
		{"incorrect", "45", ""},
		{"empty", "", ""},
		{"intAsString", `qwe\4\5`, "qwe45"},
		{"intAsStringNTimes", `qwe\45`, "qwe44444"},
		{"backslashAsString", `qwe\\5`, `qwe\\\\\`},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := unpack(testCase.input)

			if res != testCase.expected {
				t.Errorf("got %s want %s", res, testCase.expected)
				t.Error(err)
			}
		})
	}
}
