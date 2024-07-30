package prismacloud

import (
	"context"

	"github.com/paloaltonetworks/prisma-cloud-go/permission_group"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePrismaPermissionGroup(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prismacloud_permission_group",
		Description: "List of available permission groups.",
		Get: &plugin.GetConfig{
			Hydrate:    getPrismaPermissionGroup,
			KeyColumns: plugin.SingleColumn("id"),
		},
		List: &plugin.ListConfig{
			Hydrate: listPrismaPermissionGroups,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The unique identifier for the permission group.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "The name of the permission group.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "The description of the permission group.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "type",
				Description: "The type of the permission group.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_modified_by",
				Description: "The user who last modified the permission group.",
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
				Description: "Indicates if the permission group accepts account groups.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "accept_resource_lists",
				Description: "Indicates if the permission group accepts resource lists.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "accept_code_repositories",
				Description: "Indicates if the permission group accepts code repositories.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "custom",
				Description: "Indicates if the permission group is custom.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "deleted",
				Description: "Indicates if the permission group has been deleted.",
				Type:        proto.ColumnType_BOOL,
			},

			{
				Name:        "associated_roles",
				Description: "The roles associated with the permission group.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "features",
				Description: "The features associated with the permission group.",
				Type:        proto.ColumnType_JSON,
			},

			// Steampipe standard column
			{
				Name:        "title",
				Description: "Title of the permission group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		}),
	}
}

//// LIST FUNCTION

func listPrismaPermissionGroups(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_permission_group.listPrismaPermissionGroups", "connection_error", err)
		return nil, err
	}

	permissions, err := permission_group.List(conn)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_permission_group.listPrismaPermissionGroups", "api_error", err)
		return nil, err
	}

	for _, permission := range permissions {

		d.StreamListItem(ctx, permission)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}

	}

	return nil, nil
}

//// HYDRATE FUNCTION

func getPrismaPermissionGroup(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQualString("id")

	// Empty check
	if id == "" {
		return nil, nil
	}

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_permission_group.getPrismaPermissionGroup", "connection_error", err)
		return nil, err
	}

	permission, err := permission_group.Get(conn, id)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_permission_group.getPrismaPermissionGroup", "api_error", err)
		return nil, err
	}

	return permission, nil
}
