use tokoku;
drop table barang;
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
nama_pegawai int,
create_at datetime DEFAULT CURRENT_TIMESTAMP,
update_at datetime DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`),
CONSTRAINT FK_nama_pegawai FOREIGN KEY(nama_pegawai) REFERENCES pegawai(id)
);

CREATE TABLE barang (
id int NOT NULL AUTO_INCREMENT,
nama_barang varchar(100) NOT NULL,
stok_barang int(15),
deskripsi varchar(255) NOT NULL,
nama_pegawai int,
create_at datetime DEFAULT CURRENT_TIMESTAMP,
update_at datetime DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`),
CONSTRAINT fk_pegawai FOREIGN KEY(nama_pegawai) REFERENCES pegawai(id)
);

CREATE TABLE transaksi (
id int NOT NULL AUTO_INCREMENT,
total_qty int DEFAULT NULL,
tanggal_transaksi datetime DEFAULT NULL,
nama_pegawai int,
nama_barang int,
nama_customer int,
create_at datetime DEFAULT CURRENT_TIMESTAMP,
update_at datetime DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`),
CONSTRAINT fk_nama FOREIGN KEY(nama_pegawai) REFERENCES pegawai(id),
CONSTRAINT FK_nama_cust FOREIGN KEY(nama_customer) REFERENCES customer(id),
CONSTRAINT FK_nama_barang FOREIGN KEY(nama_barang) REFERENCES barang(id)
);
