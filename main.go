package main

import (
	"fmt"
	"os"
	"os/exec"
	"tokoku/admin"
	"tokoku/barang"
	config "tokoku/config"
	"tokoku/customer"
	"tokoku/pegawai"
	"tokoku/transaksi"
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
	var authBrgMenu = barang.AuthMenu{DB: conn}
	var authCustMenu = customer.AuthMenu{DB: conn}
	var authTransMenu = transaksi.AuthMenu{DB: conn}
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
						fmt.Print("=========Program TOKOKU=========")
						fmt.Print("\nPILIHAN anda:\n1. Tambah Pegawai \n2. Delete Transaksi \n3. Delete Customer \n4. Delete Barang \n5. Delete Pegawai \n9. Logout\n")
						fmt.Println("=========Masukkan Pilihan Anda=========")
						var pilihan int
						fmt.Scanln(&pilihan)

						callClear()
						switch pilihan {
						case 1:
							{
								fmt.Print("=========Program TOKOKU=========")
								fmt.Print("\n=========Menu Menambahkan Pegawai=========")
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
								fmt.Print("\n=========Menu Menghapus Transaksi=========")
								var deleteTrans transaksi.Transaksi
								fmt.Println("\nmasukkan id transaksi yang akan diHapus :")
								fmt.Scanln(&deleteTrans.ID)
								res, err := authTransMenu.DeleteTransaksi(deleteTrans)
								if err != nil {
									fmt.Println(err.Error())
								}
								if res {
									fmt.Println("Sukses menghapus Transaksi")
								} else {
									fmt.Println("Gagal menghapus Transaksi")
								}
								fmt.Println("=========Data Transaksi=========")
							}
						case 3:
							{
								fmt.Print("=========Program TOKOKU=========")
								fmt.Print("\n=========Menu Menghapus Customer=========")
								var deleteCust customer.Customer
								fmt.Println("\nmasukkan id customer yang akan diHapus :")
								fmt.Scanln(&deleteCust.Id)
								res, err := authCustMenu.DeleteCustomer(deleteCust)
								if err != nil {
									fmt.Println(err.Error())
								}
								if res {
									fmt.Println("Sukses menghapus Customer")
								} else {
									fmt.Println("Gagal menghapus Customer")
								}
								fmt.Println("=========Data Customer=========")
							}
						case 4:
							{
								fmt.Print("=========Program TOKOKU=========")
								fmt.Print("\n=========Menu Menghapus Barang=========")
								var deleteBarang barang.Barang
								fmt.Println("\nmasukkan id barang yang akan diHapus :")
								fmt.Scanln(&deleteBarang.Id)
								res, err := authBrgMenu.Deletebarang(deleteBarang)
								if err != nil {
									fmt.Println(err.Error())
								}
								if res {
									fmt.Println("Sukses menghapus Barang")
								} else {
									fmt.Println("Gagal menghapus Barang")
								}
								fmt.Println("=========Data Barang=========")
							}
						case 5:
							{
								fmt.Print("=========Program TOKOKU=========")
								fmt.Print("\n=========Menu Menghapus Pegawai=========")
								var deletePegawai pegawai.Pegawai
								fmt.Println("\nmasukkan id Pegawai yang akan diHapus :")
								fmt.Scanln(&deletePegawai.ID)
								res, err := authMenu.DeletePegawai(deletePegawai)
								if err != nil {
									fmt.Println(err.Error())
								}
								if res {
									fmt.Println("Sukses menghapus Pegawai")
								} else {
									fmt.Println("Gagal menghapus Pegawai")
								}
								fmt.Println("=========Data Pegawai=========")
							}
						case 9:
							callClear()
							isRunning2 = false
						}
					}
				} else if res.ID > 1 {
					if res.ID > 1 {
						fmt.Println("Sukses Login", "selamat datang", Username)
						var isRunning2 bool = true
						for isRunning2 {
							fmt.Print("=========Program TOKOKU=========")
							fmt.Print("\nPILIHAN anda:\n1. Tambah Customer \n2. Tambah Barang \n3. Edit Informasi \n4. Update Stock Barang \n5. Transaksi \n6. Cetak Nota Transaksi \n9. Logout\n")
							fmt.Println("=========Masukkan Pilihan Anda=========")
							var choice2 int
							fmt.Scanln(&choice2)

							callClear()
							switch choice2 {
							case 1:
								{
									fmt.Print("=========Program TOKOKU=========")
									fmt.Print("\n=========Menu Menambahkan Customer=========")
									var newUser customer.Customer
									fmt.Print("\nMasukkan nama : ")
									fmt.Scanln(&newUser.Nama_Customer)
									fmt.Print("\nMasukkan ID Pegawai : ")
									fmt.Scanln(&newUser.Nama_Pegawai)
									res, err := authCustMenu.Customer(newUser)
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
							case 2:
								{
									fmt.Print("=========Program TOKOKU=========")
									fmt.Print("\n=========Menu Menambahkan Barang=========")
									var newBarang barang.Barang
									fmt.Print("\nMasukkan nama barang : ")
									fmt.Scanln(&newBarang.Nama_Barang)
									fmt.Print("\nMasukkan jumlah Stock : ")
									fmt.Scanln(&newBarang.Stock)
									fmt.Print("\nMasukkan deskripsi barang : ")
									fmt.Scanln(&newBarang.Deskripsi)
									fmt.Print("\nMasukkan ID Pegawai : ")
									fmt.Scanln(&newBarang.Nama_Pegawai)
									res, err := authBrgMenu.Barang(newBarang)
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
							case 3:
								{
									fmt.Print("=========Program TOKOKU=========")
									fmt.Print("\n=========Menu Edit Informasi Barang=========")
									var updateBarang barang.Barang
									fmt.Println("\nmasukkan id barang yang akan diedit :")
									fmt.Scanln(&updateBarang.Id)
									fmt.Println("\nmasukkan Informasi terbaru")
									fmt.Scanln(&updateBarang.Deskripsi)
									res, err := authBrgMenu.EditBarang(updateBarang)
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
							case 4:
								{
									fmt.Print("=========Program TOKOKU=========")
									fmt.Print("\n=========Menu Update Stock Barang=========")
									var updateBarang barang.Barang
									fmt.Println("\nmasukkan id barang yang akan diedit :")
									fmt.Scanln(&updateBarang.Id)
									fmt.Println("\nmasukkan Jumlah stok terbaru")
									fmt.Scanln(&updateBarang.Stock)

									res, err := authBrgMenu.UpdateBarang(updateBarang)
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
									var isRun bool = true
									for isRun {
										fmt.Print("=========Program TOKOKU=========")
										fmt.Print("\nPILIHAN anda:\n1. Tambah Transaksi \n2. Tambah jumlah Barang \n0. HOME\n")
										fmt.Println("=========Masukkan Pilihan Anda=========")
										var choice2 int
										fmt.Scanln(&choice2)

										callClear()
										switch choice2 {
										case 1:
											{
												fmt.Print("=========Program TOKOKU=========")
												fmt.Print("\n=========Menu Tambah Transaksi=========")
												var newTransaksi transaksi.Transaksi
												fmt.Println("\nMasukkan ID Customer: ")
												fmt.Scanln(&newTransaksi.ID_Customer)
												fmt.Println("\nMasukkan ID Pegawai: ")
												fmt.Scanln(&newTransaksi.ID_Pegawai)

												res, err := authTransMenu.AddTransaksi(newTransaksi)
												if err != nil {
													fmt.Println(err.Error())
												}
												if res {
													fmt.Println("TRANSAKSI SUKSES")
												} else {
													fmt.Println("TRANSAKSI GAGAL")
												}
												fmt.Println("=========Transaksi=========")
											}
										case 2:
											{
												fmt.Print("=========Program TOKOKU=========")
												fmt.Print("\n=========Menu Tambah Transaksi=========")
												var newTransaksi transaksi.Transaksi
												fmt.Println("\nMasukkan ID Transaksi: ")
												fmt.Scanln(&newTransaksi.ID)
												fmt.Println("\nMasukkan ID Barang: ")
												fmt.Scanln(&newTransaksi.ID_Barang)
												fmt.Println("\nJumlah Barang: ")
												fmt.Scanln(&newTransaksi.Total_Qty)

												res, err := authTransMenu.AddQTY(newTransaksi)
												if err != nil {
													fmt.Println(err.Error())
												}
												if res {
													fmt.Println("TRANSAKSI SUKSES")
												} else {
													fmt.Println("TRANSAKSI GAGAL")
												}
												fmt.Println("=========Transaksi=========")
											}
										case 0:
											{
												callClear()
												isRun = false

											}
										}
									}
								}
							case 6:
								{
									fmt.Print("=========Program TOKOKU=========")
									var newTransaksi transaksi.Transaksi
									fmt.Println("\n=========Nota Transaksi=========")
									// fmt.Print("\n No Transaksi			:", "<", newTransaksi.ID, ">")
									// fmt.Print("\n Tanggal Transaksi		:", "<", newTransaksi.Tanggal_Transaksi, ">")
									// fmt.Print("\n Kasir Transaksi		:", "<", newTransaksi.ID_Pegawai, ">")
									// fmt.Print("\n Barang Transaksi		:", "<", newTransaksi.ID_Barang, ">")
									// fmt.Print("\n Jumlah Barang			:", "<", newTransaksi.Total_Qty, ">")
									fmt.Print("\n Customer			:", "<", newTransaksi.ID_Customer, ">")
									fmt.Println("\n=========Nota Transaksi TOKOKU=========")

									res, err := authTransMenu.CetakNota(newTransaksi)
									if err != nil {
										fmt.Println(err.Error())
									}
									if res {
										fmt.Println("TRANSAKSI SUKSES")
									} else {
										fmt.Println("TRANSAKSI GAGAL")
									}

									fmt.Println("=========Transaksi=========")
								}
							case 9:
								callClear()
								isRunning2 = false
							}
						}
					} else {
						fmt.Println("Gagal Login!", Username, "\nsilahkan cek nama dan password anda kembali")
					}
				} else {
					fmt.Println("Gagal Login!", Username, "\nsilahkan cek nama dan password anda kembali")
				}
			}
		case 0:
			{
				callClear()
				isRunning = false

			}
		}

	}
	fmt.Println("TERIMAKASIH SUDAH MENGGUNAKAN PROGRAM TOKOKU")
}
