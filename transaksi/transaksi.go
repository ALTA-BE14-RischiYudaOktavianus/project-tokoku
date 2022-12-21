package transaksi

import (
	"database/sql"
	"errors"
	"log"
)

type Transaksi struct {
	ID                int
	Total_Qty         int
	Tanggal_Transaksi string
	ID_Pegawai        int
	ID_Barang         int
	ID_Customer       int
}

type AuthMenu struct {
	DB *sql.DB
}

// func (am *AuthMenu) Stock(kurangQty  string) bool {
// 	var kuantitas Transaksi
// 	var transfer barang.Barang
// 	if kuantitas.Total_Qty > transfer.Stock {
// 		fmt.Println("Stock Anda Tidak Mencukupi")
// 	}

// 	var terima Transaksi
// 	Stock := transfer.Stock - kuantitas.Total_Qty
// 	Stock2 := terima.Total_Qty + kuantitas.Total_Qty

// 	ress, err := addQry.Exec(Stock, transfer.Stock)
// 	if err != nil {
// 		log.Println("Gagal line 69", err.Error())
// 	}
// 	affRows, err := ress.RowsAffected()

// 	_, err = addQry.Exec(Stock2, terima.Total_Qty)
// 	if err != nil {
// 		log.Println("Gagal line 74", err.Error())
// 	}
// 	res := am.DB.QueryRow("SELECT id FROM pegawai where nama_pegawai = ?", name)

// 	var idExist int
// 	err := res.Scan(&idExist)
// 	if err != nil {
// 		log.Println("Result scan error", err.Error())
// 		return false
// 	}
// 	return true

// }

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

func (am *AuthMenu) CetakNota(newCetak Transaksi) (bool, error) {
	addQry, err := am.DB.Prepare("SELECT barang_has_transaksi.total_qty, barang_has_transaksi.barang_id,  transaksi.id, customer.nama_cust as pelanggan, transaksi.tanggal_transaksi,transaksi.id_pegawai,p.nama_pegawai as pegawai FROM transaksi INNER JOIN customer on customer.id = transaksi.id_customer INNER JOIN pegawai p on transaksi.id_pegawai = p.id Left join barang_has_transaksi on barang_has_transaksi.transaksi_id = transaksi.id WHERE transaksi.id_customer = ?")
	if err != nil {
		log.Println("Select Cetak prepare", err.Error())
		return false, errors.New("prepare Select Cetak error")
	}

	res, err := addQry.Exec(newCetak.ID_Customer)
	if err != nil {
		log.Println("Select cetak", err.Error())
		return false, errors.New("Select Cetak error")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after Select Cetak", err.Error())
		return false, errors.New("after Select cetak error")
	}

	if affRows <= 0 {
		log.Println("No record affected")
		return true, errors.New("No record")
	}

	return true, nil
}
