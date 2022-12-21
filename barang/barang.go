package barang

import (
	"database/sql"
	"errors"
	"log"
)

type Barang struct {
	Id           int
	Nama_Barang  string
	Stock        int
	Deskripsi    string
	Nama_Pegawai int
}

type AuthMenu struct {
	DB *sql.DB
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

func (am *AuthMenu) EditBarang(editBarang Barang) (bool, error) {

	addQry, err := am.DB.Prepare("UPDATE barang set deskripsi=?  where id= ?")
	// addQry, err := am.DB.Prepare("UPDATE barang set nama_barang, stok_barang, deskripsi, nama_pegawai  where id= ?")
	if err != nil {
		log.Println("Update barang prepare", err.Error())
		return false, errors.New("prepare Edit barang error")
	}

	res, err := addQry.Exec(editBarang.Deskripsi, editBarang.Id)
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
func (am *AuthMenu) UpdateBarang(editBarang Barang) (bool, error) {

	addQry, err := am.DB.Prepare("UPDATE barang set stok_barang=? where id= ?")
	// addQry, err := am.DB.Prepare("UPDATE barang set nama_barang, stok_barang, deskripsi, nama_pegawai  where id= ?")
	if err != nil {
		log.Println("Update barang prepare", err.Error())
		return false, errors.New("prepare Edit barang error")
	}

	res, err := addQry.Exec(editBarang.Stock, editBarang.Id)
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
func (am *AuthMenu) Barang(newBarang Barang) (bool, error) {

	registerQry, err := am.DB.Prepare("INSERT INTO barang(nama_barang,stok_barang,deskripsi,id_pegawai) VALUES (?,?,?,?)")
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

func (am *AuthMenu) Deletebarang(deleteBarang Barang) (bool, error) {

	registerQry, err := am.DB.Prepare("DELETE FROM barang WHERE id=?")
	if err != nil {
		log.Println("prepare delete barang ", err.Error())
		return false, errors.New("prepare statement delete barang error")
	}

	if am.DuplicateBarang(deleteBarang.Nama_Barang) {
		log.Println("duplicated information")
		return false, errors.New("nama barang sudah tidak digunakan")
	}

	// menjalankan query dengan parameter tertentu
	res, err := registerQry.Exec(deleteBarang.Id)
	if err != nil {
		log.Println("delete barang ", err.Error())
		return false, errors.New("delete barang error")
	}
	// Cek berapa baris yang terpengaruh query diatas
	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after Delete barang ", err.Error())
		return false, errors.New("error setelah Delete")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}

	return true, nil
}
