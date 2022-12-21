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
total_qty int DEFAULT NULL,
tanggal_transaksi datetime DEFAULT NULL,
id_pegawai int,
id_barang int,
id_customer int,
create_at datetime DEFAULT CURRENT_TIMESTAMP,
update_at datetime DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`),
CONSTRAINT fk_nama FOREIGN KEY(id_pegawai) REFERENCES pegawai(id),
CONSTRAINT FK_nama_cust FOREIGN KEY(id_customer) REFERENCES customer(id),
CONSTRAINT FK_nama_barang FOREIGN KEY(id_barang) REFERENCES barang(id)
);

CREATE TABLE IF NOT EXISTS `tokoku`.`barang_has_transaksi` (
  `barang_id` INT NOT NULL,
  `transaksi_id` INT NOT NULL,
  `total_qty` INT NOT NULL,
  PRIMARY KEY (`barang_id`, `transaksi_id`),
  INDEX `fk_barang_has_transaksi_transaksi1_idx` (`transaksi_id` ASC) VISIBLE,
  INDEX `fk_barang_has_transaksi_barang1_idx` (`barang_id` ASC) VISIBLE,
  CONSTRAINT `fk_barang_has_transaksi_barang1`
    FOREIGN KEY (`barang_id`)
    REFERENCES `tokoku`.`barang` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_barang_has_transaksi_transaksi1`
    FOREIGN KEY (`transaksi_id`)
    REFERENCES `tokoku`.`transaksi` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;

INSERT INTO pegawai(nama_pegawai, password)
Values("admin", "admin");
