package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Tic Tac Toe Game")
	hashes := strings.Repeat("#", 10)
	fmt.Println(hashes)
	fmt.Println("Player1: X || Computer: O")
}
