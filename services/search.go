package services

import "fmt"

func SequentialSearch() {
	data := LoadData()
	var jenis string
	fmt.Print("Cari jenis: ")
	fmt.Scan(&jenis)

	for _, w := range data {
		if w.Jenis == jenis {
			fmt.Println(w)
		}
	}
}

func BinarySearch() {
	data := LoadData()
	SelectionSort()

	var nama string
	fmt.Print("Cari nama: ")
	fmt.Scan(&nama)

	low, high := 0, len(data)-1
	for low <= high {
		mid := (low + high) / 2
		if data[mid].Nama == nama {
			fmt.Println(data[mid])
			return
		} else if data[mid].Nama < nama {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("Tidak ditemukan")
}