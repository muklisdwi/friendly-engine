package main

import "fmt"

func main() {
	fmt.Println("Fungsi dan Metode")
	o1 := Orang{
		Nama:    "Fulan",
		Lokasi:  "Bandung",
		Umur:    23,
		Bekerja: true,
		Dagang:  true,
	}

	IniFungsi(o1)
	cekDagang := CekOrangDagang(o1)
	fmt.Println("cekDagang =", cekDagang)

	o1.IniMetode()
	fmt.Println("Cek Dagang o1 =", o1.CekDagang())
	fmt.Println("Sebelum", o1)
	o1.UbahLokasi("Jakarta")
	fmt.Println("Sesudah", o1)
}

type Orang struct {
	Nama    string
	Lokasi  string
	Umur    int
	Bekerja bool
	Dagang  bool
}

func (o Orang) IniMetode() {
	fmt.Println(o)
}

func (o Orang) CekDagang() bool {
	return o.Dagang
}

func (o *Orang) UbahLokasi(s string) {
	o.Lokasi = s
	fmt.Println("Lokasi", o.Nama, "sekarang di", o.Lokasi)
}

func IniFungsi(o Orang) {
	fmt.Println(o)
}

func CekOrangDagang(o Orang) bool {
	return o.Dagang
}
