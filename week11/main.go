package main

import "fmt"

func main() {
	// var primes [3]int
	// primes[0] = 2
	// primes[1] = 3
	// primes[2] = 5

	// fmt.Println(primes, primes[1])

	// var primes [3]int = [3]int{2, 3, 5} // primes 배열을 배열 리터럴로 초기화
	// fmt.Println(primes, primes[1])

	primes := [3]int{2, 3, 5} // primes 배열을 배열 리터럴로 초기화
	fmt.Println(primes, primes[1])

	test := [5]bool{true, true, true}
	fmt.Println(test[3]) //boolean 타입의 제로 값, false

}