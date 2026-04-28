package services

func SelectionSort() {
	data := LoadData()
	n := len(data)

	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if data[j].Nama < data[min].Nama {
				min = j
			}
		}
		data[i], data[min] = data[min], data[i]
	}

	SaveData(data)
}