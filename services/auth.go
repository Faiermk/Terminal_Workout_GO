package services

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tubes/models"
)

var userFilePath = "data/login.txt"

func loadUsers() []models.User {
	var data []models.User

	file, err := os.Open(userFilePath)
	if err != nil {
		return data
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		if len(parts) < 3 {
			continue
		}
		id, _ := strconv.Atoi(parts[0])
		data = append(data, models.User{
			ID:		id,
			Username: parts[1],
			Password: parts[2],
		})
	}
	return data
}

func saveUsers(users []models.User) {
	file, _ := os.Create(userFilePath)
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, u :=	range users {
	line := fmt.Sprintf("%d|%s|%s\n", u.ID, u.Username, u.Password)
		writer.WriteString(line)
	}
	writer.Flush()
}

func Register() {
	users := loadUsers()
	var Username, Password string

	fmt.Print("Username: ")
	fmt.Scan(&Username)

	for _, u := range users {
		if u.Username == Username {
			fmt.Println("USERNAME SUDAH DIPAKAI!")
			return
		}
	}

	fmt.Print("Password: ")
	fmt.Scan(&Password)

	newUser := models.User{
		ID:		len(users) + 1,
		Username: Username,
		Password: Password,
	}
	users = append(users, newUser)
	saveUsers(users)
	fmt.Println("REGISTRASI BERHASIL! SILAKAN LOGIN.")
}

func Login() *models.User {
	users := loadUsers()
	var Username, Password string

	fmt.Print("Username: ")
	fmt.Scan(&Username)
	fmt.Print("Password: ")
	fmt.Scan(&Password)

	for i, u := range users {
		if u.Username == Username && u.Password == Password {
			fmt.Printf("SELAMAT DATANG, %s!\n", u.Username)
			return &users[i]
		}
	}

	fmt.Println("USERNAME TIDAK DITEMUKAN ATAU PASSWORD SALAH!")
	return nil
}