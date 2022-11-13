package main

import "fmt"

// we do not have ; here!!
func var_main() {
	//Variables
	/*
		var lol string = "salam" => specify explicitly the type , is not a must
		var lol = "salam" => its fine to declare like this, no warnnings
		lol := "salam" => it is equivalant to the above line , := means that you are defining a var for the first time
	*/
	var lol string = "salam"
	var lol2 = "salam"
	lol3 := "salammm"
	lol3 = "na"
	fmt.Println(lol)
	fmt.Println(lol2)
	fmt.Println(lol3)
	a, b := 2, 3
	fmt.Println(a, b)

}
