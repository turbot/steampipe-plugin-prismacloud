package prismacloud

import (
	"context"

	"github.com/paloaltonetworks/prisma-cloud-go/cloud/account"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePrismacloudAccount(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prismacloud_account",
		Description: "List all cloud accounts onboarded onto the Prisma Cloud platform.",
		List: &plugin.ListConfig{
			Hydrate: listPrismacloudAccounts,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "account_id",
				Description: "The unique identifier for the account.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "The name of the account.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "cloud_type",
				Description: "The type of cloud (e.g., AWS, Azure, GCP).",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "account_type",
				Description: "The type of the account.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "enabled",
				Description: "Indicates if the account is enabled.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "last_modified_ts",
				Description: "The timestamp of the last modification.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("LastModifiedTs").Transform(transform.NullIfZeroValue).Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "last_modified_by",
				Description: "The user who last modified the account.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "storage_scan_enabled",
				Description: "Indicates if the storage scan is enabled.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "protection_mode",
				Description: "The protection mode of the account.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "ingestion_mode",
				Description: "The ingestion mode of the account.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "status",
				Description: "The status of the account.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "number_of_child_accounts",
				Description: "The number of child accounts associated with the account.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "added_on",
				Description: "The timestamp when the account was added.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "group_ids",
				Description: "The IDs of the groups associated with the account.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "groups",
				Description: "The groups associated with the account.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "account_details",
				Description: "The account details based on cloud type.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getAccountDetails,
				Transform:   transform.FromValue(),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		}),
	}
}

//// LIST FUNCTION

func listPrismacloudAccounts(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_account.listPrismacloudAccounts", "connection_error", err)
		return nil, err
	}

	accounts, err := account.List(conn)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_account.listPrismacloudAccounts", "api_error", err)
		return nil, err
	}

	for _, account := range accounts {

		d.StreamListItem(ctx, account)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}

	}

	return nil, nil
}

//// HYDRATE FUNCTION

func getAccountDetails(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	accountData := h.Item.(account.Account)

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_account.getAccountDetails", "connection_error", err)
		return nil, err
	}

	account, err := account.Get(conn, accountData.CloudType, accountData.AccountId)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_account.getAccountDetails", "api_error", err)
		return nil, err
	}

	if account != nil {
		return account, nil
	}

	return nil, nil
}
