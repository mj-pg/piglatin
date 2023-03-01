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
			want:  "eggay",
		},
		{
			input: "always",
			want:  "alwaysay",
		},
		{
			input: "Hello World I'm a gopher",
			want:  "elloHay orldWay I'may ay ophergay",
		},
	}

	for _, tc := range tests {
		got := Translate(tc.input)
		if got != tc.want {
			t.Errorf("input: %v, want: %v, got: %v", tc.input, tc.want, got)
		}
	}
}
