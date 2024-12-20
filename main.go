package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type User struct {
	Username string
	Password string
	UserType string
}

var users = []User{
	{"admin", "admin", "admin"},
	{"user", "user", "customer"},
	{"user2", "user2", "customer"},
}

var currentUser *User

func main() {
	for {
		fmt.Println("Hoş geldiniz!\n0 - Admin Girişi\n1 - Müşteri Girişi\n2 - Çıkış")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 0:
			adminGiriş()
		case 1:
			customerGiriş()
		case 2:
			fmt.Println("Programdan çıkılıyor...")
			return
		default:
			fmt.Println("Geçersiz seçim!")
		}
	}
}

func customerGiriş() {
	fmt.Print("Müşteri kullanıcı adı: ")
	username := inputText()
	fmt.Print("Müşteri şifresi: ")
	password := inputText()
	for _, user := range users {
		if user.Username == username && user.Password == password && user.UserType == "customer" {
			currentUser = &user
			logEntry("Müşteri Girişi", true)
			for customerMenu() {
			}
			return
		}
	}
	logEntry("Hatalı Müşteri Girişi", false)
	fmt.Println("Giriş başarısız!")
}

func adminMenu() bool {
	fmt.Println("\nAdmin Menüsü\n1 - Müşteri Ekle\n2 - Müşteri Sil\n3 - Logları Görüntüle\n4 - Çıkış (Admin Paneli)")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		Customerekle()
	case 2:
		Customersil()
	case 3:
		displayLogs()
	case 4:
		fmt.Println("Admin panelinden çıkılıyor...")
		return false
	default:
		fmt.Println("Geçersiz seçim!")
	}
	return true
}

func customerMenu() bool {
	fmt.Println("\nMüşteri Menüsü\n1 - Profil Görüntüle\n2 - Çıkış (Müşteri Paneli)")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		profil()
	case 2:
		fmt.Println("Müşteri çıkış yapıyor...")
		return false
	default:
		fmt.Println("Geçersiz seçim!")
	}
	return true
}

func adminGiriş() {
	fmt.Print("Admin kullanıcı adı: ")
	username := inputText()
	fmt.Print("Admin şifresi: ")
	password := inputText()
	for _, user := range users {
		if user.Username == username && user.Password == password && user.UserType == "admin" {
			logEntry("Admin Girişi", true)
			for adminMenu() {
			}
			return
		}
	}
	logEntry("Hatalı Admin Girişi", false)
	fmt.Println("Giriş başarısız!")
}

func logEntry(action string, status bool) {
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Log dosyası oluşturulamadı:", err)
		return
	}
	defer f.Close()
	statusText := "Başarılı"
	if !status {
		statusText = "Başarısız"
	}
	entry := fmt.Sprintf("%s - %s: %s\n", time.Now().Format("2006-01-02 15:04:05"), action, statusText)
	f.WriteString(entry)
}

func inputText() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func Customerekle() {
	fmt.Print("Eklenecek müşteri kullanıcı adı: ")
	username := inputText()
	fmt.Print("Eklenecek müşteri şifresi: ")
	password := inputText()
	users = append(users, User{Username: username, Password: password, UserType: "customer"})
	fmt.Println("Müşteri eklendi.")
	logEntry("Müşteri Ekleme: "+username, true)
}

func Customersil() {
	fmt.Print("Silinecek müşteri kullanıcı adı: ")
	username := inputText()
	for i, user := range users {
		if user.Username == username && user.UserType == "customer" {
			users = append(users[:i], users[i+1:]...)
			fmt.Println("Müşteri silindi.")
			logEntry("Müşteri Silme: "+username, true)
			return
		}
	}
	fmt.Println("Müşteri bulunamadı.")
	logEntry("Müşteri Silme: "+username, false)
}

func displayLogs() {
	f, err := os.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Log dosyası okunamadı:", err)
		return
	}
	fmt.Println(string(f))
	logEntry("Log Görüntüleme", true)
	fmt.Println("\nDevam etmek için bir tuşa basın...")
	inputText()
}

func profil() {
	if currentUser != nil {
		fmt.Printf("Kullanıcı Adı: %s\nŞifre: ******\n", currentUser.Username)
		logEntry("Profil Görüntüleme: "+currentUser.Username, true)
	} else {
		fmt.Println("Profil bilgileri bulunamadı.")
	}
}
