package main

import (
	"fmt"
	"log"
)

func main() {
	// variable menu
	var menuLogin int = 1


	for menuLogin != 0 {
		fmt.Println("1. Login Sebagai Admin")
		fmt.Println("2. Login Sebagai Pegawai")
		fmt.Println("9. Exit")
		fmt.Scanln(&menuLogin)
	}
}