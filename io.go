package main

import (
	"bufio"
	"os"
	"strings"
)

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
