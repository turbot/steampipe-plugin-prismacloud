package prismacloud

import (
	"context"

	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/api"
	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/model"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePrismaComplianceBreakdownSectionSummary(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prismacloud_compliance_breakdown_section_summary",
		Description: "List all available compliance breakdown statistics of section summary.",
		List: &plugin.ListConfig{
			Hydrate:    listPrismaComplianceBreakdownSectionSummary,
			KeyColumns: commonComplianceBreakdownKeyQualColumns(),
		},
		Columns: complianceBreakdownCommonFilterColumns([]*plugin.Column{
			{
				Name:        "requirement_id",
				Description: "The unique identifier for the section.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "requirement_name",
				Description: "The unique identifier for the section.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "The unique identifier for the section.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "The name of the section.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "failed_resources",
				Description: "The number of failed compliance section scanned resources.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "passed_resources",
				Description: "The number of passed compliance section scanned resources.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "total_resources",
				Description: "The total number of compliance section scanned resources.",
				Type:        proto.ColumnType_INT,
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the section summary.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		}),
	}
}

type BreakdownComplianceSectionSummary struct {
	RequirementId   string
	RequirementName string
	model.SectionSummary
}

//// LIST FUNCTION

func listPrismaComplianceBreakdownSectionSummary(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_compliance_breakdown_section_summary.listPrismaComplianceBreakdownSectionSummary", "connection_error", err)
		return nil, err
	}

	query := buildComplianceBreakdownStatisticQueryParameter(ctx, d)

	postures, err := api.LisComplianceBreakdownStatistics(conn, query)
	if err != nil {

		plugin.Logger(ctx).Error("prismacloud_compliance_breakdown_section_summary.listPrismaComplianceBreakdownSectionSummary", "api_error", err)
		return nil, err
	}

	for _, requirement := range postures.RequirementSummaries {
		for _, section := range requirement.SectionSummaries {
			d.StreamListItem(ctx, BreakdownComplianceSectionSummary{requirement.ID, requirement.Name, section})

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}
