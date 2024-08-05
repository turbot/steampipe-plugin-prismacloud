package prismacloud

import (
	"context"
	"fmt"
	"net/url"
	"reflect"

	"github.com/iancoleman/strcase"
	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/api"
	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/model"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"github.com/turbot/steampipe-plugin-sdk/v5/query_cache"
)

func tablePrismacloudInventoryAssetView(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prismacloud_inventory_asset_view",
		Description: "Prisma Cloud inventory asset view.",
		List: &plugin.ListConfig{
			Hydrate: listPrismacloudInventoryAssetView,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "account_name", Require: plugin.Optional},
				{Name: "cloud_type_name", Require: plugin.Optional},
				{Name: "region_name", Require: plugin.Optional},
				{Name: "resource_type_name", Require: plugin.Optional},
				{Name: "service_name", Require: plugin.Optional},
				{Name: "group_by", Require: plugin.Optional, CacheMatch: query_cache.CacheMatchExact},
				{Name: "compliance_requirement_name", Require: plugin.Optional, CacheMatch: query_cache.CacheMatchExact},
				{Name: "compliance_standard_name", Require: plugin.Optional, CacheMatch: query_cache.CacheMatchExact},
				{Name: "scan_status", Require: plugin.Optional, CacheMatch: query_cache.CacheMatchExact},
			},
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "group_by",
				Description: "The table content depends on the groupBy query parameter. Default value is 'cloud.service'. Possible values are: 'cloudType', 'cloud.account', 'cloud.region', 'cloud.service', and 'resource.type'.",
				Type:        proto.ColumnType_STRING,
				Default:     "cloud.service",
				Transform:   transform.FromQual("group_by"),
			},
			{
				Name:        "account_id",
				Description: "The ID of the cloud account.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "account_name",
				Description: "The name of the cloud account.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getColumnValueForAnyGroupByResult,
				Transform:   transform.FromField("account_name"),
			},
			{
				Name:        "allow_drill_down",
				Description: "Indicates if it's possible to drill down further.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "cloud_type_name",
				Description: "The name of the cloud type.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getColumnValueForAnyGroupByResult,
				Transform:   transform.FromField("cloud_type_name"),
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
				Name:        "critical_severity_failed_resources",
				Description: "The number of resources whose highest policy failure is critical.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "critical_vulnerability_failed_resources",
				Description: "The number of resources that failed with critical vulnerabilities.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "failed_resources",
				Description: "The number of failed resources.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "high_severity_failed_resources",
				Description: "The number of resources that failed high severity policies.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "high_vulnerability_failed_resources",
				Description: "The number of resources that failed with high vulnerabilities.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "informational_severity_failed_resources",
				Description: "The number of resources whose highest policy failure is informational.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "low_severity_failed_resources",
				Description: "The number of resources whose highest policy failure is low.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "low_vulnerability_failed_resources",
				Description: "The number of resources that failed with low vulnerabilities.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "medium_severity_failed_resources",
				Description: "The number of resources whose highest policy failure is medium.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "medium_vulnerability_failed_resources",
				Description: "The number of resources that failed with medium vulnerabilities.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "passed_resources",
				Description: "The number of passed resources.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "region_name",
				Description: "The name of the cloud region.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getColumnValueForAnyGroupByResult,
				Transform:   transform.FromField("region_name"),
			},
			{
				Name:        "resource_type_name",
				Description: "The name of the resource type.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getColumnValueForAnyGroupByResult,
				Transform:   transform.FromField("resource_type_name"),
			},
			{
				Name:        "service_name",
				Description: "The name of the cloud service.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getColumnValueForAnyGroupByResult,
				Transform:   transform.FromField("service_name"),
			},
			{
				Name:        "total_resources",
				Description: "The total number of resources.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "total_vulnerability_failed_resources",
				Description: "The total number of resources that failed with vulnerabilities.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "unscanned_resources",
				Description: "The total number of unscanned resources.",
				Type:        proto.ColumnType_INT,
			},
		}),
	}
}

func listPrismacloudInventoryAssetView(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_inventory_asset_explorer.listPrismacloudInventoryAssetView", "connection_error", err)
		return nil, err
	}

	groupBy := "cloud.service"
	if d.EqualsQualString("group_by") != "" {
		groupBy = d.EqualsQualString("group_by")
	}
	query := url.Values{
		"groupBy": []string{groupBy},
	}
	if d.EqualsQualString("account_name") != "" {
		query.Set("cloud.account", d.EqualsQualString("account_name"))
	}
	if d.EqualsQualString("service_name") != "" {
		query.Set("cloud.service", d.EqualsQualString("service_name"))
	}
	if d.EqualsQualString("cloud_type_name") != "" {
		query.Set("cloud.type", d.EqualsQualString("cloud_type_name"))
	}
	if d.EqualsQualString("region_name") != "" {
		query.Set("cloud.region", d.EqualsQualString("region_name"))
	}
	if d.EqualsQualString("resource_type_name") != "" {
		query.Set("resource.type", d.EqualsQualString("resource_type_name"))
	}
	if d.EqualsQualString("compliance_standard_name") != "" {
		query.Set("policy.complianceStandard", d.EqualsQualString("compliance_standard_name"))
	}
	if d.EqualsQualString("compliance_requirement_name") != "" {
		query.Set("policy.complianceRequirement", d.EqualsQualString("compliance_requirement_name"))
	}
	if d.EqualsQualString("scan_status") != "" {
		query.Set("scan.status", d.EqualsQualString("scan_status"))
	}

	resp, err := api.ListInventoryAsset(conn, query)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_inventory_asset_explorer.listPrismacloudInventoryAssetView", "api_error", err)
		return nil, err
	}

	for _, view := range resp.GroupedAggregate {

		d.StreamListItem(ctx, view)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTION

// This table fetches data based on the optional `group_by` qualifiers.
// Column values are populated based on the specified query parameters.
// For example, if you execute `select * from prismacloud_inventory_asset_view where group_by = 'resource.type' and account_name = 'gcsl-p-shared-base-2b3c - 801735715038'`,
// the API returns results. However, Steampipe applies a level of filtering that results in empty rows.
// This happens because the `account_name` column will not be populated when `group_by = 'resource.type'` is used as a query parameter.
func getColumnValueForAnyGroupByResult(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	data, ok := h.Item.(model.GroupedAggregateAsset)
	if !ok {
		plugin.Logger(ctx).Error("getColumnValueForAnyGroupByResult", "unexpected type for h.Item")
		return nil, fmt.Errorf("unexpected type for h.Item: %T", h.Item)
	}

	output := make(map[string]interface{})
	val := reflect.ValueOf(data)

	// Ensure the input is a struct or pointer to a struct
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		plugin.Logger(ctx).Error("getColumnValueForAnyGroupByResult", "input is not a struct")
		return nil, fmt.Errorf("input is not a struct: %T", data)
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldName := field.Name
		snakeKey := strcase.ToSnake(fieldName)
		output[snakeKey] = val.Field(i).Interface()
	}

	for _, columnName := range d.QueryContext.Columns {
		value, exists := output[columnName]
		if !exists || value == nil || (exists && fmt.Sprintf("%v", value) == "") {
			plugin.Logger(ctx).Debug("getColumnValueForAnyGroupByResult", "missing or empty column", "columnName", columnName)
			output[columnName] = d.EqualsQualString(columnName)
		}
	}

	return output, nil
}
