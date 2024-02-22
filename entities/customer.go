package entities

import "time"

type Customer struct {
	Id_pesanan    int
	Tanggal       time.Time
	Nomor_polisi  string
	Nama_motor    string
	Jenis_pesanan string
}
