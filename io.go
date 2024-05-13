package main

import (
	"bufio"
	"fmt"
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
				if domain != "@blockLists." {
					app.CustomDNSMap[domain] = ip
				} else {
					app.BlockLists = strings.Split(ip, ",")
					for i := range len(app.BlockLists) {
						app.BlockLists[i] = strings.TrimSpace(app.BlockLists[i])
					}
				}
			} else {
				app.ExternalDNSProvider = ip
			}

		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return loadBlocklists(app)
}

func loadBlocklists(app *app) error {
	for i := range len(app.BlockLists) {
		fmt.Println("Loading: " + app.BlockLists[i])
		err := loadBlocklist(app.BlockLists[i], app)
		if err != nil {
			return err
		}
	}
	return nil
}

func loadBlocklist(filename string, app *app) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) == 2 {
			domain := strings.TrimSpace(parts[1]) + "."
			ip := strings.TrimSpace(parts[0])
			app.CustomDNSMap[domain] = ip
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
