package main

import "fmt"

type A struct {
        year int
}

func (a A) Greet() { fmt.Println("Hello GolangUK", a.year) }

type B struct {
        A
}

func (b B) Greet() { fmt.Println("Welcome to GolangUK", b.year) }

func main() {
        var a A
        a.year = 2016
        var b B
        b.year = 2016
        a.Greet() // Hello GolangUK 2016
        b.Greet() // Welcome to GolangUK 2016



		var octo OctoCat
        fmt.Println(octo.Legs()) // 5
        octo.PrintLegs()         // I have 4 legs
}



// Open / Closed Principle :: Bertrand Meyer, Object-Oriented Software Construction

type Cat struct {
	Name string
}

func (c Cat) Legs() int { return 4 }

func (c Cat) PrintLegs() {
	fmt.Printf("I have %d legs\n", c.Legs())
}

type OctoCat struct {
	Cat
}

func (o OctoCat) Legs() int { return 5 }


// Barbara Liskov
//https://levelup.gitconnected.com/practical-solid-in-golang-liskov-substitution-principle-e0d2eb9dd39
//A great rule of thumb for Go is accept interfaces, return structs.