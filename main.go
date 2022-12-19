package main

import (
	"fmt"
	"os"
	"os/exec"
	"tokoku/admin"
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
	var authAdmMenu = admin.AuthMenu{DB: conn}
	var isRunning bool = true
	for isRunning {
		fmt.Print("=========Program TOKOKU=========")
		fmt.Print("\nPILIHAN Menu:\n1. Login Sebagai Admin \n2. Login Sebagai Pegawai \n0. Exit\n")
		fmt.Println("=========Masukkan menu Pilihan Anda=========")
		var choice int
		fmt.Scanln(&choice)
		callClear()
		switch choice {
		case 1:
			{
				fmt.Print("=========Program TOKOKU=========")
				var Admin pegawai.Pegawai
				fmt.Print("Masukkan nama : ")
				fmt.Scanln(&Admin.Username)
				fmt.Print("Masukkan password : ")
				fmt.Scanln(&Admin.Password)
				res, idLogged, err := authMenu.Login(Admin)
				if err != nil {
					fmt.Println(err.Error())
				}
				if res {
					fmt.Println("Sukses Login", "selamat datang", Admin.Username)
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
								authAdmMenu.Id = idLogged
							}
						case 9:
							callClear()
							isRunning2 = false
						}
					}
				} else {
					fmt.Println("Gagal Login!", Admin.Username, "\nsilahkan cek nama dan pasword anda kembali")
				}
			}
		case 2:
			{
				var Pegawai pegawai.Pegawai
				fmt.Print("Masukkan nama : ")
				fmt.Scanln(&Pegawai.Username)
				fmt.Print("Masukkan password : ")
				fmt.Scanln(&Pegawai.Password)
				res, idLogged, err := authMenu.Login(Pegawai)
				if err != nil {
					fmt.Println(err.Error())
				}
				if res {
					fmt.Println("Sukses Login", "selamat datang", Pegawai.Username)
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
								authAdmMenu.Id = idLogged
								// var newActivity activity.Activity
								// reader := bufio.NewReader(os.Stdin) //standard input
								// fmt.Print("Judul kegiatan : ")
								// text, _ := reader.ReadString('\n')
								// newActivity.Title = text
								// fmt.Print("Lokasi : ")
								// location, _ := reader.ReadString('\n')
								// newActivity.Location = location
								// newActivity.ID = idLogged
								// res, err := authActvMenu.AddActivity(newActivity)
								// if err != nil {
								// 	fmt.Println(err.Error())
								// }
								// if res {
								// 	fmt.Println("Sukses tambah Kegiatan :", text)
								// } else {
								// 	fmt.Println("Gagal tambah data")
								// }
								fmt.Println("==================")
							}
						case 9:
							callClear()
							isRunning2 = false
						}
					}
				} else {
					fmt.Println("Gagal Login!", Pegawai.Username, "\nsilahkan cek nama dan pasword anda kembali")
				}
			}
		case 0:
			{
				callClear()
				isRunning = false

			}
		}

	}
	fmt.Println("TERIMAKASIH SUDAH MENGGUNAKAN ACTIVITY")
}
