package prismacloud

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// isNotFoundError:: function which returns an ErrorPredicate for Azure API calls
func isNotFoundError(notFoundErrors []string) plugin.ErrorPredicateWithContext {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {

		// Added to support in not found errors
		for _, pattern := range notFoundErrors {
			if strings.Contains(err.Error(), pattern) {
				return true
			}
		}
		return false
	}
}
