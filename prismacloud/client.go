package prismacloud

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"

	prismacloud "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func connect(_ context.Context, d *plugin.QueryData) (*prismacloud.Client, error) {
	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "prismacloud"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*prismacloud.Client), nil
	}

	c := prismacloud.Client{}

	prismacloudConfig := GetConfig(d.Connection)

	// Return error if the minimum credential is not provided
	if (prismacloudConfig.Username == nil || prismacloudConfig.Password == nil) && (prismacloudConfig.Token == nil) {
		return nil, fmt.Errorf("'username', 'password' and 'customer_name' or 'token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	// Logging
	if len(c.Logging) == 0 {
		if len(prismacloudConfig.Logging) > 0 {
			c.Logging = make(map[string]bool)
			for key, val := range prismacloudConfig.Logging {
				c.Logging[key] = val
			}
		} else {
			c.Logging = map[string]bool{"LogAction": true} // Default to Logging action "true"
		}
	}

	//// Request timeout
	if c.Timeout == 0 {
		if prismacloudConfig.Timeout != nil {
			c.Timeout = int(*prismacloudConfig.Timeout)
		} else {
			// Default to 180 second
			c.Timeout = 180
		}
	}
	if c.Timeout < 0 {
		return nil, fmt.Errorf("invalid timeout")
	}

	//// Port
	if c.Port == 0 {
		if prismacloudConfig.Port != nil {
			c.Port = int(*prismacloudConfig.Port)
		}
	}
	if c.Port > 65535 || c.Port < 0 {
		return nil, fmt.Errorf("invalid port number")
	}

	//// Protocol
	if c.Protocol == "" {
		if prismacloudConfig.Protocol != nil {
			c.Protocol = *prismacloudConfig.Protocol
		} else {
			// Default to https
			c.Protocol = "https"
		}
	}
	if c.Protocol != "http" && c.Protocol != "https" {
		return nil, fmt.Errorf("invalid protocol")
	}

	//// URL
	if c.Url == "" && prismacloudConfig.Url != nil {
		c.Url = *prismacloudConfig.Url
	}
	if strings.HasPrefix(c.Url, "http://") || strings.HasPrefix(c.Url, "https://") {
		return nil, fmt.Errorf("specify protocol using the Protocol param, not as the URL")
	}
	c.Url = strings.TrimRight(c.Url, "/")
	if c.Url == "" {
		return nil, fmt.Errorf("prismacloud Cloud URL is not set")
	}

	//// Username/Password/Customer Name
	if c.Username == "" && prismacloudConfig.Username != nil {
		c.Username = *prismacloudConfig.Username
	}

	if c.Password == "" && prismacloudConfig.Password != nil {
		c.Password = *prismacloudConfig.Password
	}

	if c.CustomerName == "" && prismacloudConfig.CustomerName != nil {
		c.CustomerName = *prismacloudConfig.CustomerName
	}

	//// SSL Certificate Verification
	skipSslCertVerification := prismacloudConfig.SkipSslCertVerification
	if c.Transport == nil {
		c.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: c.SkipSslCertVerification,
			},
			Proxy: http.ProxyFromEnvironment,
		}
	}
	if skipSslCertVerification != nil {
		c.Transport.TLSClientConfig.InsecureSkipVerify = *skipSslCertVerification
	}

	//// JWT token
	if c.JsonWebToken == "" && prismacloudConfig.Token != nil {
		c.JsonWebToken = *prismacloudConfig.Token
	}

	//// Disable Reconnect
	if prismacloudConfig.DisableReconnect != nil {
		c.DisableReconnect = *prismacloudConfig.DisableReconnect
	}

	//// Maximum number of retries
	if prismacloudConfig.MaxRetries != nil {
		c.MaxRetries = *prismacloudConfig.MaxRetries
	} else {
		c.MaxRetries = 9 // Default to 9
	}

	//// Number of retries for API requests.
	if prismacloudConfig.Retries != nil {
		c.Retries = *prismacloudConfig.Retries
	} else {
		c.Retries = 3 // Default to 3
	}

	if prismacloudConfig.RetryMaxDelay != nil {
		c.RetryMaxDelay = *prismacloudConfig.RetryMaxDelay
	} else {
		c.RetryMaxDelay = 5000 // Default to 5000 milliseconds
	}

	err := c.Initialize("")
	if err != nil {
		return nil, fmt.Errorf("error in initialize client: %v+", err)
	}

	return &c, nil
}
