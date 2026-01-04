package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func simulateLogs() {
	file, _ := os.Create("server.log")
	defer file.Close()

	writer := bufio.NewWriter(file)
	statuses := []string{"200 OK", "404 Not Found", "500 Internal Error", "403 Forbidden"}
	ips := []string{"192.168.1.10", "10.0.0.5", "172.16.0.22", "8.8.8.8"}

	fmt.Println("Generating massive log file...")
	for i := 0; i < 100000; i++ {
		ip := ips[rand.Intn(len(ips))]
		status := statuses[rand.Intn(len(statuses))]

		// Write line: "TIMESTAMP | IP | STATUS"
		line := fmt.Sprintf("%s | %s | %s\n", time.Now().Format(time.RFC3339), ip, status)
		writer.WriteString(line)
	}
	writer.Flush()
	fmt.Println("Done! 'server.log' created.")
}

func analyseLogs() {
	file, err := os.Open("server.log")
	if err != nil {
		fmt.Println("Error opening the file", err)
		return
	}

	defer file.Close()
	errorMap := make(map[string]int)
	scanner := bufio.NewScanner(file)
	start := time.Now()

	for scanner.Scan() {
		line := scanner.Text()
		if contains(line, "500 Internal Error") {
			ip := extractIP(line)
			errorMap[ip] = errorMap[ip] + 1
		}
	}

	duration := time.Since(start)

	// MY own report made right now
	fmt.Println("Analysis took", duration)
	fmt.Println("These are the attackers we found :")

	for ip, count := range errorMap {
		fmt.Printf("IP: %s has crashed %d times \n", ip, count)
	}
}

func extractIP(line string) string {
	return ""
}

func contains(s string, substr string) bool {
	return false
}
