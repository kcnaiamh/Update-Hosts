package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Input for domain name and IP address
	var domain, ip string
	fmt.Print("Enter the domain name: ")
	fmt.Scanln(&domain)
	fmt.Print("Enter the IP address: ")
	fmt.Scanln(&ip)

	// Hosts file paths
	windowsHostsFile := "/mnt/c/Windows/System32/drivers/etc/hosts"
	linuxHostsFile := "/etc/hosts"

	// Entry to be added
	entry := fmt.Sprintf("%s %s\n", ip, domain)

	// Writing entry to Windows hosts file
	err := appendToFile(windowsHostsFile, entry)
	if err != nil {
		fmt.Printf("Failed to write to Windows hosts file: %v\n", err)
	} else {
		fmt.Printf("Successfully written to Windows hosts file.\n")
	}

	// Writing entry to Linux hosts file
	err = appendToFile(linuxHostsFile, entry)
	if err != nil {
		fmt.Printf("Failed to write to Linux hosts file: %v\n", err)
	} else {
		fmt.Printf("Successfully written to Linux hosts file.\n")
	}
}

// appendToFile appends the provided text to the specified file
func appendToFile(filePath, text string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(text)
	if err != nil {
		return err
	}
	writer.Flush()

	return nil
}

