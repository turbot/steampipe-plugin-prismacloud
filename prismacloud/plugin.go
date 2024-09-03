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
		ConnectionKeyColumns: []plugin.ConnectionKeyColumn{
			{
				Name:    "email",
				Hydrate: getCurrentUserEmail,
			},
		},
		TableMap: map[string]*plugin.Table{
			"prismacloud_account":    tablePrismacloudAccount(ctx),
			"prismacloud_alert":      tablePrismacloudAlert(ctx),
			"prismacloud_alert_rule": tablePrismacloudAlertRule(ctx),
			"prismacloud_compliance_breakdown_requirement_summary": tablePrismacloudComplianceBreakdownRequirementSummary(ctx),
			"prismacloud_compliance_breakdown_statistic":           tablePrismacloudComplianceBreakdownStatistic(ctx),
			"prismacloud_compliance_breakdown_summary":             tablePrismacloudComplianceBreakdownSummary(ctx),
			"prismacloud_compliance_requirement":                   tablePrismacloudComplianceRequirement(ctx),
			"prismacloud_compliance_standard":                      tablePrismacloudComplianceStandard(ctx),
			"prismacloud_iam_permission":                           tablePrismacloudIAMPermission(ctx),
			"prismacloud_iam_role":                                 tablePrismacloudIAMRole(ctx),
			"prismacloud_iam_user":                                 tablePrismacloudIAMUser(ctx),
			"prismacloud_inventory_api_endpoint":                   tablePrismacloudInventoryAPIEndpoint(ctx),
			"prismacloud_inventory_asset_explorer":                 tablePrismacloudInventoryAssetExplorer(ctx),
			"prismacloud_inventory_asset_view":                     tablePrismacloudInventoryAssetView(ctx),
			"prismacloud_inventory_workload":                       tablePrismacloudInventoryWorkload(ctx),
			"prismacloud_inventory_workload_container_image":       tablePrismacloudInventoryWorkloadContainerImage(ctx),
			"prismacloud_inventory_workload_host":                  tablePrismacloudInventoryWorkloadHost(ctx),
			"prismacloud_permission_group":                         tablePrismacloudPermissionGroup(ctx),
			"prismacloud_policy":                                   tablePrismacloudPolicy(ctx),
			"prismacloud_prioritized_vulnerability":                tablePrismacloudPrioritizedVulnerability(ctx),
			"prismacloud_report":                                   tablePrismacloudReport(ctx),
			"prismacloud_resource":                                 tablePrismacloudResource(ctx),
			"prismacloud_trusted_alert_ip":                         tablePrismacloudTrustedAlertIp(ctx),
			"prismacloud_vulnerability_asset":                      tablePrismacloudVulnerabilityAsset(ctx),
			"prismacloud_vulnerability_burndown":                   tablePrismacloudVulnerabilityBurndown(ctx),
			"prismacloud_vulnerability_overview":                   tablePrismacloudVulnerabilityOverview(ctx),
		},
	}
	return p
}
