package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func getLocalIP(interfaceName string) {
	cmd := exec.Command("ipconfig", "getifaddr", interfaceName)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Failed to retrieve IP address for interface:", interfaceName)
		log.Fatal(err)
	}
	fmt.Println("Local IP address:", string(out))
}

func listNetworkInterfaces() []string {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal("Failed to retrieve network interfaces:", err)
	}

	var interfaceNames []string
	for _, iface := range interfaces {
		interfaceNames = append(interfaceNames, iface.Name)
		fmt.Println(iface.Name)
	}
	return interfaceNames
}

func chooseInterface() string {
	interfaces := listNetworkInterfaces()
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

func getLocalIPWithInterfaceCheck() {
	interfaceName := "en0"
	cmd := exec.Command("ipconfig", "getifaddr", interfaceName)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Interface does not exist or failed to retrieve local IP address.")
		fmt.Println("Available network interfaces:")
		interfaceName = chooseInterface()
	} else {
		fmt.Println("Local IP address for en0:", string(out))
		return
	}

	// Retry getting the local IP with the chosen interface
	getLocalIP(interfaceName)
}

func getPublicIP() {
	response, err := http.Get("https://api.ipify.org")
	if err != nil {
		log.Fatal("Failed to retrieve public IP address:", err)
	}
	defer response.Body.Close()
	ip, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Failed to read public IP address:", err)
	}
	fmt.Println("Public IP address:", string(ip))
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "jv",
		Short: "JV Command Line Tool",
		Long:  `A CLI tool to fetch IP addresses, both local and public.`,
	}

	var ipCmd = &cobra.Command{
		Use:   "ip",
		Short: "IP related commands",
	}

	var localCmd = &cobra.Command{
		Use:   "local",
		Short: "Get the local IP address",
		Run: func(cmd *cobra.Command, args []string) {
			getLocalIPWithInterfaceCheck()
		},
	}

	var publicCmd = &cobra.Command{
		Use:   "public",
		Short: "Get the public IP address",
		Run: func(cmd *cobra.Command, args []string) {
			getPublicIP()
		},
	}

	rootCmd.AddCommand(ipCmd)
	ipCmd.AddCommand(localCmd)
	ipCmd.AddCommand(publicCmd)

	viper.BindPFlags(rootCmd.Flags())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
