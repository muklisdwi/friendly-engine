package main

import "fmt"

func main() {
	fmt.Println("Struktur Data")
	fmt.Println("")
	strukturData()
}

type TipeData struct {
	Angka int
	Nama  string
	Valid bool
}

func strukturData() {
	IniArray := [5]int{1, 2, 3, 4, 5}
	IniSlice := []string{"satu", "dua", "tiga", "empat", "lima"}
	IniMap := make(map[string]int)

	fmt.Printf("IniArray = %v \ndengan tipe %T\n\n", IniArray, IniArray)
	fmt.Printf("IniSlice = %v \ndengan tipe %T\n\n", IniSlice, IniSlice)

	for i := 0; i < len(IniArray); i++ {
		IniMap[IniSlice[i]] = IniArray[i]
	}

	fmt.Printf("IniSlice = %v \ndengan tipe %T\n\n", IniMap, IniMap)

	TD1 := TipeData{
		Angka: 10,
		Nama:  "Sepuluh",
		Valid: true,
	}

	fmt.Printf("TD1 = %v \ndengan tipe %T \n\n", TD1, TD1)

}
