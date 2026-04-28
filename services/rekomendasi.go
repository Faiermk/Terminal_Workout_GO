package services

import "fmt"

func RekomendasiWorkout() {
	data := LoadData()
	count := make(map[string]int)

	for _, w := range data {
		count[w.Nama]++
	}

	max := 0
	rekom := ""

	for k, v := range count {
		if v > max {
			max = v
			rekom = k
		}
	}

	fmt.Println("Rekomendasi:", rekom)
}