package barang

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type Barang struct {
	Id           int
	Nama_Barang  string
	Stock        int
	Deskripsi    string
	Nama_Pegawai string
}

type AuthMenu struct {
	DB *sql.DB
}

// Fungsi pengecekan apakah ada barang yang sama atau tidak
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

// Fungsi untuk mengedit barang
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

// Fungsi untuk update stock barang
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

// Fungsi untuk menambahkan barang
func (am *AuthMenu) AddBarang(newBarang Barang) (bool, error) {

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

// Fungsi delete barang di database
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

// Fungsi mencari barang di sistem
func (am *AuthMenu) SearchBarang(id int) (liatBarang []Barang) {
	var strBarang string
	rows, e := am.DB.Query(
		`SELECT id,
		nama_barang,
		stok_barang, deskripsi, id_pegawai
		FROM barang;`)

	if e != nil {
		log.Println(e)
		return
	}

	liatBarang = make([]Barang, 0)
	for rows.Next() {
		row := Barang{}
		rows.Scan(&row.Id, &row.Nama_Barang, &row.Stock, &row.Deskripsi, &row.Nama_Pegawai)
		strBarang += fmt.Sprintf("ID: %d %s (%d) (%s) <%s>\n", row.Id, row.Nama_Barang, row.Stock, row.Deskripsi, row.Nama_Pegawai)
		liatBarang = append(liatBarang, row)
	}
	return liatBarang
}

func (am *AuthMenu) DisplayBarang() ([]Barang, error) {
	var strBarang string
	rows, e := am.DB.Query(
		`SELECT b.id "ID Barang", b.nama_barang "Nama barang", b.stok_barang "Stock Barang", b.deskripsi "Deskripsi Barang", p.nama_pegawai "Nama Pegawai"
		FROM barang b
		JOIN pegawai p on b.id_pegawai = p.id;`)

	lihatBarang := make([]Barang, 0)

	if e != nil {
		log.Println(e)
		return lihatBarang, e
	}

	for rows.Next() {
		row := Barang{}
		rows.Scan(&row.Id, &row.Nama_Barang, &row.Stock, &row.Deskripsi, &row.Nama_Pegawai)
		strBarang += fmt.Sprintf("ID: %d %s (%d) (%s) <%s>\n", row.Id, row.Nama_Barang, row.Stock, row.Deskripsi, row.Nama_Pegawai)
		lihatBarang = append(lihatBarang, row)
	}
	return lihatBarang, nil
}

func (am *AuthMenu) Data(id int) ([]Barang, string, error) {
	var (
		selectBarangQry *sql.Rows
		err             error
		strBarang       string
	)
	if id == 0 {
		selectBarangQry, err = am.DB.Query(`
	 SELECT b.id ,b.id_pegawai ,b.nama_barang "Nama Barang" ,b.stok_barang ,b.deskripsi,p.nama_pegawai 'Nama Pegawai'
	 FROM barang b
	 JOIN pegawai p ON p.id = b.id_pegawai;`)
	} else {
		selectBarangQry, err = am.DB.Query(`
		SELECT b.id ,b.id_pegawai ,b.nama_barang "Nama Barang" ,b.stok_barang ,b.deskripsi ,p.nama_pegawai 'Nama Pegawai'
		FROM barang b, pegawai p
		WHERE b.id = ?;`, id)
	}

	if err != nil {
		log.Println("select barang", err.Error())
		return nil, strBarang, errors.New("select barang error")
	}
	arrBarang := []Barang{}
	for selectBarangQry.Next() {
		var tmp Barang
		err = selectBarangQry.Scan(&tmp.Id, &tmp.Nama_Barang, &tmp.Stock, &tmp.Deskripsi, &tmp.Nama_Pegawai)
		if err != nil {
			log.Println("Loop through rows, using Scan to assign column data to struct fields", err.Error())
			return arrBarang, strBarang, err
		}
		strBarang += fmt.Sprintf("ID: %d %s (%d) (%s) <%s>\n", tmp.Id, tmp.Nama_Barang, tmp.Stock, tmp.Deskripsi, tmp.Nama_Pegawai)
		arrBarang = append(arrBarang, tmp)
	}
	return arrBarang, strBarang, nil
}
