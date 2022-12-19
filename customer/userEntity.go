package entity

type User struct {
	Id          int
	Nama        string
	Email       string
	Phone       string
	Alamat      string
	Foto_profil string
	Kata_sandi  string
}
type Top_up struct {
	Id_Tp        int
	Id_user      string
	Phone        string
	Jumlah_TopUP int
}

type TopUp struct {
	Phone      string
	Jumlah_TUP int
}
type Transfer struct {
	Id_Tf       int
	id_penerima int
	Jumlah_TF   int
	id_pengirim int
}

type Transfers struct {
	Jumlah_TF int
	Phone     string
}
