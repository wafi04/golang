package main
import (
    "fmt"
    "math"
)

func main() {
	var name string
	var age int8

	// Input dengan fmt.Scan() memiliki potensi masalah dengan whitespace
	fmt.Printf("Masukkan nama: ")
	fmt.Scanln(&name)  // Gunakan Scanln untuk membaca input dengan spasi

	fmt.Printf("Masukkan umur: ")
	fmt.Scanln(&age)

	// Pengecekan genap/ganjil
	if age % 2 == 0 {
		fmt.Println("Angka genap")
	} else {
		fmt.Println("Angka ganjil")
	}

	fmt.Println("Nama:", name)
	fmt.Println("Umur:", age)

	// Deklarasi array
	array := [4]int{1, 2, 3, 4}
	var a int8 = 127     // Maks int8
    var b int16 = 32767  // Maks int16
    var c int32 = math.MaxInt32  // Maks int32
    var d int64 = math.MaxInt64  // Maks int64

    fmt.Printf("int8  maks: %d\n", a)
    fmt.Printf("int16 maks: %d\n", b)
    fmt.Printf("int32 maks: %d\n", c)
    fmt.Printf("int64 maks: %d\n", d)

	// Perulangan dengan for
	for i := 0; i < len(array); i++ {
		fmt.Printf("Elemen ke-%d: %d\n", i, array[i])
	}

	// Alternatif perulangan dengan range
	for index, value := range array {
		fmt.Printf("Index: %d, Nilai: %d\n", index, value)
	}

	fmt.Println("Array lengkap:", array)
}