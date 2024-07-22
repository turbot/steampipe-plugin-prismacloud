package prismacloud

import (
	"context"

	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePrismaComplianceBreakdownStatistic(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prismacloud_compliance_breakdown_statistic",
		Description: "List all available compliance breakdown statistics.",
		List: &plugin.ListConfig{
			Hydrate:    listPrismaComplianceBreakdownStatistics,
			KeyColumns: commonComplianceBreakdownKeyQualColumns(),
		},
		Columns: complianceBreakdownCommonFilterColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "Name of the Compliance Standard/Requirement/Section.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "assigned_policies",
				Description: "Number of policies assigned to the Compliance Standard/Requirement/Section.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "critical_severity_failed_resources",
				Description: "Number of Compliance Standard/Requirement/Section scanned resources failing critical severity policies.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "default",
				Description: "Indicates if it is a system default.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "description",
				Description: "Description of the Compliance Standard/Requirement/Section.",
				Type:        proto.ColumnType_STRING,
			},

			{
				Name:        "failed_resources",
				Description: "Number of failing Compliance Standard/Requirement/Section scanned resources.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "high_severity_failed_resources",
				Description: "Number of Compliance Standard/Requirement/Section scanned resources failing high severity policies.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "id",
				Description: "ID of the Compliance Standard/Requirement/Section.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "informational_severity_failed_resources",
				Description: "Number of Compliance Standard/Requirement/Section scanned resources failing informational severity policies.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "low_severity_failed_resources",
				Description: "Number of Compliance Standard/Requirement/Section scanned resources failing low severity policies.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "medium_severity_failed_resources",
				Description: "Number of Compliance Standard/Requirement/Section scanned resources failing medium severity policies.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "passed_resources",
				Description: "Number of passing Compliance Standard/Requirement/Section scanned resources.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "total_resources",
				Description: "Total number of Compliance Standard/Requirement/Section scanned resources.",
				Type:        proto.ColumnType_INT,
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the compliance posture.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		}),
	}
}

//// LIST FUNCTION

func listPrismaComplianceBreakdownStatistics(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_compliance_breakdown_statistic.listPrismaComplianceBreakdownStatistics", "connection_error", err)
		return nil, err
	}

	query := buildComplianceBreakdownStatisticQueryParameter(ctx, d)

	postures, err := api.LisComplianceBreakdownStatistics(conn, query)
	if err != nil {

		plugin.Logger(ctx).Error("prismacloud_compliance_breakdown_statistic.listPrismaComplianceBreakdownStatistics", "api_error", err)
		return nil, err
	}

	for _, posture := range postures.ComplianceDetails {

		d.StreamListItem(ctx, posture)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}

	}

	return nil, nil
}
