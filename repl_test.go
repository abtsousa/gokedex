package main

import "testing"

type Cases struct {
	input    string
	expected []string
}

func TestCleanInput(t *testing.T) {
	cases := []Cases{
		{"  hello   world   ", []string{"hello", "world"}},
		{"hello world", []string{"hello", "world"}},
		{"Charmander Bulbasaur PIKACHU", []string{"charmander", "bulbasaur", "pikachu"}},
	}

	for no, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("[%v] mismatched array length: %v /// expected %v", no, actual, c.expected)
			return
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("[%v] mismatched word: %v /// expected %v", no, actual, c.expected)
				return
			}
		}
	}
}
