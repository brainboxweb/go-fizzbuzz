package main

import "testing"

var tests = []struct {
	input    int
	expected string
}{
	{1, ""},
	{2, ""},
	{3, "Fizz"},
	{4, ""},
	{5, "Buzz"},
	{6, "Fizz"},
	{7, ""},
	{8, ""},
	{9, "Fizz"},
	{10, "Buzz"},
	{11, ""},
	{12, "Fizz"},
	{13, ""},
	{14, ""},
	{15, "FizzBuzz"},
	{16, ""},
	{17, ""},
	{18, "Fizz"},
	{19, ""},
	{20, "Buzz"},
	{21, "Fizz"},
	{22, ""},
	{23, ""},
	{24, "Fizz"},
	{25, "Buzz"},
	{26, ""},
	{27, "Fizz"},
	{28, ""},
	{29, ""},
	{30, "FizzBuzz"},
	{31, ""},
}

func TestFizzBuzz(t *testing.T) {
	for _, test := range tests {
		if actual := FizzBuzz(test.input); actual != test.expected {
			t.Errorf("Convert(%d) = %q, expected %q.",
				test.input, actual, test.expected)
		}
	}
}
