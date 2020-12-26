package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	b := englishBot{}
	printGreeting(b)

	s := spanishBot{}
	printGreeting(s)
}

func (eb englishBot) getGreeting() string {
	return "Hello"
}

func (eb englishBot) AnotherMethod() {
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
