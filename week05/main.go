package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Tiem
	now = time.Now()
	year := now.Year()
	var month string = now.Month().String()
	fmt.Println(year, month, now.Hour(), now.Minute(), now.Second())
}
