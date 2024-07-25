package prismacloud

import (
	"context"

	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePrismaInventoryWorkload(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prismacloud_inventory_workload",
		Description: "Query Prisma Cloud Inventory Workload summary.",
		List: &plugin.ListConfig{
			Hydrate: listPrismaInventoryWorkloads,
		},
		Columns: []*plugin.Column{
			{
				Name:        "container_images_build",
				Description: "Number of container images in the build stage.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("ContainerImages.Stages.Build"),
			},
			{
				Name:        "container_images_deploy",
				Description: "Number of container images in the deploy stage.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("ContainerImages.Stages.Deploy"),
			},
			{
				Name:        "container_images_run",
				Description: "Number of container images in the run stage.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("ContainerImages.Stages.Run"),
			},
			{
				Name:        "container_images_vulnerable",
				Description: "Number of vulnerable container images.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("ContainerImages.Vulnerable"),
			},
			{
				Name:        "container_images_cloud_providers",
				Description: "Cloud providers for container images.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ContainerImages.CloudProviders"),
			},
			{
				Name:        "hosts_total",
				Description: "Total number of hosts.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Hosts.Total"),
			},
			{
				Name:        "hosts_vulnerable",
				Description: "Number of vulnerable hosts.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Hosts.Vulnerable"),
			},
			{
				Name:        "hosts_cloud_providers",
				Description: "Cloud providers for hosts.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Hosts.CloudProviders"),
			},
		},
	}
}

func listPrismaInventoryWorkloads(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_inventory_workload.listPrismaInventoryWorkloads", "connection_error", err)
		return nil, err
	}

	resp, err := api.GetInventoryWorkloads(conn.JsonWebToken)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_inventory_workload.listPrismaInventoryWorkloads", "api_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, resp)

	return nil, nil
}
