# Dasar-dasar Pemrograman Go

[[ Daftar Isi ]](README.md) [[ Awal ]](../README.md)

## Struktur Program Go

## Program Aplikasi Sederhana - 1 File `Executable` Utama

Suatu aplikasi `executable` (artinya bisa dijalankan secara langsung oleh sistem operasi) mempunyai struktur seperti yang terlihat pada listing berikut ini:

~~~go
/*
	aplikasi.go

	Contoh program sederhana untuk menjelaskan
	struktur program Go untuk aplikasi executable

	(c) bpdp.name

*/

// Program Go diawali dengan nama pake.
// Paket untuk aplikasi executable selalu berada
// pada paket main.
package main

// pustaka standar yang diperlukan
// Jika hanya satu:
// import "fmt"
// Jika lebih dari satu:
import (
	"fmt"
	"os"
)

// "Fungsi" merupakan satuan terintegrasi dari
// program Go, selalu diberi nama "main" untuk
// aplikasi executable.
func main() {

	// ini adalah kode sumber / program Go
	// akan dijelaskan lebih lanjut, abaikan
	// jika belum paham
	var (
		user    string
		homeDir string
		goHome  string
	)

	user = os.Getenv("USER")
	homeDir = os.Getenv("HOME")
	goHome = os.Getenv("GOROOT")

	fmt.Printf("Halo %s", user)
	fmt.Printf("\nHome anda di %s", homeDir)
	fmt.Printf("\nAnda menggunakan Go di %s", goHome)
	fmt.Printf("\n")

}
~~~

## Pustaka / Package


## Tipe Data Dasar

### Angka / Numerik

|  Tipe     | Arti | Jangkauan |
|-----------|------|-----------|
| uint8     | unsigned  8-bit integer | 0 sampai  255 |
| uint16    | unsigned 16-bit integer | 0 sampai 65535 |
| uint32    | unsigned 32-bit integer | 0 sampai 4294967295 |
| uint64    | unsigned 64-bit integer | 0 sampai 18446744073709551615 |
|  |  |  |
| int8      | signed  8-bit integer | -128 sampai 127 |
| int16     | signed 16-bit integer | -32768 sampai 32767 |
| int32     | signed 32-bit integer | -2147483648 sampai 2147483647 |
| int64     | signed 64-bit integer | -9223372036854775808 sampai 9223372036854775807 |
|  |  |  |
| float32   | IEEE-754 32-bit floating-point |  |
| float64   | IEEE-754 64-bit floating-point |  |
|  |  |  |
| complex64  | bilangan kompleks dengan float32 riil dan imajiner | ~ |
| complex128 | bilangan kompleks dengan float64 riil dan imajiner | ~ |
|  |  |  |
| byte | alias dari uint8 |  |
| rune | alias dari int32 |  |


### String

### Boolean

## Variabel dan Konstanta

### Variabel

### Konstanta

Konstanta dimaksudkan untuk menampung data yang tidak akan berubah-ubah. Konstanta dideklarasikan menggunakan kata kunci *const*. Konstant bisa bertipe *character*, string, boolean, atau numerik.

## Pointer

## Struktur Kendali

### Perulangan dengan `for`

### Seleksi Kondisi

#### Pernyataan `if`


#### Pernyataan `switch`

[[ Daftar Isi ]](README.md) [[ Awal ]](../README.md)

