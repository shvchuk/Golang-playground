package main

import (
	"fmt"
)

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4b23e",
		"white": "#fffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
