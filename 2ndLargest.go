package main

import (
	"fmt"
)

func main(){
    var a,b,c,d,l,sl int= 0,0,0,0,0,0;
	fmt.Println("Enter the first number")
	fmt.Scanln(&a)
	fmt.Println("Enter the second number")
	fmt.Scanln(&b)
	fmt.Println("Enter the third number")
	fmt.Scanln(&c)
	fmt.Println("Enter the fourth number")
	fmt.Scanln(&d)
	if(a>l){
		l = a
	}
	if b>l{
		sl=l
		l=b
	}
	if b>sl&&b<l{
		sl = b
	}
	if c>l {
		sl=l
		l=c
	}
	if c>sl&&c<l{
		sl = c
	}
	if d>l {
		sl = l
		l =d
	}
	if d>sl&&d<l{
        sl = d
	}
	fmt.Printf("Second largest number is %v",sl)
}