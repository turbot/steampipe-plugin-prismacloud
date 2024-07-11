package prisma

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"

	prisma "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func connect(_ context.Context, d *plugin.QueryData) (*prisma.Client, error) {
	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "prisma"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*prisma.Client), nil
	}

	c := prisma.Client{}

	prismaConfig := GetConfig(d.Connection)

	// Return error if the minimum credential is not provided
	if (prismaConfig.Username == nil || prismaConfig.Password == nil ) && (prismaConfig.Token == nil) {
		return nil, fmt.Errorf("'username', 'password' and 'customer_name' or 'token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	// Logging
	if len(c.Logging) == 0 {
		if len(prismaConfig.Logging) > 0 {
			c.Logging = make(map[string]bool)
			for key, val := range prismaConfig.Logging {
				c.Logging[key] = val
			}
		} else {
			c.Logging = map[string]bool{"LogAction": true} // Default to Logging action "true"
		}
	}

	//// Request timeout
	if c.Timeout == 0 {
		if prismaConfig.Timeout != nil {
			c.Timeout = int(*prismaConfig.Timeout)
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
		if prismaConfig.Port != nil {
			c.Port = int(*prismaConfig.Port)
		}
	}
	if c.Port > 65535 || c.Port < 0 {
		return nil, fmt.Errorf("invalid port number")
	}

	//// Protocol
	if c.Protocol == "" {
		if prismaConfig.Protocol != nil {
			c.Protocol = *prismaConfig.Protocol
		} else {
			// Default to https
			c.Protocol = "https"
		}
	}
	if c.Protocol != "http" && c.Protocol != "https" {
		return nil, fmt.Errorf("invalid protocol")
	}

	//// URL
	if c.Url == "" && prismaConfig.Url != nil {
		c.Url = *prismaConfig.Url
	}
	if strings.HasPrefix(c.Url, "http://") || strings.HasPrefix(c.Url, "https://") {
		return nil, fmt.Errorf("specify protocol using the Protocol param, not as the URL")
	}
	c.Url = strings.TrimRight(c.Url, "/")
	if c.Url == "" {
		return nil, fmt.Errorf("prisma Cloud URL is not set")
	}

	//// Username/Password/Customer Name
	if c.Username == "" && prismaConfig.Username != nil {
		c.Username = *prismaConfig.Username
	}

	if c.Password == "" && prismaConfig.Password != nil {
		c.Password = *prismaConfig.Password
	}

	if c.CustomerName == "" && prismaConfig.CustomerName != nil {
		c.CustomerName = *prismaConfig.CustomerName
	}

	//// SSL Certificate Verification
	skipSslCertVerification := prismaConfig.SkipSslCertVerification
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
	if c.JsonWebToken == "" && prismaConfig.Token != nil {
		c.JsonWebToken = *prismaConfig.Token
	}

	//// Disable Reconnect
	if prismaConfig.DisableReconnect != nil {
		c.DisableReconnect = *prismaConfig.DisableReconnect
	}

	//// Maximum number of retries
	if prismaConfig.MaxRetries != nil {
		c.MaxRetries = *prismaConfig.MaxRetries
	} else {
		c.MaxRetries = 9 // Default to 9
	}

	//// Number of retries for API requests.
	if prismaConfig.Retries != nil {
		c.Retries = *prismaConfig.Retries
	} else {
		c.Retries = 3 // Default to 3
	}

	if prismaConfig.RetryMaxDelay != nil {
		c.RetryMaxDelay = *prismaConfig.RetryMaxDelay
	} else {
		c.RetryMaxDelay = 5000 // Default to 5000 milliseconds
	}

	err := c.Initialize("")
	if err != nil {
		return nil, fmt.Errorf("error in initialize client: %v+", err)
	}

	return &c, nil
}
