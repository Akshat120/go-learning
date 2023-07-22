package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#228B22",
		"white": "#ffffff",
	}

	// var colors map[string]string

	// colors := make(map[int]string)
	// colors[10] = "#ffffff"

	// delete(colors, 10)

	printMap(colors)
	// fmt.Printf("%+v\n", colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}
