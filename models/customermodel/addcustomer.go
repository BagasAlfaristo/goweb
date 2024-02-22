package customermodel

//import "net/http"

import (
	"database/sql"
	"net/http"
)

/*func AddCustomer(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "admin:admin@tcp(localhost:3306)/bengkel")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tmpl := template.Must(template.ParseFiles("views/customer/add.html"))

	tmpl.Execute(w, nil)

	//log.Fatal(http.ListenAndServe(":8080", nil))
}*/

func AddHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Nomor_polisi := r.Form.Get("Nomor_polisi")
	Nama_motor := r.Form.Get("Nama_motor")
	Jenis_pesanan := r.Form.Get("Jenis_pesanan")

	// Masukkan data ke dalam database
	db, err := sql.Open("mysql", "admin:admin@tcp(localhost:3306)/bengkel")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO customer (nomor_polisi, nama_motor, jenis_pesanan) VALUES (?, ?, ?)", Nomor_polisi, Nama_motor, Jenis_pesanan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Tampilkan pesan sukses
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

/* func AddCustomer() {
	db := config.Dbconnection()
	defer db.Close()

	/*Input Data Customer
	fmt.Println("Nomor Polisi : ")
	fmt.Scanln(&nomor_polisi)

	fmt.Println("Nama Motor : ")
	fmt.Scanln(&nama_motor)

	scanner := bufio.NewScanner(os.Stdin)

	// Minta input pertama dari pengguna

	fmt.Print("Nama : ")
	if scanner.Scan() {
		Name = scanner.Text()
	} else {
		log.Fatal("Gagal membaca input pertama:", scanner.Err())
	}

	// Minta input kedua dari pengguna

	fmt.Print("Age : ")
	if scanner.Scan() {
		Age = scanner.Text()
	} else {
		log.Fatal("Gagal membaca input kedua:", scanner.Err())
	}

	insertQuery := "INSERT INTO people(Name, Age) VALUES(?, ?)"

	stmt, err := db.Prepare(insertQuery)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Err 2")
	}
	defer stmt.Close()

	// Eksekusi pernyataan dengan nilai input
	_, err = stmt.Exec(Name, Age)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Err 3")
	}

	fmt.Println("Berhasil Input Customer!")
}*/
