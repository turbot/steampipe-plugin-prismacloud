package api

import (
	"net/url"

	prismacloud "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/model"
)

// Resource Scan Info V2 - GET
// https://pan.dev/prisma-cloud/api/cspm/get-resource-scan-info-v-2/
// Query parameter:
//
//	query := url.Values{
//			"cloud.type": []string{"aws"},
//	}
func ListInventoryAssetExplorer(c *prismacloud.Client, query url.Values) (*model.PrismaCloudAssetExplorer, error) {
	c.Log(prismacloud.LogAction, "(get) list of %s", "asses explorers")

	var explorers model.PrismaCloudAssetExplorer

	if _, err := c.Communicate("GET", []string{"v2", "resource", "scan_info"}, query, nil, &explorers); err != nil {
		return nil, err
	}

	return &explorers, nil
}
