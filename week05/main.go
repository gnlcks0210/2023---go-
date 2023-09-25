package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Input score: ")
	reader := bufio.NewReader(os.Stdin)
	inputScore, err := reader.ReadString('\n') //option2
	if err != nill {
		log.Fatal(err)
	}
	fmt.Println(inputScore)
}
