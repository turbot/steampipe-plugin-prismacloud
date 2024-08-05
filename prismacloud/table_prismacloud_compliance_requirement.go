package prismacloud

import (
	"context"

	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/api"
	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/model"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePrismacloudComplianceRequirement(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prismacloud_compliance_requirement",
		Description: "List all available compliance requirement.",
		Get: &plugin.GetConfig{
			Hydrate:    getPrismacloudComplianceRequirement,
			KeyColumns: plugin.SingleColumn("id"),
		},
		List: &plugin.ListConfig{
			ParentHydrate: listPrismacloudComplianceStandards,
			Hydrate:       listPrismacloudComplianceRequirements,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "compliance_id", Require: plugin.Optional},
			},
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the requirement.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "compliance_id",
				Description: "The unique identifier for the compliance standard.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ComplianceID"),
			},
			{
				Name:        "id",
				Description: "The unique identifier for the requirement.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromGo(),
			},
			{
				Name:        "created_by",
				Description: "The user who created the requirement.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_on",
				Description: "The timestamp when the requirement was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("CreatedOn").Transform(transform.NullIfZeroValue).Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "description",
				Description: "The description of the requirement.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_modified_by",
				Description: "The user who last modified the requirement.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_modified_on",
				Description: "The timestamp when the requirement was last modified.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("LastModifiedOn").Transform(transform.NullIfZeroValue).Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "policies_assigned_count",
				Description: "The number of policies assigned to the requirement.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "requirement_id",
				Description: "The unique identifier for the requirement within the compliance standard.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("RequirementID"),
			},
			{
				Name:        "standard_name",
				Description: "The name of the compliance standard.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "system_default",
				Description: "Indicates if the requirement is a system default.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "view_order",
				Description: "The order in which the requirement should be viewed.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "requirement_sections",
				Description: "All compliance requirement sections for the specified compliance requirement.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPrismacloudComplianceRequirementSections,
				Transform:   transform.FromValue(),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the compliance requirement.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		}),
	}
}

//// LIST FUNCTION

func listPrismacloudComplianceRequirements(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	standardCompliance := h.Item.(*model.ComplianceStandard)

	complianceId := standardCompliance.ID

	// Restrict API call with given compliance ID
	if d.EqualsQualString("compliance_id") != "" && d.EqualsQualString("compliance_id") != complianceId {
		return nil, nil
	}

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_compliance_requirement.listPrismacloudComplianceRequirements", "connection_error", err)
		return nil, err
	}

	requirements, err := api.ListComplianceRequirements(conn, complianceId)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_compliance_requirement.listPrismacloudComplianceRequirements", "api_error", err)
		return nil, err
	}

	for _, requirement := range requirements {
		// The API is not returning the complianceId, though it has defined in schema.
		requirement.ComplianceID = complianceId
		d.StreamListItem(ctx, requirement)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTION

func getPrismacloudComplianceRequirement(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQualString("id")

	// Empty check
	if id == "" {
		return nil, nil
	}

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_compliance_requirement.getPrismacloudComplianceRequirement", "connection_error", err)
		return nil, err
	}

	requirement, err := api.GetComplianceRequirement(conn, id)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_compliance_requirement.getPrismacloudComplianceRequirement", "api_error", err)
		return nil, err
	}

	if requirement != nil {
		return requirement, nil
	}

	return nil, nil
}

func getPrismacloudComplianceRequirementSections(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	requirement := h.Item.(*model.ComplianceRequirement)

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_compliance_requirement.getPrismacloudComplianceRequirement", "connection_error", err)
		return nil, err
	}

	requirementSessions, err := api.ListComplianceRequirementSections(conn, requirement.ID)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_compliance_requirement.getPrismacloudComplianceRequirement", "api_error", err)
		return nil, err
	}

	if requirementSessions != nil {
		return requirementSessions, nil
	}

	return nil, nil
}
