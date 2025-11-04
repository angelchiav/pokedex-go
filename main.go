package main

import (
	"fmt"
	"regexp"
)

func cleanInput(text string) []string {
	re := regexp.MustCompile(`[A-Za-zÀ-ÿ]+`)
	return re.FindAllString(text, -1)
}

func main() {
	fmt.Println(cleanInput("La mamá de la mamá"))
}
