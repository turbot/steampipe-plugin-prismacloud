package prisma

import (
	"context"

	"github.com/paloaltonetworks/prisma-cloud-go/alert/rule"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePrismaAlertRule(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prisma_alert_rule",
		Description: "List all information for prima cloud alert rules.",
		List: &plugin.ListConfig{
			Hydrate: listPrismaAlertRules,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "Name of the rule.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "policy_scan_config_id",
				Description: "ID of the policy scan configuration.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "Description of the rule.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "enabled",
				Description: "Whether the rule is enabled.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "scan_all",
				Description: "Whether to scan all policies.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "last_modified_on",
				Description: "Timestamp of the last modification.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("LastModifiedOn").Transform(transform.NullIfZeroValue).Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "last_modified_by",
				Description: "User who last modified the rule.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "allow_auto_remediate",
				Description: "Whether auto-remediation is allowed.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "delay_notification_ms",
				Description: "Delay for notifications in milliseconds.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "notify_on_open",
				Description: "Whether to notify when alert is open.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "notify_on_snoozed",
				Description: "Whether to notify when alert is snoozed.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "notify_on_dismissed",
				Description: "Whether to notify when alert is dismissed.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "notify_on_resolved",
				Description: "Whether to notify when alert is resolved.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "owner",
				Description: "Owner of the rule.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "open_alerts_count",
				Description: "Count of open alerts.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "read_only",
				Description: "Whether the rule is read-only.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "deleted",
				Description: "Whether the rule is deleted.",
				Type:        proto.ColumnType_BOOL,
			},

			// JSON fields
			{
				Name:        "notification_channels",
				Description: "Channels for notifications.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "notification_config",
				Description: "Notification configuration for the alert rule.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "policy_labels",
				Description: "Labels associated with the policies.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "excluded_policies",
				Description: "Policies that are excluded.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "target",
				Description: "Target configuration.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "policies",
				Description: "List of policies.",
				Type:        proto.ColumnType_JSON,
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the rule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

//// LIST FUNCTION

func listPrismaAlertRules(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prisma_alert_rule.listPrismaAlertRule", "connection_error", err)
		return nil, err
	}
	rules, err := rule.List(conn)
	if err != nil {
		plugin.Logger(ctx).Error("prisma_alert_rule.listPrismaAlertRule", "api_error", err)
		return nil, err
	}
	for _, rule := range rules {

		d.StreamListItem(ctx, rule)
		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}

	}
	return nil, nil
}
