package main

import (
	"fmt"
	"math/cmplx"
)

const Trx = 42.1786

var accountId, name, number  = "001724","ijeoma",34

var(
	ToBe bool = false
	sk rune = 12
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

var (
	nj int = 32
	gh float64 = float64(nj)
)

func add(x, y int) int {
	return x + y
}

func swap(x string, y string) (string,string){
	return y, x
}

func kl(x int,y int) (int,string){
	jane := "frugality"
	return (y*x ),jane
}

func namedReturnValues(s int, p int) (x,y int){
	x = s + 2
	y = p + 5
	return
}

func main() {
	
	fmt.Println(add(19,2))
	a,b := swap("hi","mom")
	fmt.Println(a,b)
	fmt.Println(kl(2,4))
	fmt.Println(namedReturnValues(1,2))
	fmt.Println(accountId,name,number)
	fmt.Printf("the type %T has value: %v \n", ToBe,ToBe)
	fmt.Printf("the type %T has value: %v \n", sk,sk)
	fmt.Printf("the type %T has value: %v \n", z,z)
	fmt.Printf("the type %T has value: %v \n", gh,gh)
	fmt.Println("the constant is ", Trx)
}