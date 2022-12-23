package transaksi

import (
	"database/sql"
	"errors"
	"fmt"
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

type Nota struct {
	IdNota       int
	NamaCustomer string
	NamaPegawai  string
	NamaBarang   string
	Kuantiti     int
	TanggalTransaksi string
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

// Fungsi ini menambahkan transaksi baru ke database. Fungsi ini menerima sebuah struct Transaksi sebagai input
// dan mengembalikan nilai boolean yang menandakan sukses atau gagalnya operasi dan sebuah objek error. 
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

// Fungsi ini menambahkan jumlah barang ke dalam transaksi yang ada di database. Fungsi ini menerima sebuah struct Transaksi
// sebagai input dan mengembalikan nilai boolean yang menandakan sukses atau gagalnya operasi dan sebuah objek error.
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

// Fungsi ini menghapus sebuah transaksi dari database. Fungsi ini menerima sebuah struct Transaksi sebagai input
// dan mengembalikan nilai boolean yang menandakan sukses atau gagalnya operasi dan sebuah objek error.
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

// Fungsi ini mencetak nota transaksi dari database. Fungsi ini menerima sebuah struct Nota sebagai input
// dan mengembalikan sebuah slice dari struct Nota yang berisi data transaksi dan sebuah objek error.
func (am *AuthMenu) CetakNota(newCetak Nota) ([]Nota, error) {
	addQry, err := am.DB.Prepare(
		`SELECT c.nama_cust "Customer", p.nama_pegawai "Kasir", b.nama_barang "Barang", bht.total_qty "Jumlah", t.create_at "Tanggal Transaksi"
			FROM barang_has_transaksi bht 
			JOIN barang b on b.id = bht.barang_id
			JOIN transaksi t on t.id = bht.transaksi_id
			JOIN pegawai p on p.id = t.id_pegawai
			JOIN customer c on c.id = t.id_customer
			WHERE bht.transaksi_id = ?;`)
	if err != nil {
		log.Println("Select Cetak prepare", err.Error())
		return nil, errors.New("prepare Select Cetak error")
	}

	rows, err := addQry.Query(newCetak.IdNota)
	if err != nil {
		log.Println("Select cetak", err.Error())
		return nil, errors.New("select cetak error")
	}
	transaksi := []Nota{}
	for rows.Next() {
		trans := Nota{}
		err = rows.Scan(&trans.NamaCustomer, &trans.NamaPegawai, &trans.NamaBarang, &trans.Kuantiti, &trans.TanggalTransaksi)
		if err != nil {
			log.Println("error Loop baris untuk memasukkan data", err.Error())
			return transaksi, err
		}
		transaksi = append(transaksi, trans)
	}
	return transaksi, nil

}

// Fungsi ini menampilkan data transaksi dari database. Fungsi ini menerima sebuah integer sebagai input
// dan mengembalikan sebuah slice dari struct Transaksi yang berisi data transaksi.
func (am *AuthMenu) SearchTrans(id int) (liatTrans []Transaksi) {
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

	liatTrans = make([]Transaksi, 0)
	for rows.Next() {
		row := Transaksi{}
		rows.Scan(&row.ID, &row.Total_Qty, &row.Tanggal_Transaksi, &row.ID_Pegawai, &row.ID_Barang, &row.ID_Customer)
		strBarang += fmt.Sprintf("ID: %d %d %s (%d) (%d) <%d>\n", row.ID, row.Total_Qty, row.Tanggal_Transaksi, row.ID_Pegawai, row.ID_Barang, row.ID_Customer)
		liatTrans = append(liatTrans, row)
	}
	return liatTrans
}
