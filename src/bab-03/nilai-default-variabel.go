package main

import "fmt"

func main() {

	// unsigned-integer
	var defUint8 uint8
	var defUint16 uint16
	var defUint32 uint32
	var defUint64 uint64
	var defUint uint

	// signed-integer
	var defInt8 int8
	var defInt16 int16
	var defInt32 int32
	var defInt64 int64
	var defInt int

	// string
	var defString string

	// floating-point
	var defFloat32 float32
	var defFloat64 float64

	// complex
	var defComplex64 complex64
	var defComplex128 complex128

	// alias
	var defByte byte
	var defRune rune

	fmt.Println("\nNilai default untuk uint8 = ", defUint8)
	fmt.Println("Nilai default untuk uint16 = ", defUint16)
	fmt.Println("Nilai default untuk uint32 = ", defUint32)
	fmt.Println("Nilai default untuk uint64 = ", defUint64)
	fmt.Println("Nilai default untuk uint = ", defUint)

	fmt.Println("\nNilai default untuk int8 = ", defInt8)
	fmt.Println("Nilai default untuk int16 = ", defInt16)
	fmt.Println("Nilai default untuk int32 = ", defInt32)
	fmt.Println("Nilai default untuk int63 = ", defInt64)
	fmt.Println("Nilai default untuk int = ", defInt)

	fmt.Println("\nNilai default untuk string = ", defString)

	fmt.Println("\nNilai default untuk float32 = ", defFloat32)
	fmt.Println("Nilai default untuk float64 = ", defFloat64)

	fmt.Println("\nNilai default untuk complex64 = ", defComplex64)
	fmt.Println("Nilai default untuk complex128 = ", defComplex128)

	fmt.Println("\nNilai default untuk byte = ", defByte)
	fmt.Println("Nilai default untuk rune = ", defRune)

}
