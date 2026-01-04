# File_audit_system
A lightweight, memory-efficient CLI tool written in Go to audit server logs. It utilizes buffered I/O to process massive datasets with minimal memory footprint.It is purely backend friendly and as of now i have used this project to use the maximum possible usage of go without channels and mutexes


# GoLogAudit

A high-performance command-line tool designed to parse and analyze large-scale server logs. 

Built with Go, this project demonstrates **stream processing** techniques to handle large files (e.g., 100,000+ lines) without loading the entire dataset into RAM. It effectively simulates how production-grade agents (like Datadog or Splunk forwarders) process data on resource-constrained servers.

## 🚀 Key Features
* **Memory Efficiency:** Uses `bufio.Scanner` to stream data line-by-line, ensuring $O(1)$ memory usage regardless of file size.
* **Performance:** Capable of parsing and aggregating statistics from massive log files in milliseconds.
* **Data Aggregation:** utilizes Hash Maps to perform real-time frequency analysis of error codes and IP addresses.
* **Zero-Dependency:** Written in pure Go (Standard Library) with no external packages.

## 🛠️ Technical Concepts Applied
* **Buffered I/O:** Efficient file reading using `os` and `bufio` packages.
* **String Slicing:** Manual parsing logic to extract data points without regex overhead.
* **Resource Management:** Proper use of `defer` for safe file handle cleanup.
* **Algorithmic Logic:** Custom implementation of substring search and aggregation loops.

## 💻 How to Run
1. Clone the repository
2. Run the simulation to generate a dummy 10MB log file:
   ```bash
   go run main.go
