package main

import (
	"bufio"
	"os"
	"strings"
)

func loadConfigFromFile(filename string, app *app) error {
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
			if domain != "@externalDNS." {
				app.CustomDNSMap[domain] = ip
			} else {
				app.ExternalDNSProvider = ip
			}

		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
