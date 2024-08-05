package prismacloud

import (
	"context"

	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"github.com/turbot/steampipe-plugin-sdk/v5/query_cache"
)

func tablePrismacloudVulnerabilitiyBurndown(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prismacloud_vulnerabilitiy_burndown",
		Description: "The burndown summary of vulnerabilitiy.",
		List: &plugin.ListConfig{
			Hydrate: listPrismacloudVulnerabilitiyBurndown,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "asset_type", Require: plugin.Required, CacheMatch: query_cache.CacheMatchExact},
				{Name: "life_cycle", Require: plugin.Required, CacheMatch: query_cache.CacheMatchExact},
				{Name: "severities", Require: plugin.Required, CacheMatch: query_cache.CacheMatchExact},
			},
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "asset_type",
				Description: "The type of asset. Possible values are: iac, package, deployedImage, serverlessFunction, host, registryImage, vmImage.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("asset_type"),
			},
			{
				Name:        "life_cycle",
				Description: "The life cycle stage of the asset. Possible values are: code, build, deploy, run.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("life_cycle"),
			},
			{
				Name:        "severities",
				Description: "The severities of the asset. Possible values are: low, medium, high, critical.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("severities"),
			},
			{
				Name:        "day_num",
				Description: "Count down of the day backwards from present day.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "total_count",
				Description: "Number of vulnerabilities in the given day.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "remediated_count",
				Description: "Number of vulnerabilities remediated for the given day.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "epoch_timestamp",
				Description: "Time up to which the entry was recorded.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("EpochTimestamp").Transform(transform.NullIfZeroValue).Transform(transform.UnixMsToTimestamp),
			},
		}),
	}
}

//// LIST FUNCTION

func listPrismacloudVulnerabilitiyBurndown(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_vulnerabilitiy_burndown.getPrismacloudVulnerabilitiyBurndown", "connection_error", err)
		return nil, err
	}

	query := buildBurndownVulnerabilitiesQueryParameter(ctx, d)

	vulnerability, err := api.ListVulnerabilityBurndown(conn, query)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_vulnerabilitiy_burndown.getPrismacloudVulnerabilitiyBurndown", "api_error", err)
		return nil, err
	}

	for _, burndown := range vulnerability {
		d.StreamListItem(ctx, burndown)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
