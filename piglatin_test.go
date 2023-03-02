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
			input: "one two three four o'clock hasn't fr-end 12$ 4friend",
			want:  "oneway otway eethray ourfay o'clockway asn'thay -endfray 12$ 4iendfray",
		},
		{
			input: "Hello World?! I'm a gopher 4 5yrs. My email: joe@g.com ,phone #hash",
			want:  "elloHay orldWay?! I'mway away ophergay 4 5yrsay. Myay emailway: oe@g.comjay ,onephay #ashhay",
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
			want:  "12yrsay",
		},
		{
			input: "4ever",
			want:  "4everay",
		},
		{
			input: "front$-end",
			want:  "ont$-endfray",
		},
		{
			input: "front-end",
			want:  "ont-endfray",
		},
		{
			input: "front's",
			want:  "ont'sfray",
		},
	}

	for _, tc := range tests {
		got := pigLatinize(tc.input)
		if got != tc.want {
			t.Errorf("input: %v, want: %v, got: %v", tc.input, tc.want, got)
		}
	}
}

func TestTrimNonLetters(t *testing.T) {
	type test struct {
		input string
		left  string
		mid   string
		right string
	}

	tests := []test{
		{
			input: "hello?",
			left:  "",
			mid:   "hello",
			right: "?",
		},
		{
			input: "hello?!!",
			left:  "",
			mid:   "hello",
			right: "?!!",
		},
		{
			input: "#twitter",
			left:  "#",
			mid:   "twitter",
			right: "",
		},
		{
			input: "[[100",
			left:  "[[100",
			mid:   "",
			right: "",
		},
		{
			input: "(tag)",
			left:  "(",
			mid:   "tag",
			right: ")",
		},
		{
			input: "{{tag}}}",
			left:  "{{",
			mid:   "tag",
			right: "}}}",
		},
		{
			input: "{{tag}}a}",
			left:  "{{",
			mid:   "tag}}a",
			right: "}",
		},
		{
			input: "none",
			left:  "",
			mid:   "none",
			right: "",
		},
		{
			input: "????",
			left:  "????",
			mid:   "",
			right: "",
		},
	}

	for _, tc := range tests {
		left, mid, right := trimNonLetters(tc.input)
		if left != tc.left ||
			mid != tc.mid ||
			right != tc.right {
			t.Errorf("input: %v, want: %q %q %q, got: %q %q %q",
				tc.input, tc.left, tc.mid, tc.right, left, mid, right)
		}
	}
}
