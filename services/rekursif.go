package services

import (
	"fmt"
	"tubes/models"
)

func TampilWorkoutRekursif(data []models.Workout, index int) {
	if index == len(data) {
		return
	}

	w := data[index]

	fmt.Printf("ID: %d | %s | %s | %s | %d menit | %d kalori | %s\n",
		w.ID, w.Tanggal, w.Nama, w.Jenis, w.Durasi, w.Kalori, w.Catatan)

	TampilWorkoutRekursif(data, index+1)
}

func TotalKaloriRekursif(data []models.Workout, index int) int {
	if index == len(data) {
		return 0
	}

	return data[index].Kalori + TotalKaloriRekursif(data, index+1)
}

func TotalDurasiRekursif(data []models.Workout, index int) int {
	if index == len(data) {
		return 0
	}

	return data[index].Durasi + TotalDurasiRekursif(data, index+1)
}

func MenuRekursif() {
	data := LoadData()

	if len(data) == 0 {
		fmt.Println("Data workout masih kosong.")
		return
	}

	fmt.Println("\n-----DATA WORKOUT REKURSIF-----")
	TampilWorkoutRekursif(data, 0)

	totalKalori := TotalKaloriRekursif(data, 0)
	totalDurasi := TotalDurasiRekursif(data, 0)

	fmt.Println("\n-----HASIL REKURSIF-----")
	fmt.Println("Total kalori terbakar :", totalKalori, "kalori")
	fmt.Println("Total durasi workout  :", totalDurasi, "menit")
}