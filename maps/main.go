package main

import "fmt"

func main() {
	// var colors map[string]string

	// colors := make(map[int]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#abc",
		"white": "#ffffff",
	}

	colors["yellow"] = "asdasd"
	delete(colors, "yellow")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
