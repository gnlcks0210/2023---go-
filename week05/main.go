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
	inputScore, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	if inputScore >= 90 { //mismatched types string and untyped int
		grade := "A grade!"
	} else {
		gread := "BCDE grad~"
	}

}
