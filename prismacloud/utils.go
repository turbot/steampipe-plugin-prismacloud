package prismacloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// Common filter columns for the Compliance Breakdown
func complianceBreakdownCommonFilterColumns(columns []*plugin.Column) []*plugin.Column {
	return append(columns, []*plugin.Column{
		{
			Name:        "account_id",
			Description: "The unique identifier for the account.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromQual("account_id"),
		},
		{
			Name:        "cloud_type",
			Description: "The type of cloud (e.g., AWS, Azure, GCP).",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromQual("cloud_type"),
		},
		{
			Name:        "cloud_region",
			Description: "The region of the cloud where the resource is located.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromQual("cloud_region"),
		},
		{
			Name:        "policy_compliance_standard_name",
			Description: "The name of the compliance standard associated with the policy.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromQual("policy_compliance_standard_name"),
		},
		{
			Name:        "policy_compliance_requirement_name",
			Description: "The name of the compliance requirement associated with the policy.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromQual("policy_compliance_requirement_name"),
		},
		{
			Name:        "policy_compliance_section_id",
			Description: "The ID of the compliance section associated with the policy.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromQual("policy_compliance_section_id"),
		},
	}...)
}

// Common key columns for the the Compliance Breakdown
func commonComplianceBreakdownKeyQualColumns() plugin.KeyColumnSlice {
	commonKeyQualsCol := plugin.AnyColumn([]string{
		"account_id",
		"cloud_type",
		"cloud_region",
		"policy_compliance_standard_name",
		"policy_compliance_requirement_name",
		"policy_compliance_section_id",
	})

	return commonKeyQualsCol
}

// Build input query parameter for the Get Compliance Statistics Breakdown API call
func buildComplianceBreakdownStatisticQueryParameter(_ context.Context, d *plugin.QueryData) url.Values {
	queryParameter := make(url.Values)

	for columnName, qual := range d.Quals {
		if qual != nil {
			operator := ""
			var val interface{}
			for _, qu := range qual.Quals {
				operator = qu.Operator
				val = qu.Value.GetStringValue()
			}
			switch columnName {
			case "account_id":
				if operator == "=" {
					queryParameter["cloud.account"] = []string{fmt.Sprint(val)}
				}
			case "cloud_type":
				if operator == "=" {
					queryParameter["cloud.type"] = []string{fmt.Sprint(val)}
				}
			case "cloud_region":
				if operator == "=" {
					queryParameter["cloud.region"] = []string{fmt.Sprint(val)}
				}
			case "policy_compliance_standard_name":
				if operator == "=" {
					queryParameter["policy.complianceStandard"] = []string{fmt.Sprint(val)}
				}
			case "policy_compliance_requirement_name":
				if operator == "=" {
					queryParameter["policy.complianceRequirement"] = []string{fmt.Sprint(val)}
				}
			case "policy_compliance_section_id":
				if operator == "=" {
					queryParameter["policy.complianceSection"] = []string{fmt.Sprint(val)}
				}
			}
		}
	}

	return queryParameter
}

// Build input query parameter for the Get Prioritized Vulnerabilities API call
func buildPrioritizedVulnerabilitiesQueryParameter(_ context.Context, d *plugin.QueryData) url.Values {
	queryParameter := make(url.Values)

	for columnName, qual := range d.Quals {
		if qual != nil {
			operator := ""
			var val interface{}
			for _, qu := range qual.Quals {
				operator = qu.Operator
				val = qu.Value.GetStringValue()
			}
			switch columnName {
			case "asset_type":
				if operator == "=" {
					queryParameter["asset_type"] = []string{fmt.Sprint(val)}
				}
			case "life_cycle":
				if operator == "=" {
					queryParameter["life_cycle"] = []string{fmt.Sprint(val)}
				}
			}
		}
	}

	return queryParameter
}