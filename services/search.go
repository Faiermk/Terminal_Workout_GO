package services

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Sequential Search (berdasarkan jenis workout)
func SequentialSearch() {
	data := LoadData()
	var jenis string
	found := false

	if len(data) == 0 {
		fmt.Println("Data kosong")
		return
	}

	// kasih pilihan ke user
	fmt.Println("\nJenis yang tersedia:")
	fmt.Println("1. Strength")
	fmt.Println("2. Cardio")
	fmt.Println("3. Flexibility")
	fmt.Println("4. Balance")
	fmt.Println("5. HIIT")

	jenisList := []string{"Strength", "Cardio", "Flexibility", "Balance", "HIIT"}
	var pilihan int

	fmt.Print("\nPilih jenis workout (1-5): ")
	fmt.Scan(&pilihan)
	fmt.Scanln()

	if pilihan < 1 || pilihan > len(jenisList) {
		fmt.Println("Pilihan tidak valid")
		return
	}

	jenis = jenisList[pilihan-1]


	// fmt.Print("\nCari jenis workout: ")
	// fmt.Scan(&jenis)

	fmt.Println("\n-----HASIL PENCARIAN-----")
	for _, w := range data {
		if strings.EqualFold(w.Jenis, jenis) {
			fmt.Printf("ID: %d | %s | %s | %s | %d menit | %d kalori | %s\n",
				w.ID, w.Tanggal, w.Nama, w.Jenis, w.Durasi, w.Kalori, w.Catatan)
			found = true
		}
	}

	if !found {
		fmt.Println("Data tidak ditemukan")
	}
}

// Binary Search (berdasarkan nama workout)
func BinarySearch() {
	data := LoadData()
	n := len(data)

	if n == 0 {
		fmt.Println("Data kosong")
		return
	}

	// tampilkan data dulu
	fmt.Println("\n-----DAFTAR WORKOUT-----")
	for _, w := range data {
		fmt.Printf("- %s\n", w.Nama)
	}

	// sorting lokal (wajib untuk binary)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if strings.ToLower(data[j].Nama) < strings.ToLower(data[min].Nama) {
				min = j
			}
		}
		data[i], data[min] = data[min], data[i]
	}

	reader := bufio.NewReader(os.Stdin)

	var nama string
	fmt.Print("\nCari nama workout (contoh: Push Up): ")
	// fmt.Scan(&nama)
	nama, _ = reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	low, high := 0, n-1

	for low <= high {
		mid := (low + high) / 2

		if strings.EqualFold(data[mid].Nama, nama) {
			fmt.Println("\n-----DATA DITEMUKAN-----")
			fmt.Printf("ID: %d | %s | %s | %s | %d menit | %d kalori | %s\n",
				data[mid].ID, data[mid].Tanggal, data[mid].Nama,
				data[mid].Jenis, data[mid].Durasi, data[mid].Kalori, data[mid].Catatan)
			return
		} else if strings.ToLower(data[mid].Nama) < strings.ToLower(nama) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	fmt.Println("Data tidak ditemukan")
}