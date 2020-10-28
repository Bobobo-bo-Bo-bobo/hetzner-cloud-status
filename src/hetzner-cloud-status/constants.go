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

const hetznerAllServersURL = "https://api.hetzner.cloud/v1/servers?sort=name"
const hetznerAllVolumesURL = "https://api.hetzner.cloud/v1/volumes?sort=name"

var statusServerMap = map[string]aurora.Value{
	"deleting":     aurora.Bold(aurora.Red("deleting")),
	"initializing": aurora.Cyan("initialising"),
	"migrating":    aurora.Blink(aurora.Green("migrating")),
	"off":          aurora.Red("off"),
	"rebuilding":   aurora.Blue("rebuilding"),
	"running":      aurora.Bold(aurora.Green("running")),
	"starting":     aurora.Faint(aurora.Green("starting")),
	"stopping":     aurora.Faint(aurora.Yellow("stopping")),
	"unknown":      aurora.Bold(aurora.Blink(aurora.Magenta("unknown"))),
}

var statusVolumeMap = map[string]aurora.Value{
	"creating":  aurora.Cyan("creating"),
	"available": aurora.Bold(aurora.Green("available")),
}

var mapSI = []string{
	"",
	"k",
	"M",
	"G",
	"T",
	"P",
}
