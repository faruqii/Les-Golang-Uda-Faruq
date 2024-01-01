package main

import "fmt"

func main() {
	const exitCommand = "STOP"
	var siswa []string

	for {
		fmt.Println("Masukkan Nama Siswa (STOP untuk berhenti)")
		var namaSiswa string
		fmt.Scan(&namaSiswa)

		if namaSiswa == exitCommand {
			fmt.Println("Berhasil Keluar Program")
			break
		}

		siswa = append(siswa, namaSiswa)
	}

	for _, siswa := range siswa {
		fmt.Println(siswa)
	}
}