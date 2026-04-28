package services

import (
	"fmt"
	"tubes/models"
)

func TambahWorkout() {
	data := LoadData()
	var w models.Workout

	w.ID = len(data) + 1
	fmt.Print("Tanggal: "); fmt.Scan(&w.Tanggal)
	fmt.Print("Nama: "); fmt.Scan(&w.Nama)
	fmt.Print("Jenis: "); fmt.Scan(&w.Jenis)
	fmt.Print("Durasi: "); fmt.Scan(&w.Durasi)
	fmt.Print("Kalori: "); fmt.Scan(&w.Kalori)
	fmt.Print("Catatan: "); fmt.Scan(&w.Catatan)

	data = append(data, w)
	SaveData(data)

	fmt.Println("Data berhasil ditambahkan!")
}

func LihatWorkout() {
	data := LoadData()
	for _, w := range data {
		fmt.Println(w)
	}
}

func UpdateWorkout() {
	data := LoadData()
	var id int
	fmt.Print("Masukkan ID: ")
	fmt.Scan(&id)

	for i := range data {
		if data[i].ID == id {
			fmt.Print("Nama baru: "); fmt.Scan(&data[i].Nama)
			SaveData(data)
			fmt.Println("Berhasil update!")
			return
		}
	}
	fmt.Println("Data tidak ditemukan")
}

func HapusWorkout() {
	data := LoadData()
	var id int
	fmt.Print("Masukkan ID: ")
	fmt.Scan(&id)

	for i := range data {
		if data[i].ID == id {
			data = append(data[:i], data[i+1:]...)
			SaveData(data)
			fmt.Println("Berhasil hapus!")
			return
		}
	}
	fmt.Println("Data tidak ditemukan")
}