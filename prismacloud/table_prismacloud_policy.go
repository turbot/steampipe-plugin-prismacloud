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
)

func tablePrismaPolicy(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prismacloud_policy",
		Description: "List of available policies in Prisma Cloud.",
		Get: &plugin.GetConfig{
			Hydrate:    getPrismaPolicy,
			KeyColumns: plugin.SingleColumn("policy_id"),
		},
		List: &plugin.ListConfig{
			Hydrate: listPrismaPolicies,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "cloud_type", Require: plugin.Optional},
				{Name: "severity", Require: plugin.Optional},
				{Name: "policy_type", Require: plugin.Optional},
				{Name: "enabled", Require: plugin.Optional},
				{Name: "policy_mode", Require: plugin.Optional},
				{Name: "remediable", Require: plugin.Optional},
				{Name: "name", Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
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
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "created_by",
				Description: "The user who created the policy.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_modified_on",
				Description: "The timestamp of the last modification.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "last_modified_by",
				Description: "The user who last modified the policy.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "rule_last_modified_on",
				Description: "The timestamp of the last modification to the rule.",
				Type:        proto.ColumnType_INT,
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
				Hydrate:     getPrismaOpenAlertCountForPolicy,
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
		},
	}
}

//// LIST FUNCTION

func listPrismaPolicies(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_policy.listPrismaPolicies", "connection_error", err)
		return nil, err
	}

	query := buildPrismacloudListPolicyInputQuery(d)

	policies, err := policy.List(conn, query)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_policy.listPrismaPolicies", "api_error", err)
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

func getPrismaPolicy(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQualString("policy_id")

	// Empty check
	if id == "" {
		return nil, nil
	}

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_policy.getPrismaPolicy", "connection_error", err)
		return nil, err
	}

	policy, err := policy.Get(conn, id)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_policy.getPrismaPolicy", "api_error", err)
		return nil, err
	}

	return policy, nil
}

func getPrismaOpenAlertCountForPolicy(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
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
		plugin.Logger(ctx).Error("prismacloud_policy.getPrismaOpenAlertCountForPolicy", "connection_error", err)
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
		plugin.Logger(ctx).Error("prismacloud_policy.getPrismaOpenAlertCountForPolicy", "api_error", err)
		return nil, err
	}

	policies = append(policies, results.Policies...)

	for results.NextPageToken != "" {
		req["nextPageToken"] = results.NextPageToken
		results, err = api.GetAlertCountOfPolicies(conn, req)
		if err != nil {
			plugin.Logger(ctx).Error("prismacloud_policy.getPrismaOpenAlertCountForPolicy", "paging_error", err)
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
		"cloud_type":  "cloud.type",
		"severity":    "policy.severity",
		"policy_type": "policy.type",
		"enabled":     "policy.enabled",
		"policy_mode": "policy.mode",
		"remediable":  "policy.remediable",
		"name":        "policy.name",
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
