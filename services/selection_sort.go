package services

import (
	"fmt"
	"strings"
)


func SelectionSort() {
	data := LoadData()
	n := len(data)

	if n == 0 {
		fmt.Println("Data kosong, tidak bisa sorting.")
		return
	}

	for i := 0; i < n-1; i++ {
		min := i

		for j := i + 1; j < n; j++ {
			if strings.ToLower(data[j].Nama) < strings.ToLower(data[min].Nama) {
				min = j
			}
		}

		data[i], data[min] = data[min], data[i]
	}

	SaveData(data)

	fmt.Println("Data berhasil diurutkan berdasarkan Nama Workout menggunakan Selection Sort")
	fmt.Println("\n-----HASIL SELECTION SORT-----")

	for _, w := range data {
		fmt.Printf("ID: %d | %s | %s | %s | %d menit | %d kalori | %s\n",
			w.ID, w.Tanggal, w.Nama, w.Jenis, w.Durasi, w.Kalori, w.Catatan)
	}
}