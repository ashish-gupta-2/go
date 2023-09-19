package main

import "fmt"

func main(){

	// len = 3 cap = 6
	// slice := make([] int , 3,6)

	// slice1:= [] int {4,5}

	n1 := []int{10, 20, 30, 40}
    n1 = append(n1, 100)
    fmt.Println(len(n1), cap(n1))
	n1 = append(n1, 100)
    fmt.Println(len(n1), cap(n1))

}