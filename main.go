package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/masterzen/winrm"
)

//InitWinRMShell InitWinRMShell
func InitWinRMShell(username string, password string, ipAddress string, port string) (*winrm.Client, error) {
	intPort, _ := strconv.Atoi(port)
	endpoint := winrm.NewEndpoint(ipAddress, intPort, true, true, nil, nil, nil, 0)
	params := winrm.DefaultParameters
	params.TransportDecorator = func() winrm.Transporter { return &winrm.ClientNTLM{} }

	client, err := winrm.NewClientWithParameters(endpoint, username, password, params)
	if err != nil {
		return nil, err
	}
	return client, nil

}

//VerifyWinRM VerifyWinRM
func VerifyWinRM(username string, password string, ipAddress string, port string) bool {
	intPort, _ := strconv.Atoi(port)
	endpoint := winrm.NewEndpoint(ipAddress, intPort, true, true, nil, nil, nil, 0)

	params := winrm.DefaultParameters
	params.TransportDecorator = func() winrm.Transporter { return &winrm.ClientNTLM{} }

	client, err := winrm.NewClientWithParameters(endpoint, username, password, params)
	if err != nil {
		return false
	}

	stdout, _, _, _ := client.RunWithString("hostname", "")
	if strings.TrimSpace(stdout) == "" {
		return false
	}
	return true
}

func main() {
	// Credentials
	username := "testuser"
	password := "securepassword"
	ip := "securepassword"
	port := "5986"
	command := "hostname"

	// Connect
	connection, _ := InitWinRMShell(username, password, ip, port)
	sdtcheck := VerifyWinRM(username, password, ip, port)

	if sdtcheck == false {
		fmt.Println("Bağlantı sağlanamadı")
		os.Exit(1)
	}

	// Run command
	// encoder := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder()
	// encoded, _ := encoder.String(command)
	// command = base64.StdEncoding.EncodeToString([]byte(encoded))

	stdout, stderr, _, err := connection.RunWithString("powershell.exe -command "+command, "")
	// stdout, stderr, _, err := connection.RunWithString("powershell.exe -encodedCommand "+command, "")
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	} else if stderr != "" {
		fmt.Println("stderr")
		fmt.Println(stderr)
	} else {
		fmt.Println("stdout")
	fmt.Println(stdout)
	}

}
