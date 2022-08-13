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

//fungsi untuk menampilkan semua data tamu
func tampil (pesan string) response{
	db, err := koneksi()
	if err != nill {
		return response{
			Status: false,
			Pesan: "Gagal koneksi: "+err.Error(),
			Data: []tamu{}
		}
	}
	defer db.Close()
	dataTamu, err := db.Query("SELECT * FROM `tamu`")
	if err != nill {
		return response{
			Status: false,
			Pesan: "Query error: "+err.Error(),
			Data: []tamu{}
		}
	}
	defer dataTamu.Close()
	var hasil []tamu
	for dataTamu.Next(){
		var tm = tamu{}
		var err = dataTamu.Scan(&tm.Uuid,&tm.NamaLengkap,&tm.Domisili,&tm.CreatedAt,&tm.UpdatedAt)
		if err != nill {
			return response{
				Status: false,
				Pesan: "Gagal baca data: "+err.Error(),
				Data: []tamu{}
			}
		}
		hasil = append(hasil, tm)
	}
	return response{
		Status: true,
		Pesan: pesan,
		Data: hasil
	}
}


//fungsi untuk menampilkan data tamu berdasarkan Uuid
func tampilFilterBerdasarkanUuid (pesan string, uuid string) response{
	db, err := koneksi()
	if err != nill {
		return response{
			Status: false,
			Pesan: "Gagal koneksi: "+err.Error(),
			Data: []tamu{}
		}
	}
	defer db.Close()
	dataTamu, err := db.Query("SELECT * FROM `tamu` WHERE Uuid=?",uuid)
	if err != nill {
		return response{
			Status: false,
			Pesan: "Query error: "+err.Error(),
			Data: []tamu{}
		}
	}
	defer dataTamu.Close()
	var hasil []tamu
	for dataTamu.Next(){
		var tm = tamu{}
		var err = dataTamu.Scan(&tm.Uuid,&tm.NamaLengkap,&tm.Domisili,&tm.CreatedAt,&tm.UpdatedAt)
		if err != nill {
			return response{
				Status: false,
				Pesan: "Gagal baca data tamu dengan Uuid "+uuid+ " :"+err.Error(),
				Data: []tamu{}
			}
		}
		hasil = append(hasil, tm)
	}
	return response{
		Status: true,
		Pesan: pesan,
		Data: hasil
	}
	
}

//fungsi untuk menambahkan data tamu
func tambah (uuid string, namaLengkap string, domisili string, createdAt string) response{
	db, err := koneksi()
	if err != nill {
		return response{
			Status: false,
			Pesan: "Gagal koneksi: "+err.Error(),
			Data: []tamu{}
		}
	}
	defer db.Close()
	_, err := db.Query("INSERT INTO `tamu`(`uuid`, `nama_lengkap`, `domisili`, `created_at`, `updated_at`) VALUES (?,?,?,?)",uuid, namaLengkap,domisili,createdAt)
	if err != nill {
		return response{
			Status: false,
			Pesan: "Query insert error: "+err.Error(),
			Data: []tamu{}
		}
	}
	return response{
		Status: true,
		Pesan: "Berhasil tambah data tamu "+namaLengkap,
		Data: []tamu{}
	}
}

//fungsi untuk mengubah data tamu
func ubah (uuid string, namaLengkap string, domisili string, updatedAt string) response{
	db, err := koneksi()
	if err != nill {
		return response{
			Status: false,
			Pesan: "Gagal koneksi: "+err.Error(),
			Data: []tamu{}
		}
	}
	defer db.Close()
	_, err := db.Query("UPDATE `tamu` SET `nama_lengkap`=?,`domisili`=?,`updated_at`=? WHERE uuid=?", namaLengkap,domisili,updatedAt,uuid)
	if err != nill {
		return response{
			Status: false,
			Pesan: "Query update error: "+err.Error(),
			Data: []tamu{}
		}
	}
	return response{
		Status: true,
		Pesan: "Berhasil ubah data tamu "+uuid,
		Data: []tamu{}
	}
}