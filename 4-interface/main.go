package main

import "fmt"

func main() {
	fmt.Println("Interface")

	Harimau1 := Harimau{
		Hewan: Hewan{
			Habitat: "Hutan",
			Makanan: "Daging",
			Berat:   300.00,
			Panjang: 1.50,
		},
	}

	Kucing1 := Kucing{
		Hewan: Hewan{
			Habitat: "Rumah",
			Makanan: "Wiskas",
			Berat:   5.00,
			Panjang: 0.50,
		},
		Jenis: "Anggora",
	}

	BunyiHewan(Harimau1)
	BunyiHewan(Kucing1)

}

func BunyiHewan(sh SuaraHewan) {
	sh.Suara()
}

type SuaraHewan interface {
	Suara()
}

type Hewan struct {
	Habitat string
	Makanan string
	Berat   float32
	Panjang float32
}

type Harimau struct {
	Hewan
}

func (h Harimau) Suara() {
	fmt.Println("Mengaummm.")
}

type Kucing struct {
	Hewan
	Jenis string
}

func (k Kucing) Suara() {
	fmt.Println("Meowww..")
}
