package main

import "fmt"

func main(){
	fmt.Println("Enter the number")
	var a int
	fmt.Scanln(&a);
	count:=0
	for ;a!=0;a=a/10 {
		count++
	}
	fmt.Print(count)
}