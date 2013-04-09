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
	// akan dijelaskan lebih lanjut, no need to
	// care a lot right now.
	var (
		user string
		homeDir string
		goHome string
	)

	user = os.Getenv("USER")
	homeDir = os.Getenv("HOME")
	goHome = os.Getenv("GOROOT")

	fmt.Printf("Halo %s", user)
	fmt.Printf("\nHome anda di %s", homeDir)
	fmt.Printf("\nAnda menggunakan Go di %s", goHome)
	fmt.Printf("\n")

}
