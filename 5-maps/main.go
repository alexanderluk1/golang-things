package main

import "fmt"

func main() {
	colors := map[string]string {
		"red": "ff0000",
		"green": "1250ja",
	}

	emptyColor := make(map[string]string)
	emptyColor["white"] = "ff00000"
	emptyColor["green"] = "ff19519"

	delete(emptyColor, "white")

	fmt.Println(colors)
	fmt.Println(emptyColor)
}