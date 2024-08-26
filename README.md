# Update Hosts Tool

This is a simple Go-based tool that allows you to append a domain name and its corresponding IP address to the `hosts` file on both Windows and Linux systems.
Built for CTF and Pentesting purposes. 

## Prerequisites

- **Go Language**: Ensure you have Go installed on your machine. You can download it from [golang.org](https://golang.org/).
- **Permissions**: You might need elevated permissions to modify system files such as `/etc/hosts` on Linux or `C:\Windows\System32\drivers\etc\hosts` on Windows.

## Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/kcnaiamh/update-hosts-tool.git
   cd update-hosts-tool
   ```
2. **Build the Tool**:
   ```
   go build -o update_hosts update_hosts.go
   ```

Put the binary in your windows home directory. Start an administrative terminal. Go to WSL. Run the binary with sudo.
