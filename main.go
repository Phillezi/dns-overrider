package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/miekg/dns"
)

var CustomDNSMap = make(map[string]string)

type dnsHandler struct{}

func (h *dnsHandler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	msg := new(dns.Msg)
	msg.SetReply(r)
	msg.Authoritative = true

	for _, question := range r.Question {
		fmt.Printf("Received query: %s\n", question.Name)

		ip, ok := CustomDNSMap[question.Name]
		if ok {
			fmt.Println("Domain found in custom DNS map")
			answers := createARecord(question.Name, ip)
			msg.Answer = append(msg.Answer, answers...)
		} else {
			fmt.Println("Domain not found in custom DNS map. Fetching from Google DNS.")
			answers, err := fetchFromGoogleDNS(question.Name)
			if err != nil {
				fmt.Println("Failed to fetch from Google DNS:", err)
				answers := createNXDOMAINRecord(question.Name)
				msg.Answer = append(msg.Answer, answers...)
			} else {
				msg.Answer = append(msg.Answer, answers...)
			}
		}
	}

	w.WriteMsg(msg)
}

func createARecord(name, ip string) []dns.RR {
	rr, err := dns.NewRR(fmt.Sprintf("%s A %s", name, ip))
	if err != nil {
		fmt.Printf("Error creating A record: %v\n", err)
		return nil
	}
	return []dns.RR{rr}
}

func createNXDOMAINRecord(name string) []dns.RR {
	rr, err := dns.NewRR(fmt.Sprintf("%s SOA nonexistent-domain.com. hostmaster.nonexistent-domain.com. 1 10800 3600 604800 3600", name))
	if err != nil {
		fmt.Printf("Error creating NXDOMAIN record: %v\n", err)
		return nil
	}
	return []dns.RR{rr}
}

func fetchFromGoogleDNS(name string) ([]dns.RR, error) {
	c := new(dns.Client)
	m := new(dns.Msg)
	m.SetQuestion(name, dns.TypeA)
	r, _, err := c.Exchange(m, "8.8.8.8:53")
	if err != nil {
		return nil, err
	}
	return r.Answer, nil
}

func loadCustomDNSMapFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) == 2 {
			domain := strings.TrimSpace(parts[0]) + "."
			ip := strings.TrimSpace(parts[1])
			CustomDNSMap[domain] = ip
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

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
