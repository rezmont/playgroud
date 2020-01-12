package main

import "fmt"

func Hello(p string) string {
	return fmt.Sprintf("Hello %s", p)
}

func main() {
	fmt.Println(Hello("world"))
}
