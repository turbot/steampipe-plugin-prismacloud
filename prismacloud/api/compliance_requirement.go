package api


import (
	prismacloud "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/model"
)

// List Compliance Requirements
// https://pan.dev/prisma-cloud/api/cspm/get-requirements/
func ListComplianceRequirements(c *prismacloud.Client, complianceStandardId string) ([]*model.ComplianceRequirement, error) {
	c.Log(prismacloud.LogAction, "list of %s", "compliance requirements")

	var requirements []*model.ComplianceRequirement
	if _, err := c.Communicate("GET", []string{"compliance", complianceStandardId, "requirement"}, nil, nil, &requirements); err != nil {
		return nil, err
	}

	return requirements, nil
}

// Get Compliance Standard by ID
// https://pan.dev/prisma-cloud/api/cspm/get-requirement-by-id/
func GetComplianceRequirement(c *prismacloud.Client, requirementId string) (*model.ComplianceRequirement, error) {
	c.Log(prismacloud.LogAction, "get %s", "compliance requirement")

	var requirement *model.ComplianceRequirement
	if _, err := c.Communicate("GET", []string{"compliance", "requirement", requirementId}, nil, nil, &requirement); err != nil {
		return nil, err
	}

	return requirement, nil
}

// List Compliance Requirement Sections
// https://pan.dev/prisma-cloud/api/cspm/get-sections/
func ListComplianceRequirementSections(c *prismacloud.Client, requirementId string) ([]*model.ComplianceRequirementSection, error) {
	c.Log(prismacloud.LogAction, "list of %s", "compliance requirements")

	var sections []*model.ComplianceRequirementSection
	if _, err := c.Communicate("GET", []string{"compliance", requirementId, "section"}, nil, nil, &sections); err != nil {
		return nil, err
	}

	return sections, nil
}
