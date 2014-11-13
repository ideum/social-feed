package main

import "testing"

var truncationTestCases = []struct{ test, expected string }{
	{
		"This is short.",
		"This is short.",
	},
	{
		"This is a long phrase. It has a lot of words.",
		"This is a long...",
	},
}

func TestTruncatingText(t *testing.T) {
	for _, testCase := range truncationTestCases {
		actual := truncate(testCase.test, 15)
		if testCase.expected != actual {
			t.Errorf("\nExpected: \"%s\"\nActual:   \"%s\"", testCase.expected, actual)
		}
	}
}
