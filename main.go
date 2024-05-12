package main

import (
	"fmt"

	"github.com/miekg/dns"
)

var CustomDNSMap = make(map[string]string)

type dnsHandler struct{}

func main() {
	if err := loadCustomDNSMapFromFile("override.conf"); err != nil {
		fmt.Printf("Error loading custom DNS mappings: %v\n", err)
		return
	}

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
