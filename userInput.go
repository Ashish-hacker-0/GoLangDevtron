package main

import (
	"fmt"
	"math"
)

func main(){
	var principle, intrestRate, timePeriod, compoundIntrest float64
	fmt.Print("Enter the principle")
	fmt.Scanln(&principle)
	fmt.Print("Enter the Intrest Rate")
	fmt.Scanln(&intrestRate)
	fmt.Print("Enter the time period")
	fmt.Scanln(&timePeriod)

	compoundIntrest=  principle*math.Pow((1+intrestRate/100),timePeriod)
	fmt.Print(compoundIntrest)
}