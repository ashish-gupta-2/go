package main

import "fmt"


func main(){

	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Diana",
		lastName:  "Zimmermann",
		age:       30,
	}

	fmt.Println(diana)


	
	bestBook := Book{"1984 by George Orwell", 10.2, false}


	fmt.Println(bestBook.float64)

	//Go provides the & (ampersand) operator also known as the address of operator.
	lang := "Golang"
	fmt.Println(&lang) // -> 0xc00010a040

	var x int = 2                                                                  // -> int value
	ptr := &x                                                                      // -> pointer to int
	fmt.Printf("ptr is of type %T with value %v and address %p\n", ptr, *ptr, &ptr) 

}

type Book struct {
	string
	float64
	bool
}
