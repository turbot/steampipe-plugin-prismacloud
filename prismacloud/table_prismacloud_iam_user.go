package prismacloud

import (
	"context"

	"github.com/paloaltonetworks/prisma-cloud-go/user/profile"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePrismaIAMUser(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prismacloud_iam_user",
		Description: "List all available users and service accounts.",
		List: &plugin.ListConfig{
			Hydrate: listPrismaIAMUsers,
		},
		Columns: []*plugin.Column{
			{
				Name:        "display_name",
				Description: "Display name of the user.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DisplayName"),
			},
			{
				Name:        "last_name",
				Description: "Last name of the user.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("LastName"),
			},
			{
				Name:        "account_type",
				Description: "Type of the account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("AccountType"),
			},
			{
				Name:        "username",
				Description: "Username associated with the profile.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "first_name",
				Description: "First name of the user.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("FirstName"),
			},
			{
				Name:        "email",
				Description: "Email address of the user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "access_keys_allowed",
				Description: "Whether access keys are allowed for the user.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("AccessKeysAllowed"),
			},
			{
				Name:        "access_key_expiration",
				Description: "Expiration time of the access key.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("AccessKeyExpiration"),
			},
			{
				Name:        "access_key_name",
				Description: "Name of the access key.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("AccessKeyName"),
			},
			{
				Name:        "default_role_id",
				Description: "Default role ID assigned to the user.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DefaultRoleId"),
			},
			{
				Name:        "enable_key_expiration",
				Description: "Whether key expiration is enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("EnableKeyExpiration"),
			},
			{
				Name:        "time_zone",
				Description: "Time zone of the user.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("TimeZone"),
			},
			{
				Name:        "enabled",
				Description: "Whether the profile is enabled.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "last_login_ts",
				Description: "Timestamp of the last login.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("LastLoginTs"),
			},
			{
				Name:        "last_modified_by",
				Description: "Identifier of the user who last modified the profile.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("LastModifiedBy"),
			},
			{
				Name:        "last_modified_ts",
				Description: "Timestamp of the last modification.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("LastModifiedTs"),
			},
			{
				Name:        "access_keys_count",
				Description: "Count of access keys associated with the user.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("AccessKeysCount"),
			},
			{
				Name:        "role_ids",
				Description: "List of role IDs assigned to the user.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("RoleIds"),
			},
			{
				Name:        "roles",
				Description: "Roles assigned to the user.",
				Type:        proto.ColumnType_JSON,
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the profile.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DisplayName"),
			},
		},
	}
}

//// LIST FUNCTION

func listPrismaIAMUsers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_iam_user.listPrismaIAMUsers", "connection_error", err)
		return nil, err
	}

	users, err := profile.List(conn)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_iam_user.listPrismaIAMUsers", "api_error", err)
		return nil, err
	}

	for _, user := range users {

		d.StreamListItem(ctx, user)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}

	}

	return nil, nil
}
