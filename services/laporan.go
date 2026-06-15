package services

import "fmt"

func Laporan() {
	data := LoadData()

	fmt.Println("10 Data Terakhir:")
	start := len(data) - 10
	if start < 0 {
		start = 0
	}

	for i := start; i < len(data); i++ {
		fmt.Printf("ID: %d | %s | %s | %s | %d menit | %d kalori | %s\n",
			data[i].ID, data[i].Tanggal, data[i].Nama, data[i].Jenis,
			data[i].Durasi, data[i].Kalori, data[i].Catatan)
		// fmt.Println(data[i])
	}

	total := 0
	for _, w := range data {
		total += w.Kalori
	}

	fmt.Println("Total Kalori:", total, "kcal")
}