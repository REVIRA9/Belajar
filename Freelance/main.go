package main

import (
	"Freelance/function"
	"fmt"
)

func main(){
	fmt.Println("hello")
	function.Menu()

	for i:=0;i<=10;i++{
		fmt.Println(i)
	}
	function.Menu()
}