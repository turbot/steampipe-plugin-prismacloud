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
		Description: "The overview summary of vulnerabilitiy.",
		List: &plugin.ListConfig{
			Hydrate: getPrismacloudVulnerabilitiyOverview,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "total_vulnerable_runtime_assets",
				Description: "The name of the standard.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("OverviewSummary.TotalVulnerableRuntimeAssets"),
			},
			{
				Name:        "total_vulnerabilitiesin_runtime",
				Description: "The unique identifier for the standard.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("OverviewSummary.TotalVulnerabilitiesinRuntime"),
			},
			{
				Name:        "total_remediated_in_runtime",
				Description: "The number of policies assigned to the standard.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("OverviewSummary.TotalRemediatedInRuntime"),
			},
			{
				Name:        "values",
				Description: "Indicates if the standard is a system default.",
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
