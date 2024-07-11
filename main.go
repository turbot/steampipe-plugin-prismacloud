package main

import (
	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: prismacloud.Plugin})
}
