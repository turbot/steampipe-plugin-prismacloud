package prismacloud

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin/transform"
)

func commonColumns(cols []*plugin.Column) []*plugin.Column {
	return append([]*plugin.Column{
		{
			Name:        "email",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getCurrentUserEmail,
			Description: "Email address of the current session user.",
			Transform:   transform.FromValue(),
		},
	}, cols...)
}
