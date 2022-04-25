package main

import (
	"fmt"
)

func main() {
	fmt.Println("Error Handling")
	fmt.Println("")

	// newError := errors.New("uji error")
	// PanicIfError(newError)
	PanicIfError(nil)
	fmt.Println("OK")
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
