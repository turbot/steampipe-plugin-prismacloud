package api

import (
	prismacloud "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/model"
)

// My profile
// https://pan.dev/prisma-cloud/api/cspm/get-my-profile/
func GetCurrentUserProfile(c *prismacloud.Client) (*model.UserProfile, error) {
	c.Log(prismacloud.LogAction, "get %s", "my profile")

	var profile *model.UserProfile
	if _, err := c.Communicate("GET", []string{"user", "me"}, nil, nil, &profile); err != nil {
		return nil, err
	}

	return profile, nil
}
