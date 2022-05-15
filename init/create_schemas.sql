 CREATE TABLE user(
    id int UNSIGNED AUTO_INCREMENT primary key,
    nama varchar(40),
    email  varchar(50),
    password varchar(50),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
    );
CREATE TABLE siswa(
    id int UNSIGNED AUTO_INCREMENT primary key,
    nama varchar(40),
    ttl  varchar(30),
    alamat varchar(50),
    no_wali varchar(15),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
    );

CREATE TABLE kelas(
    id int UNSIGNED AUTO_INCREMENT primary key,
    wali varchar(40),
    nama  varchar(20),
    tahun_ajaran varchar(20),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
    );

CREATE TABLE nilai_raport(
    id int UNSIGNED AUTO_INCREMENT primary key,
    id_siswa   int UNSIGNED,
    id_kelas   int UNSIGNED,
    IPA        int,           
	IPS        int,            
	MTK        int,            
	PKN        int,            
	B_INDO     int,            
	B_INGGRIS  int,           
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
    );
ALTER TABLE nilai_raport ADD FOREIGN KEY (id_siswa) REFERENCES siswa(id);

ALTER TABLE nilai_raport ADD FOREIGN KEY (id_kelas) REFERENCES kelas(id);