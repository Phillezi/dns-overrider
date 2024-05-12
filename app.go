package main

import "github.com/miekg/dns"

type dnsHandler struct{ *app }
type app struct {
	CustomDNSMap        map[string]string
	ExternalDNSProvider string
	Handler             *dnsHandler
	Server              *dns.Server
}

func (a *app) initialize() error {
	a.ExternalDNSProvider = "8.8.8.8"
	a.CustomDNSMap = make(map[string]string)
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
