package main

import (
	"log"
	"os"

	"github.com/miekg/dns"
)

type dnsHandler struct{ *app }
type app struct {
	CustomDNSMap        map[string]string
	ExternalDNSProvider string
	Handler             *dnsHandler
	Server              *dns.Server
	Queries             *log.Logger
	Info                *log.Logger
	Error               *log.Logger
	Warn                *log.Logger
	BlockLists          []string
	DNSResponses        map[string]dns.Msg
}

func (a *app) initialize() error {
	a.ExternalDNSProvider = "8.8.8.8"
	a.CustomDNSMap = make(map[string]string)
	a.DNSResponses = make(map[string]dns.Msg)
	a.Queries = log.New(os.Stdout, "QUERY: ", log.Ldate|log.Ltime|log.Lshortfile)
	a.Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	a.Error = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	a.Warn = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	a.Handler = new(dnsHandler)
	a.Handler.app = a
	a.Server = &dns.Server{
		Addr:    ":53",
		Net:     "udp",
		Handler: a.Handler,
		UDPSize: 65535,
	}
	return loadConfigFromFile("override.conf", a)
}

func (a *app) start() error {
	return a.Server.ListenAndServe()
}
