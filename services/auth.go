package services

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tubes/models"
)

// path file penyimpanan data user
var userFilePath = "data/login.txt"

// loadUsers baca semua data user dari file login.txt
// dan kembaliin sebagai slice of User
func loadUsers() []models.User {
	var data []models.User

	// buka file login.txt
	file, err := os.Open(userFilePath)
	if err != nil {
		return data
	}
	defer file.Close()

	// baca file baris per baris
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// pemisah tiap baris berdasarkan delimiter "|"
		parts := strings.Split(line, "|")
		if len(parts) < 3 {
			continue
		}

		// konversi ID dari string ke int
		id, _ := strconv.Atoi(parts[0])
		
		// buat menambahkan user ke slice
		data = append(data, models.User{
			ID:		id,
			Username: parts[1],
			Password: parts[2],
		})
	}
	return data
}

// saveUsers untuk menyimpan semua data user ke file login.txt
func saveUsers(users []models.User) {
	file, err := os.Create(userFilePath)
	if err != nil {
		fmt.Println("Gagal menyimpan data:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	// ini untuk ngatur format usernya (format: ID|Username|Password)
	for _, u :=	range users {
		line := fmt.Sprintf("%d|%s|%s\n", u.ID, u.Username, u.Password)
		writer.WriteString(line)
	}
	writer.Flush()
}

// Register disini tempat proses pendaftaran user baru
func Register() {
	users := loadUsers()
	var Username, Password string

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Username: ")
	Username, _ = reader.ReadString('\n')
	Username = strings.TrimSpace(Username)

	if Username == "" {
		fmt.Println("Username tidak boleh kosong!")
		return
	}

	// cek apakah username sudah dipakai tanpa membedakan huruf besar/kecil
	for _, u := range users {
		if strings.EqualFold(u.Username, Username) {
			fmt.Println("USERNAME SUDAH DIPAKAI!")
			return
		}
	}

	fmt.Print("Password: ")
	Password, _ = reader.ReadString('\n')
	Password = strings.TrimSpace(Password)

	if Password == "" {
		fmt.Println("Password tidak boleh kosong!")
		return
	}

	newUser := models.User{
		ID:       len(users) + 1,
		Username: Username,
		Password: Password,
	}

	users = append(users, newUser)
	saveUsers(users)
	fmt.Println("REGISTRASI BERHASIL! SILAKAN LOGIN.")
}

// Login melewati proses autentikasi user
// kembaliin pointer ke User jika berhasil, nil jika gagal
func Login() *models.User {
	users := loadUsers()

	var Username, Password string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Username: ")
	Username, _ = reader.ReadString('\n')
	Username = strings.TrimSpace(Username)

	fmt.Print("Password: ")
	Password, _ = reader.ReadString('\n')
	Password = strings.TrimSpace(Password)

	// username tidak case-sensitive, password tetap harus sama persis
	for i, u := range users {
		if strings.EqualFold(u.Username, Username) && u.Password == Password {
			fmt.Printf("SELAMAT DATANG, %s!\n", u.Username)
			return &users[i]
		}
	}

	fmt.Println("USERNAME TIDAK DITEMUKAN ATAU PASSWORD SALAH!")
	return nil
}