package api

import (
	"net/url"

	prismacloud "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/model"
)

// Get Permissions V4
// https://pan.dev/prisma-cloud/api/cspm/permission-search-v-4/
func ListIAMPermissions(c *prismacloud.Client, query url.Values, req map[string]interface{}) (*model.PermissionResponse, error) {
	c.Log(prismacloud.LogAction, "list %s", "IAM permission")

	var prems model.PermissionResponse
	if _, err := c.Communicate("POST", []string{"iam", "api", "v4", "search", "permission"}, query, req, &prems); err != nil {
		return nil, err
	}

	return &prems, nil
}
