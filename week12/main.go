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

	test := [3]string{"inha", "go", "student"} //배열 리터럴을 이용해서 test 배열 생성
	//tests := test[0:4]//invalid argument: index 4 out of bounds [0:4]
	tests := test[:2] //tests := test[0:2]
	tests2 := test[1:]
	tests2[0] = "python"
	fmt.Println(tests2)
	fmt.Println(tests, len(tests))
	fmt.Println(test)
}
