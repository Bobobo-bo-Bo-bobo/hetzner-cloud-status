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
	var parsedVolumes HetznerAllVolumes
	var useProxy string

	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}

	token, err := readSingleLine(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	environment := getEnvironment()

	// Like curl the lower case variant has precedence.
	httpsproxy, lcfound := environment["https_proxy"]
	httpsProxy, ucfound := environment["HTTPS_PROXY"]

	if lcfound {
		useProxy = httpsproxy
	} else if ucfound {
		useProxy = httpsProxy
	}

	// get list of instances
	servers, err := httpRequest(hetznerAllServersURL, "GET", nil, nil, useProxy, token)
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

	// get list of shared storage volumes
	volumes, err := httpRequest(hetznerAllVolumesURL, "GET", nil, nil, useProxy, token)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if volumes.StatusCode != http.StatusOK {
		fmt.Fprintln(os.Stderr, volumes.Status)
		os.Exit(1)
	}

	err = json.Unmarshal(volumes.Content, &parsedVolumes)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	printAllServers(parsedServers, parsedVolumes)
}
