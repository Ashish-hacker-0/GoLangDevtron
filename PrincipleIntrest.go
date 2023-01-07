package main

import (
	"fmt"
	"math"
)

func main(){
	var p,r,n float64=1000,7,3
	ci := (float64)(p*math.Pow(1+r/100,n))
	fmt.Print(ci);

}