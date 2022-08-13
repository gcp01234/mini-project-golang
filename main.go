/*
Mini project untuk pelatihan Golang PROA
Author: Gita Citra Puspita
Dibuat pada: 12 Agustus 2022
*/

package main

import (
	"fmt"
	"strconv"
	"net/http"
	"html/template"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/",http.StripPrefix("/assets",fs))
	http.HandleFunc("/",kontroler)
	fmt.Println("Server berjalan di port 8080 ...")
	http.ListenAndServe(":8080",nil)
}

//struktur data tabel tamu
type tamu struct{
	Id int
	NamaLengkap string
	Domisili string
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
			Data: []tamu{},
		}
	}
	defer db.Close()
	dataTamu, err := db.Query("SELECT id, nama_lengkap, domisili FROM `tamu`")
	if err != nil {
		return response{
			Status: false,
			Pesan: "Query error: "+err.Error(),
			Data: []tamu{},
		}
	}
	defer dataTamu.Close()
	var hasil []tamu
	for dataTamu.Next(){
		var tm = tamu{}
		var err = dataTamu.Scan(&tm.Id,&tm.NamaLengkap,&tm.Domisili)
		if err != nil {
			return response{
				Status: false,
				Pesan: "Gagal baca data: "+err.Error(),
				Data: []tamu{},
			}
		}
		hasil = append(hasil, tm)
	}
	return response{
		Status: true,
		Pesan: pesan,
		Data: hasil,
	}
}


//fungsi untuk menampilkan data tamu berdasarkan Id
func tampilFilterBerdasarkanId (id int) response{
	db, err := koneksi()
	if err != nil {
		return response{
			Status: false,
			Pesan: "Gagal koneksi: "+err.Error(),
			Data: []tamu{},
		}
	}
	defer db.Close()
	dataTamu, err := db.Query("SELECT id, nama_lengkap, domisili FROM `tamu` WHERE id=?",id)
	if err != nil {
		return response{
			Status: false,
			Pesan: "Query error: "+err.Error(),
			Data: []tamu{},
		}
	}
	defer dataTamu.Close()
	var hasil []tamu
	for dataTamu.Next(){
		var tm = tamu{}
		var err = dataTamu.Scan(&tm.Id,&tm.NamaLengkap,&tm.Domisili)
		if err != nil {
			return response{
				Status: false,
				Pesan: "Gagal baca data tamu dengan Id "+string(id)+ " :"+err.Error(),
				Data: []tamu{},
			}
		}
		hasil = append(hasil, tm)
	}
	return response{
		Status: true,
		Pesan: "Berhasil tampilkan data tamu!",
		Data: hasil,
	}
	
}

//fungsi untuk menambahkan data tamu
func tambah (namaLengkap string, domisili string) response{
	db, err := koneksi()
	if err != nil {
		return response{
			Status: false,
			Pesan: "Gagal koneksi: "+err.Error(),
			Data: []tamu{},
		}
	}
	defer db.Close()
	_, err = db.Query("INSERT INTO `tamu`( `nama_lengkap`, `domisili`) VALUES (?,?)", namaLengkap,domisili)
	if err != nil {
		return response{
			Status: false,
			Pesan: "Query insert error: "+err.Error(),
			Data: []tamu{},
		}
	}
	return response{
		Status: true,
		Pesan: "Berhasil tambah data tamu "+namaLengkap,
		Data: []tamu{},
	}
}

//fungsi untuk mengubah data tamu
func ubah (id int, namaLengkap string, domisili string) response{
	db, err := koneksi()
	if err != nil {
		return response{
			Status: false,
			Pesan: "Gagal koneksi: "+err.Error(),
			Data: []tamu{},
		}
	}
	defer db.Close()
	_, err = db.Query("UPDATE `tamu` SET `nama_lengkap`=?,`domisili`=? WHERE id=?", namaLengkap,domisili,id)
	if err != nil {
		return response{
			Status: false,
			Pesan: "Query update error: "+err.Error(),
			Data: []tamu{},
		}
	}
	return response{
		Status: true,
		Pesan: "Berhasil ubah data tamu "+string(id),
		Data: []tamu{},
	}
}

func hapus (id int) response{
	db, err := koneksi()
	if err != nil {
		return response{
			Status: false,
			Pesan: "Gagal koneksi: "+err.Error(),
			Data: []tamu{},
		}
	}
	defer db.Close()
	_, err = db.Query("DELETE FROM `tamu` WHERE id=?", id)
	if err != nil {
		return response{
			Status: false,
			Pesan: "Query delete error: "+err.Error(),
			Data: []tamu{},
		}
	}
	return response{
		Status: true,
		Pesan: "Berhasil hapus data tamu "+string(id),
		Data: []tamu{},
	}
}

func kontroler (w http.ResponseWriter, r *http.Request){
	var tampilHtml, err1 = template.ParseFiles("template/tampil.html")
	if err1 != nil{
		fmt.Println(err1.Error())
		return
	}

	var tambahHtml, err2 = template.ParseFiles("template/tambah.html")
	if err2 != nil{
		fmt.Println(err2.Error())
		return
	}

	var ubahHtml, err3 = template.ParseFiles("template/ubah.html")
	if err3 != nil{
		fmt.Println(err3.Error())
		return
	}

	var hapusHtml, err4 = template.ParseFiles("template/hapus.html")
	if err4 != nil{
		fmt.Println(err4.Error())
		return
	}

	switch r.Method {
		case "GET":
			aksi := r.URL.Query()["aksi"]
			if (len(aksi)==0) {
				tampilHtml.Execute(w, tampil("Berhasil tampilkan semua data!"))
			}else if aksi[0] == "tambah" {
				tambahHtml.Execute(w, nil)
			}else if aksi[0] == "ubah" {
				id := r.URL.Query()["id"]
				i, _ := strconv.Atoi(id[0])
				ubahHtml.Execute(w, tampilFilterBerdasarkanId (i))
			} else if aksi[0] == "hapus" {
				id := r.URL.Query()["id"]
				i ,_ := strconv.Atoi(id[0])
				hapusHtml.Execute(w, tampilFilterBerdasarkanId (i))
			} else{
				tampilHtml.Execute(w, tampil("Berhasil tampilkan semua data!"))
			}
		case "POST":
			var err = r.ParseForm()
			if err != nil{
				fmt.Fprint(w,"Maaf, terjadi kesalahan: ", err)
				return
			}
			var id string = r.FormValue("id")
			i,_ := strconv.Atoi(id)
			var namaLengkap string = r.FormValue("namaLengkap")
			var domisili string = r.FormValue("domisili")
			var aksi = r.URL.Path
			if aksi == "/tambah"{
				var hasil = tambah(namaLengkap,domisili)
				tampilHtml.Execute(w, tampil(hasil.Pesan))
			} else if aksi == "/ubah" {
				var hasil = ubah(i,namaLengkap,domisili)
				tampilHtml.Execute(w, tampil(hasil.Pesan))
			}else if aksi == "/hapus" {
				var hasil = hapus(i)
				tampilHtml.Execute(w, tampil(hasil.Pesan))
			}else{
				tampilHtml.Execute(w, tampil("Berhasil tampilkan semua data!"))
			}
		default:
			fmt.Fprint(w,"Maaf, Method yang di dukung hanya GET dan POST!")
	}
}