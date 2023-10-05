package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	// Pengaturan URL dan data autentikasi
	fortigateURL := "https://192.168.1.99:10443/remote/saml/login"
	username := "your_username"
	password := "your_password"

	// Membuat body permintaan
	data := fmt.Sprintf(`{"user": "%s", "passwd": "%s"}`, username, password)
	payload := strings.NewReader(data)

	// Membuat permintaan POST ke FortiGate
	req, err := http.NewRequest("POST", fortigateURL, payload)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Add("Content-Type", "application/json")

	// Melakukan permintaan HTTP
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer res.Body.Close()

	// Membaca respon dari FortiGate
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Menampilkan respon dari FortiGate
	fmt.Println("Response:", string(body))
}
