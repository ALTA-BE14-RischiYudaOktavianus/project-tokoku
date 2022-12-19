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

func (am *AuthMenu) UpdateBarang(db *sql.DB, update entity.Barang) (sql.Result, error) {
	// res := db.QueryRow("SELECT Id_user,Nama_user,phone,alamat,foto_profil from user where id_user=?", id)
	// var barisUser entity.User
	var query = "UPDATE barang set nama_barang = ?, stock_barang = ?, deskripsi = ?, nama_pegawai = ?  where id_barang = ?"
	statement, errPrepare := am.DB.Prepare(query)
	if errPrepare != nil {
		log.Fatal("erorr prepare update", errPrepare.Error())
	}
	result, errExec := statement.Exec(update.Nama_Barang, update.Stock, update.Deskripsi, update.Nama_Pegawai, update.Id)

	if errExec != nil {
		log.Fatal("erorr Exec update", errExec.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("berhasil")
		} else {
			fmt.Println("gagal")
		}
	}
	return result, nil
}

func (am *AuthMenu) Customer(db *sql.DB, cust entity.Customer, Id int) (entity.Customer, error) {
	usr := db.QueryRow("SELECT id, nama_cust,nama_pegawai from customer where id=?", cust.Id)

	var rowUser entity.Customer
	errscan := usr.Scan(&rowUser.Id, &rowUser.Nama_Customer, &rowUser.Id)
	fmt.Println(rowUser.Id)
	var query = "INSERT INTO customer(nama_cust) VALUES (?)"
	statement, errPrepare := am.DB.Prepare(query)
	if errPrepare != nil {
		log.Fatal("erorr prepare insert", errPrepare.Error())

	}

	result, errExec := statement.Exec(Id, cust.Nama_Customer)
	if errExec != nil {
		log.Fatal("erorr Exec insert", errExec.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("data customer berhasil ditambahkan")
		} else {
			fmt.Println("data gagal ditambahkan")
		}
	}

	if errscan != nil {
		if errscan == sql.ErrNoRows {
			log.Fatal("error scan", errscan.Error())
		}
	}
	return rowUser, nil

}

func (am *AuthMenu) Barang(db *sql.DB, barang entity.Barang, Id int) (entity.Barang, error) {
	usr := db.QueryRow("SELECT id, nama_barang, stock_barang, deskripsi, nama_pegawai from barang where id=?", barang.Nama_Barang)

	var rowUser entity.Barang
	errscan := usr.Scan(&rowUser.Id, &rowUser.Nama_Barang, &rowUser.Stock)
	fmt.Println(rowUser.Id)
	var query = "INSERT INTO barang(id,nama_barang,stock_barang,nama_pegawai) VALUES (?,?,?,?)"
	statement, errPrepare := am.DB.Prepare(query)
	if errPrepare != nil {
		log.Fatal("erorr prepare insert", errPrepare.Error())

	}

	result, errExec := statement.Exec(Id, barang.Nama_Barang)
	if errExec != nil {
		log.Fatal("erorr Exec insert", errExec.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("barang yang diinputkan berhasil ditambahkan")
		} else {
			fmt.Println("penginputan anda gagal")
		}
	}

	if errscan != nil {
		if errscan == sql.ErrNoRows {
			log.Fatal("error scan", errscan.Error())
		}
	}
	return rowUser, nil

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
