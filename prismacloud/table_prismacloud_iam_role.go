package prismacloud

import (
	"context"

	"github.com/paloaltonetworks/prisma-cloud-go/user/role"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePrismaIAMRole(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prismacloud_iam_role",
		Description: "List all available roles for the users.",
		List: &plugin.ListConfig{
			Hydrate: listPrismaIAMRoles,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the role.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "The unique identifier for the role.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "The description of the role.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "role_type",
				Description: "The type of the role.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_modified_by",
				Description: "The user who last modified the role.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_modified_ts",
				Description: "The timestamp when the role was last modified.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("LastModifiedTs").Transform(transform.NullIfZeroValue).Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "restrict_dismissal_access",
				Description: "Whether the role restricts dismissal access.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "account_group_ids",
				Description: "A list of account group IDs associated with the role.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "resource_list_ids",
				Description: "A list of resource list IDs associated with the role.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "code_repository_ids",
				Description: "A list of code repository IDs associated with the role.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "associated_users",
				Description: "A list of user IDs associated with the role.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "account_groups",
				Description: "A list of account groups associated with the role.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "additional_attributes",
				Description: "Additional attributes associated with the role.",
				Type:        proto.ColumnType_JSON,
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the role.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

//// LIST FUNCTION

func listPrismaIAMRoles(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_user_role.listPrismaIAMRoles", "connection_error", err)
		return nil, err
	}

	roles, err := role.List(conn)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_user_role.listPrismaIAMRoles", "api_error", err)
		return nil, err
	}

	for _, role := range roles {

		d.StreamListItem(ctx, role)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}

	}

	return nil, nil
}
