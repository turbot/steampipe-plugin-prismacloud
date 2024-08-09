package prismacloud

import (
	"context"

	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// Note: You need vulnerabilityDashboard feature with View permission to access this endpoint. Verify if your permission group includes this feature using the Get Permission Group by ID endpoint. You can also check this in the Prisma Cloud console by ensuring that Dashboard > Vulnerability is enabled.

func tablePrismacloudVulnerabilitiyOverview(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prismacloud_vulnerabilitiy_overview",
		Description: "Provides an overview summary of vulnerabilities in the environment.",
		List: &plugin.ListConfig{
			Hydrate: getPrismacloudVulnerabilitiyOverview,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "total_vulnerable_runtime_assets",
				Description: "The total number of runtime assets that are vulnerable.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("OverviewSummary.TotalVulnerableRuntimeAssets"),
			},
			{
				Name:        "total_vulnerabilitiesin_runtime",
				Description: "The total number of vulnerabilities identified in runtime.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("OverviewSummary.TotalVulnerabilitiesinRuntime"),
			},
			{
				Name:        "total_remediated_in_runtime",
				Description: "The total number of vulnerabilities that have been remediated in runtime.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("OverviewSummary.TotalRemediatedInRuntime"),
			},
			{
				Name:        "values",
				Description: "Additional details related to the vulnerability overview.",
				Type:        proto.ColumnType_JSON,
			},
		}),
	}
}

//// LIST FUNCTION

func getPrismacloudVulnerabilitiyOverview(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_vulnerabilitiy_overview.getPrismacloudVulnerabilitiyOverview", "connection_error", err)
		return nil, err
	}

	vulnerability, err := api.GetVulnerabilityOverview(conn)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_vulnerabilitiy_overview.getPrismacloudVulnerabilitiyOverview", "api_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, vulnerability)

	return nil, nil
}
