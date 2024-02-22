package customermodel

import (
	"goweb/config"
	"goweb/entities"
	"log"
	"time"
)

func GetAll() []entities.Customer {

	rows, err := config.DB.Query("SELECT id_pesanan, tanggal, nomor_polisi, nama_motor, jenis_pesanan FROM customer")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var customers []entities.Customer

	for rows.Next() {
		var customer entities.Customer
		var tanggalDB []uint8 // Menggunakan []uint8 untuk menampung tanggal dari database

		if err := rows.Scan(&customer.Id_pesanan, &tanggalDB, &customer.Nomor_polisi, &customer.Nama_motor, &customer.Jenis_pesanan); err != nil {
			log.Fatal(err)
		}

		// Konversi dari []uint8 ke string
		tanggalString := string(tanggalDB)

		// Konversi dari string ke time.Time
		tanggal, err := time.Parse("2006-01-02 15:04:05", tanggalString)
		if err != nil {
			log.Fatal(err)
		}

		customer.Tanggal = tanggal

		customers = append(customers, customer)
	}

	return customers
}
