package main

import "fmt"

func main(){

	str := "test"
 
	// byte =  uint8 and rune = int32  
	run := []rune(str)
	byt := []byte(str)
	var rev string
	var rev1 string

	for i:=len(run)-1 ; i >=0  ; i-- {
		rev = rev + string(run[i])
		rev1 = rev1+ string(byt[i])

		fmt.Println(byt[i])
	}
	fmt.Println(rev)
	fmt.Println(rev1)
}