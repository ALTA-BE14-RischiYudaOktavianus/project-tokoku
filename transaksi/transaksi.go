package transaksi

import (
	"database/sql"
	"errors"
	"log"
)

type Transaksi struct {
	ID                int
	Total_Qty         string
	Tanggal_Transaksi string
	ID_Pegawai        string
	ID_Barang         string
	ID_Customer       string
}

type AuthMenu struct {
	DB *sql.DB
}

func (am *AuthMenu) AddTransaksi(newTransaksi Transaksi) (bool, error) {
	addQry, err := am.DB.Prepare("INSERT into transaksi (id_pegawai, id_customer) VALUES (?, ?)")
	if err != nil {
		log.Println("Insert transaksi prepare", err.Error())
		return false, errors.New("prepare Insert transaksi error")
	}

	res, err := addQry.Exec(newTransaksi.ID_Pegawai, newTransaksi.ID_Customer)
	if err != nil {
		log.Println("insert transaksi", err.Error())
		return false, errors.New("Insert transaksi error")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after insert transaksi", err.Error())
		return false, errors.New("after Insert transaksi error")
	}

	if affRows <= 0 {
		log.Println("No record affected")
		return true, errors.New("No record")
	}

	return true, nil
}
func (am *AuthMenu) AddQTY(newQty Transaksi) (bool, error) {
	addQry, err := am.DB.Prepare("INSERT into barang_has_transaksi (barang_id, transaksi_id, total_qty) VALUES (?, ?, ?)")
	if err != nil {
		log.Println("Insert QTY prepare", err.Error())
		return false, errors.New("prepare Insert QTY error")
	}

	res, err := addQry.Exec(newQty.ID_Barang, newQty.ID, newQty.Total_Qty)
	if err != nil {
		log.Println("insert QTY", err.Error())
		return false, errors.New("Insert QTY error")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after insert QTY", err.Error())
		return false, errors.New("after Insert QTY error")
	}

	if affRows <= 0 {
		log.Println("No record affected")
		return true, errors.New("No record")
	}

	return true, nil
}
func (am *AuthMenu) DeleteTransaksi(deleteTransaksi Transaksi) (bool, error) {

	registerQry, err := am.DB.Prepare("DELETE FROM transaksi WHERE id=?")
	if err != nil {
		log.Println("prepare delete transaksi ", err.Error())
		return false, errors.New("prepare statement delete transaksi error")
	}

	// menjalankan query dengan parameter tertentu
	res, err := registerQry.Exec(deleteTransaksi.ID)
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
