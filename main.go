package main

import (
	"crypto/rand"
  	"fmt"
	 log "github.com/sirupsen/logrus"
)

func randomMAC() (string, error) {
        bytes := make([]byte, 6)
        if _, err := rand.Read(bytes); err != nil {
    		return "", err
  	}
	bytes[0] |= 2

	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", bytes[0], bytes[1], bytes[2], bytes[3], bytes[4], bytes[5]), nil
}

func main() {
	MACAddr, err := randomMAC()
	if err != nil {
		log.Fatalf("Failed to generate random MAC address.")
	}
	fmt.Println(MACAddr)
}
