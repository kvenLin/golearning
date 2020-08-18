package main

import "fmt"

func updateSlice(arr []int) {
	arr[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s := arr[2:6]
	//fmt.Println(s)
	//fmt.Println("arr[2:6] = ", arr[2:6])
	//fmt.Println("arr[:6] = ", arr[:6])
	//fmt.Println("arr[2:] = ", arr[2:])
	//fmt.Println("arr[:] = ", arr[:])
	//
	//updateSlice(arr[:])
	//fmt.Println(arr)
	//
	//fmt.Printf("s=%v, len(s1)=%d, cap(s)=%d \n", s, len(s), cap(s))
	//fmt.Println("ReSlice:")
	//s = s[:6]
	//fmt.Println(s)
	//
	//k := s[:4]
	//fmt.Printf("k=%v, len(k)=%d, cap(k)=%d \n", k, len(k), cap(k))
	//fmt.Println("k[:6] = ", k[:6])
	s3 := append(s, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3, s4, s5 = ", s3, s4, s5)

	fmt.Println("arr = ", arr)
}
