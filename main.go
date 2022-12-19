package main

import (
	"fmt"
	// "log"
)

func main() {
	// variable menu
	var menuLogin int = 1

	for menuLogin != 0 {
		fmt.Println("1. Login Sebagai Admin")
		fmt.Println("2. Login Sebagai Pegawai")
		fmt.Println("0. Exit")
		fmt.Scanln(&menuLogin)
		if menuLogin == 1 {
			fmt.Println("LOGIN SEBAGAI ADMIN")
			fmt.Println("Masukkan Nama: ")
			fmt.Scanln()
			fmt.Println("Masukkan Password: ")
			fmt.Scanln()
		} else if menuLogin == 2 {
			fmt.Println("LOGIN SEBAGAI PEGAWAI")
			fmt.Println("Masukkan Nama: ")
			fmt.Scanln()
			fmt.Println("Masukkan Password: ")
			fmt.Scanln()
		} else if menuLogin == 0 {
			break
		}
	}
}
