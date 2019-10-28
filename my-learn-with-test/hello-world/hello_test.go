package main

import "testing"

func TestHello(t *testing.T) {
	testCases := []struct {
		name     string
		language string
		want     string
	}{
		{
			name: "reza",
			want: "Hello, reza!",
		},
		{
			name:     "Pippa",
			language: "spanish",
			want:     "Hola, Pippa!",
		},
		{
			name:     "Penny",
			language: "Spanish",
			want:     "Hola, Penny!",
		},
		{
			name:     "Brandi",
			language: "Japaneese",
			want:     "Hello, Brandi!",
		},
		{
			name: "",
			want: "Hello, World!",
		},
		{
			name:     "",
			language: "spanish",
			want:     "Hola, World!",
		},
	}
	for _, tC := range testCases {
		t.Run("", func(t *testing.T) {
			got := hello(tC.name, tC.language)
			if tC.want != got {
				t.Errorf("wanted '%s', but got '%s'", tC.want, got)
			}
		})
	}
}
