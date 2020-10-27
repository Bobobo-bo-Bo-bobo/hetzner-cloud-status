package main

import (
	"fmt"
	"math"
	"os"

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

func printAllServers(a HetznerAllServer) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Status", "Type", "Cores", "RAM", "HDD", "Zone", "Data center", "Traffic in", "Traffic out", "Traffic %"})

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
		status, found := statusMap[srv.Status]
		if !found {
			status = aurora.White(srv.Status)
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
		})
	}

	table.Render()
}
