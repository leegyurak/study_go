package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(varName string) (length int, uppercase string) {
	defer fmt.Println("finish!") // defer: 함수가 종료되고 실행되는 코드를 정의 가능
	length = len(varName)
	uppercase = strings.ToUpper(varName)
	// return length, uppercase  (not go style code)
	return //go는 naked return이라는 개념을 사용해서 파라미터에 정의된 변수를 자동으로 리턴해줌
}

func main() {
	const constName string = "leegyurak" // constant
	var varName string = "leegyurak"     // variable, == varName := "leegyurak" only function and variable

	length, uppercase := lenAndUpper(varName)

	fmt.Println(length, uppercase)
}
