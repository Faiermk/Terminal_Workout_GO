package services

import (
	"fmt"
	"strings"
)

// Fungsi untuk memfilter data workout berdasarkan tanggal
func FilterTanggal() {
	data := LoadData()

	// Cek apakah data workout masih kosong
	if len(data) == 0 {
		fmt.Println("Belum ada data workout.")
		return
	}

	fmt.Println("\nFILTER BY TANGGAL")
	fmt.Print("Masukkan tanggal (DD-MM-YYYY): ")

	var tanggal string
	fmt.Scan(&tanggal)
	fmt.Scanln() // untuk buang newline setelah input tanggal

	found := false
	fmt.Println("\n=== HASIL FILTER ===")

	// Menampilkan data yang tanggalnya sesuai input
	for _, w := range data {
		if strings.EqualFold(w.Tanggal, tanggal) {
			fmt.Printf("ID: %d | %s | %s | %s | %d menit | %d kalori | %s\n",
				w.ID, w.Tanggal, w.Nama, w.Jenis, w.Durasi, w.Kalori, w.Catatan)
			found = true
		}
	}

	// Jika tidak ada data yang cocok
	if !found {
		fmt.Println("Tidak ada workout di tanggal tersebut.")
	}
}