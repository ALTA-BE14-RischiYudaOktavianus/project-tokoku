package main

import (
	"fmt"
	"os"
	"os/exec"

	"tokoku/admin"
	config "tokoku/config"
	"tokoku/entity"
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
						fmt.Print("\nPILIHAN anda:\n1. Tambah Pegawai \n2. Tambah Customer \n3. Tambah Barang \n4. Edit Barang \n5. Update Barang \n6. Transaksi \n7. Delete Pegawai \n8. Delete Barang \n9. Logout\n")
						fmt.Println("=========Masukkan Pilihan Anda=========")
						var pilihan int
						fmt.Scanln(&pilihan)

						callClear()
						switch pilihan {
						case 1:
							{
								fmt.Print("=========Program TOKOKU=========")
								var newUser admin.Pegawai
								fmt.Print("\nMasukkan nama : ")
								fmt.Scanln(&newUser.Nama_Pegawai)
								fmt.Print("Masukkan password : ")
								fmt.Scanln(&newUser.Password)
								res, err := authAdmMenu.Register(newUser)
								if err != nil {
									fmt.Println(err.Error())
								}
								if res {
									fmt.Println("Sukses mendaftarkan data")
								} else {
									fmt.Println("Gagal mendaftarkan data")
								}
								fmt.Println("=========Data Pegawai=========")

							}
						case 2:
							{
								fmt.Print("=========Program TOKOKU=========")
								var newUser entity.Customer
								fmt.Print("\nMasukkan nama : ")
								fmt.Scanln(&newUser.Nama_Customer)
								fmt.Print("\nMasukkan ID Pegawai : ")
								fmt.Scanln(&newUser.Nama_Pegawai)
								res, err := authAdmMenu.Customer(newUser)
								if err != nil {
									fmt.Println(err.Error())
								}
								if res {
									fmt.Println("Sukses mendaftarkan customer")
								} else {
									fmt.Println("Gagal mendaftarkan customer")
								}
								fmt.Println("=========Data Customer=========")
							}
						case 3:
							{
								fmt.Print("=========Program TOKOKU=========")
								var newBarang entity.Barang
								fmt.Print("\nMasukkan nama barang : ")
								fmt.Scanln(&newBarang.Nama_Barang)
								fmt.Print("\nMasukkan jumlah Stock : ")
								fmt.Scanln(&newBarang.Stock)
								fmt.Print("\nMasukkan deskripsi barang : ")
								fmt.Scanln(&newBarang.Deskripsi)
								fmt.Print("\nMasukkan ID Pegawai : ")
								fmt.Scanln(&newBarang.Nama_Pegawai)
								res, err := authAdmMenu.Barang(newBarang)
								if err != nil {
									fmt.Println(err.Error())
								}
								if res {
									fmt.Println("Sukses mendaftarkan Barang")
								} else {
									fmt.Println("Gagal mendaftarkan Barang")
								}
								fmt.Println("=========Data Barang=========")
							}
						case 4:
							{
								fmt.Print("=========Program TOKOKU=========")
								var updateBarang entity.Barang
								fmt.Println("\nmasukkan id barang yang akan diedit :")
								fmt.Scanln(&updateBarang.Id)
								fmt.Println("\nmasukkan Informasi terbaru")
								fmt.Scanln(&updateBarang.Deskripsi)
								res, err := authAdmMenu.EditBarang(updateBarang)
								if err != nil {
									fmt.Println(err.Error())
								}
								if res {
									fmt.Println("Sukses mengUpdate Barang")
								} else {
									fmt.Println("Gagal mengUpdate Barang")
								}
								fmt.Println("=========Data Barang=========")
							}
						case 5:
							{
								fmt.Print("=========Program TOKOKU=========")
								var updateBarang entity.Barang
								fmt.Println("\nmasukkan id barang yang akan diedit :")
								fmt.Scanln(&updateBarang.Id)
								fmt.Println("\nmasukkan Jumlah stok terbaru")
								fmt.Scanln(&updateBarang.Stock)

								res, err := authAdmMenu.UpdateBarang(updateBarang)
								if err != nil {
									fmt.Println(err.Error())
								}
								if res {
									fmt.Println("Sukses mengUpdate Barang")
								} else {
									fmt.Println("Gagal mengUpdate Barang")
								}
								fmt.Println("=========Data Barang=========")
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
							fmt.Print("=========Program TOKOKU=========")
							fmt.Print("\nPILIHAN anda:\n1.Tambah Customer \n2. Tambah Barang \n4. Edit Barang \n5. Update Barang \n6. Transaksi \n9. Logout\n")
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
