package pegawai

import (
	"database/sql"
	"errors"
	"log"
)

type Pegawai struct {
	ID       int
	Username string
	Password string
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
	registerQry, err := am.DB.Prepare("INSERT INTO pegawai (username, password) values (?,?)")
	if err != nil {
		log.Println("prepare insert user ", err.Error())
		return false, errors.New("prepare statement insert user error")
	}

	if am.Duplicate(newUser.Username) {
		log.Println("duplicated information")
		return false, errors.New("nama sudah digunakan")
	}

	// menjalankan query dengan parameter tertentu
	res, err := registerQry.Exec(newUser.Username, newUser.Password)
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
func (am *AuthMenu) Ceklogin(name string) bool {
	res := am.DB.QueryRow("SELECT id FROM pegawai where username = ?", name)
	var idExist int
	err := res.Scan(&idExist)
	if err != nil {
		log.Println("Result scan error", err.Error())
		return false
	}
	return true
}

func (am *AuthMenu) Login(newUser Pegawai) (bool, int, error) {
	res := am.DB.QueryRow("SELECT id FROM pegawai where username = ? and password = ?", newUser.Username, newUser.Password)
	var idExist int
	err := res.Scan(&idExist)
	if err != nil {
		if err.Error() != "sql error" {
			log.Println("Result scan error", err.Error())
			return false, 0, errors.New("result scan errorz")
		}
	}
	if idExist > 0 {
		return true, idExist, nil
	}
	return false, idExist, errors.New("username and password is salah")
}
