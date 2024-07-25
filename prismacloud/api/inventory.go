package api

import (
	prismacloud "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/model"
)

func ListInventoryDiscoveredAPI(c *prismacloud.Client, req map[string]interface{}) (*model.InventoryDiscoveredAPIResponse, error) {
	c.Log(prismacloud.LogAction, "list of %s", "inventory api endpoints")
	// https://api.anz.prismacloud.io/waas-api-discovery/api/v1/discovered-api
	var apis model.InventoryDiscoveredAPIResponse
	if _, err := c.Communicate("POST", []string{"waas-api-discovery", "api", "v1", "discovered-api"}, nil, req, &apis); err != nil {
		return nil, err
	}

	return &apis, nil
}
