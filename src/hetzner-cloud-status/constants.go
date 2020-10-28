package main

import (
	"github.com/logrusorgru/aurora"
)

const name = "hetzner-cloud-status"
const version = "1.0.0-20201028"
const userAgent = name + "/" + version

const versionText = `%s version %s
Copyright (C) 2020 by Andreas Maus <maus@ypbind.de>
This program comes with ABSOLUTELY NO WARRANTY.

%s is distributed under the Terms of the GNU General
Public License Version 3. (http://www.gnu.org/copyleft/gpl.html)

Build with go version: %s
`

const hetznerAllServers = "https://api.hetzner.cloud/v1/servers?sort=name"

var statusMap = map[string]aurora.Value{
	"running":      aurora.Bold(aurora.Green("running")),
	"initializing": aurora.Cyan("initialising"),
	"starting":     aurora.Faint(aurora.Green("starting")),
	"stopping":     aurora.Faint(aurora.Yellow("stopping")),
	"off":          aurora.Red("off"),
	"deleting":     aurora.Bold(aurora.Red("deleting")),
	"migrating":    aurora.Blink(aurora.Green("migrating")),
	"rebuilding":   aurora.Blue("rebuilding"),
	"unknown":      aurora.Bold(aurora.Blink(aurora.Magenta("unknown"))),
}

var mapSI = []string{
	"",
	"k",
	"M",
	"G",
	"T",
	"P",
}
