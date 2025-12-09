package main

import "fmt"

func main() {
	colors := map[string]string {
		"red": "ff0000",
		"green": "1250ja",
		"white": "55501a",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}