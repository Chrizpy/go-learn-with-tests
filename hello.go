package main

import "fmt"

const englishHelloPrefix = "Hello "
const exclamationMarkSuffix = "!"

func Hello(name string) string {
	if name == "" {
		name = "world"
	}

	return englishHelloPrefix + name + exclamationMarkSuffix
}

func main() {
	fmt.Println(Hello("world"))
}
