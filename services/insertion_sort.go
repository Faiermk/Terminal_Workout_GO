package services

import (
	"fmt"
)

func InsertionSort() {
	data := LoadData()
	n := len(data)

	if n == 0 {
		fmt.Println("Data kosong, tidak bisa sorting.")
		return
	}

	var pilihan int

	fmt.Println("-----PILIH INSERTION SORT-----")
	fmt.Println("1. Berdasarkan ID")
	fmt.Println("2. Berdasarkan Durasi Workout")
	fmt.Println("3. Berdasarkan Kalori")
	fmt.Print("Pilih (1-3): ")
	fmt.Scan(&pilihan)
	fmt.Scanln()

	switch pilihan {

	case 1:
		for i := 1; i < n; i++ {
			temp := data[i]
			j := i - 1

			for j >= 0 && data[j].ID > temp.ID {
				data[j+1] = data[j]
				j--
			}

			data[j+1] = temp
		}
		fmt.Println("Data berhasil diurutkan berdasarkan ID menggunakan Insertion Sort")

	case 2:
		for i := 1; i < n; i++ {
			temp := data[i]
			j := i - 1

			for j >= 0 && data[j].Durasi > temp.Durasi {
				data[j+1] = data[j]
				j--
			}

			data[j+1] = temp
		}
		fmt.Println("Data berhasil diurutkan berdasarkan Durasi Workout menggunakan Insertion Sort")

	case 3:
		for i := 1; i < n; i++ {
			temp := data[i]
			j := i - 1

			for j >= 0 && data[j].Kalori > temp.Kalori {
				data[j+1] = data[j]
				j--
			}

			data[j+1] = temp
		}
		fmt.Println("Data berhasil diurutkan berdasarkan Kalori menggunakan Insertion Sort")

	default:
		fmt.Println("Pilihan tidak valid")
		return
	}

	SaveData(data)

	fmt.Println("\n-----HASIL INSERTION SORT-----")
	for _, w := range data {
		fmt.Printf("ID: %d | %s | %s | %s | %d menit | %d kalori | %s\n",
			w.ID, w.Tanggal, w.Nama, w.Jenis, w.Durasi, w.Kalori, w.Catatan)
	}
}
