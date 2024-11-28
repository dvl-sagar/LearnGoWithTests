package main

import "fmt"


const french = "French"
const spanish = "Spanish"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string{
	if name == "" {
		name = "World"
	}
	switch language{
	case spanish:
		return spanishHelloPrefix+name
	case french:
		return frenchHelloPrefix+name
	}
	return englishHelloPrefix + name
}
func main() {
	fmt.Println(Hello("world", "english"))
}
