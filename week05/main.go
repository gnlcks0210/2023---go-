package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Input score: ")
	reader := bufio.NewReader(os.Stdin)
	inputScore, err := reader.ReadString('\n') //err declared and not used
	fmt.Println(inputScore)
}
