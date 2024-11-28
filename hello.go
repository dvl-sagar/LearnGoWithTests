package main

import "fmt"

const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string{
	if name == "" {
		name = "World"
	}
	if language == spanish{
		return spanishHelloPrefix + name
	}
	return englishHelloPrefix + name
}
func main() {
	fmt.Println(Hello("world", "english"))
}
