package main

import (
	"github.com/turbot/steampipe-plugin-prisma/prisma"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: prisma.Plugin})
}
