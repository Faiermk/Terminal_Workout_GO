package services

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tubes/models"
)

func generateID(data []models.Workout) int {
	maxID := 0
	for _, w := range data {
		if w.ID > maxID {
			maxID = w.ID
		}
	}
	return maxID + 1
}

func TambahWorkout() {
	data := LoadData()
	var w models.Workout

	reader := bufio.NewReader(os.Stdin)

	w.ID = generateID(data)

    reader.ReadString('\n')

	fmt.Print("Tanggal (DD-MM-YYYY): ")
    w.Tanggal, _ = reader.ReadString('\n')
    w.Tanggal = strings.TrimSpace(w.Tanggal)

	fmt.Print("Nama Workout: ")
    w.Nama, _ = reader.ReadString('\n')
    w.Nama = strings.TrimSpace(w.Nama)
	
	jenisList := []string{"Strength", "Cardio", "Flexibility", "Balance", "HIIT"}
	var pilihan int

	for {
		fmt.Println("Pilih Jenis:")
		for i, j := range jenisList {
			fmt.Printf("%d. %s\n", i+1, j)
		}

		fmt.Print("Pilih (1-5): ")
		fmt.Scan(&pilihan)

		if pilihan >= 1 && pilihan <= len(jenisList) {
			w.Jenis = jenisList[pilihan-1]
			break
		}

		fmt.Println("Pilihan tidak valid, coba lagi!\n")
	}

	fmt.Print("Durasi (menit): ")
	fmt.Scan(&w.Durasi)

	fmt.Print("Kalori: ")
	fmt.Scan(&w.Kalori)


	fmt.Print("Catatan: ")
	reader.ReadString('\n')
	catatan, _ := reader.ReadString('\n')
	w.Catatan = strings.TrimSpace(catatan)

	data = append(data, w)
	SaveData(data)

	fmt.Println("Data berhasil ditambahkan!")
}

func LihatWorkout() {
	data := LoadData()

	fmt.Println("\n=== DATA WORKOUT ===")
	if len(data) == 0 {
		fmt.Println("Belum ada data.")
		return
	}

	for _, w := range data {
		fmt.Printf("ID: %d | %s | %s | %s | %d menit | %d kalori | %s\n",
			w.ID, w.Tanggal, w.Nama, w.Jenis, w.Durasi, w.Kalori, w.Catatan)
	}
}

func UpdateWorkout() {
	data := LoadData()
	var id int

	fmt.Print("Masukkan ID yang ingin diupdate: ")
	fmt.Scan(&id)

	for i := range data {
		if data[i].ID == id {

			fmt.Print("Nama baru: ")
			fmt.Scan(&data[i].Nama)

			fmt.Print("Durasi baru: ")
			fmt.Scan(&data[i].Durasi)

			fmt.Print("Kalori baru: ")
			fmt.Scan(&data[i].Kalori)

			SaveData(data)
			fmt.Println("Data berhasil diupdate!")
			return
		}
	}

	fmt.Println("Data tidak ditemukan")
}

func HapusWorkout() {
	data := LoadData()
	var id int

	fmt.Print("Masukkan ID yang ingin dihapus: ")
	fmt.Scan(&id)

	for i := range data {
		if data[i].ID == id {
			data = append(data[:i], data[i+1:]...)
			SaveData(data)
			fmt.Println("Data berhasil dihapus!")
			return
		}
	}

	fmt.Println("Data tidak ditemukan")
}