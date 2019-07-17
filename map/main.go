package main

import "fmt"

func main() {
	// 1 of 3 ways of initializing a Map value
	// var colors map[string]string

	// 2 of 3 ways of initializing a Map value
	// colors := make(map[string]string)

	// how to add a value to the map
	// there is no dot notation for keys for maps and have to use square brackets.
	// Also the correct key type as to be used so if the key type was int we would use colors[10]
	// colors["white"] = "#ffffff"

	// how to delete a key/value from map
	// delete(colors, "white")

	// 3 of 3 ways of initializing a Map value
	// maps are similar to objects, dictionaries, hashes except the keys have to be the same type and the values have to be the same type
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
