package main

import (
	"testing"
)

func TestTranslate(t *testing.T) {
	type test struct {
		input string
		want  string
	}

	tests := []test{
		{
			input: "one two three four",
			want:  "oneway otway eethray ourfay",
		},
		{
			input: "Hello World?! I'm a gopher 4 5yrs. My email: joe@g.com #hashtag",
			want:  "elloHay orldWay?! I'may ay ophergay 4 5yrs. Myay emailway: joe@g.com #hashtag",
		},
	}

	for _, tc := range tests {
		got := translate(tc.input)
		if got != tc.want {
			t.Errorf("input: %v, want: %v, got: %v", tc.input, tc.want, got)
		}
	}
}

func TestPigLatinize(t *testing.T) {
	type test struct {
		input string
		want  string
	}

	tests := []test{
		{
			input: "too",
			want:  "ootay",
		},
		{
			input: "simple",
			want:  "implesay",
		},
		{
			input: "floor",
			want:  "oorflay",
		},
		{
			input: "pry",
			want:  "pryay",
		},
		{
			input: "egg",
			want:  "eggway",
		},
		{
			input: "always",
			want:  "alwaysway",
		},
		{
			input: "aeiou",
			want:  "aeiouway",
		},
		{
			input: "flu",
			want:  "uflay",
		},
		{
			input: "12yrs",
			want:  "12yrs",
		},
		{
			input: "front$-end",
			want:  "front$-end",
		},
		{
			input: "front-end",
			want:  "ont-endfray",
		},
		{
			input: "front's end",
			want:  "front's end",
		},
	}

	for _, tc := range tests {
		got := pigLatinize(tc.input)
		if got != tc.want {
			t.Errorf("input: %v, want: %v, got: %v", tc.input, tc.want, got)
		}
	}
}
