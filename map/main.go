package main

import "fmt"

func main() {
	// var colours map[string]string
	colours := make(map[string]string)

	colours = map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	colours["white"] = "#ffffff"

	printMap(colours)

	delete(colours, "green")

	fmt.Println(colours)
}

func printMap(c map[string]string) {
	for colour, hex := range c {
		fmt.Println("Hex code for", colour, "is", hex)
	}

}
