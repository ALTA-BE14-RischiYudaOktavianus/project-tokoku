package admin

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"tokoku/entity"
)

type Admin struct {
	ID       int
	Nama     string
	Password string
}

type Pegawai struct {
	ID           int
	Nama_Pegawai string
	Password     string
}
type AuthMenu struct {
	DB *sql.DB
}

// func NewAuthMenu() *AuthMenu {
// 	cfg := config.ReadConfig()
// 	conn := config.ConnectSQL(*cfg)
// 	return &AuthMenu{DB: conn}
// }

func (am *AuthMenu) Duplicate(name string) bool {
	res := am.DB.QueryRow("SELECT id FROM pegawai where nama_pegawai = ?", name)
	var idExist int
	err := res.Scan(&idExist)
	if err != nil {
		log.Println("Result scan error", err.Error())
		return false
	}
	return true
}
func (am *AuthMenu) DuplicateBarang(name string) bool {
	res := am.DB.QueryRow("SELECT id FROM barang where nama_barang = ?", name)
	var idExist int
	err := res.Scan(&idExist)
	if err != nil {
		log.Println("Result scan error", err.Error())
		return false
	}
	return true
}
func (am *AuthMenu) DuplicateCustomer(name string) bool {
	res := am.DB.QueryRow("SELECT id FROM customer where nama_cust = ?", name)
	var idExist int
	err := res.Scan(&idExist)
	if err != nil {
		log.Println("Result scan error", err.Error())
		return false
	}
	return true
}
func (am *AuthMenu) Register(newUser Pegawai) (bool, error) {
	// menyiapakn query untuk insert
	registerQry, err := am.DB.Prepare("INSERT INTO pegawai (nama_pegawai, password) values (?,?)")
	if err != nil {
		log.Println("prepare insert user ", err.Error())
		return false, errors.New("prepare statement insert user error")
	}

	if am.Duplicate(newUser.Nama_Pegawai) {
		log.Println("duplicated information")
		return false, errors.New("nama sudah digunakan")
	}

	// menjalankan query dengan parameter tertentu
	res, err := registerQry.Exec(newUser.Nama_Pegawai, newUser.Password)
	if err != nil {
		log.Println("insert user ", err.Error())
		return false, errors.New("insert user error")
	}
	// Cek berapa baris yang terpengaruh query diatas
	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after insert user ", err.Error())
		return false, errors.New("error setelah insert")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}

	return true, nil
}

func (am *AuthMenu) Login(nama_pegawai string, password string) (Pegawai, error) {
	loginQry, err := am.DB.Prepare("SELECT id FROM pegawai WHERE username = ? AND password = ?")
	if err != nil {
		log.Println("prepare insert user ", err.Error())
		return Pegawai{}, errors.New("prepare statement insert user error")
	}

	row := loginQry.QueryRow(nama_pegawai, password)

	if row.Err() != nil {
		log.Println("login query ", row.Err().Error())
		return Pegawai{}, errors.New("tidak bisa login, data tidak ditemukan")
	}
	res := Pegawai{}
	err = row.Scan(&res.ID)

	if err != nil {
		log.Println("after login query ", err.Error())
		return Pegawai{}, errors.New("tidak bisa login, kesalahan setelah error")
	}

	res.Nama_Pegawai = nama_pegawai

	return res, nil
}
func (am *AuthMenu) EditBarang(editBarang entity.Barang) (bool, error) {

	addQry, err := am.DB.Prepare("UPDATE barang set deskripsi=?, nama_pegawai=?  where id= ?")
	// addQry, err := am.DB.Prepare("UPDATE barang set nama_barang, stok_barang, deskripsi, nama_pegawai  where id= ?")
	if err != nil {
		log.Println("Update barang prepare", err.Error())
		return false, errors.New("prepare Edit barang error")
	}

	res, err := addQry.Exec(editBarang.Id, editBarang.Deskripsi, editBarang.Nama_Pegawai)
	if err != nil {
		log.Println("Update barang", err.Error())
		return false, errors.New("Update Barang error")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after Update Barang", err.Error())
		return false, errors.New("after Update Barang error")
	}

	if affRows <= 0 {
		log.Println("No record affected")
		return true, errors.New("No record")
	}

	return true, nil
}
func (am *AuthMenu) UpdateBarang(editBarang entity.Barang) (bool, error) {

	addQry, err := am.DB.Prepare("UPDATE barang set stok_barang=?, nama_pegawai=?  where id= ?")
	// addQry, err := am.DB.Prepare("UPDATE barang set nama_barang, stok_barang, deskripsi, nama_pegawai  where id= ?")
	if err != nil {
		log.Println("Update barang prepare", err.Error())
		return false, errors.New("prepare Edit barang error")
	}

	res, err := addQry.Exec(editBarang.Id, editBarang.Nama_Barang, editBarang.Stock, editBarang.Deskripsi, editBarang.Nama_Pegawai)
	if err != nil {
		log.Println("Update barang", err.Error())
		return false, errors.New("Update Barang error")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after Update Barang", err.Error())
		return false, errors.New("after Update Barang error")
	}

	if affRows <= 0 {
		log.Println("No record affected")
		return true, errors.New("No record")
	}

	return true, nil
}
func (am *AuthMenu) Customer(newUser entity.Customer) (bool, error) {
	// menyiapakn query untuk insert
	registerQry, err := am.DB.Prepare("INSERT INTO customer (nama_cust, nama_pegawai) values (?,?)")
	if err != nil {
		log.Println("prepare insert cust ", err.Error())
		return false, errors.New("prepare statement insert cust error")
	}

	if am.DuplicateCustomer(newUser.Nama_Customer) {
		log.Println("duplicated information")
		return false, errors.New("nama sudah digunakan")
	}

	// menjalankan query dengan parameter tertentu
	res, err := registerQry.Exec(newUser.Nama_Customer, newUser.Nama_Pegawai)
	if err != nil {
		log.Println("insert cust ", err.Error())
		return false, errors.New("insert cust error")
	}
	// Cek berapa baris yang terpengaruh query diatas
	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after insert cust ", err.Error())
		return false, errors.New("error setelah insert")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}

	return true, nil
}

func (am *AuthMenu) Barang(newBarang entity.Barang) (bool, error) {

	registerQry, err := am.DB.Prepare("INSERT INTO barang(nama_barang,stok_barang,deskripsi,nama_pegawai) VALUES (?,?,?,?)")
	if err != nil {
		log.Println("prepare insert barang ", err.Error())
		return false, errors.New("prepare statement insert barang error")
	}

	if am.DuplicateBarang(newBarang.Nama_Barang) {
		log.Println("duplicated information")
		return false, errors.New("nama barang sudah digunakan")
	}

	// menjalankan query dengan parameter tertentu
	res, err := registerQry.Exec(newBarang.Nama_Barang, newBarang.Stock, newBarang.Deskripsi, newBarang.Nama_Pegawai)
	if err != nil {
		log.Println("insert barang ", err.Error())
		return false, errors.New("insert barang error")
	}
	// Cek berapa baris yang terpengaruh query diatas
	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after insert barang ", err.Error())
		return false, errors.New("error setelah insert")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}

	return true, nil
}

func (am *AuthMenu) DeleteBarang(db *sql.DB, barang entity.Barang, Id int) (entity.Barang, error) {
	usr := db.QueryRow("SELECT id, nama_barang, stock_barang, deskripsi, nama_pegawai from barang where id=?", barang.Nama_Barang)

	var rowUser entity.Barang
	errscan := usr.Scan(&rowUser.Id, &rowUser.Nama_Barang, &rowUser.Stock)
	fmt.Println(rowUser.Id)
	var query = "DELETE FROM barang(id,nama_barang,stock_barang,nama_pegawai) VALUES (?,?,?,?)"
	statement, errPrepare := am.DB.Prepare(query)
	if errPrepare != nil {
		log.Fatal("erorr prepare delete", errPrepare.Error())

	}

	result, errExec := statement.Exec(Id, barang.Nama_Barang)
	if errExec != nil {
		log.Fatal("erorr Exec delete", errExec.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("barang yang berhasil dihapus")
		} else {
			fmt.Println("gagal dihapus")
		}
	}

	if errscan != nil {
		if errscan == sql.ErrNoRows {
			log.Fatal("error scan", errscan.Error())
		}
	}
	return rowUser, nil

}
