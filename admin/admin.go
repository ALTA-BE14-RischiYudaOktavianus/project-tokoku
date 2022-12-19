package controllers

import (
	"database/sql"
	"fmt"
	"log"
)

func Register(db *sql.DB, newUser entity.User) (sql.Result, error) {

	var query = "INSERT INTO user(nama_user,email,phone,alamat,foto_profil,kata_sandi) VALUES (?,?,?,?,?,?)"
	statement, errPrepare := db.Prepare(query)

	if errPrepare != nil {
		log.Fatal("erorr prepare insert", errPrepare.Error())

	}
	result, errExec := statement.Exec(newUser.Nama, newUser.Email, newUser.Phone, newUser.Alamat, newUser.Foto_profil, newUser.Kata_sandi)
	if errExec != nil {
		log.Fatal("erorr Exec insert", errExec.Error())
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
func LoginUser(db *sql.DB, user entity.User) (entity.User, error) {
	statm := db.QueryRow("SELECT Id_user,phone,kata_sandi FROM user WHERE phone = ? AND kata_sandi = ?", user.Phone, user.Kata_sandi)

	var row entity.User
	errs := statm.Scan(&row.Id, &row.Phone, &row.Kata_sandi)
	//  bcrypt.GenerateFromPassword(, bcrypt.DefaultCost))

	if errs != nil {
		log.Fatal("Maaf No Telfon atau Password salah ")
	}
	return row, nil
}

func Readsdata(db *sql.DB, id int) (entity.User, error) {
	res := db.QueryRow("SELECT Id_user,Nama_user,phone,alamat,foto_profil from user where id_user=?", id)
	var barisUser entity.User

	errscan := res.Scan(&barisUser.Id, &barisUser.Nama, &barisUser.Phone, &barisUser.Alamat, &barisUser.Foto_profil)

	if errscan != nil {
		if errscan == sql.ErrNoRows {
			log.Fatal("error scan", errscan.Error())
		}
	}
	return barisUser, nil
}
func UpdateUser(db *sql.DB, update entity.User) (sql.Result, error) {
	// res := db.QueryRow("SELECT Id_user,Nama_user,phone,alamat,foto_profil from user where id_user=?", id)
	// var barisUser entity.User
	var query = "UPDATE user set Nama_user = ?, email = ?, phone = ?,alamat = ? ,kata_sandi = ?  where Id_user = ?"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("erorr prepare update", errPrepare.Error())
	}
	result, errExec := statement.Exec(update.Nama, update.Email, update.Phone, update.Alamat, update.Kata_sandi, update.Id)

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

func Top_up(db *sql.DB, total entity.Top_up, Id int) (entity.User, error) {
	usr := db.QueryRow("SELECT Id_user, Nama_user,phone from user where phone=?", total.Phone)

	var rowUser entity.User
	errscan := usr.Scan(&rowUser.Id, &rowUser.Nama, &rowUser.Phone)
	fmt.Println(rowUser.Id)
	var query = "INSERT INTO topup(users_Id,Jumlah_Topup) VALUES (?,?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("erorr prepare insert", errPrepare.Error())

	}

	result, errExec := statement.Exec(Id, total.Jumlah_TopUP)
	if errExec != nil {
		log.Fatal("erorr Exec insert", errExec.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("berhasil")
		} else {
			fmt.Println("gagal")
		}
	}

	if errscan != nil {
		if errscan == sql.ErrNoRows {
			log.Fatal("error scan", errscan.Error())
		}
	}
	return rowUser, nil

}
func TfUser(db *sql.DB, total entity.Transfers, id int) (entity.User, error) {
	usr := db.QueryRow("SELECT Id_user, Nama_user,phone from user where phone=?", total.Phone)
	// usr2 := db.QueryRow("SELECT Id_user, Nama_user,phone from user where id=?", id)
	// var barisUser entity.User
	// errscan2:= usr2.Scan(&barisUser.Id,&barisUser.Nama)
	var rowUser entity.User
	errscan := usr.Scan(&rowUser.Id, &rowUser.Nama, &rowUser.Phone)
	fmt.Println(rowUser.Id)
	var query = "INSERT INTO transfers(pengirim_id,Jumlah_TF,penerima_id) VALUES (?,?,?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("erorr prepare insert", errPrepare.Error())

	}

	result, errExec := statement.Exec(id, total.Jumlah_TF, rowUser.Id)
	if errExec != nil {
		log.Fatal("erorr Exec insert", errExec.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("berhasil")
		} else {
			fmt.Println("gagal")
		}
	}

	if errscan != nil {
		if errscan == sql.ErrNoRows {
			log.Fatal("error scan", errscan.Error())
		}
	}
	return rowUser, nil

}
func History_Top_upUser(db *sql.DB, find entity.TopUp, Id int) (entity.Top_up, error) {
	usr := db.QueryRow("SELECT Id_user, Nama_user,phone from user where phone=?", find.Phone)
	var rowUser entity.Top_up

	errscan := usr.Scan(&rowUser.Id_Tp, &rowUser.Id_user, &rowUser.Phone)
	var query = "INSERT INTO topup(users_Id,Jumlah_Topup) VALUES (?,?)"
	statement, errPrepare := db.Prepare(query)

	if errPrepare != nil {
		log.Fatal("erorr prepare insert", errPrepare.Error())

	}
	result, errExec := statement.Exec(Id, find.Jumlah_TUP)
	if errExec != nil {
		log.Fatal("erorr Exec insert", errExec.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("berhasil")
		} else {
			fmt.Println("gagal")
		}
	}

	if errscan != nil {
		if errscan == sql.ErrNoRows {
			log.Fatal("error scan", errscan.Error())
		}
	}
	return rowUser, nil
}

func Search_Profil(db *sql.DB, search entity.User) (entity.User, error) {
	statm := db.QueryRow("SELECT Id_user, Nama_user, Phone, Alamat, Kata_sandi FROM user WHERE phone = ?", search.Phone)

	var row entity.User
	errs := statm.Scan(&row.Id, &row.Nama, &row.Phone, &row.Alamat, &row.Kata_sandi)
	if errs != nil {
		log.Fatal("Error Search ", errs.Error())
	}
	return row, nil
}
