package main

import "fmt"

func main(){
	s := sum()
	s(1)
	s(2)
	s(3)

	pov := adder()
	for i:= 1; i<= 5; i++{
		fmt.Println(pov(i))
	}
}

func sum() func(int){
	sum := 0
	 return func(i int) {
		sum += i
		fmt.Println(sum)
	 }
}


func adder() func(int)int{
	sum := 0
	 return func(i int)int {
		sum += i
		return sum
	 }
}