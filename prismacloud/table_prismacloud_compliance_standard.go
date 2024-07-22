package prismacloud

import (
	"context"

	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePrismaComplianceStandard(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prismacloud_compliance_standard",
		Description: "List all available compliance standard.",
		Get: &plugin.GetConfig{
			Hydrate:    getPrismaComplianceStandard,
			KeyColumns: plugin.SingleColumn("id"),
		},
		List: &plugin.ListConfig{
			Hydrate: listPrismaComplianceStandards,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the standard.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "The unique identifier for the standard.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromGo(),
			},
			{
				Name:        "policies_assigned_count",
				Description: "The number of policies assigned to the standard.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "system_default",
				Description: "Indicates if the standard is a system default.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "created_on",
				Description: "The timestamp when the standard was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("CreatedOn").Transform(transform.NullIfZeroValue).Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "created_by",
				Description: "The user who created the standard.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "The description of the standard.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_modified_by",
				Description: "The user who last modified the standard.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_modified_on",
				Description: "The timestamp when the standard was last modified.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("LastModifiedOn").Transform(transform.NullIfZeroValue).Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "cloud_type",
				Description: "The types of cloud environments the standard applies to.",
				Type:        proto.ColumnType_JSON,
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the compliance standard.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

//// LIST FUNCTION

func listPrismaComplianceStandards(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_compliance_standard.listPrismaComplianceStandards", "connection_error", err)
		return nil, err
	}

	standards, err := api.ListComplianceStandards(conn)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_compliance_standard.listPrismaComplianceStandards", "api_error", err)
		return nil, err
	}

	for _, standard := range standards {

		d.StreamListItem(ctx, standard)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}

	}

	return nil, nil
}

//// HYDRATE FUNCTION

func getPrismaComplianceStandard(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQualString("id")

	// Empty check
	if id == "" {
		return nil, nil
	}

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_compliance_standard.getPrismaComplianceStandard", "connection_error", err)
		return nil, err
	}

	standard, err := api.GetComplianceStandard(conn, id)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_compliance_standard.getPrismaComplianceStandard", "api_error", err)
		return nil, err
	}

	if standard != nil {
		return standard, nil
	}

	return nil, nil
}
