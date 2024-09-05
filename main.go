package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "start":
		startApache()
	case "stop":
		stopApache()
	case "restart":
		restartApache()
	case "help":
		printHelp()
	default:
		fmt.Println("Invalid command. Use 'help' for usage information.")
		os.Exit(1)
	}
}

func startApache() {
	fmt.Println("Starting Apache server...")
	cmd := exec.Command("sudo", "systemctl", "start", "apache2")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error starting Apache:", err)
		os.Exit(1)
	}
	fmt.Println("Apache server started successfully.")
}

func stopApache() {
	fmt.Println("Stopping Apache server...")
	cmd := exec.Command("sudo", "systemctl", "stop", "apache2")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error stopping Apache:", err)
		os.Exit(1)
	}
	fmt.Println("Apache server stopped successfully.")
}

func restartApache() {
	fmt.Println("Restarting Apache server...")
	cmd := exec.Command("sudo", "systemctl", "restart", "apache2")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error restarting Apache:", err)
		os.Exit(1)
	}
	fmt.Println("Apache server restarted successfully.")
}

func printHelp() {
	fmt.Println("Usage: apache-control <command>")
	fmt.Println("\nCommands:")
	fmt.Println("  start    Start the Apache server")
	fmt.Println("  stop     Stop the Apache server")
	fmt.Println("  restart  Restart the Apache server")
	fmt.Println("  help     Show this help message")
}