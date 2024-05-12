package main

import (
	"fmt"

	"github.com/miekg/dns"
)

var CustomDNSMap = make(map[string]string)
var ExternalDNSProvider = "8.8.8.8"

type dnsHandler struct{}

func main() {
	if err := loadConfigFromFile("override.conf"); err != nil {
		fmt.Printf("Error loading custom DNS mappings: %v\n", err)
		return
	}

	fmt.Println("External DNS Provider: " + ExternalDNSProvider)
	fmt.Println(CustomDNSMap)

	handler := new(dnsHandler)
	server := &dns.Server{
		Addr:    ":53",
		Net:     "udp",
		Handler: handler,
		UDPSize: 65535,
	}

	fmt.Println("Starting DNS server on port 53")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Failed to start server: %s\n", err.Error())
	}
}
