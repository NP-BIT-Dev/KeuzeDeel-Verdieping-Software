package main

import "fmt"



func main(){
	primes := [6]int{2, 3, 5, 7, 11, 13}


	//Slice that includes the first element "1" = 3 in primes and excludes the last element "4" = 11 in primes
	var s []int = primes[1:4]
	fmt.Println(s)
}