package main

import (
	"fmt"
	"strings"
)

const (
	world        = "World"
	spansih      = "spanish"
	spanishHello = "Hola"
	englishHello = "Hello"
)

func hello(s, lang string) string {
	if s == "" {
		s = world
	}

	var hello string
	switch strings.ToLower(lang) {
	case spansih:
		hello = spanishHello
	default:
		hello = englishHello
	}

	return hello + ", " + s + "!"
}

func main() {
	fmt.Println(hello("reza", "english"))
}
