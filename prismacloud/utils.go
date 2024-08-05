package prismacloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/api"
	"github.com/turbot/steampipe-plugin-prismacloud/prismacloud/model"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"github.com/turbot/steampipe-plugin-sdk/v5/query_cache"
)

// Common filter columns for the Compliance Breakdown
func complianceBreakdownCommonFilterColumns(columns []*plugin.Column) []*plugin.Column {
	return append(columns, []*plugin.Column{
		{
			Name:        "account_id",
			Description: "The unique identifier for the account.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "account_name",
			Description: "The unique identifier for the account.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "cloud_type",
			Description: "The type of cloud (e.g., AWS, Azure, GCP).",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "cloud_region",
			Description: "The region of the cloud where the resource is located.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromQual("cloud_region"),
		},
		{
			Name:        "policy_compliance_standard_name",
			Description: "The name of the compliance standard associated with the policy.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromQual("policy_compliance_standard_name"),
		},
		{
			Name:        "policy_compliance_requirement_name",
			Description: "The name of the compliance requirement associated with the policy.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromQual("policy_compliance_requirement_name"),
		},
		{
			Name:        "policy_compliance_section_id",
			Description: "The ID of the compliance section associated with the policy.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromQual("policy_compliance_section_id"),
		},
	}...)
}

// Common key columns for the the Compliance Breakdown
func commonComplianceBreakdownKeyQualColumns() plugin.KeyColumnSlice {
	commonKeyQualsCol := plugin.KeyColumnSlice{
		{Name: "account_name", Require: plugin.Optional},
		{Name: "cloud_type", Require: plugin.Optional},
		{Name: "cloud_region", Require: plugin.Optional, CacheMatch: query_cache.CacheMatchExact},
		{Name: "policy_compliance_standard_name", Require: plugin.Optional, CacheMatch: query_cache.CacheMatchExact},
		{Name: "policy_compliance_requirement_name", Require: plugin.Optional, CacheMatch: query_cache.CacheMatchExact},
		{Name: "policy_compliance_section_id", Require: plugin.Optional, CacheMatch: query_cache.CacheMatchExact},
	}

	return commonKeyQualsCol
}

// Build input query parameter for the Get Compliance Statistics Breakdown API call
func buildComplianceBreakdownStatisticQueryParameter(_ context.Context, d *plugin.QueryData, queryParameter url.Values) url.Values {

	for columnName, qual := range d.Quals {
		if qual != nil {
			operator := ""
			var val interface{}
			for _, qu := range qual.Quals {
				operator = qu.Operator
				val = qu.Value.GetStringValue()
			}
			switch columnName {
			case "cloud_type":
				if operator == "=" {
					queryParameter["cloud.type"] = []string{fmt.Sprint(val)}
				}
			case "cloud_region":
				if operator == "=" {
					queryParameter["cloud.region"] = []string{fmt.Sprint(val)}
				}
			case "policy_compliance_standard_name":
				if operator == "=" {
					queryParameter["policy.complianceStandard"] = []string{fmt.Sprint(val)}
				}
			case "policy_compliance_requirement_name":
				if operator == "=" {
					queryParameter["policy.complianceRequirement"] = []string{fmt.Sprint(val)}
				}
			case "policy_compliance_section_id":
				if operator == "=" {
					queryParameter["policy.complianceSection"] = []string{fmt.Sprint(val)}
				}
			}
		}
	}

	return queryParameter
}

// Build input query parameter for the Get Prioritized Vulnerabilities API call
func buildPrioritizedVulnerabilitiesQueryParameter(_ context.Context, d *plugin.QueryData) url.Values {
	queryParameter := make(url.Values)

	for columnName, qual := range d.Quals {
		if qual != nil {
			operator := ""
			var val interface{}
			for _, qu := range qual.Quals {
				operator = qu.Operator
				val = qu.Value.GetStringValue()
			}
			switch columnName {
			case "asset_type":
				if operator == "=" {
					queryParameter["asset_type"] = []string{fmt.Sprint(val)}
				}
			case "life_cycle":
				if operator == "=" {
					queryParameter["life_cycle"] = []string{fmt.Sprint(val)}
				}
			}
		}
	}

	return queryParameter
}

// Build input query parameter for the Get Vulnerabilities Burndown API call
func buildBurndownVulnerabilitiesQueryParameter(_ context.Context, d *plugin.QueryData) url.Values {
	queryParameter := make(url.Values)

	for columnName, qual := range d.Quals {
		if qual != nil {
			operator := ""
			var val interface{}
			for _, qu := range qual.Quals {
				operator = qu.Operator
				val = qu.Value.GetStringValue()
			}
			switch columnName {
			case "asset_type":
				if operator == "=" {
					queryParameter["asset_type"] = []string{fmt.Sprint(val)}
				}
			case "life_cycle":
				if operator == "=" {
					queryParameter["life_cycle"] = []string{fmt.Sprint(val)}
				}
			case "severities":
				if operator == "=" {
					queryParameter["severities"] = []string{fmt.Sprint(val)}
				}
			}
		}
	}

	return queryParameter
}

// Build input query parameter for the Get Vulnerabilities Asset API call
func buildVulnerabilityAssetsQueryParameter(_ context.Context, d *plugin.QueryData) url.Values {
	queryParameter := make(url.Values)

	for columnName, qual := range d.Quals {
		if qual != nil {
			operator := ""
			var val interface{}
			for _, qu := range qual.Quals {
				operator = qu.Operator
				val = qu.Value.GetStringValue()
			}
			switch columnName {
			case "asset_type":
				if operator == "=" {
					queryParameter["asset_type"] = []string{fmt.Sprint(val)}
				}
			case "life_cycle":
				if operator == "=" {
					queryParameter["life_cycle"] = []string{fmt.Sprint(val)}
				}
			case "severities":
				if operator == "=" {
					queryParameter["severities"] = []string{fmt.Sprint(val)}
				}
			}
		}
	}

	return queryParameter
}

// Connection key quals
// if the caching is required other than per connection, build a cache key for the call and use it in Memoize
// since getCurrentUserProfile is a call, caching should be per connection
var getCurrentUserProfileMemoized = plugin.HydrateFunc(getCurrentUserProfileUncached).Memoize(memoize.WithCacheKeyFunction(getCurrentUserProfileCacheKey))

// Build a cache key for the call to getCurrentUserProfile.
func getCurrentUserProfileCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "getPrismacloudCurrentUserProfile"
	return key, nil
}

func getCurrentUserEmail(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (any, error) {
	userInfo, err := getCurrentUserProfileMemoized(ctx, d, h)
	if err != nil {
		return nil, err
	}

	userDetails := userInfo.(*model.UserProfile)

	return userDetails.Email, nil
}

func getCurrentUserProfileUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getCurrentUserProfileUncached", "connection_error", err)
		return nil, err
	}

	profile, err := api.GetCurrentUserProfile(conn)
	if err != nil {
		plugin.Logger(ctx).Error("getCurrentUserProfileUncached", "api_error", err)
		return nil, err
	}

	return profile, nil
}
