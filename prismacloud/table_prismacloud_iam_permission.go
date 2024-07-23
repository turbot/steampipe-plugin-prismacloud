package prismacloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/api"
	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/model"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePrismaIAMPermission(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prismacloud_iam_permission",
		Description: "List all available permission for the accounts.",
		List: &plugin.ListConfig{
			Hydrate: listPrismaIAMPermissions,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "permission_query", Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The unique identifier for the permission.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "source_public",
				Description: "Indicates if the source is public.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "source_cloud_type",
				Description: "The type of cloud for the source.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "source_cloud_account",
				Description: "The cloud account for the source.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "source_cloud_region",
				Description: "The cloud region for the source.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "source_cloud_service_name",
				Description: "The cloud service name for the source.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "source_resource_name",
				Description: "The resource name for the source.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "source_resource_type",
				Description: "The resource type for the source.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "source_resource_id",
				Description: "The resource ID for the source.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "source_cloud_resource_uai",
				Description: "The cloud resource UAI for the source.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "source_idp_service",
				Description: "The IDP service for the source.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "source_idp_domain",
				Description: "The IDP domain for the source.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "source_idp_email",
				Description: "The IDP email for the source.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "source_idp_user_id",
				Description: "The IDP user ID for the source.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "source_idp_username",
				Description: "The IDP username for the source.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "source_idp_group",
				Description: "The IDP group for the source.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "source_idp_uai",
				Description: "The IDP UAI for the source.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "dest_cloud_type",
				Description: "The type of cloud for the destination.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "dest_cloud_account",
				Description: "The cloud account for the destination.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "dest_cloud_region",
				Description: "The cloud region for the destination.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "dest_cloud_service_name",
				Description: "The cloud service name for the destination.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "dest_resource_name",
				Description: "The resource name for the destination.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "dest_resource_type",
				Description: "The resource type for the destination.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "dest_resource_id",
				Description: "The resource ID for the destination.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "dest_cloud_resource_uai",
				Description: "The cloud resource UAI for the destination.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "granted_by_cloud_type",
				Description: "The cloud type granting the access.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "granted_by_cloud_policy_id",
				Description: "The cloud policy ID granting the access.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "granted_by_cloud_policy_name",
				Description: "The cloud policy name granting the access.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "granted_by_cloud_policy_type",
				Description: "The cloud policy type granting the access.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "granted_by_cloud_policy_uai",
				Description: "The cloud policy UAI granting the access.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "granted_by_cloud_policy_account",
				Description: "The cloud policy account granting the access.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "granted_by_cloud_entity_id",
				Description: "The cloud entity ID granting the access.",

				Type: proto.ColumnType_STRING,
			},
			{
				Name:        "granted_by_cloud_entity_name",
				Description: "The cloud entity name granting the access.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "granted_by_cloud_entity_type",
				Description: "The cloud entity type granting the access.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "granted_by_cloud_entity_account",
				Description: "The cloud entity account granting the access.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "granted_by_cloud_entity_uai",
				Description: "The cloud entity UAI granting the access.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "granted_by_level_type",
				Description: "The level type granting the access.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "granted_by_level_id",
				Description: "The level ID granting the access.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "granted_by_level_name",
				Description: "The level name granting the access.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "granted_by_level_uai",
				Description: "The level UAI granting the access.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_access_date",
				Description: "The date of last access.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_access_status",
				Description: "The status of the last access.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "accessed_resources_count",
				Description: "The count of accessed resources.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "effective_action_name",
				Description: "The name of the effective action.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "exceptions",
				Description: "The list of exceptions.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "wild_card_dest_cloud_resource_name",
				Description: "Indicates if the destination cloud resource name is a wildcard.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "permission_query",
				Description: "The query used.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Query"),
			},
			{
				Name:        "response_id",
				Description: "The unique identifier for the response.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("QueryId"),
			},
			{
				Name:        "saved",
				Description: "Indicates if the response is saved.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "name",
				Description: "The name of the search.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "time_range",
				Description: "The time range for the search.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "search_type",
				Description: "The type of the search.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "The description of the search.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "cloud_type",
				Description: "The type of cloud.",
				Type:        proto.ColumnType_STRING,
			},

			// Steampipe standard column
			{
				Name:        "title",
				Description: "The title of the search.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

type IAMPerm struct {
	Query       string
	QueryId     string
	Saved       bool
	Name        string
	TimeRange   map[string]interface{}
	SearchType  string
	Description string
	CloudType   string
	model.Item
}

//// LIST FUNCTION

func listPrismaIAMPermissions(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_iam_permission.listPrismaIAMPermissions", "connection_error", err)
		return nil, err
	}

	maxLimit := int32(10000)
	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)
		if limit < maxLimit {
			maxLimit = limit
		}
	}

	query := url.Values{
		"limit": []string{fmt.Sprint(maxLimit)},
	}

	// https://docs.prismacloud.io/en/classic/rql-reference/rql-reference/iam-query/iam-query-examples#id565e9de4-815d-4794-a3c3-7aecb6d9fb91
	req := map[string]interface{}{
		"query": "config from iam where dest.cloud.resource.name = '*'", // Default to all permissions
		"groupByFields": []string{"source",
			"sourceCloudAccount",
			"grantedByEntity",
			"entityCloudAccount",
			"grantedByPolicy",
			"policyCloudAccount",
			"grantedByLevel",
			"action",
			"destination",
			"destCloudAccount",
			"lastAccess"},
	}

	if d.EqualsQualString("permission_query") != "" {
		req["query"] = d.EqualsQualString("permission_query")
	}

	resp, err := api.ListIAMPermissions(conn, query, req)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_iam_permission.listPrismaIAMPermissions", "api_error", err)
		return nil, err
	}

	for _, perm := range resp.Data.Items {

		d.StreamListItem(ctx, IAMPerm{resp.Query, resp.Id, resp.Saved, resp.Name, resp.TimeRange, resp.SearchType, resp.Description, resp.CloudType, perm})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}

	}

	for resp.Data.NextPageToken != "" {
		req["nextPageToken"] = resp.Data.NextPageToken
		resp, err = api.ListIAMPermissions(conn, query, req)
		if err != nil {
			plugin.Logger(ctx).Error("prismacloud_iam_permission.listPrismaIAMPermissions", "api_paging_error", err)
			return nil, err
		}

		for _, perm := range resp.Data.Items {

			d.StreamListItem(ctx, IAMPerm{resp.Query, resp.Id, resp.Saved, resp.Name, resp.TimeRange, resp.SearchType, resp.Description, resp.CloudType, perm})

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}

		}
	}

	return nil, nil
}
