package prismacloud

import (
	"context"

	alertIp "github.com/paloaltonetworks/prisma-cloud-go/trusted-alert-ip"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePrismaTrustedAlertIp(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "prismacloud_trusted_alert_ip",
		Description: "List of trusted alert IPs in Prisma Cloud.",
		Get: &plugin.GetConfig{
			Hydrate:    getPrismaTrustedAlertIp,
			KeyColumns: plugin.SingleColumn("uuid"),
		},
		List: &plugin.ListConfig{
			Hydrate: listPrismaTrustedAlertIps,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the trusted alert IP.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "uuid",
				Description: "The unique identifier of trusted alert IP.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("UUID"),
			},
			{
				Name:        "cidr_count",
				Description: "The number of CIRDs.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "cidrs",
				Description: "The CIDR blocks of trusted alert IP.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("CIDRS"),
			},

			// Steampipe standard column
			{
				Name:        "title",
				Description: "Title of the permission group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

//// LIST FUNCTION

func listPrismaTrustedAlertIps(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_trusted_alert_ip.listPrismaTrustedAlertIps", "connection_error", err)
		return nil, err
	}

	alertIps, err := alertIp.List(conn)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_trusted_alert_ip.listPrismaTrustedAlertIps", "api_error", err)
		return nil, err
	}

	for _, alertIp := range alertIps {

		d.StreamListItem(ctx, alertIp)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}

	}

	return nil, nil
}

//// HYDRATE FUNCTION

func getPrismaTrustedAlertIp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQualString("uuid")

	// Empty check
	if id == "" {
		return nil, nil
	}

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_trusted_alert_ip.getPrismaTrustedAlertIp", "connection_error", err)
		return nil, err
	}

	alertIp, err := alertIp.Get(conn, id)
	if err != nil {
		plugin.Logger(ctx).Error("prismacloud_trusted_alert_ip.getPrismaTrustedAlertIp", "api_error", err)
		return nil, err
	}

	return alertIp, nil
}
