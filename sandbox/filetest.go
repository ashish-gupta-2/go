package main

// import the 2 modules we need
import (
	"fmt"
	"io/ioutil"
)

func main() {
	// read in the contents of the localfile.data
	mydata := []byte("All the data I wish to write to a file")

	// the WriteFile method returns an error if unsuccessful
	err := ioutil.WriteFile("myfile.data", mydata, 0777)
	// handle this error
	if err != nil {
		// print it out
		fmt.Println(err)
	}

	data, err := ioutil.ReadFile("myfile.data")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(data))

}
