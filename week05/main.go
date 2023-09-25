package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Input score: ")
	reader := bufio.NewReader(os.Stdin)
	inputScore, _ := reader.ReadString('\n') // Ignore the error return value with the brank
	fmt.Println(inputScore)
}
