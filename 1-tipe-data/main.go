package main

import "fmt"

func main() {
	fmt.Printf("Tipe Data\n\n")
	tipeData()
}

func tipeData() {
	var IniString string
	var IniInt int
	var IniFloat float32
	var IniBool bool

	IniString = "Contoh Tipe Data"
	IniInt = 100
	IniFloat = 10.5
	IniBool = true

	fmt.Printf("IniString = %s \ndengan tipe %T.\n\n", IniString, IniString)
	fmt.Printf("IniInt = %d \ndengan tipe %T.\n\n", IniInt, IniInt)
	fmt.Printf("IniFloat = %.2f \ndengan tipe %T.\n\n", IniFloat, IniFloat)
	fmt.Printf("IniBool = %t \ndengan tipe %T.\n", IniBool, IniBool)
}
