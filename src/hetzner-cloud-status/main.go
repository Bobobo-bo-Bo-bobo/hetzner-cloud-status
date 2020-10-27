package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
)

func usageHeader() {
	fmt.Printf(versionText, name, version, name, runtime.Version())
}

func usage() {
	usageHeader()
	fmt.Printf("\nUsage: hetzner-cloud-status <api_token_file>\n")
}

func main() {
	var parsedServers HetznerAllServer

	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}

	token, err := readSingleLine(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	servers, err := httpRequest(hetznerAllServers, "GET", nil, nil, token)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if servers.StatusCode != http.StatusOK {
		fmt.Fprintln(os.Stderr, servers.Status)
		os.Exit(1)
	}

	err = json.Unmarshal(servers.Content, &parsedServers)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	printAllServers(parsedServers)
}
