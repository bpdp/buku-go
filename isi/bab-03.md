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

// Program Go diawali dengan nama paket.
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

Untuk menjalankan kode sumber di atas, ikuti langkah-langkah berikut:

### Tanpa Proses Kompilasi

~~~bash
$ go run aplikasi.go 
Halo bpdp
Home anda di /home/bpdp
Anda menggunakan Go di /home/bpdp/software/go-dev-tools/go/go1.4.1
~~~

### Mengkompilasi Menjadi *Binary Executable*

~~~bash
go build aplikasi.go 
$ ls -la
total 1544
drwxr-xr-x 2 bpdp bpdp    4096 Jan 20 15:35 .
drwxr-xr-x 5 bpdp bpdp    4096 Dec 24 20:18 ..
-rwxr-xr-x 1 bpdp bpdp 1564752 Jan 20 15:35 aplikasi
-rw-r--r-- 1 bpdp bpdp     900 Dec 10 11:32 aplikasi.go
$ ./aplikasi 
Halo bpdp
Home anda di /home/bpdp
Anda menggunakan Go di /home/bpdp/software/go-dev-tools/go/go1.4.1
$ file aplikasi
aplikasi: ELF 32-bit LSB executable, Intel 80386, version 1 (SYSV), statically linked, not stripped
$
~~~


## Pustaka / Package


## Tipe Data Dasar

### Angka / Numerik

Untuk tipe numerik, pada dasarnya kita bisa menggunakan bilangan bulat (*integer*) dan bilangan pecahan (*floating-point*). Bilangan bulat terdiri atas bilangan bertanda (*signed* - int) dan bilangan tak-bertanda (*unsigned* - uint). Berikut ini adalah daftar lengkap dari tipe data numerik tersebut:


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

Selain definisi di atas, Go juga mempunyai alias penyebutan yang implementasinya tergantung pada arsitektur komputer yang digunakan:

| Tipe | Arti |
|------|------|
| uint | arsitektur 32 atau 64 bit |
| int  | mempunyai ukuran yang sama dengan uint |
| uintptr | bilangan bulat tak bertanda untuk menyimpan nilai pointer |

### String

### Boolean

## Variabel dan Konstanta

### Variabel


```go
// nilai-default-variabel.go
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
~~~

Hasil eksekusi:

```bash
$ go run nilai-default-variabel.go

Nilai default untuk uint8 =  0
Nilai default untuk uint16 =  0
Nilai default untuk uint32 =  0
Nilai default untuk uint64 =  0
Nilai default untuk uint =  0

Nilai default untuk int8 =  0
Nilai default untuk int16 =  0
Nilai default untuk int32 =  0
Nilai default untuk int63 =  0
Nilai default untuk int =  0

Nilai default untuk string =  

Nilai default untuk float32 =  0
Nilai default untuk float64 =  0

Nilai default untuk complex64 =  (0+0i)
Nilai default untuk complex128 =  (0+0i)

Nilai default untuk byte =  0
Nilai default untuk rune =  0
$
~~~

### Konstanta

Konstanta dimaksudkan untuk menampung data yang tidak akan berubah-ubah. Konstanta dideklarasikan menggunakan kata kunci *const*. Konstant bisa bertipe *character*, string, boolean, atau numerik.

## Pointer

## Struktur Kendali

### Perulangan dengan `for`

### Seleksi Kondisi

#### Pernyataan `if`


#### Pernyataan `switch`

[[ Daftar Isi ]](README.md) [[ Awal ]](../README.md)

