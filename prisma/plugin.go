package prisma

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-prisma",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		DefaultTransform: transform.FromCamel(),
		TableMap: map[string]*plugin.Table{
			"prisma_account":          tablePrismaAccount(ctx),
			"prisma_alert_rule":       tablePrismaAlertRule(ctx),
			"prisma_permission_group": tablePrismaPermissionGroup(ctx),
			"prisma_policy":           tablePrismaPolicy(ctx),
			"prisma_report":           tablePrismaReport(ctx),
			"prisma_resource":         tablePrismaResource(ctx),
			"prisma_trusted_alert_ip": tablePrismaTrustedAlertIp(ctx),
			"prisma_user":             tablePrismaUser(ctx),
		},
	}
	return p
}
