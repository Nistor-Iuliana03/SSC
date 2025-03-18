package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)

func main() {
	// Obține lista interfețelor de rețea
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatalf("Eroare la găsirea interfețelor: %v", err)
	}

	// Afișează fiecare interfață
	fmt.Println("Interfețe de rețea disponibile:")
	for _, device := range devices {
		fmt.Printf("Nume: %s\n", device.Name)
		if len(device.Description) > 0 {
			fmt.Printf("Descriere: %s\n", device.Description)
		}
		fmt.Println("Adrese:")
		for _, address := range device.Addresses {
			fmt.Printf("- IP: %s, Masca: %s\n", address.IP, address.Netmask)
		}
		fmt.Println()
	}
}
