package customermodel

import (
	"goweb/config"
	"goweb/entities"
	"log"
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
		if err := rows.Scan(&customer.Id_pesanan, &customer.Tanggal, &customer.Nomor_polisi, &customer.Nama_motor, &customer.Jenis_pesanan); err != nil {
			log.Fatal(err)
		}
		customers = append(customers, customer)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return customers
}
