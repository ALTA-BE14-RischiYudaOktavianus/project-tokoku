use tokoku;
drop table transaksi;
CREATE TABLE pegawai (
id int NOT NULL AUTO_INCREMENT,
nama_pegawai varchar(100) NOT NULL,
password varchar(255) DEFAULT NULL,
create_at datetime DEFAULT CURRENT_TIMESTAMP,
update_at datetime DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`)
);

CREATE TABLE customer (
id int NOT NULL AUTO_INCREMENT,
nama_cust varchar(100) NOT NULL,
id_pegawai int,
create_at datetime DEFAULT CURRENT_TIMESTAMP,
update_at datetime DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`),
CONSTRAINT FK_id_pegawai FOREIGN KEY(id_pegawai) REFERENCES pegawai(id)
);

CREATE TABLE barang (
id int NOT NULL AUTO_INCREMENT,
nama_barang varchar(100) NOT NULL,
stok_barang int(15),
deskripsi varchar(255) NOT NULL,
id_pegawai int,
create_at datetime DEFAULT CURRENT_TIMESTAMP,
update_at datetime DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`),
CONSTRAINT fk_pegawai FOREIGN KEY(id_pegawai) REFERENCES pegawai(id)
);

CREATE TABLE transaksi (
id int NOT NULL AUTO_INCREMENT,
tanggal_transaksi datetime DEFAULT CURRENT_TIMESTAMP,
id_pegawai int,
id_customer int,
create_at datetime DEFAULT CURRENT_TIMESTAMP,
update_at datetime DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`),
CONSTRAINT fk_nama FOREIGN KEY(id_pegawai) REFERENCES pegawai(id),
CONSTRAINT FK_nama_cust FOREIGN KEY(id_customer) REFERENCES customer(id)
);

CREATE TABLE `barang_has_transaksi` (
  `barang_id` int NOT NULL,
  `transaksi_id` int NOT NULL,
  `total_qty` int NOT NULL,
  PRIMARY KEY (`barang_id`,`transaksi_id`),
  KEY `fk_barang_has_transaksi_transaksi1_idx` (`transaksi_id`),
  KEY `fk_barang_has_transaksi_barang1_idx` (`barang_id`),
  CONSTRAINT `fk_barang_has_transaksi_barang1` FOREIGN KEY (`barang_id`) REFERENCES `barang` (`id`),
  CONSTRAINT `fk_barang_has_transaksi_transaksi1` FOREIGN KEY (`transaksi_id`) REFERENCES `transaksi` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO pegawai(nama_pegawai, password)
Values("admin", "admin");

SELECT t.id_pegawai "kasir", t.id, t.nama_cust "Pelanggan"
FROM barang_has_transaksi bht
JOIN transaksi t on t.tanggal_transaksi = bht.transaksi_id
JOIN barang b on b.nama_barang = bht.transaksi_id
JOIN customer c on c.nama_cust = bht.transaksi_id;

SELECT t.id as "Transaksi yang dilakukan", c.nama_cust "Pelanggan", b.nama_barang "barang yang beli", h.total_qty "kuantitas", p.nama_pegawai "kasir", t.tanggal_transaksi "tanggal"
FROM customer c, transaksi t, barang b, barang_has_transaksi h, pegawai p
join barang_has_transaksi h on h.transaksi_id = t.id
join barang_has_transaksi h on h.barang_id = b.nama_barang
join transaksi t on t.id_pegawai = p.nama_pegawai
join transaksi t on t.id_customer = c.nama_cust
WHERE h = t.id =c.id = p.id = b.id;

SELECT * FROM tokoku.barang_has_transaksi;
SELECT t.id as "Transaksi yang dilakukan", c.nama_cust "Pelanggan", b.nama_barang "barang yang beli", bht.total_qty "kuantitas", p.nama_pegawai "kasir" 
FROM transaksi t
join transaksi t on t.id_pegawai = p.nama_pegawai
join transaksi t on t.id_customer = c.nama_cust
WHERE t.id = 1;

SELECT pegawai.id, customer.id ,transaksi.tanggal_transaksi FROM transaksi
INNER JOIN pegawai on  pegawai.nama_pegawai =   transaksi.id_pegawai
INNER JOIN customer on  customer.nama_cust =   transaksi.id_customer
where transaksi.id =1
union
SELECT barang.id, transaksi.id , barang_has_transaksi.total_qty FROM barang_has_transaksi
INNER JOIN barang on  barang.nama_barang =   barang_has_transaksi.barang_id
INNER JOIN transaksi on  transaksi.id =  barang_has_transaksi.transaksi_id
where barang_id = transaksi.id;

SELECT pegawai.id, pegawai.nama_pegawai,  customer.id, customer.nama_cust, transaksi.tanggal_transaksi FROM transaksi
INNER JOIN customer on customer.id = transaksi.id_customer
Left join pegawai on pegawai.id = transaksi.id_pegawai
where transaksi.id=1;

SELECT barang_has_transaksi.total_qty, barang_has_transaksi.barang_id,  transaksi.id, customer.nama_cust as pelanggan, transaksi.tanggal_transaksi,transaksi.id_pegawai,p.nama_pegawai as pegawai
FROM transaksi
INNER JOIN customer on customer.id = transaksi.id_customer
INNER JOIN pegawai p on transaksi.id_pegawai = p.id
Left join barang_has_transaksi on barang_has_transaksi.transaksi_id = transaksi.id
WHERE transaksi.id_customer = 1;

SELECT barang_has_transaksi.total_qty as jumlah_barang, barang_has_transaksi.barang_id,  transaksi.id as no_transaksi, customer.nama_cust as pelanggan, transaksi.tanggal_transaksi,transaksi.id_pegawai,p.nama_pegawai as pegawai
FROM transaksi
INNER JOIN customer on customer.id = transaksi.id_customer
INNER JOIN pegawai p on transaksi.id_pegawai = p.id
Left join barang_has_transaksi on barang_has_transaksi.transaksi_id = transaksi.id
WHERE transaksi.id_customer = 1;

SELECT b.id ,b.id_pegawai ,b.nama_barang "Nama Barang" ,b.stok_barang ,b.deskripsi,p.nama_pegawai 'Nama Pegawai'
FROM barang b
JOIN pegawai p ON p.id = b.id_pegawai;
SELECT b.id ,b.id_pegawai ,b.nama_barang "Nama Barang" ,b.stok_barang ,b.deskripsi ,p.nama_pegawai 'Nama Pegawai'
FROM barang b, pegawai p
WHERE b.id = 1;