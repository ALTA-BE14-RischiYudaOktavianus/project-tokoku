package customer

import (
	"database/sql"
	"errors"
	"log"
)

type Customer struct {
	Id            int
	Nama_Customer string
	Nama_Pegawai  string
}
type AuthMenu struct {
	DB *sql.DB
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

func (am *AuthMenu) Customer(newUser Customer) (bool, error) {
	// menyiapakn query untuk insert
	registerQry, err := am.DB.Prepare("INSERT INTO customer (nama_cust, id_pegawai) values (?,?)")
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

func (am *AuthMenu) DeleteCustomer(deleteCustomer Customer) (bool, error) {

	registerQry, err := am.DB.Prepare("DELETE FROM customer WHERE id=?")
	if err != nil {
		log.Println("prepare delete customer ", err.Error())
		return false, errors.New("prepare statement delete customer error")
	}

	if am.DuplicateCustomer(deleteCustomer.Nama_Customer) {
		log.Println("duplicated information")
		return false, errors.New("nama customer sudah tidak digunakan")
	}

	// menjalankan query dengan parameter tertentu
	res, err := registerQry.Exec(deleteCustomer.Id)
	if err != nil {
		log.Println("delete customer ", err.Error())
		return false, errors.New("delete customer error")
	}
	// Cek berapa baris yang terpengaruh query diatas
	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after Delete customer ", err.Error())
		return false, errors.New("error setelah Delete")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}

	return true, nil
}
