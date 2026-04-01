package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Pobieranie nazwy hosta (Hostname)
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "Nieznany"
	}

	// Pobieranie adresu IP
	var ip string
	addrs, err := net.LookupIP(hostname)
	if err == nil {
		for _, addr := range addrs {
			// Sprawdzanie czy to adres IPv4
			if ipv4 := addr.To4(); ipv4 != nil && !addr.IsLoopback() {
				ip = ipv4.String()
				break
			}
		}
	}

	// Pobieranie wersji aplikacji ze zmiennej srodowiskowej
	version := os.Getenv("APP_VERSION")
	if version == "" {
		version = "Brak wersji"
	}

	// Dynamiczne generowanie strony HTML
	fmt.Fprintf(w, "<h1>Lab 5 - Aplikacja Go</h1><p>IP: %s</p><p>Hostname: %s</p><p>Wersja: %s</p>", ip, hostname, version)
}

func main() {
	// Rejestracja funkcji obslugujacej zapytania
	http.HandleFunc("/", handler)
	log.Println("Aplikacja dziala na porcie 8080...")
	// Uruchomienie serwera aplikacji
	log.Fatal(http.ListenAndServe(":8080", nil))
}