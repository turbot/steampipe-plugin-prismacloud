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
)

func tablePrismaComplianceBreakdownRequirementSummary(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prismacloud_compliance_breakdown_requirement_summary",
		Description: "List all available compliance breakdown statistics of requirement summary.",
		List: &plugin.ListConfig{
			ParentHydrate: listPrismaAccounts,
			Hydrate:       listPrismaComplianceBreakdownRequirementSummary,
			KeyColumns:    commonComplianceBreakdownKeyQualColumns(),
		},
		Columns: commonColumns(complianceBreakdownCommonFilterColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The unique identifier for the requirement.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "The name the requirement.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "section_summaries",
				Description: "The summary of the sections.",
				Type:        proto.ColumnType_JSON,
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the section summary.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("RequirementName"),
			},
		})),
	}
}

type BreakdownComplianceSectionSummary struct {
	AccountName      string
	AccountId        string
	CloudType        string
	Id               string
	Name             string
	SectionSummaries []model.SectionSummary
}

//// LIST FUNCTION

func listPrismaComplianceBreakdownRequirementSummary(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	account := h.Item.(account.Account)

	if d.EqualsQualString("account_name") != "" && d.EqualsQualString("account_name") != account.Name {
		return nil, nil
	}

	if d.EqualsQualString("cloud_type") != "" && d.EqualsQualString("cloud_type") != account.CloudType {
		return nil, nil
	}

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_compliance_breakdown_requirement_summary.listPrismaComplianceBreakdownRequirementSummary", "connection_error", err)
		return nil, err
	}

	// For any of the query parameter it the returning the same row. However, the query param is required to make the the API call do hardcoded the value, and added it as a optional qual.
	// If we are not specifying any of the parameter the API doesn't return any result.
	query := url.Values{
		"cloud.account": []string{account.Name},
	}

	query = buildComplianceBreakdownStatisticQueryParameter(ctx, d, query)

	postures, err := api.LisComplianceBreakdownStatistics(conn, query)
	if err != nil {

		plugin.Logger(ctx).Error("prismacloud_compliance_breakdown_requirement_summary.listPrismaComplianceBreakdownRequirementSummary", "api_error", err)
		return nil, err
	}

	for _, requirement := range postures.RequirementSummaries {
		d.StreamListItem(ctx, BreakdownComplianceSectionSummary{account.Name, account.AccountId, account.CloudType, requirement.ID, requirement.Name, requirement.SectionSummaries})

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
