package prismacloud

import (
	"context"

	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"github.com/turbot/steampipe-plugin-sdk/v5/query_cache"
)

func tablePrismaVulnerabilitiyAsset(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prismacloud_vulnerabilitiy_asset",
		Description: "The asset summary of vulnerabilitiy.",
		List: &plugin.ListConfig{
			Hydrate: listPrismaVulnerabilitiyAsset,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "asset_type", Require: plugin.Optional},
				{Name: "life_cycle", Require: plugin.Optional},
				{Name: "severities", Require: plugin.Optional, CacheMatch: query_cache.CacheMatchExact},
			},
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "asset_type",
				Description: "The type of asset. Possible values are: iac, package, deployedImage, serverlessFunction, host, registryImage, vmImage.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "life_cycle",
				Description: "The life cycle stage of the asset. Possible values are: code, build, deploy, run.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Stage"),
			},
			{
				Name:        "severities",
				Description: "The severities of the asset. Possible values are: low, medium, high, critical.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("severities"),
			},
			{
				Name:        "total_vulnerabilities",
				Description: "The total number of vulnerabilities.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "total_assets",
				Description: "The total number of assets.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "stats",
				Description: "Statistics of the vulnerable assets.",
				Type:        proto.ColumnType_JSON,
			},
		}),
	}
}

//// LIST FUNCTION

func listPrismaVulnerabilitiyAsset(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_vulnerabilitiy_asset.getPrismaVulnerabilitiyAsset", "connection_error", err)
		return nil, err
	}

	query := buildVulnerabilityAssetsQueryParameter(ctx, d)

	assets, err := api.ListVulnerabilityAssets(conn, query)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_vulnerabilitiy_asset.getPrismaVulnerabilitiyAsset", "api_error", err)
		return nil, err
	}

	for _, asset := range assets.Value {
		d.StreamListItem(ctx, asset)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
