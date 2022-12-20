package main

import (
	"fmt"
	"os"
	"os/exec"

	// "tokoku/admin"
	config "tokoku/config"
	"tokoku/pegawai"
	// _customer _"tokoku/customer"
	// _transaksi _"tokoku/transaksi"
)

func callClear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	var cfg = config.ReadConfig()
	var conn = config.ConnectSQL(*cfg)
	var authMenu = pegawai.AuthMenu{DB: conn}
	// var authAdmMenu = admin.AuthMenu{DB: conn}
	var isRunning bool = true
	for isRunning {
		fmt.Print("=========Program TOKOKU=========")
		fmt.Print("\nPILIHAN Menu:\n1. Login \n0. Exit\n")
		fmt.Println("=========Masukkan menu Pilihan Anda=========")
		var choice int
		fmt.Scanln(&choice)
		callClear()
		switch choice {
		case 1:
			{
				fmt.Print("=========Program TOKOKU=========")
				var Username, Password string
				fmt.Print("\nMasukkan nama : ")
				fmt.Scanln(&Username)
				fmt.Print("Masukkan password : ")
				fmt.Scanln(&Password)
				res, err := authMenu.Login(Username, Password)
				if err != nil {
					fmt.Println(err.Error())
				}
				if res.ID == 1 {
					fmt.Println("Sukses Login", "selamat datang", Username)
					var isRunning2 bool = true
					for isRunning2 {
						fmt.Print("=========Program Activity Sederhana=========")
						fmt.Print("\nPILIHAN anda:\n1. Tambah Pegawai \n2. Tambah Barang \n3. Edit Barang \n4. Update Barang \n5. Transaksi \n6. Delete Pegawai \n7. Delete Barang \n9. Logout\n")
						fmt.Println("=========Masukkan Pilihan Anda=========")
						var pilihan int
						fmt.Scanln(&pilihan)

						callClear()
						switch pilihan {
						case 1:
							{

							}
						case 9:
							callClear()
							isRunning2 = false
						}
					}
				} else if res.ID > 1 {
					var Username, Password string
					fmt.Print("Masukkan nama : ")
					fmt.Scanln(&Username)
					fmt.Print("Masukkan password : ")
					fmt.Scanln(&Password)
					res, err := authMenu.Login(Username, Password)
					if err != nil {
						fmt.Println(err.Error())
					}
					if res.ID > 0 {
						fmt.Println("Sukses Login", "selamat datang", Username)
						var isRunning2 bool = true
						for isRunning2 {
							fmt.Print("=========Program Activity Sederhana=========")
							fmt.Print("\nPILIHAN anda:\n1. Tambah Aktivitas \n9. Logout\n")
							fmt.Println("=========Masukkan Pilihan Anda=========")
							var choice2 int
							fmt.Scanln(&choice2)

							callClear()
							switch choice2 {
							case 1:
								{
									fmt.Println("=========Program Activity Sederhana=========")

									fmt.Println("==================")
								}
							case 9:
								callClear()
								isRunning2 = false
							}
						}
					} else {
						fmt.Println("Gagal Login!", Username, "\nsilahkan cek nama dan pasword anda kembali")
					}
				} else {
					fmt.Println("Gagal Login!", Username, "\nsilahkan cek nama dan pasword anda kembali")
				}
			}
		case 0:
			{
				callClear()
				isRunning = false

			}
		}

	}
	fmt.Println("TERIMAKASIH SUDAH MENGGUNAKAN TOKOKU")
}
