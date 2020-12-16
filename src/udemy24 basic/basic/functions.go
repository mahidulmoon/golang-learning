package main

import "fmt"

func main(){
	list := [] float64{22,33,44,55,66}

	average := average(list...)

	fmt.Println("Average:",average)
}

func average(listdata ...float64) float64 {
	total := 0.0

	for _,v := range listdata{
		total += v
	}

	return total/float64(len(listdata))

}