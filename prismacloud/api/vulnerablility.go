package api

import (
	"net/url"

	prismacloud "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/model"
)

// Get Vulnerability Overview
// https://pan.dev/prisma-cloud/api/cspm/vulnerability-dashboard-overview-v-2/
func GetVulnerabilityOverview(c *prismacloud.Client) (*model.VulnerableOverviewSummary, error) {
	c.Log(prismacloud.LogAction, "get %s", "vulnerabilities overview")

	var vulnerabilities model.VulnerableOverviewSummary
	if _, err := c.Communicate("GET", []string{"uve", "api", "v2", "dashboard", "vulnerabilities", "overview"}, nil, nil, &vulnerabilities); err != nil {
		return nil, err
	}

	return &vulnerabilities, nil
}

// Get Prioritized Vulnerabilities
// https://pan.dev/prisma-cloud/api/cspm/prioritised-vulnerability-v-4/
func GetPrioritizedVulnerability(c *prismacloud.Client, query url.Values) (*model.PrioritizedVulnerabilitySummary, error) {
	c.Log(prismacloud.LogAction, "get of %s", "prioritized vulnerabilities")

	var vulnerabilities model.PrioritizedVulnerabilitySummary
	if _, err := c.Communicate("GET", []string{"uve", "api", "v4", "dashboard", "vulnerabilities", "prioritised"}, query, nil, &vulnerabilities); err != nil {
		return nil, err
	}

	return &vulnerabilities, nil
}


// Get Vulnerabilities Burndown
// https://pan.dev/prisma-cloud/api/cspm/get-burndown/
func ListVulnerabilityBurndown(c *prismacloud.Client, query url.Values) ([]*model.VulnerabilityBurndownSummary, error) {
	c.Log(prismacloud.LogAction, "list of %s", "vulnerabilities burldown")

	var burndownSummary []*model.VulnerabilityBurndownSummary
	if _, err := c.Communicate("GET", []string{"uve", "api", "v2", "dashboard", "vulnerabilities", "burndown"}, query, nil, &burndownSummary); err != nil {
		return nil, err
	}

	return burndownSummary, nil
}

// Get Vulnerable Assets
// https://pan.dev/prisma-cloud/api/cspm/vulnerable-assets/
func ListVulnerabilityAssets(c *prismacloud.Client, query url.Values) (*model.VulnerableAssets, error) {
	c.Log(prismacloud.LogAction, "list of %s", "vulnerability assets")

	var assets *model.VulnerableAssets

	if _, err := c.Communicate("GET", []string{"uve", "api", "v1", "dashboard", "vulnerabilities", "vulnerableAsset"}, query, nil, &assets); err != nil {
		return nil, err
	}

	return assets, nil
}
