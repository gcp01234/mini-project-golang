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
	if err != nil{
		return nil, err
	}
	return db, nil
}

//fungsi untuk menampilkan semua data tamu
func tampil (pesan string) response{
	db, err := koneksi()
	if err != nil {
		return response{
			Status: false,
			Pesan: "Gagal koneksi: "+err.Error(),
			Data: []tamu{}
		}
	}
	defer db.Close()
	dataTamu, err := db.Query("SELECT * FROM `tamu`")
	if err != nil {
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
		if err != nil {
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
func tampilFilterBerdasarkanUuid (uuid string) response{
	db, err := koneksi()
	if err != nil {
		return response{
			Status: false,
			Pesan: "Gagal koneksi: "+err.Error(),
			Data: []tamu{}
		}
	}
	defer db.Close()
	dataTamu, err := db.Query("SELECT * FROM `tamu` WHERE Uuid=?",uuid)
	if err != nil {
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
		if err != nil {
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
		Pesan: "Berhasil tampilkan data tamu!",
		Data: hasil
	}
	
}

//fungsi untuk menambahkan data tamu
func tambah (uuid string, namaLengkap string, domisili string, createdAt string) response{
	db, err := koneksi()
	if err != nil {
		return response{
			Status: false,
			Pesan: "Gagal koneksi: "+err.Error(),
			Data: []tamu{}
		}
	}
	defer db.Close()
	_, err := db.Query("INSERT INTO `tamu`(`uuid`, `nama_lengkap`, `domisili`, `created_at`, `updated_at`) VALUES (?,?,?,?)",uuid, namaLengkap,domisili,createdAt)
	if err != nil {
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
	if err != nil {
		return response{
			Status: false,
			Pesan: "Gagal koneksi: "+err.Error(),
			Data: []tamu{}
		}
	}
	defer db.Close()
	_, err := db.Query("UPDATE `tamu` SET `nama_lengkap`=?,`domisili`=?,`updated_at`=? WHERE uuid=?", namaLengkap,domisili,updatedAt,uuid)
	if err != nil {
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

func hapus (uuid string) response{
	db, err := koneksi()
	if err != nil {
		return response{
			Status: false,
			Pesan: "Gagal koneksi: "+err.Error(),
			Data: []tamu{}
		}
	}
	defer db.Close()
	_, err := db.Query("DELETE FROM `tamu` WHERE uuid=?", uuid)
	if err != nil {
		return response{
			Status: false,
			Pesan: "Query delete error: "+err.Error(),
			Data: []tamu{}
		}
	}
	return response{
		Status: true,
		Pesan: "Berhasil hapus data tamu "+uuid,
		Data: []tamu{}
	}
}

func kontroler (w http.ResponseWriter, r *http.Request){
	var tampilHtml, err1 = template.ParseFiles("template/tampil.html")
	if err1 != nil{
		fmt.Println(err1.Error())
		return nil
	}

	var tambahHtml, err2 = template.ParseFiles("template/tambah.html")
	if err2 != nil{
		fmt.Println(err2.Error())
		return nil
	}

	var ubahHtml, err3 = template.ParseFiles("template/ubah.html")
	if err3 != nil{
		fmt.Println(err3.Error())
		return nil
	}

	var hapusHtml, err4 = template.ParseFiles("template/hapus.html")
	if err4 != nil{
		fmt.Println(err4.Error())
		return nil
	}

	switch r.Method {
		case "GET":
			aksi := r.URL.Query()["aksi"]
			if (len(aksi)==0) {
				tampilHtml.Execute(w, tampil("Berhasil tampilkan semua data!"))
			}else if aksi[0] == "tambah" {
				tambahHtml.Execute(w, nil)
			}else if aksi[0] == "ubah" {
				uuid := r.URL.Query()["uuid"]
				ubahHtml.Execute(w, tampilFilterBerdasarkanUuid (uuid))
			} else if aksi[0] == "hapus" {
				uuid := r.URL.Query()["uuid"]
				hapusHtml.Execute(w, tampilFilterBerdasarkanUuid (uuid))
			} else{
				tampilHtml.Execute(w, tampil("Berhasil tampilkan semua data!"))
			}
		case "POST":
			var err = r.ParseForm()
			if err != nil{
				fmt.Fprint(w,"Maaf, terjadi kesalahan: ", err)
				return nil
			}
			var uuid string = r.FormValue("uuid")
			var namaLengkap string = r.FormValue("namaLengkap")
			var domisili string = r.FormValue("domisili")
			var updateAt string = "2022-08-01 00:00:00"
			var createdAt string = "2022-08-01 00:00:00"
			var aksi = r.URL.Path
			if aksi == "/tambah"{
				var hasil = tambah(uuid,namaLengkap,domisili,createdAt)
				tampilHtml.Execute(w, tampil(hasil.Pesan))
			} else if aksi == "/ubah" {
				var hasil = ubah(uuid,namaLengkap,domisili,updateAt)
				tampilHtml.Execute(w, tampil(hasil.Pesan))
			}else if aksi == "/hapus" {
				var hasil = hapus(uuid)
				tampilHtml.Execute(w, tampil(hasil.Pesan))
			}else{
				tampilHtml.Execute(w, tampil("Berhasil tampilkan semua data!"))
			}
		default:
			fmt.Fprint(w,"Maaf, Method yang di dukung hanya GET dan POST!")
	}
}