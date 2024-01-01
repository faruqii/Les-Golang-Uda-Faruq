package main

import "fmt"

func main() {

	// Ask User How Many Number They Want to Input
	// Then Ask User to Input the Number
	// Then Display all the Number
	fmt.Println("Selamat Datang di Program Input Number")
	fmt.Println("Masukkan Jumlah Angka yang ingin diinput: ")
	var jumlah int
	fmt.Scan(&jumlah)

	var numbers []int //list

	for i := 0; i < jumlah; i++ {
		fmt.Printf("Masukkan Angka ke-%d: ", i+1)
		var number int
		fmt.Scan(&number)

		numbers = append(numbers, number)
	}

	fmt.Println("Berikut adalah angka yang kamu input: ")
	for _, number := range numbers {
		fmt.Print(number, " ")
	}

}
