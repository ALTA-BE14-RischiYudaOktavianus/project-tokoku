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

type Barang struct {
	Id           int
	Nama_Barang  string
	Stock        int
	Deskripsi    string
	Nama_Pegawai int
}

type Customer struct {
	Id            int
	Nama_Customer string
	Nama_Pegawai  string
}

type User_Activity struct {
	Id_User     int
	Id_activity int
}
