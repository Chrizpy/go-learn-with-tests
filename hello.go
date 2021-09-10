package main

import "fmt"

const spanish = "Spanish"
const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Hola "
const exclamationMarkSuffix = "!"

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == spanish {
		return spanishHelloPrefix + name + exclamationMarkSuffix
	}

	return englishHelloPrefix + name + exclamationMarkSuffix
}

func main() {
	fmt.Println(Hello("world", ""))
}
