package prismacloud

import (
	"context"

	"github.com/paloaltonetworks/prisma-cloud-go/report"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePrismacloudReport(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prismacloud_report",
		Description: "List of available alert and compliance reports.",
		Get: &plugin.GetConfig{
			Hydrate:    getPrismacloudReport,
			KeyColumns: plugin.SingleColumn("id"),
		},
		List: &plugin.ListConfig{
			Hydrate: listPrismacloudReports,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The unique identifier for the report.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "The name of the report.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "type",
				Description: "The type of the report.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "cloud_type",
				Description: "The type of cloud (e.g., AWS, Azure, GCP).",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "compliance_standard_id",
				Description: "The ID of the compliance standard associated with the report.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "target",
				Description: "The target configuration of the report.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "status",
				Description: "The status of the report.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_on",
				Description: "The timestamp when the report was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("CreatedOn").Transform(transform.NullIfZeroValue).Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "created_by",
				Description: "The user who created the report.",
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
				Description: "The user who last modified the report.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "next_schedule",
				Description: "The timestamp of the next scheduled run.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("NextSchedule").Transform(transform.NullIfZeroValue).Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "last_scheduled",
				Description: "The timestamp of the last scheduled run.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("LastSchedule").Transform(transform.NullIfZeroValue).Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "total_instance_count",
				Description: "The total number of instances in the report.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "counts",
				Description: "Counts of various metrics in the report.",
				Type:        proto.ColumnType_JSON,
			},

			// Steampipe standard column
			{
				Name:        "title",
				Description: "Title of the report.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		}),
	}
}

//// LIST FUNCTION

func listPrismacloudReports(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_report.listPrismacloudReports", "connection_error", err)
		return nil, err
	}

	reports, err := report.List(conn)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_report.listPrismacloudReports", "api_error", err)
		return nil, err
	}

	for _, report := range reports {

		d.StreamListItem(ctx, report)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}

	}

	return nil, nil
}

//// HYDRATE FUNCTION

func getPrismacloudReport(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQualString("id")

	// Empty check
	if id == "" {
		return nil, nil
	}

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_report.getPrismacloudReport", "connection_error", err)
		return nil, err
	}

	report, err := report.Get(conn, id)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_report.getPrismacloudReport", "api_error", err)
		return nil, err
	}

	return report, nil
}
