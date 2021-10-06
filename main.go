package main

import (
	"fmt"
	"os"
)

type memberGo struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

type validation struct {
	errString string
}

func main() {
	var ab memberGo
	var err validation

	idMember := os.Args[0]

	//var masukkanMenu string
	// fmt.Println("masukkan Nomor peserta", masukkanMenu)
	// fmt.Scanf("%s", &masukkanMenu)
	switch idMember {
	case "GLNG016ONL001":
		ab.Nama = "Mas Dendhy N"
		ab.Alamat = "kalcit"
		ab.Alasan = "master go"
		ab.Pekerjaan = "Software Engineer"
	case "GLNG016ONL002":
		ab.Nama = "Haris Ardianto"
		ab.Alamat = "Bekasi"
		ab.Alasan = "Belajar go"
		ab.Pekerjaan = "Software Engineer"
	case "GLNG016ONL003":
		ab.Nama = "Hidayat Aji"
		ab.Alamat = "Depok"
		ab.Alasan = "master go"
		ab.Pekerjaan = "Programmer"
	case "GLNG016ONL004":
		ab.Nama = "M Alif Pratama"
		ab.Alamat = "Tendean"
		ab.Alasan = "Belajar"
		ab.Pekerjaan = "Backend dev"
	case "GLNG016ONL005":
		ab.Nama = "M Reyhan"
		ab.Alamat = "jakarta"
		ab.Alasan = "master go"
		ab.Pekerjaan = "Software Engineer"
	case "GLNG016ONL006":
		ab.Nama = "Imam Setia U N"
		ab.Alamat = "jakarta"
		ab.Alasan = "master go"
		ab.Pekerjaan = "Software Engineer"
	case "GLNG016ONL007":
		ab.Nama = "Rezan Lantang R"
		ab.Alamat = "jakarta"
		ab.Alasan = "master go"
		ab.Pekerjaan = "Software Engineer"
	case "GLNG016ONL008":
		ab.Nama = "Mohtar Nurwahid"
		ab.Alamat = "jakarta"
		ab.Alasan = "master go"
		ab.Pekerjaan = "Software Engineer"
	case "GLNG016ONL009":
		ab.Nama = "Dwi Ahmad H"
		ab.Alamat = "jakarta"
		ab.Alasan = "master go"
		ab.Pekerjaan = "Software Engineer"
	case "GLNG016ONL010":
		ab.Nama = "Barikly M"
		ab.Alamat = "jakarta"
		ab.Alasan = "master go"
		ab.Pekerjaan = "Software Engineer"
	case "GLNG016ONL011":
		ab.Nama = "Seven"
		ab.Alamat = "jakarta"
		ab.Alasan = "master go"
		ab.Pekerjaan = "Software Engineer"
	case "GLNG016ONL012":
		ab.Nama = "Arnold"
		ab.Alamat = "jakarta"
		ab.Alasan = "master go"
		ab.Pekerjaan = "Software Engineer"
	case "GLNG016ONL013":
		ab.Nama = "Beni"
		ab.Alamat = "jakarta"
		ab.Alasan = "master go"
		ab.Pekerjaan = "Software Engineer"
	case "GLNG016ONL014":
		ab.Nama = "Ivan"
		ab.Alamat = "jakarta"
		ab.Alasan = "master go"
		ab.Pekerjaan = "Software Engineer"
	default:
		err.errString = "Kode Member Tidak ditemukan"
		ab.Nama = "-"
		ab.Alamat = "-"
		ab.Alasan = "-"
		ab.Pekerjaan = "-"

		fmt.Println(err.errString)
	}

	fmt.Println("Nama: ", ab.Nama)
	fmt.Println("Alamat: ", ab.Alamat)
	fmt.Println("Alasan: ", ab.Alasan)
	fmt.Println("Pekerjaan: ", ab.Pekerjaan)
}
