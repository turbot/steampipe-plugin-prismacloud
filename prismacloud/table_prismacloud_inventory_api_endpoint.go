package prismacloud

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePrismaInventoryAPIEndpoint(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prismacloud_inventory_api_endpoint",
		Description: "Query Prisma Cloud Inventory API Endpoint.",
		List: &plugin.ListConfig{
			Hydrate: listPrismaInventoryAPIEndpoints,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "asset_id",
				Description: "The unique identifier for the asset.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("AssetID"),
			},
			{
				Name:        "api_path",
				Description: "The API path.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("APIPath"),
			},
			{
				Name:        "http_method",
				Description: "The HTTP method.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("HTTPMethod"),
			},
			{
				Name:        "api_server",
				Description: "The API server URL.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("APIServer"),
			},
			{
				Name:        "hits",
				Description: "The number of hits.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "workloads",
				Description: "The workloads associated with the asset.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "service_name",
				Description: "The name of the service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ServiceName"),
			},
			{
				Name:        "cloud_type",
				Description: "The type of cloud.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "region",
				Description: "The region of the asset.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "account_id",
				Description: "The account ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("AccountID"),
			},
			{
				Name:        "account_name",
				Description: "The account name.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "discovery_method",
				Description: "The method of discovery.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "inspection_type",
				Description: "The type of inspection.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_changed",
				Description: "The timestamp when the asset was last changed.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("LastChanged").Transform(transform.NullIfZeroValue).Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "last_observed",
				Description: "The timestamp when the asset was last observed.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("LastObserved").Transform(transform.NullIfZeroValue).Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "total",
				Description: "The total number of items.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Total"),
			},
			{
				Name:        "count",
				Description: "The count of items.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Count"),
			},
			{
				Name:        "path_risk_factors",
				Description: "The risk factors associated with the path.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("PathRiskFactors"),
			},

			// Steampipe standard column
			{
				Name:        "title",
				Description: "Title of the asset.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		}),
	}
}

func listPrismaInventoryAPIEndpoints(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_inventory_api_endpoint.listPrismaInventoryAPIEndpoints", "connection_error", err)
		return nil, err
	}

	maxLimit := int32(100)
	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)
		if limit < maxLimit {
			maxLimit = limit
		}
	}

	req := map[string]interface{}{
		"limit": fmt.Sprint(maxLimit),
		"orderBy":"assetId",
		"orderDirection":"desc",
	}

	resp, err := api.ListInventoryDiscoveredAPI(conn, req)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_inventory_api_endpoint.listPrismaInventoryAPIEndpoints", "api_error", err)
		return nil, err
	}

	for _, member := range resp.Members {

		d.StreamListItem(ctx, member)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	for resp.NextPageToken != nil {
		req["nextPageToken"] = *resp.NextPageToken
		resp, err = api.ListInventoryDiscoveredAPI(conn, req)
		if err != nil {
			plugin.Logger(ctx).Error("prismacloud_inventory_api_endpoint.listPrismaInventoryAPIEndpoints", "paging_error", err)
			return nil, err
		}

		for _, member := range resp.Members {

			d.StreamListItem(ctx, member)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}
