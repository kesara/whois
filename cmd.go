// Whois command line program
package main

import (
	"fmt"
	"github.com/kesara/whois/whois"
	"os"
)

const (
	WHOIS_SERVER = "whois.iana.org"
	WHOIS_PORT   = "43"
)

func main() {
	if len(os.Args) > 4 || len(os.Args) < 2 {
		fmt.Println(fmt.Sprintf(
			"Usage: %s <domain> [host [port]]", os.Args[0]))
		os.Exit(1)
	}

	var domain string
	var host whois.WhoisServer

	host.Server = WHOIS_SERVER
	host.Port = WHOIS_PORT

	domain = os.Args[1]
	if len(os.Args) > 2 {
		host.Server = os.Args[2]
	}
	if len(os.Args) > 3 {
		host.Port = os.Args[3]
	}

	response, err := whois.Query(host, domain)
	if err != nil {
		fmt.Println(fmt.Sprintln("Error: %s", err))
		os.Exit(1)
	}

	fmt.Println(response)
	os.Exit(0)
}
