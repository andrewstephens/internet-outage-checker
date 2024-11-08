package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	defaultCheckInterval = 10
	defaultLogFile       = "connection_log.txt"
	testURL              = "http://www.google.com"
	timeout              = 10 * time.Second
)

type Config struct {
	checkInterval time.Duration
	logFilePath   string
	printOutput   bool
}

func parseFlags() *Config {
	intervalPtr := flag.Int("interval", defaultCheckInterval, "Check interval in seconds")
	flag.IntVar(intervalPtr, "i", defaultCheckInterval, "Check interval in seconds (shorthand)")

	logFilePtr := flag.String("logfile", defaultLogFile, "Log file path")
	flag.StringVar(logFilePtr, "l", defaultLogFile, "Log file path (shorthand)")

	printOutputPtr := flag.Bool("print", false, "Print output to console")
	flag.BoolVar(printOutputPtr, "p", false, "Print output to console (shorthand)")

	flag.Parse()

	return &Config{
		checkInterval: time.Duration(*intervalPtr) * time.Second,
		logFilePath:   *logFilePtr,
		printOutput:   *printOutputPtr,
	}
}

func setupLogger(logFilePath string) (*os.File, *log.Logger, error) {
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open log file: %w", err)
	}
	return file, log.New(file, "", log.LstdFlags), nil
}

func checkConnection() bool {
	client := &http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(testURL)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

func logStatus(logger *log.Logger, printOutput bool, connected, lastStatus bool) bool {
	if !connected && lastStatus {
		logger.Println("Internet connection lost")
		if printOutput {
			fmt.Println("Internet connection lost")
		}
	} else if connected && !lastStatus {
		logger.Println("Internet connection restored")
		if printOutput {
			fmt.Println("Internet connection restored")
		}
	}
	return connected
}

func main() {
	config := parseFlags()

	file, logger, err := setupLogger(config.logFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if config.printOutput {
		fmt.Printf("Monitoring internet connection. Check interval: %v, Log file: %s\n",
			config.checkInterval, config.logFilePath)
	}

	lastStatus := false
	for {
		connected := checkConnection()
		lastStatus = logStatus(logger, config.printOutput, connected, lastStatus)
		time.Sleep(config.checkInterval)
	}
}
