package main

import "fmt"


func main(){
	arr :=[5]int{}

	fmt.Println(len(arr))
	fmt.Println(cap(arr))

	sl1 := arr[:]
	sl1 = append(sl1, 3)

	fmt.Println(arr)
	fmt.Println(sl1)
	fmt.Println(len(sl1))
	fmt.Println(cap(sl1))

	sl1 = append(sl1, 3)

	sl1 = append(sl1, 3)

	sl1 = append(sl1, 3)

	sl1 = append(sl1, 3)
	fmt.Println(sl1)
	fmt.Println(len(sl1))
	fmt.Println(cap(sl1))
	sl1 = append(sl1, 3)

	fmt.Println(sl1)
	fmt.Println(len(sl1))
	fmt.Println(cap(sl1))
}