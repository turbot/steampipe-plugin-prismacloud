package prismacloud

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type prismacloudCloudConfig struct {
	Url                     *string         `hcl:"url"`
	Username                *string         `hcl:"username,optional"`
	Password                *string         `hcl:"password,optional"`
	CustomerName            *string         `hcl:"customer_name,optional"`
	Protocol                *string         `hcl:"protocol,optional"`
	Port                    *int32          `hcl:"port,optional"`
	Timeout                 *int32          `hcl:"timeout,optional"`
	SkipSslCertVerification *bool           `hcl:"skip_ssl_cert_verification,optional"`
	Logging                 map[string]bool `hcl:"logging,optional"`
	DisableReconnect        *bool           `hcl:"disable_reconnect,optional"`
	MaxRetries              *int            `hcl:"max_retries,optional"`
	RetryMaxDelay           *int            `hcl:"retry_max_delay,optional"`
	Retries                 *int            `hcl:"retries,optional"`
	Token                   *string         `hcl:"token,optional"`
}

func ConfigInstance() interface{} {
	return &prismacloudCloudConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) prismacloudCloudConfig {
	if connection == nil || connection.Config == nil {
		return prismacloudCloudConfig{}
	}
	config, _ := connection.Config.(prismacloudCloudConfig)
	return config
}
