package api

import (
	prismacloud "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/model"
)

// Get Alert Count of Policies
// https://pan.dev/prisma-cloud/api/cspm/alert-policy-list/
func GetAlertCountOfPolicies(c *prismacloud.Client, req map[string]interface{}) (*model.AlertCount, error) {
	c.Log(prismacloud.LogAction, "get %s", "Alert Count of Policies")

	var alertCounts model.AlertCount
	if _, err := c.Communicate("POST", []string{"alert", "v1", "policy"}, nil, req, &alertCounts); err != nil {
		return nil, err
	}

	return &alertCounts, nil
}
