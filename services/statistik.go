package services

import (
	"fmt"
	"tubes/models"
)

// cariMaksDurasi nyari workout dengan durasi terlama secara rekursif
// cara kerjanya: bandingin satu-satu dari index sekarang sampai habis
// kalau durasi data[index] lebih besar dari maks, ganti maksnya
func cariMaksDurasi(data []models.Workout, index int, maks models.Workout) models.Workout {
	// base case: kalau udah sampe akhir data, return hasilnya
	if index == len(data) {
		return maks
	}

	// kalau durasi data sekarang lebih besar, jadiin dia maks yang baru
	if data[index].Durasi > maks.Durasi {
		maks = data[index]
	}
	// panggil diri sendiri dengan index berikutnya (rekursif)
	return cariMaksDurasi(data, index+1, maks)
}

// cariMaksKalori nyari workout dengan kalori terbanyak secara rekursif
// sama konsepnya kayak cariMaksDurasi, bedanya yang dibandingkan field Kalori
func cariMaksKalori(data []models.Workout, index int, maks models.Workout) models.Workout {
	// base case: kalau udah sampe akhir data, return hasilnya
	if index == len(data) {
		return maks
	}

	// kalau kalori data sekarang lebih besar, jadiin dia maks yang baru
	if data[index].Kalori > maks.Kalori {
		maks = data[index]
	}
	// panggil diri sendiri dengan index berikutnya (rekursif)
	return cariMaksKalori(data, index+1, maks)
}

// StatistikWorkout nampilin menu statistik dan manggil fungsi yang sesuai
func StatistikWorkout() {
	// load semua data workout dari file
	data := LoadData()

	// load semua data workout dari file
	if len(data) == 0 {
		fmt.Println("Belum ada data workout.")
		return
	}

	fmt.Println("\n=== STATISTIK WORKOUT ===")
	fmt.Println("1. Workout dengan durasi terlama")
	fmt.Println("2. Workout dengan kalori terbanyak")
	fmt.Println("3. Keduanya")
	fmt.Print("Pilih: ")

	var pilih int
	fmt.Scan(&pilih)
	fmt.Scanln() // untuk buang newline setelah input angka

	switch pilih {
	case 1:
		// mulai rekursif dari index 1, dengan data[0] sebagai nilai awal maks
		hasil := cariMaksDurasi(data, 1, data[0])
		fmt.Println("\nWorkout Durasi Terlama:")
		fmt.Printf("ID: %d | %s | %s | %s | %d menit | %d kalori | %s\n",
			hasil.ID, hasil.Tanggal, hasil.Nama, hasil.Jenis,
			hasil.Durasi, hasil.Kalori, hasil.Catatan)

	case 2:
		// mulai rekursif dari index 1, dengan data[0] sebagai nilai awal maks
		hasil := cariMaksKalori(data, 1, data[0])
		fmt.Println("\nWorkout Kalori Terbanyak:")
		fmt.Printf("ID: %d | %s | %s | %s | %d menit | %d kalori | %s\n",
			hasil.ID, hasil.Tanggal, hasil.Nama, hasil.Jenis,
			hasil.Durasi, hasil.Kalori, hasil.Catatan)

	case 3:
		// jalanin keduanya sekaligus
		hasilDurasi := cariMaksDurasi(data, 1, data[0])
		hasilKalori := cariMaksKalori(data, 1, data[0])

		fmt.Println("\nWorkout Durasi Terlama:")
		fmt.Printf("ID: %d | %s | %s | %s | %d menit | %d kalori | %s\n",
			hasilDurasi.ID, hasilDurasi.Tanggal, hasilDurasi.Nama, hasilDurasi.Jenis,
			hasilDurasi.Durasi, hasilDurasi.Kalori, hasilDurasi.Catatan)

		fmt.Println("\nWorkout Kalori Terbanyak:")
		fmt.Printf("ID: %d | %s | %s | %s | %d menit | %d kalori | %s\n",
			hasilKalori.ID, hasilKalori.Tanggal, hasilKalori.Nama, hasilKalori.Jenis,
			hasilKalori.Durasi, hasilKalori.Kalori, hasilKalori.Catatan)

	default:
		fmt.Println("Pilihan tidak valid.")
	}
}
