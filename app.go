package main

import "fmt"

func main() {
	var m string = "1234"
	var n *string = &m
	var b string = *n
	*n = "324"
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(b)
	fmt.Println(*n)

}
