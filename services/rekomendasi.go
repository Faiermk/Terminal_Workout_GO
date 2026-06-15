package services

import (
	"fmt"
	"strings"
)


func RekomendasiWorkout() {
	data := LoadData()

	if len(data) == 0 {
		fmt.Println("Data workout masih kosong.")
		return
	}

	frekuensi := make(map[string]int)
	namaAsli := make(map[string]string)

	for _, w := range data {
		namaKey := strings.ToLower(strings.TrimSpace(w.Nama))

		if namaKey == "" {
			continue
		}

		frekuensi[namaKey]++

		if namaAsli[namaKey] == "" {
			namaAsli[namaKey] = w.Nama
		}
	}

	var rekomendasi string
	var jumlahTerbanyak int

	for nama, jumlah := range frekuensi {
		if jumlah > jumlahTerbanyak {
			jumlahTerbanyak = jumlah
			rekomendasi = nama
		}
	}

	if rekomendasi == "" {
		fmt.Println("Belum ada data workout yang bisa direkomendasikan.")
		return
	}

	fmt.Println("-----REKOMENDASI WORKOUT-----")
	fmt.Println("Workout yang paling sering dilakukan:", namaAsli[rekomendasi])
	fmt.Println("Jumlah dilakukan:", jumlahTerbanyak, "kali")
	fmt.Println("Rekomendasi: Pertahankan latihan", namaAsli[rekomendasi], "karena paling sering dilakukan.")
}