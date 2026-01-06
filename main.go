package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

const (
	Latency = 1 * time.Millisecond
)

type SafeCounter struct {
	mu     sync.Mutex
	counts map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counts[key]++
}

func (c *SafeCounter) value() map[string]int {
	c.mu.Lock()
	defer c.mu.Unlock()
	clone := make(map[string]int)
	for k, v := range c.counts {
		clone[k] = v
	}

	return clone
}

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

func procesLine(line string) string {
	time.Sleep(Latency)
	if contains(line, "500 Internal Error") {
		return extractIP(line)
	}
	return ""
}

func runSequential() time.Duration {
	file, _ := os.Open("server.log")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	counts := make(map[string]int)

	start := time.Now()

	fmt.Println("This is Sequential")
	for scanner.Scan() {
		line := scanner.Text()
		ip := procesLine(line)
		if ip != "" {
			counts[ip]++
		}

	}
	return time.Since(start)

}

func runConcurrent() time.Duration {
	file, _ := os.Open("server.log")
	defer file.Close()

	counter := &SafeCounter{counts: make(map[string]int)}
	jobs := make(chan string, 50)

	var wg sync.WaitGroup

	workers := 10
	fmt.Println("This is Concurrent")
	start := time.Now()
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for line := range jobs {
				ip := procesLine(line)
				if ip != "" {
					counter.Inc(ip)

				}
			}
		}(i)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		jobs <- scanner.Text()
	}
	close(jobs)
	wg.Wait()

	return time.Since(start)
}

func extractIP(line string) string {
	fp := -1 //
	sp := -1 // to find and distinguish between first pipe and second pipe

	for i, char := range line {
		if char == '|' {
			if fp == -1 {
				fp = i
			} else {
				sp = i
				break
			}

		}
	}

	if fp != -1 && sp != -1 {
		return line[fp+2 : sp-1]
	}

	return ""

}

func contains(s string, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func main() {
	simulateLogs()
	seqTime := runSequential()
	fmt.Printf("Sequential workflow   Done in %v\n\n", seqTime)

	concTime := runConcurrent()
	fmt.Printf("Concurrent workflow   Done in %v\n\n", concTime)

}
