package main

import "fmt"

func main(){
	var a  = 5
	var is = true 
	for i:=2 ;i*i<a;i++ {
        if a%i==0 {
			is = false
			break
		}
	}
	if is {
		fmt.Printf("%v is prime", a)
	}else{
		fmt.Printf("%v is not prime",a)
	}
}