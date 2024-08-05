package api

import (
	prismacloud "github.com/paloaltonetworks/prisma-cloud-go"
	"net/url"
	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/model"
)


// Get Compliance Statistics Breakdown
// https://pan.dev/prisma-cloud/api/cspm/get-compliance-posture-v-2/
// Query parameter:
// query := url.Values{
//		"cloud.type": []string{"aws"},
// }
func LisComplianceBreakdownStatistics(c *prismacloud.Client, query url.Values) (*model.ComplianceData, error) {
	c.Log(prismacloud.LogAction, "(get) list of %s", "compliance postures")

	var postures model.ComplianceData
	if _, err := c.Communicate("GET", []string{"v2", "compliance", "posture"}, query, nil, &postures); err != nil {
		return nil, err
	}

	return &postures, nil
}
