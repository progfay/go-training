package main

import "fmt"

func ReturnNonZeroValueWithoutReturnStatement() (value int64) {
	defer func() {
		recover()
	}()

	value = 1
	panic("")
}

func main() {
	fmt.Println(ReturnNonZeroValueWithoutReturnStatement())
}
