package services

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tubes/models"
)

var filePath = "data/workout.txt"

func LoadData() []models.Workout {
	var data []models.Workout

	file, err := os.Open(filePath)
	if err != nil {
		return data
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")

		if len(parts) < 7 {
			continue
		}

		id, _ := strconv.Atoi(parts[0])
		durasi, _ := strconv.Atoi(parts[4])
		kalori, _ := strconv.Atoi(parts[5])

		w := models.Workout{
			ID:      id,
			Tanggal: parts[1],
			Nama:    parts[2],
			Jenis:   parts[3],
			Durasi:  durasi,
			Kalori:  kalori,
			Catatan: parts[6],
		}

		data = append(data, w)
	}

	return data
}

func SaveData(data []models.Workout) {
	file, _ := os.Create(filePath)
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, w := range data {
		line := fmt.Sprintf("%d|%s|%s|%s|%d|%d|%s\n",
			w.ID, w.Tanggal, w.Nama, w.Jenis, w.Durasi, w.Kalori, w.Catatan)
		writer.WriteString(line)
	}

	writer.Flush()
}