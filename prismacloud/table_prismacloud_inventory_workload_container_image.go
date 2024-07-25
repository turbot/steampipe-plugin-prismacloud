package prismacloud

import (
	"context"

	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePrismaInventoryWorkloadContainerImage(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prismacloud_inventory_workload_container_image",
		Description: "Query Prisma Cloud Inventory Workload Container Image.",
		List: &plugin.ListConfig{
			Hydrate: listPrismaInventoryWorkloadContainerImages,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the container image.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "uai_id",
				Description: "The unique identifier of the UAI.",
				Type:        proto.ColumnType_STRING,
				Transform: transform.FromField("UaiID"),
			},
			{
				Name:        "running_containers",
				Description: "The number of running containers.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "scanPassed",
				Description: "Indicates if the scan passed.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "base",
				Description: "Indicates if the image is a base image.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "relatedImages",
				Description: "The number of related images.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "stages",
				Description: "The stages of the container image.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "vuln_funnel",
				Description: "The vulnerability funnel details.",
				Type:        proto.ColumnType_JSON,
			},
		},
	}
}

func listPrismaInventoryWorkloadContainerImages(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_inventory_workload_container_image.listPrismaInventoryWorkloadContainerImages", "connection_error", err)
		return nil, err
	}

	resp, err := api.GetInventoryWorkloadContainerImages(conn.JsonWebToken, "")
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_inventory_workload_container_image.listPrismaInventoryWorkloadContainerImages", "api_error", err)
		return nil, err
	}

	for _, image := range resp.Value {

		d.StreamListItem(ctx, image)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	for resp.NextPageToken != "" {
		resp, err = api.GetInventoryWorkloadContainerImages(conn.JsonWebToken, resp.NextPageToken)
		if err != nil {
			plugin.Logger(ctx).Error("prismacloud_inventory_workload_container_image.listPrismaInventoryWorkloadContainerImages", "paging_error", err)
			return nil, err
		}

		for _, image := range resp.Value {
			d.StreamListItem(ctx, image)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
	}
	return nil, nil
}
