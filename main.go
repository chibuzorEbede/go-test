package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func swap(x string, y string) (string,string){
	return y, x
}

func main() {
	
	fmt.Println(add(19,2))
	a,b := swap("hi","mom")
	fmt.Println(a,b)
}