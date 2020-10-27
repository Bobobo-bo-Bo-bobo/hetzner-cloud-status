package main

import (
	"net/http"
)

// HTTPResult - result of the http_request calls
type HTTPResult struct {
	URL        string
	StatusCode int
	Status     string
	Header     http.Header
	Content    []byte
}

// HetznerAllServer - list of all hetzner servers
type HetznerAllServer struct {
	Server []HetznerServer `json:"servers"`
}

// HetznerServer - single server
type HetznerServer struct {
	ID              uint64            `json:"id"`
	Name            string            `json:"name"`
	Status          string            `json:"status"`
	ServerType      HetznerServerType `json:"server_type"`
	Datacenter      HetznerDatacenter `json:"datacenter"`
	Image           HetznerImage      `json:"image"`
	IncludedTraffic uint64            `json:"included_traffic"`
	IncomingTraffic uint64            `json:"ingoing_traffic"`
	OutgoingTraffic uint64            `json:"outgoing_traffic"`
	// ...
}

// HetznerServerType - server type
type HetznerServerType struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Cores       uint64  `json:"cores"`
	Memory      float64 `json:"memory"`
	Disk        uint64  `json:"disk"`
}

// HetznerDatacenter - datacenter
type HetznerDatacenter struct {
	ID          uint64                    `json:"id"`
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	Location    HetznerDatacenterLocation `json:"location"`
}

// HetznerDatacenterLocation - datacenter location
type HetznerDatacenterLocation struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Country     string  `json:"country"`
	City        string  `json:"city"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	NetworkZone string  `json:"network_zone"`
}

// HetznerImage - image info
type HetznerImage struct {
	ID          uint64 `json:"id"`
	Type        string `json:"type"`
	Status      string `json:"status"`
	Name        string `json:"name"`
	Description string `json:"description"`
	OSFlavor    string `json:"os_flavor"`
	OSVersion   string `json:"os_version"`
}
