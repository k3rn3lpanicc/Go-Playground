package main

import "fmt"

func foo() (int, int) {
	return 2, 3
}

func multiple_return_main() {
	a, b := foo()
	fmt.Println(a, b)
}
