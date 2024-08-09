package prismacloud

import (
	"context"

	resource "github.com/paloaltonetworks/prisma-cloud-go/resource-list"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePrismacloudResource(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prismacloud_resource",
		Description: "List of available resources in Prisma Cloud.",
		Get: &plugin.GetConfig{
			Hydrate:    getPrismacloudResource,
			KeyColumns: plugin.SingleColumn("id"),
		},
		List: &plugin.ListConfig{
			Hydrate: listPrismacloudResources,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The unique identifier for the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "The name of the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "The description of the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "type",
				Description: "The type of the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_modified_by",
				Description: "The user who last modified the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_modified_ts",
				Description: "The timestamp of the last modification.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("LastModifiedTs").Transform(transform.NullIfZeroValue).Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "accept_account_groups",
				Description: "Indicates if the resource accepts account groups.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "accept_resource_lists",
				Description: "Indicates if the resource accepts resource lists.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "accept_code_repositories",
				Description: "Indicates if the resource accepts code repositories.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "custom",
				Description: "Indicates if the resource is custom.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "deleted",
				Description: "Indicates if the resource has been deleted.",
				Type:        proto.ColumnType_BOOL,
			},

			{
				Name:        "associated_roles",
				Description: "The roles associated with the resource.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "features",
				Description: "The features associated with the resource.",
				Type:        proto.ColumnType_JSON,
			},

			// Steampipe standard column
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		}),
	}
}

//// LIST FUNCTION

func listPrismacloudResources(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_resource.listPrismacloudResources", "connection_error", err)
		return nil, err
	}

	resources, err := resource.List(conn)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_resource.listPrismacloudResources", "api_error", err)
		return nil, err
	}

	for _, resource := range resources {

		d.StreamListItem(ctx, resource)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}

	}

	return nil, nil
}

//// HYDRATE FUNCTION

func getPrismacloudResource(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQualString("id")

	// Empty check
	if id == "" {
		return nil, nil
	}

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_resource.getPrismacloudResource", "connection_error", err)
		return nil, err
	}

	resource, err := resource.Get(conn, id)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_resource.getPrismacloudResource", "api_error", err)
		return nil, err
	}

	return resource, nil
}
