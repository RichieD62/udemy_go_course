package main

import "fmt"

func main() {
	// var colors map[string]string
	// colors := make(map[string]string) // lines 6 and 8 are the same thing

	colors := map[string]string{
		"red":   "#ff0000",
		"blue":  "#0000ff",
		"green": "#008000",
		"white": "#ffffff",
	}

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("The hex code for", color, "is", hex)
	}
}

// colors["white"] = "#ffffff" // adds a key:value pair of white:#ffffff
// delete(colors, "white") // deletes a key:value pair from a map
