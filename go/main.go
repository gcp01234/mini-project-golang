/*
Mini project untuk pelatihan Golang PROA
Author: Gita Citra Puspita
Dibuat pada: 12 Agustus 2022
*/

package main

import (
	"fmt"
	"net/http"
	"html/template"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main(){

	fmt.Println("Cek.")

}

//struktur data tabel tamu
type tamu struct{
	Uuid string
	NamaLengkap string
	Domisili string
	CreatedAt string
	UpdatedAt string
}

//struktur data respon
type response struct{
	Status bool
	Pesan string
	Data []tamu
}
//fungsi untuk koneksi ke database mysql
func koneksi() (*sql.DB, error){
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/undangan")
	if err != nill{
		return nill, err
	}
	return db, nill
}

