package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/logrusorgru/aurora"

	"github.com/olekukonko/tablewriter"
)

func formatTraffic(i float64) string {
	if i < 1024 {
		return fmt.Sprintf("%.0f", i)
	}
	idx := int(math.Floor(math.Log(i) / math.Log(1024.)))
	v := i / math.Pow(1024., float64(idx))
	return fmt.Sprintf("%.3f%s", v, mapSI[idx])
}

func printVolume(v HetznerVolume) string {
	status, found := statusVolumeMap[v.Status]
	if !found {
		status = aurora.White(v.Status)
	}

	return fmt.Sprintf("%s (%s, %d GB)", v.Name, status, v.Size)
}

func printAllServers(a HetznerAllServer, v HetznerAllVolumes) {
	var serverMap = make(map[uint64]HetznerServer)
	var volumeMap = make(map[uint64]HetznerVolume)
	var vols []string

	// map server ID to server object
	for _, srv := range a.Server {
		serverMap[srv.ID] = srv
	}

	// map volume ID to volume object
	for _, vol := range v.Volume {
		volumeMap[vol.ID] = vol
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Status", "Type", "Cores", "RAM", "HDD", "Zone", "Data center", "Traffic in", "Traffic out", "Traffic %", "Volumes"})

	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)

	for _, srv := range a.Server {
		status, found := statusServerMap[srv.Status]
		if !found {
			status = aurora.White(srv.Status)
		}

		// discard data
		vols = nil
		for _, vid := range srv.Volumes {
			vol, found := volumeMap[vid]
			if found {
				vols = append(vols, printVolume(vol))
			} else {
				// should not happen!
				vols = append(vols, "?")
			}
		}

		volumes := "-"
		if len(vols) > 0 {
			volumes = strings.Join(vols, ", ")
		}

		table.Append([]string{
			srv.Name,
			fmt.Sprintf("%s", status),
			srv.ServerType.Description,
			fmt.Sprintf("%d", srv.ServerType.Cores),
			fmt.Sprintf("%.0f GB", srv.ServerType.Memory),
			fmt.Sprintf("%d GB", srv.ServerType.Disk),
			srv.Datacenter.Location.NetworkZone,
			fmt.Sprintf("%s (%s, %s)", srv.Datacenter.Description, srv.Datacenter.Location.Country, srv.Datacenter.Location.City),
			formatTraffic(float64(srv.IncomingTraffic)),
			formatTraffic(float64(srv.OutgoingTraffic)),
			fmt.Sprintf("%3.2f%%", 100.0*float64(srv.IncomingTraffic+srv.OutgoingTraffic)/float64(srv.IncludedTraffic)),
			volumes,
		})
	}

	table.Render()
}
