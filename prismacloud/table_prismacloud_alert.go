package prismacloud

import (
	"context"
	"time"

	"github.com/paloaltonetworks/prisma-cloud-go/alert"
	"github.com/paloaltonetworks/prisma-cloud-go/timerange"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePrismaAlert(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prismacloud_alert",
		Description: "List all information for prima cloud alerts.",
		Get: &plugin.GetConfig{
			Hydrate:    getPrismaAlert,
			KeyColumns: plugin.SingleColumn("id"),
		},
		List: &plugin.ListConfig{
			Hydrate: listPrismaAlerts,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "alert_time", Require: plugin.Optional, Operators: []string{"=", ">=", "<=", ">", "<"}},
				{Name: "status", Require: plugin.Optional, Operators: []string{"="}},
				{Name: "policy_id", Require: plugin.Optional, Operators: []string{"="}},
				{Name: "policy_type", Require: plugin.Optional, Operators: []string{"="}},
				{Name: "policy_remediable", Require: plugin.Optional, Operators: []string{"=", "<>"}},
			},
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The unique identifier for the alert.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status",
				Description: "The current status of the alert.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "first_seen",
				Description: "The timestamp when the alert was first seen.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("FirstSeen").Transform(transform.NullIfZeroValue).Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "last_seen",
				Description: "The timestamp when the alert was last seen.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("LastSeen").Transform(transform.NullIfZeroValue).Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "alert_time",
				Description: "The timestamp when the alert was triggered.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("AlertTime").Transform(transform.NullIfZeroValue).Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "event_occurred",
				Description: "The timestamp when the event occurred.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "triggered_by",
				Description: "The entity that triggered the alert.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "alert_count",
				Description: "The count of how many times the alert was triggered.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "history",
				Description: "The history of the alert.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "policy_id",
				Description: "The ID of the policy associated with the alert.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Policy.Id"),
			},
			{
				Name:        "policy_type",
				Description: "The type of the policy associated with the alert.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Policy.Type"),
			},
			{
				Name:        "policy_remediable",
				Description: "If the policy associated with the alert is remediable.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Policy.Remediable"),
			},
			{
				Name:        "policy_system_default",
				Description: "If the policy associated with the alert is system default.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Policy.SystemDefault"),
			},
			{
				Name:        "risk_detail",
				Description: "The risk details associated with the alert.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Risk"),
			},
			{
				Name:        "resource",
				Description: "The resource associated with the alert.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "investigate_options",
				Description: "Options for investigating the alert.",
				Type:        proto.ColumnType_JSON,
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the alert.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Id"),
			},
		}),
	}
}

//// LIST FUNCTION

func listPrismaAlerts(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_alert.listPrismaAlerts", "connection_error", err)
		return nil, err
	}

	// https://pan.dev/prisma-cloud/api/cspm/get-alerts-v-2/
	// Limiting the results
	maxLimit := int32(10000)
	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)
		if limit < maxLimit {
			maxLimit = limit
		}
	}

	// Default to last 6 months
	// The API is returning a huge amount of data for last 1 year. So set the default time range of 6 months.
	timeRange := timerange.Absolute{
		Start: time.Now().AddDate(0, -6, 0).Second() * 1000,
		End:   time.Now().Second(),
	}
	st, et := getAlertStartTImeAndSearchEndTime(d)
	if st != 0 {
		timeRange.Start = st
	}
	if et != 0 {
		timeRange.End = et
	}

	req := alert.Request{
		Limit:     int(maxLimit),
		Detailed:  true,
		Offset:    0,
		PageToken: "",
		TimeRange: timerange.TimeRange{
			Value: timeRange,
		},
	}

	filter := getAlertFilter(d)
	if len(filter) > 0 {
		req.Filters = filter
	}

	alerts, err := alert.List(conn, req)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_alert.listPrismaAlerts", "api_error", err)
		return nil, err
	}
	for _, alert := range alerts.Data {

		d.StreamListItem(ctx, alert)
		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}

	}

	for alerts.PageToken != "" {
		req.Offset = req.Offset + alerts.Total
		req.PageToken = alerts.PageToken

		alerts, err = alert.List(conn, req)
		if err != nil {
			plugin.Logger(ctx).Error("prismacloud_alert.listPrismaAlerts", "api_paging_error", err)
			return nil, err
		}
		for _, alert := range alerts.Data {

			d.StreamListItem(ctx, alert)
			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}

		}
	}
	return nil, nil
}

//// HYDRATE FUNCTION

func getPrismaAlert(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQualString("id")

	// Empty check
	if id == "" {
		return nil, nil
	}

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_alert.getPrismaAlert", "connection_error", err)
		return nil, err
	}

	alert, err := alert.Get(conn, id)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_alert.getPrismaAlert", "api_error", err)
		return nil, err
	}

	return alert, nil
}

//// UTILITY FUNCTION

// Build the filter parameter

func getAlertFilter(keyQuals *plugin.QueryData) []alert.Filter {
	var filter []alert.Filter

	qualsMap := map[string]string{
		"status":            "alert.status",
		"policy_id":         "policy.id",
		"policy_type":       "policy.type",
		"policy_remediable": "policy.remediable",
	}

	for columnName, filterValue := range qualsMap {
		if keyQuals.Quals[columnName] != nil {
			if columnName == "policy_remediable" {
				f := alert.Filter{
					Name:     filterValue,
					Operator: "=",
					Value:    "",
				}
				for _, q := range keyQuals.Quals[columnName].Quals {
					if q.Operator == "=" {
						f.Value = "true"
					}
					if q.Operator == "<>" {
						f.Value = "false"
					}
				}
				filter = append(filter, f)
				continue
			}
			f := alert.Filter{
				Name:     filterValue,
				Operator: "=",
				Value:    "",
			}
			for _, q := range keyQuals.Quals[columnName].Quals {
				if q.Operator == "=" {
					f.Value = q.Value.GetStringValue()
				}
			}

			filter = append(filter, f)
		}
	}
	return filter
}

func getAlertStartTImeAndSearchEndTime(keyQuals *plugin.QueryData) (int, int) {

	st, et := 0, 0

	if keyQuals.Quals["alert_time"] != nil && !(len(keyQuals.Quals["alert_time"].Quals) > 1) {
		for _, q := range keyQuals.Quals["period_start"].Quals {
			t := q.Value.GetTimestampValue().AsTime()
			switch q.Operator {
			case "=", ">=", ">":
				st = t.Second() * 1000
				et = time.Now().Second() * 1000
			case "<", "<=":
				et = t.Second() * 1000
				st = t.AddDate(0, -3, 0).Second() * 1000
			}
		}
	}

	return st, et
}
