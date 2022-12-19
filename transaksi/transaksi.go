package transaksi

import (
	"database/sql"
	"errors"
	"log"
)

type Transaksi struct {
	ID       int
	Title    string
	Location string
}

type TransaksiMenu struct {
	DB *sql.DB
}

func (tm *TransaksiMenu) AddTransaksi(newTransaksi Transaksi) (bool, error) {
	addQry, err := tm.DB.Prepare("INSERT into activity (id_user, title, location, create_at) VALUES (?, ?, ?, now());")
	if err != nil {
		log.Println("Insert activity prepare", err.Error())
		return false, errors.New("prepare Insert activity error")
	}

	res, err := addQry.Exec(newTransaksi.ID, newTransaksi.Title, newTransaksi.Location)
	if err != nil {
		log.Println("insert activity", err.Error())
		return false, errors.New("Insert activity error")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after insert activity", err.Error())
		return false, errors.New("after Insert activity error")
	}

	if affRows <= 0 {
		log.Println("No record affected")
		return true, errors.New("No record")
	}

	return true, nil
}
