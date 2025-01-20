package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("Pokedex > ")
		in.Scan()
		in := strings.Fields(strings.ToLower(in.Text()))
		fmt.Printf("Your command was: %s\n", in[0])
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
