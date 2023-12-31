package main

import "fmt"

func main() {
	// If Else
	// Studi Kasus: Pemilu
	// Objective : Buatlah sebuah program untuk mengecek apakah orang eligible untuk memilih
	// Prasyarat: Orang yang bisa memilih harus berumur lebih dari 17 dan Mempunyai KTP Selain Itu tidak bisa memilih
	var name string
	var age int

	fmt.Println("Selamat Datang di Aplikasi Pemilu")
	fmt.Println("Masukkan Nama: ")
	fmt.Scan(&name)
	fmt.Println("Masukkan Umur: ")
	fmt.Scan(&age)

	if age >= 17 {
		var ktp string
		fmt.Println("Apakah sudah mempunyai KTP? (y/n)")
		fmt.Scan(&ktp)
		if ktp == "y" {
			fmt.Printf("Selamat %s Kamu Bisa Memilih", name)
		} else {
			fmt.Printf("Mohon Maaf %s Kamu Belum Bisa Memilih Karena Belum Mempunyai KTP", name)
		}
	} else {
		fmt.Printf("Mohon Maaf %s Kamu Belum Bisa Memilih", name)
	}

	

}
