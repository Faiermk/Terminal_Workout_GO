package main

import (
	"fmt"
	"tubes/models"
	"tubes/services"
)

func main() {
	// currentUser menyimpan data user yang sedang login
	// nilnya nil jika belum login
	var currentUser *models.User

	// loop autentikasi — terus minta login/register sampai berhasil login
	for currentUser == nil {
		fmt.Println("-----TERMINAL WORKOUT-----")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih: ")

		var pilihAuth int
		fmt.Scan(&pilihAuth)
		fmt.Scanln()

		switch pilihAuth {
			case 1:
				// Login mengembalikan pointer User jika berhasil, nil jika gagal
				currentUser = services.Login()
			case 2:
				services.Register()
			case 0: 
				return
		}
	}

	var pilih int

	// loop utama aplikasi, jadi program bakal berjalan terus selama user tidak logout
	for {
		fmt.Println("-----APLIKASI WORKOUT-----")
		fmt.Println("1. Tambah")
		fmt.Println("2. Lihat")
		fmt.Println("3. Update")
		fmt.Println("4. Hapus")
		fmt.Println("5. Search Sequential Jenis Workout")
		fmt.Println("6. Search Binary Nama Workout")
		fmt.Println("7. Selection Sort Nama Workout")
		fmt.Println("8. Insertion Sort ID/Durasi/Kalori")
		fmt.Println("9. Rekomendasi")
		fmt.Println("10. Laporan")
		fmt.Println("11. Recursive")
		fmt.Println("12. Statistik Workout")
		fmt.Println("13. Filter by Tanggal")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih (0-11): ")
		fmt.Scan(&pilih)
		fmt.Scanln()

		switch pilih {
		case 1:
			services.TambahWorkout()
		case 2:
			services.LihatWorkout()
		case 3:
			services.UpdateWorkout()
		case 4:
			services.HapusWorkout()
		case 5:
			services.SequentialSearch()
		case 6:
			services.BinarySearch()
		case 7:
			services.SelectionSort()
		case 8:
			services.InsertionSort()
		case 9:
			services.RekomendasiWorkout()
		case 10:
			services.Laporan()
		case 11:
			services.MenuRekursif()
		case 12:
			services.StatistikWorkout()
		case 13:
			services.FilterTanggal()
		case 0:
			currentUser = nil
			for currentUser == nil {
				fmt.Println("\n-----TERMINAL WORKOUT-----")
				fmt.Println("1. Login")
				fmt.Println("2. Registrasi")
				fmt.Println("0. Keluar")
				fmt.Print("Pilih (0-2): ")

				var pilihAuth int
				fmt.Scan(&pilihAuth)
				fmt.Scanln()

			switch pilihAuth {
			case 1:
				currentUser = services.Login()
			case 2:
				services.Register()
			case 0:
				return
			}
		}
	}
	}
}