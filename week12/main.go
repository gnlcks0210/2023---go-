package main

import (
	"fmt"
)

func main() {
	s := []int{0, 0, 0, 0, 0} //슬라이스 리터럴 이용해 슬라이스 생성 및 메모리 할당, 초기화 진행

	s[4] = 99
	s[2] = 91

	for _, value := range s {
		fmt.Println(value)
	}

	copyS := s[1:4]

	for _, value := range copyS {
		fmt.Println(value)
	}
}
