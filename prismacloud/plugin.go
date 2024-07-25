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
		DefaultGetConfig: &plugin.GetConfig{
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"object not found"}),
			},
		},
		TableMap: map[string]*plugin.Table{
			"prismacloud_account":    tablePrismaAccount(ctx),
			"prismacloud_alert":      tablePrismaAlert(ctx),
			"prismacloud_alert_rule": tablePrismaAlertRule(ctx),
			"prismacloud_compliance_breakdown_requirement_summary": tablePrismaComplianceBreakdownRequirementSummary(ctx),
			"prismacloud_compliance_breakdown_statistic":           tablePrismaComplianceBreakdownStatistic(ctx),
			"prismacloud_compliance_breakdown_summary":             tablePrismaComplianceBreakdownSummary(ctx),
			"prismacloud_compliance_requirement":                   tablePrismaComplianceRequirement(ctx),
			"prismacloud_compliance_standard":                      tablePrismaComplianceStandard(ctx),
			"prismacloud_iam_permission":                           tablePrismaIAMPermission(ctx),
			"prismacloud_iam_role":                                 tablePrismaIAMRole(ctx),
			"prismacloud_iam_user":                                 tablePrismaIAMUser(ctx),
			"prismacloud_inventory_asset_explorer": tablePrismaInventoryAssetExplorer(ctx),
			"prismacloud_permission_group":                         tablePrismaPermissionGroup(ctx),
			"prismacloud_policy":                                   tablePrismaPolicy(ctx),
			"prismacloud_prioritized_vulnerabilitiy":               tablePrismaPrioritizedVulnerabilitiy(ctx),
			"prismacloud_report":                                   tablePrismaReport(ctx),
			"prismacloud_resource":                                 tablePrismaResource(ctx),
			"prismacloud_trusted_alert_ip":                         tablePrismaTrustedAlertIp(ctx),
			"prismacloud_vulnerabilitiy_asset":                     tablePrismaVulnerabilitiyAsset(ctx),
			"prismacloud_vulnerabilitiy_burndown":                  tablePrismaVulnerabilitiyBurndown(ctx),
			"prismacloud_vulnerabilitiy_overview":                  tablePrismaVulnerabilitiyOverview(ctx),
		},
	}
	return p
}
