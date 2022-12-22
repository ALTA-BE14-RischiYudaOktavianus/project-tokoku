package admin

import (
	"database/sql"
	"errors"
	"log"
)

// Membuat struct untuk akun Admin
type Admin struct {
	ID       int
	Nama     string
	Password string
}

// Membuat struct untuk akun pegawai
type Pegawai struct {
	ID           int
	Nama_Pegawai string
	Password     string
}
type AuthMenu struct {
	DB *sql.DB
}

// Fungsi dibawah ini berguna untuk pengecekan duplicate
// apakah nama pegawai ada yang sama atau tidak
// memiliki input berupa nama yang di cek dari tabel pegawai dan output berupa boolean
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

// fungsi ini berguna untuk mendaftarkan akun pegawai ke database
// input yang diminta berupa nama pegawai dan memiliki dua output yaitu boolean dan pesan error ketika gagal
func (am *AuthMenu) Register(newUser Pegawai) (bool, error) {
	// menyiapkan query untuk insert ke database dengan parameter nama dan password
	registerQry, err := am.DB.Prepare("INSERT INTO pegawai (nama_pegawai, password) values (?,?)")
	if err != nil {
		log.Println("prepare insert user ", err.Error())
		return false, errors.New("prepare statement insert user error")
	}

	// fungsi cek duplicate nama dimasukkan
	if am.Duplicate(newUser.Nama_Pegawai) {
		log.Println("duplicated information")
		return false, errors.New("nama sudah digunakan")
	}

	// menjalankan query dengan parameter nama dan password
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

	// Jika query tidak ada yang berubah akan menampilkan pesan dibawah ini
	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}

	// mengembalikan return hasil fungsi
	return true, nil
}


// Dibawah ini merupakan fungsi untuk login
// Input yang dibutuhkan oleh fungsi ini adalah nama dan password
// Dan output berupa nama dan error
func (am *AuthMenu) Login(nama_pegawai string, password string) (Pegawai, error) {
	loginQry, err := am.DB.Prepare("SELECT id FROM pegawai WHERE nama_pegawai = ? AND password = ?")
	if err != nil {
		log.Println("prepare insert user ", err.Error())
		return Pegawai{}, errors.New("prepare statement insert user error")
	}

	// mengeksekusi query yang sudah kita prepare diatas menggunakan QueryRow
	row := loginQry.QueryRow(nama_pegawai, password)

	if row.Err() != nil {
		log.Println("login query ", row.Err().Error())
		return Pegawai{}, errors.New("tidak bisa login, data tidak ditemukan")
	}

	// copy kolom jika ada data yang sesuai menggunakan Scan
	res := Pegawai{}
	err = row.Scan(&res.ID)

	// Error handling ketika parameter query yang dimasukkan tidak sesuai atau tidak ada
	if err != nil {
		log.Println("after login query ", err.Error())
		return Pegawai{}, errors.New("tidak bisa login, kesalahan setelah error")
	}

	res.Nama_Pegawai = nama_pegawai

	// mengembalikan nilai res
	return res, nil
}
