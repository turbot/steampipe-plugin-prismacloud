package prismacloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"github.com/turbot/steampipe-plugin-sdk/v5/query_cache"
)

func tablePrismaInventoryAssetExplorer(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prismacloud_inventory_asset_explorer",
		Description: "Query Prisma Cloud Inventory Asset Explorer.",
		List: &plugin.ListConfig{
			Hydrate: listPrismaInventoryAssetExplorer,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "account_name", Require: plugin.Optional},
				{Name: "cloud_type", Require: plugin.Optional},
				{Name: "region_name", Require: plugin.Optional},
				{Name: "compliance_requirement_name", Require: plugin.Optional, CacheMatch: query_cache.CacheMatchExact},
				{Name: "compliance_standard_name", Require: plugin.Optional, CacheMatch: query_cache.CacheMatchExact},
				{Name: "scan_status", Require: plugin.Optional, CacheMatch: query_cache.CacheMatchExact},
			},
		},
		Columns: []*plugin.Column{
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
				Name:        "account_id",
				Description: "The unique identifier for the account.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "account_name",
				Description: "The name of the account.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "compliance_requirement_name",
				Description: "The name of the compliance requirement.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("compliance_requirement_name"),
			},
			{
				Name:        "compliance_standard_name",
				Description: "The name of the compliance standard.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("compliance_standard_name"),
			},
			{
				Name:        "scan_status",
				Description: "The scan status. Possible values are: 'passed' or 'failed'",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("scan_status"),
			},
			{
				Name:        "alert_status_critical",
				Description: "The critical alert status count.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("AlertStatus.Critical"),
			},
			{
				Name:        "alert_status_high",
				Description: "The high alert status count.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("AlertStatus.High"),
			},
			{
				Name:        "alert_status_informational",
				Description: "The informational alert status count.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("AlertStatus.Informational"),
			},
			{
				Name:        "alert_status_low",
				Description: "The low alert status count.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("AlertStatus.Low"),
			},
			{
				Name:        "alert_status_medium",
				Description: "The medium alert status count.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("AlertStatus.Medium"),
			},
			{
				Name:        "asset_type",
				Description: "The type of the asset.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "cloud_type",
				Description: "The type of cloud (e.g., AWS, Azure, GCP).",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "overall_passed",
				Description: "Indicates if the resource passed overall checks.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "region_id",
				Description: "The ID of the region.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "region_name",
				Description: "The name of the region.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "resource_config_json_available",
				Description: "Indicates if the resource config JSON is available.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "resource_details_available",
				Description: "Indicates if the resource details are available.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "rrn",
				Description: "The resource RRN.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "unified_asset_id",
				Description: "The unified asset ID.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "vulnerability_status_critical",
				Description: "The critical vulnerability status count.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("VulnerabilityStatus.Critical"),
			},
			{
				Name:        "vulnerability_status_high",
				Description: "The high vulnerability status count.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("VulnerabilityStatus.High"),
			},
			{
				Name:        "vulnerability_status_low",
				Description: "The low vulnerability status count.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("VulnerabilityStatus.Low"),
			},
			{
				Name:        "vulnerability_status_medium",
				Description: "The medium vulnerability status count.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("VulnerabilityStatus.Medium"),
			},
			{
				Name:        "timestamp",
				Description: "The timestamp of the response.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "app_names",
				Description: "The application names associated with the resource.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "scanned_policies",
				Description: "The policies that have been scanned for the resource.",
				Type:        proto.ColumnType_JSON,
			},

			// Steampipe standard column
			{
				Name:        "title",
				Description: "Title of the asset.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

func listPrismaInventoryAssetExplorer(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_inventory_asset_explorer.listPrismaInventoryAssetExplorer", "connection_error", err)
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
	if d.EqualsQualString("account_name") != "" {
		query.Set("cloud.account", d.EqualsQualString("account_name"))
	}
	if d.EqualsQualString("cloud_type") != "" {
		query.Set("cloud.type", d.EqualsQualString("cloud_type"))
	}
	if d.EqualsQualString("region_name") != "" {
		query.Set("cloud.region", d.EqualsQualString("region_name"))
	}
	if d.EqualsQualString("compliance_requirement_name") != "" {
		query.Set("policy.complianceRequirement", d.EqualsQualString("compliance_requirement_name"))
	}
	if d.EqualsQualString("scan_status") != "" {
		query.Set("scan.status", d.EqualsQualString("scan_status"))
	}

	resp, err := api.ListInventoryAssetExplorer(conn, query)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_inventory_asset_explorer.listPrismaInventoryAssetExplorer", "api_error", err)
		return nil, err
	}

	for _, resource := range resp.Resources {

		d.StreamListItem(ctx, resource)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	for resp.NextPageToken != "" {
		query.Set("pageToken", resp.NextPageToken)
		resp, err = api.ListInventoryAssetExplorer(conn, query)
		if err != nil {
			plugin.Logger(ctx).Error("prismacloud_inventory_asset_explorer.listPrismaInventoryAssetExplorer", "paging_error", err)
			return nil, err
		}

		for _, resource := range resp.Resources {

			d.StreamListItem(ctx, resource)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}
