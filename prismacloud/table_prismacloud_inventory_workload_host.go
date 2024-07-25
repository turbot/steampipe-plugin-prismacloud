package prismacloud

import (
	"context"

	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePrismaInventoryWorkloadHost(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prismacloud_inventory_workload_host",
		Description: "Query Prisma Cloud Inventory Workload Host.",
		List: &plugin.ListConfig{
			Hydrate: listPrismaInventoryWorkloadHosts,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the container image.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "The unique identifier of the host.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "uai_id",
				Description: "The unique identifier of the UAI.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("UaiID"),
			},
			{
				Name:        "vuln_funnel",
				Description: "The vulnerability funnel details.",
				Type:        proto.ColumnType_JSON,
			},

			// Steampipe standard column
			{
				Name:        "title",
				Description: "Title of the host.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

func listPrismaInventoryWorkloadHosts(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_inventory_workload_host.listPrismaInventoryWorkloadHosts", "connection_error", err)
		return nil, err
	}

	resp, err := api.GetInventoryWorkloadHosts(conn.JsonWebToken, "")
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_inventory_workload_host.listPrismaInventoryWorkloadHosts", "api_error", err)
		return nil, err
	}

	for _, host := range resp.Value {

		d.StreamListItem(ctx, host)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	for resp.NextPageToken != "" {
		resp, err = api.GetInventoryWorkloadHosts(conn.JsonWebToken, resp.NextPageToken)
		if err != nil {
			plugin.Logger(ctx).Error("prismacloud_inventory_workload_host.listPrismaInventoryWorkloadHosts", "paging_error", err)
			return nil, err
		}

		for _, host := range resp.Value {
			d.StreamListItem(ctx, host)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
	}
	return nil, nil
}
