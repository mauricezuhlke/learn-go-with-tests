package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hiya, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {

	if name == "" {
		name = "peeps"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}

	if language == french {
		return frenchHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", "English"))
}
