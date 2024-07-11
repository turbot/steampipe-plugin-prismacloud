package prisma

import (
	"context"

	"github.com/paloaltonetworks/prisma-cloud-go/policy"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePrismaPolicy(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prisma_policy",
		Description: "List of available policies in Prisma Cloud.",
		Get: &plugin.GetConfig{
			Hydrate:    getPrismaPolicy,
			KeyColumns: plugin.SingleColumn("policy_id"),
		},
		List: &plugin.ListConfig{
			Hydrate: listPrismaPolicies,
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
				Type:        proto.ColumnType_INT,
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
		plugin.Logger(ctx).Error("prisma_policy.listPrismaPolicies", "connection_error", err)
		return nil, err
	}

	policies, err := policy.List(conn, nil)
	if err != nil {
		plugin.Logger(ctx).Error("prisma_policy.listPrismaPolicies", "api_error", err)
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
		plugin.Logger(ctx).Error("prisma_policy.getPrismaPolicy", "connection_error", err)
		return nil, err
	}

	policy, err := policy.Get(conn, id)
	if err != nil {
		plugin.Logger(ctx).Error("prisma_policy.getPrismaPolicy", "api_error", err)
		return nil, err
	}

	return policy, nil
}
