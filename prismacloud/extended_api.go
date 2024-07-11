package prismacloud

import prismacloud "github.com/paloaltonetworks/prisma-cloud-go"


func LisCompliancePostures(c *prismacloud.Client) (*ComplianceData, error) {
	c.Log(prismacloud.LogAction, "(get) list of %s", "compliance postures")

	var postures ComplianceData
	if _, err := c.Communicate("GET", []string{"v2", "compliance", "posture"}, nil, nil, &postures); err != nil {
		return nil, err
	}

	return &postures, nil
}
