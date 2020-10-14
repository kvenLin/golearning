package main

import (
	"fmt"
)

func tryRecover() {
	//定义匿名函数
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic("I dont know what to do")
		}
	}()
	//panic(errors.New("this is an error"))
	panic(123)
}

func main() {
	tryRecover()
}
