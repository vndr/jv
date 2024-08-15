package ip

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// ListNetworkInterfaces lists all network interfaces available on the machine.
func ListNetworkInterfaces() []string {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal("Failed to retrieve network interfaces:", err)
	}

	var interfaceNames []string
	for _, iface := range interfaces {
		interfaceNames = append(interfaceNames, iface.Name)
	}
	return interfaceNames
}

// ChooseInterface prompts the user to choose a network interface from the list.
func ChooseInterface() string {
	interfaces := ListNetworkInterfaces()
	fmt.Println("Please choose an interface by typing its name:")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Interface: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Failed to read input:", err)
		}

		input = strings.TrimSpace(input)
		for _, iface := range interfaces {
			if input == iface {
				return input
			}
		}

		fmt.Println("Invalid interface name. Please choose again:")
	}
}

// GetLocalIP retrieves the local IP address for a given network interface.
func GetLocalIP(interfaceName string) {
	var cmd *exec.Cmd

	// Use different commands depending on the OS
	if runtime.GOOS == "darwin" {
		// macOS uses ipconfig
		cmd = exec.Command("ipconfig", "getifaddr", interfaceName)
	} else if runtime.GOOS == "linux" {
		// Linux uses ip or ifconfig
		cmd = exec.Command("ip", "addr", "show", interfaceName)
	} else {
		log.Fatalf("Unsupported OS: %s", runtime.GOOS)
	}

	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("Failed to retrieve IP address for interface: %s\n", interfaceName)
		log.Fatal(err)
	}

	// For Linux, extract the IP address from the command output
	if runtime.GOOS == "linux" {
		// Find the line containing "inet" and extract the IP address
		lines := strings.Split(string(out), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "inet ") {
				// Example format: "inet 192.168.1.5/24 brd 192.168.1.255 scope global dynamic noprefixroute wlp0s20f3"
				parts := strings.Fields(line)
				if len(parts) > 1 {
					fmt.Printf("Local IP address: %s\n", parts[1])
					return
				}
			}
		}
		log.Fatalf("Failed to parse IP address from interface: %s\n", interfaceName)
	} else {
		fmt.Printf("Local IP address: %s\n", string(out))
	}
}

// GetLocalIPWithInterfaceCheck retrieves the local IP address, allowing the user to choose an interface if needed.
func GetLocalIPWithInterfaceCheck() {
	interfaceName := "en0"
	if runtime.GOOS == "linux" {
		// You might want to use a more common interface name for Linux
		interfaceName = "eth0"
	}

	cmd := exec.Command("ipconfig", "getifaddr", interfaceName)
	if runtime.GOOS == "linux" {
		cmd = exec.Command("ip", "addr", "show", interfaceName)
	}

	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Interface does not exist or failed to retrieve local IP address.")
		fmt.Println("Available network interfaces:")
		interfaceName = ChooseInterface()
	} else {
		fmt.Printf("Local IP address for %s: %s\n", interfaceName, string(out))
		return
	}

	// Retry getting the local IP with the chosen interface
	GetLocalIP(interfaceName)
}
