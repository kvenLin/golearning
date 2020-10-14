package main

import "fmt"

func variable() {
	var a int
	var s string
	fmt.Printf("%d %q \n", a, s)
}

func main() {
	fmt.Println("hello world")
	variable()

}
