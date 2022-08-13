package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main(){
	_ , err := gorm.Open(sqlite.Open("data.db"), $gorm.Config{})
	if err != nill {
		panic ("Program error, tidak bisa terkoneksi dengan database")
	}
	fmt.Println("Koneksi database berhasil.")

}