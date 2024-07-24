package prismacloud

import (
	"context"
	"net/url"

	"github.com/paloaltonetworks/prisma-cloud-go/cloud/account"
	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/api"
	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/model"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"github.com/turbot/steampipe-plugin-sdk/v5/query_cache"
)

func tablePrismaComplianceBreakdownSummary(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prismacloud_compliance_breakdown_summary",
		Description: "List all available compliance breakdown summary.",
		List: &plugin.ListConfig{
			ParentHydrate: listPrismaAccounts,
			Hydrate:       listPrismaComplianceBreakdownSummary,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "account_name", Require: plugin.Optional},
				{Name: "cloud_type", Require: plugin.Optional},
				{Name: "cloud_region", Require: plugin.Optional, CacheMatch: query_cache.CacheMatchExact},
				{Name: "policy_compliance_standard_name", Require: plugin.Optional, CacheMatch: query_cache.CacheMatchExact},
				{Name: "policy_compliance_requirement_name", Require: plugin.Optional, CacheMatch: query_cache.CacheMatchExact},
				{Name: "policy_compliance_section_id", Require: plugin.Optional, CacheMatch: query_cache.CacheMatchExact},
			},
		},
		Columns: complianceBreakdownCommonFilterColumns([]*plugin.Column{
			{
				Name:        "critical_severity_failed_resources",
				Description: "Number of scanned compliance resources whose highest policy failure is critical.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "failed_resources",
				Description: "Number of failed scanned compliance resources.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "high_severity_failed_resources",
				Description: "Number of scanned compliance resources that failed high severity policies.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "informational_severity_failed_resources",
				Description: "Number of scanned compliance resources whose highest policy failure is informational.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "low_severity_failed_resources",
				Description: "Number of scanned compliance resources whose highest policy failure is low.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "medium_severity_failed_resources",
				Description: "Number of scanned compliance resources whose highest policy failure is medium.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "passed_resources",
				Description: "Number of passed scanned compliance resources.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "timestamp",
				Description: "Timestamp of the compliance summary.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Timestamp").Transform(transform.NullIfZeroValue).Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "total_resources",
				Description: "Total number of scanned compliance resources.",
				Type:        proto.ColumnType_INT,
			},
		}),
	}
}

type complianceBreakdownSummary struct {
	AccountName string
	AccountId   string
	CloudType   string
	model.ComplianceSummary
}

//// LIST FUNCTION

func listPrismaComplianceBreakdownSummary(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	account := h.Item.(account.Account)

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_compliance_breakdown_summary.listPrismaComplianceBreakdownSummary", "connection_error", err)
		return nil, err
	}

	if d.EqualsQualString("account_name") != "" && d.EqualsQualString("account_name") != account.Name {
		return nil, nil
	}

	if d.EqualsQualString("cloud_type") != "" && d.EqualsQualString("cloud_type") != account.CloudType {
		return nil, nil
	}

	// For any of the query parameter it the returning the same row. However, the query param is required to make the the API call do hardcoded the value.
	query := url.Values{
		"cloud.account": []string{account.Name},
	}

	if d.EqualsQualString("policy_compliance_standard_name") != "" {
		query["policy.complianceStandard"] = []string{d.EqualsQualString("policy_compliance_standard_name")}
	}
	if d.EqualsQualString("policy_compliance_requirement_name") != "" {
		query["policy.complianceRequirement"] = []string{d.EqualsQualString("policy_compliance_requirement_name")}
	}
	if d.EqualsQualString("policy_compliance_section_id") != "" {
		query["policy.complianceSection"] = []string{d.EqualsQualString("policy_compliance_section_id")}
	}

	postures, err := api.LisComplianceBreakdownStatistics(conn, query)
	if err != nil {

		plugin.Logger(ctx).Error("prismacloud_compliance_breakdown_summary.listPrismaComplianceBreakdownSummary", "api_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, complianceBreakdownSummary{account.Name, account.AccountId, account.CloudType, postures.Summary})

	return nil, nil
}
