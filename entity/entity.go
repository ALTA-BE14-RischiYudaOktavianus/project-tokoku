package entity

type Pegawai struct {
	ID       int
	Username string
	Password string
}

type Admin struct {
	ID       int
	Username string
	Password string
}

type Transaksi struct {
	Id     int
	Title  string
	Uraian string
}
