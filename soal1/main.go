package main

import (
    "fmt"
)

func main() {
    var x, y int

    fmt.Print("Masukkan dua bilangan bulat positif (x dan y):")
    fmt.Scan(&x)
    fmt.Scan(&y)

    if isFactor(x, y) {
        fmt.Println("true -", x, "habis membagi", y)
    } else {
        fmt.Println("false -", x, "tidak habis membagi", y)
    }
}

// isFactor menentukan apakah x adalah faktor dari y.
func isFactor(x, y int) bool {
    return y%x == 0
}