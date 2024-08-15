package ip

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// GetPublicIP retrieves the public IP address by making a request to a public IP service.
func GetPublicIP() {
	response, err := http.Get("https://api.ipify.org")
	if err != nil {
		log.Fatal("Failed to retrieve public IP address:", err)
	}
	defer response.Body.Close()
	ip, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Failed to read public IP address:", err)
	}
	fmt.Println("Public IP address:", string(ip))
}
