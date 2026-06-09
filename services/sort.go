package services

import "fmt"

func SelectionSort() {
	data := LoadData()
	n := len(data)

	if n == 0 {
		fmt.Println("Data kosong, tidak bisa sorting.")
		return
	}

	var pilihan int

	fmt.Println("-----PILIH SORTING-----")
	fmt.Println("1. Berdasarkan Nama")
	fmt.Println("2. Berdasarkan Durasi")
	fmt.Println("3. Berdasarkan Kalori")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilihan)

	switch pilihan {

	case 1:
		// 🔹 Sort Nama
		for i := 0; i < n-1; i++ {
			min := i
			for j := i + 1; j < n; j++ {
				if data[j].Nama < data[min].Nama {
					min = j
				}
			}
			data[i], data[min] = data[min], data[i]
		}
		fmt.Println("Data berhasil diurutkan berdasarkan Nama")

	case 2:
		// 🔹 Sort Durasi
		for i := 0; i < n-1; i++ {
			min := i
			for j := i + 1; j < n; j++ {
				if data[j].Durasi < data[min].Durasi {
					min = j
				}
			}
			data[i], data[min] = data[min], data[i]
		}
		fmt.Println("Data berhasil diurutkan berdasarkan Durasi")

	case 3:
		// 🔹 Sort Kalori
		for i := 0; i < n-1; i++ {
			min := i
			for j := i + 1; j < n; j++ {
				if data[j].Kalori < data[min].Kalori {
					min = j
				}
			}
			data[i], data[min] = data[min], data[i]
		}
		fmt.Println("Data berhasil diurutkan berdasarkan Kalori")

	default:
		fmt.Println("Pilihan tidak valid")
		return
	}

	// simpan hasil sorting
	SaveData(data)

	//untuk menampilkan hasil
	fmt.Println("\n-----HASIL SORTING-----")
	for _, w := range data {
		fmt.Printf("ID: %d | %s | %s | %s | %d menit | %d kalori | %s\n",
			w.ID, w.Tanggal, w.Nama, w.Jenis, w.Durasi, w.Kalori, w.Catatan)
	}
}