package main

import "fmt"

const helloEnglish = "Hello, %s"
const helloSpanish = "Hola, %s"
const world = "World"
const spanish = "Spanish"

func Hello(p, lang string) string {
	helloTemplate := helloEnglish
	if lang == spanish {
		helloTemplate = helloSpanish
	}

	if p != "" {
		return fmt.Sprintf(helloTemplate, p)
	}
	return fmt.Sprintf(helloTemplate, world)
}

func main() {
	fmt.Println(Hello("world", ""))
}
