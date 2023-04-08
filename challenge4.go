package main

import (
	"fmt"
	"os"
)

type data struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func (p data) printData() string {
	return fmt.Sprintf("Nama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan: %s\n", p.nama, p.alamat, p.pekerjaan, p.alasan)
}

func main() {
	p1 := data{
		nama:      "Radit",
		alamat:    "Pokonya di Bandung",
		pekerjaan: "Mahasiswa",
		alasan:    "Karena saya mau masuk Tokped",
	}
	p2 := data{
		nama:      "Fatur",
		alamat:    "Jakarta",
		pekerjaan: "Pelajar",
		alasan:    "Karena keren",
	}
	p3 := data{
		nama:      "Usop",
		alamat:    "Kapal",
		pekerjaan: "Bajak Laut",
		alasan:    "Karena saya belajar algoritma",
	}
	p4 := data{
		nama:      "Darma",
		alamat:    "Bandung",
		pekerjaan: "Ustad",
		alasan:    "Karena warnanya biru",
	}
	p5 := data{
		nama:      "Vanes",
		alamat:    "Sukabirus",
		pekerjaan: "Mahasiswa",
		alasan:    "Karena gratis",
	}
	p6 := data{
		nama:      "Adhis",
		alamat:    "Jakarta",
		pekerjaan: "Kontraktor",
		alasan:    "menurut saya itu keren",
	}
	p7 := data{
		nama:      "Hamdan",
		alamat:    "Pekanbaru",
		pekerjaan: "Pendemo",
		alasan:    "disuruh kakak",
	}
	p8 := data{
		nama:      "Abi",
		alamat:    "Malang",
		pekerjaan: "Rektor",
		alasan:    "ingin skill up",
	}
	p9 := data{
		nama:      "Anya",
		alamat:    "Jakarta",
		pekerjaan: "Pelari",
		alasan:    "karena saya sedang bosan",
	}
	p10 := data{
		nama:      "Daffa",
		alamat:    "Bogor",
		pekerjaan: "Koki",
		alasan:    "Karena saya ingin belajar hal baru",
	}

	nomor := os.Args[1]
	var p data
	switch nomor {
	case "1":
		p = p1
	case "2":
		p = p2
	case "3":
		p = p3
	case "4":
		p = p4
	case "5":
		p = p5
	case "6":
		p = p6
	case "7":
		p = p7
	case "8":
		p = p8
	case "9":
		p = p9
	case "10":
		p = p10
	default:
		fmt.Println("Nomor tidak valid")
		return
	}
	fmt.Println(p.printData())
}
