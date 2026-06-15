package services

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	var input string

	reader := bufio.NewReader(os.Stdin)

	w.ID = generateID(data)

	fmt.Print("Tanggal (DD-MM-YYYY): ")
	w.Tanggal, _ = reader.ReadString('\n')
	w.Tanggal = strings.TrimSpace(w.Tanggal)

	if w.Tanggal == "" {
		fmt.Println("Tanggal tidak boleh kosong.")
		return
	}

	fmt.Print("Nama Workout: ")
	w.Nama, _ = reader.ReadString('\n')
	w.Nama = strings.TrimSpace(w.Nama)

	if w.Nama == "" {
		fmt.Println("Nama workout tidak boleh kosong.")
		return
	}

	jenisList := []string{"Strength", "Cardio", "Flexibility", "Balance", "HIIT"}
	var pilihan int

	for {
		fmt.Println("\nPilih Jenis Workout:")
		fmt.Println("1. Strength")
		fmt.Println("2. Cardio")
		fmt.Println("3. Flexibility")
		fmt.Println("4. Balance")
		fmt.Println("5. HIIT")
		fmt.Print("Pilih (1-5): ")

		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)

		pilihan, _ = strconv.Atoi(input)

		if pilihan >= 1 && pilihan <= 5 {
			w.Jenis = jenisList[pilihan-1]
			break
		}

		fmt.Println("Pilihan tidak valid, coba lagi!")
	}

	fmt.Print("Durasi (menit): ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	durasi, err := strconv.Atoi(input)
	if err != nil || durasi <= 0 {
		fmt.Println("Durasi harus berupa angka positif.")
		return
	}
	w.Durasi = durasi

	fmt.Print("Kalori: ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	kalori, err := strconv.Atoi(input)
	if err != nil || kalori <= 0 {
		fmt.Println("Kalori harus berupa angka positif.")
		return
	}
	w.Kalori = kalori

	fmt.Print("Catatan: ")
	w.Catatan, _ = reader.ReadString('\n')
	w.Catatan = strings.TrimSpace(w.Catatan)

	data = append(data, w)
	SaveData(data)

	fmt.Println("Data workout berhasil ditambahkan!")
}

func LihatWorkout() {
	data := LoadData()

	fmt.Println("\n-----DATA WORKOUT-----")
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
	reader := bufio.NewReader(os.Stdin)

	if len(data) == 0 {
		fmt.Println("Data workout masih kosong.")
		return
	}

	fmt.Print("Masukkan ID yang ingin diupdate: ")
	inputID, _ := reader.ReadString('\n')
	inputID = strings.TrimSpace(inputID)

	id, err := strconv.Atoi(inputID)
	if err != nil {
		fmt.Println("ID harus berupa angka.")
		return
	}

	for i := range data {
		if data[i].ID == id {
			fmt.Println("\nData ditemukan:")
			fmt.Printf("ID: %d | %s | %s | %s | %d menit | %d kalori | %s\n",
				data[i].ID, data[i].Tanggal, data[i].Nama, data[i].Jenis,
				data[i].Durasi, data[i].Kalori, data[i].Catatan)

			fmt.Println("\n-----MENU UPDATE-----")
			fmt.Println("1. Update Tanggal")
			fmt.Println("2. Update Nama Workout")
			fmt.Println("3. Update Jenis Workout")
			fmt.Println("4. Update Durasi")
			fmt.Println("5. Update Kalori")
			fmt.Println("6. Update Catatan")
			fmt.Println("7. Update Semua Data")
			fmt.Print("Pilih (1-7): ")

			inputPilihan, _ := reader.ReadString('\n')
			inputPilihan = strings.TrimSpace(inputPilihan)

			pilihan, err := strconv.Atoi(inputPilihan)
			if err != nil {
				fmt.Println("Pilihan harus berupa angka.")
				return
			}

			switch pilihan {
			case 1:
				fmt.Print("Tanggal baru (DD-MM-YYYY): ")
				input, _ := reader.ReadString('\n')
				data[i].Tanggal = strings.TrimSpace(input)

			case 2:
				fmt.Print("Nama workout baru: ")
				input, _ := reader.ReadString('\n')
				data[i].Nama = strings.TrimSpace(input)

			case 3:
				data[i].Jenis = pilihJenisWorkout()

			case 4:
				fmt.Print("Durasi baru (menit): ")
				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)

				durasi, err := strconv.Atoi(input)
				if err != nil {
					fmt.Println("Durasi harus berupa angka.")
					return
				}
				data[i].Durasi = durasi

			case 5:
				fmt.Print("Kalori baru: ")
				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)

				kalori, err := strconv.Atoi(input)
				if err != nil {
					fmt.Println("Kalori harus berupa angka.")
					return
				}
				data[i].Kalori = kalori

			case 6:
				fmt.Print("Catatan baru: ")
				input, _ := reader.ReadString('\n')
				data[i].Catatan = strings.TrimSpace(input)

			case 7:
				fmt.Print("Tanggal baru (DD-MM-YYYY): ")
				input, _ := reader.ReadString('\n')
				data[i].Tanggal = strings.TrimSpace(input)

				fmt.Print("Nama workout baru: ")
				input, _ = reader.ReadString('\n')
				data[i].Nama = strings.TrimSpace(input)

				data[i].Jenis = pilihJenisWorkout()

				fmt.Print("Durasi baru (menit): ")
				input, _ = reader.ReadString('\n')
				input = strings.TrimSpace(input)
				durasi, err := strconv.Atoi(input)
				if err != nil {
					fmt.Println("Durasi harus berupa angka.")
					return
				}
				data[i].Durasi = durasi

				fmt.Print("Kalori baru: ")
				input, _ = reader.ReadString('\n')
				input = strings.TrimSpace(input)
				kalori, err := strconv.Atoi(input)
				if err != nil {
					fmt.Println("Kalori harus berupa angka.")
					return
				}
				data[i].Kalori = kalori

				fmt.Print("Catatan baru: ")
				input, _ = reader.ReadString('\n')
				data[i].Catatan = strings.TrimSpace(input)

			default:
				fmt.Println("Pilihan tidak valid.")
				return
			}

			SaveData(data)
			fmt.Println("Data berhasil diupdate!")
			return
		}
	}

	fmt.Println("Data dengan ID tersebut tidak ditemukan.")
}

func pilihJenisWorkout() string {
	jenisList := []string{"Strength", "Cardio", "Flexibility", "Balance", "HIIT"}
	var pilihan int

	fmt.Println("\nPilih Jenis Workout:")
	fmt.Println("1. Strength")
	fmt.Println("2. Cardio")
	fmt.Println("3. Flexibility")
	fmt.Println("4. Balance")
	fmt.Println("5. HIIT")
	fmt.Print("Pilih (1-5): ")
	fmt.Scan(&pilihan)
	fmt.Scanln()

	if pilihan < 1 || pilihan > len(jenisList) {
		fmt.Println("Pilihan tidak valid, jenis otomatis diisi Strength.")
		return "Strength"
	}

	return jenisList[pilihan-1]
}

func HapusWorkout() {
	data := LoadData()
	var id int

	fmt.Print("Masukkan ID yang ingin dihapus: ")
	fmt.Scan(&id)
	fmt.Scanln()

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