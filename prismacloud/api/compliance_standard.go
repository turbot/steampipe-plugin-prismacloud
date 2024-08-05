package api

import (
	prismacloud "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/model"
)

// List Compliance Standards
// https://pan.dev/prisma-cloud/api/cspm/get-all-standards/
func ListComplianceStandards(c *prismacloud.Client) ([]*model.ComplianceStandard, error) {
	c.Log(prismacloud.LogAction, "(get) list of %s", "compliance standards")

	var standards []*model.ComplianceStandard
	if _, err := c.Communicate("GET", []string{"compliance"}, nil, nil, &standards); err != nil {
		return nil, err
	}

	return standards, nil
}

// Get Compliance Standard by ID
// https://pan.dev/prisma-cloud/api/cspm/get-standards-by-id/
func GetComplianceStandard(c *prismacloud.Client, complianceStandardId string) (*model.ComplianceStandard, error) {
	c.Log(prismacloud.LogAction, "get %s", "compliance standard")

	var standard *model.ComplianceStandard
	if _, err := c.Communicate("GET", []string{"compliance", complianceStandardId}, nil, nil, &standard); err != nil {
		return nil, err
	}

	return standard, nil
}