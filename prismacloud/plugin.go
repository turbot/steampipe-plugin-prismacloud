package prismacloud

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-prismacloud",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		DefaultTransform: transform.FromCamel(),
		TableMap: map[string]*plugin.Table{
			"prismacloud_account":            tablePrismaAccount(ctx),
			"prismacloud_alert":              tablePrismaAlert(ctx),
			"prismacloud_alert_rule":         tablePrismaAlertRule(ctx),
			"prismacloud_compliance_posture": tablePrismaCompliancePosture(ctx),
			"prismacloud_permission_group":   tablePrismaPermissionGroup(ctx),
			"prismacloud_policy":             tablePrismaPolicy(ctx),
			"prismacloud_report":             tablePrismaReport(ctx),
			"prismacloud_resource":           tablePrismaResource(ctx),
			"prismacloud_trusted_alert_ip":   tablePrismaTrustedAlertIp(ctx),
			"prismacloud_user":               tablePrismaUser(ctx),
		},
	}
	return p
}
