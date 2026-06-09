package services

import "fmt"

func RekomendasiWorkout() {
	data := LoadData()

	if len(data) == 0 {
		fmt.Println("Belum ada data workout")
		return
	}

	count := make(map[string]int)

	for _, w := range data {
		count[w.Nama]++
	}

	max := 0
	var rekomList []string

	for nama, jumlah := range count {
		if jumlah > max {
			max = jumlah
			rekomList = []string{nama}
		} else if jumlah == max {
			rekomList = append(rekomList, nama)
		}
	}

	fmt.Println("\n-----REKOMENDASI WORKOUT-----")
	for _, r := range rekomList {
		fmt.Printf("- %s (%d kali dilakukan)\n", r, max)
	}
}