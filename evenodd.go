package main

import "fmt"

func main(){
	var a,b=2,5;
	if(a%2==0){
		fmt.Printf("%v is even\n", a)
		fmt.Printf("%v is odd", b)
	}else{
		fmt.Printf("%v is even\n", b)
		fmt.Printf("%v is odd", a)
	}
}