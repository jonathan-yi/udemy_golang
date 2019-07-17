package main

import "fmt"

// this interface represents all the types in this program with a function called getGreeting() that returns a string
// now all the types that belong to this interface can be referenced as type 'bot' as well
type bot interface {
	getGreeting() string
	// this means for a type to be represented by this interface you need an 1st argument of type string, 2nd argument of type int
	// and return a string and an error
	// getGreetings(string, int) (string, error)
}

type englishBot struct{}

type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi there!"
}
func (spanishBot) getGreeting() string {
	return "Hola!"
}
