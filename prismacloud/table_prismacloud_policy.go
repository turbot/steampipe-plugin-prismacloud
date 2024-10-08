package prismacloud

import (
	"context"
	"fmt"
	"strings"

	"github.com/paloaltonetworks/prisma-cloud-go/policy"
	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/api"
	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/model"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"github.com/turbot/steampipe-plugin-sdk/v5/query_cache"
)

func tablePrismacloudPolicy(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prismacloud_policy",
		Description: "List of available policies in Prisma Cloud.",
		Get: &plugin.GetConfig{
			Hydrate:    getPrismacloudPolicy,
			KeyColumns: plugin.SingleColumn("policy_id"),
		},
		List: &plugin.ListConfig{
			Hydrate: listPrismacloudPolicies,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "cloud_type", Require: plugin.Optional},
				{Name: "severity", Require: plugin.Optional},
				{Name: "policy_type", Require: plugin.Optional},
				{Name: "enabled", Require: plugin.Optional},
				{Name: "policy_mode", Require: plugin.Optional},
				{Name: "remediable", Require: plugin.Optional},
				{Name: "name", Require: plugin.Optional},
				{Name: "compliance_requirement_name", Require: plugin.Optional, CacheMatch: query_cache.CacheMatchExact},
				{Name: "compliance_standard_name", Require: plugin.Optional, CacheMatch: query_cache.CacheMatchExact},
				{Name: "compliance_section_id", Require: plugin.Optional, CacheMatch: query_cache.CacheMatchExact},
			},
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "policy_id",
				Description: "The unique identifier for the policy.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "The name of the policy.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "policy_type",
				Description: "The type of the policy.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "system_default",
				Description: "Indicates if the policy is a system default.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "policy_upi",
				Description: "The unique policy identifier.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "compliance_standard_name",
				Description: "The name of the compliance standard.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("compliance_standard_name"),
			},
			{
				Name:        "compliance_requirement_name",
				Description: "The name of the compliance requirement.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("compliance_requirement_name"),
			},
			{
				Name:        "compliance_section_id",
				Description: "The name of the compliance standard.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("compliance_section_id"),
			},
			{
				Name:        "description",
				Description: "The description of the policy.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "severity",
				Description: "The severity level of the policy.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "recommendation",
				Description: "The recommendation for the policy.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "cloud_type",
				Description: "The type of cloud (e.g., AWS, Azure, GCP).",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "enabled",
				Description: "Indicates if the policy is enabled.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "created_on",
				Description: "The timestamp when the policy was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("CreatedOn").Transform(transform.NullIfZeroValue).Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "created_by",
				Description: "The user who created the policy.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_modified_on",
				Description: "The timestamp of the last modification.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("LastModifiedOn").Transform(transform.NullIfZeroValue).Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "last_modified_by",
				Description: "The user who last modified the policy.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "rule_last_modified_on",
				Description: "The timestamp of the last modification to the rule.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("RuleLastModifiedOn").Transform(transform.NullIfZeroValue).Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "overridden",
				Description: "Indicates if the policy has been overridden.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "deleted",
				Description: "Indicates if the policy has been deleted.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "restrict_alert_dismissal",
				Description: "Indicates if alert dismissal is restricted for the policy.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "open_alerts_count",
				Description: "The number of open alerts for the policy.",
				Hydrate:     getPrismacloudOpenAlertCountForPolicy,
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "owner",
				Description: "The owner of the policy.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "policy_mode",
				Description: "The mode of the policy.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "policy_category",
				Description: "The category of the policy.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "policy_class",
				Description: "The class of the policy.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "remediable",
				Description: "Indicates if the policy is remediable.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "policy_sub_types",
				Description: "The subtypes of the policy.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "rule",
				Description: "The rule associated with the policy.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "compliance_metadata",
				Description: "The compliance metadata associated with the policy.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "remediation",
				Description: "The remediation information for the policy.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "labels",
				Description: "The labels associated with the policy.",
				Type:        proto.ColumnType_JSON,
			},

			// Steampipe standard column
			{
				Name:        "title",
				Description: "Title of the policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		}),
	}
}

//// LIST FUNCTION

func listPrismacloudPolicies(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_policy.listPrismacloudPolicies", "connection_error", err)
		return nil, err
	}

	query := buildPrismacloudListPolicyInputQuery(d)

	policies, err := policy.List(conn, query)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_policy.listPrismacloudPolicies", "api_error", err)
		return nil, err
	}

	for _, policy := range policies {
		d.StreamListItem(ctx, policy)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTION

func getPrismacloudPolicy(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQualString("policy_id")

	// Empty check
	if id == "" {
		return nil, nil
	}

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_policy.getPrismacloudPolicy", "connection_error", err)
		return nil, err
	}

	policy, err := policy.Get(conn, id)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_policy.getPrismacloudPolicy", "api_error", err)
		return nil, err
	}

	return policy, nil
}

func getPrismacloudOpenAlertCountForPolicy(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	policy := h.Item.(policy.Policy)

	type filterMap map[string]string

	filters := []filterMap{
		{"name": "alert.status", "operator": "=", "value": "open"},
		{"name": "policy.severity", "operator": "=", "value": policy.Severity},
		{"name": "policy.name", "operator": "=", "value": policy.Name},
		{"name": "timeRange.type", "operator": "=", "value": "ALERT_UPDATED"},
	}
	sortBy := []string{"alertCount:asc", "severity:desc"}
	timeRange := map[string]string{
		"type":  "to_now",
		"value": "epoch",
	}
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_policy.getPrismacloudOpenAlertCountForPolicy", "connection_error", err)
		return nil, err
	}

	req := map[string]interface{}{
		"filters":       filters,
		"sortBy":        sortBy,
		"timeRange":     timeRange,
		"nextPageToken": "",
	}
	var policies []model.Policy
	results, err := api.GetAlertCountOfPolicies(conn, req)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return 0, nil
		}
		plugin.Logger(ctx).Error("prismacloud_policy.getPrismacloudOpenAlertCountForPolicy", "api_error", err)
		return nil, err
	}

	policies = append(policies, results.Policies...)

	for results.NextPageToken != "" {
		req["nextPageToken"] = results.NextPageToken
		results, err = api.GetAlertCountOfPolicies(conn, req)
		if err != nil {
			plugin.Logger(ctx).Error("prismacloud_policy.getPrismacloudOpenAlertCountForPolicy", "paging_error", err)
			return nil, err
		}
		policies = append(policies, results.Policies...)
	}

	for _, p := range policies {
		if p.PolicyId == policy.PolicyId {
			return p.AlertCount, nil
		}
	}

	return 0, nil
}

//// UTILITY FUNCTION

// Build the list policy input param
func buildPrismacloudListPolicyInputQuery(d *plugin.QueryData) map[string]string {

	query := make(map[string]string)
	filterQuals := map[string]string{
		"compliance_standard_name":    "policy.complianceStandard",
		"compliance_requirement_name": "policy.complianceRequirement",
		"compliance_section_id":       "policy.complianceSection",
		"cloud_type":                  "cloud.type",
		"severity":                    "policy.severity",
		"policy_type":                 "policy.type",
		"enabled":                     "policy.enabled",
		"policy_mode":                 "policy.mode",
		"remediable":                  "policy.remediable",
		"name":                        "policy.name",
	}

	for columnName, qp := range filterQuals {
		if (columnName == "enabled" || columnName == "remediable") && d.EqualsQuals[columnName] != nil { // Boolean quals
			query[qp] = fmt.Sprint(d.EqualsQuals[columnName].GetBoolValue())
			continue
		}
		if d.EqualsQuals[columnName] != nil {
			query[qp] = d.EqualsQuals[columnName].GetStringValue()
		}
	}

	return query

}
