package main

import (
	"fmt"

	"github.com/miekg/dns"
)

func (h *dnsHandler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	msg := new(dns.Msg)
	msg.SetReply(r)
	msg.Authoritative = true

	for _, question := range r.Question {
		h.app.Queries.Printf("Received query: %s\n", question.Name)

		ip, ok := h.app.CustomDNSMap[question.Name]
		if ok {
			h.app.Info.Println("Domain found in custom DNS map")
			answers := createARecord(question.Name, ip, h.app)
			msg.Answer = append(msg.Answer, answers...)
		} else {
			h.app.Info.Println("Domain not found in custom DNS map. Fetching from External DNS.")
			answers, err := fetchFromExternalDNS(question.Name, h.app)
			if err != nil {
				h.app.Warn.Println("Failed to fetch from External DNS:", err)
				answers := createNXDOMAINRecord(question.Name, h.app)
				msg.Answer = append(msg.Answer, answers...)
			} else {
				msg.Answer = append(msg.Answer, answers...)
			}
		}
	}

	w.WriteMsg(msg)
}

func createARecord(name, ip string, app *app) []dns.RR {
	rr, err := dns.NewRR(fmt.Sprintf("%s A %s", name, ip))
	if err != nil {
		app.Error.Printf("Error creating A record: %v\n", err)
		return nil
	}
	return []dns.RR{rr}
}

func createNXDOMAINRecord(name string, app *app) []dns.RR {
	rr, err := dns.NewRR(fmt.Sprintf("%s SOA nonexistent-domain.com. hostmaster.nonexistent-domain.com. 1 10800 3600 604800 3600", name))
	if err != nil {
		app.Error.Printf("Error creating NXDOMAIN record: %v\n", err)
		return nil
	}
	return []dns.RR{rr}
}

func fetchFromExternalDNS(name string, app *app) ([]dns.RR, error) {
	c := new(dns.Client)
	m := new(dns.Msg)
	m.SetQuestion(name, dns.TypeA)
	r, _, err := c.Exchange(m, app.ExternalDNSProvider+":53")
	if err != nil {
		app.Error.Printf("Error fetching from external DNS: %v\n", err)
		return nil, err
	}
	return r.Answer, nil
}
