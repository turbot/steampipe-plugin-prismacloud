---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/prismacloud.svg"
brand_color: "#00BFE7"
display_name: "Prisma Cloud"
short_name: "prismacloud"
description: "Steampipe plugin for querying Prisma Cloud accounts, alerts, users, and other resources."
og_description: "Use SQL to query accounts, alerts, users, reports and more from Prisma Cloud. Open source CLI. No DB required."
og_image: "/images/plugins/turbot/prismacloud-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# Prisma Cloud + Steampipe

[Steampipe](https://steampipe.io) is an open-source platform that allows you to query cloud APIs using SQL.

[Prisma Cloud](https://www.paloaltonetworks.com/prismacloud/cloud) is a comprehensive cloud security solution by Palo Alto Networks that provides protection across the entire cloud-native technology stack.

For example:

```sql
select
  name,
  cloud_type,
  compliance_standard_id,
  status
from
  prismacloud_report;
```

```
+--------------------------+-----------+-------------------------+----------+
| name                     | cloud_type| compliance_standard_id  | status   |
+--------------------------+-----------+-------------------------+----------+
| Compliance Report        | AWS       | CIS                     | ACTIVE   |
| Vulnerability Report     | Azure     | SOC2                    | ACTIVE   |
| Configuration Audit      | GCP       | HIPAA                   | INACTIVE |
+--------------------------+-----------+-------------------------+----------+
```

## Documentation

- **[Table definitions & examples →](https://hub.steampipe.io/plugins/turbot/prismacloud/tables)**

## Get started

### Install

Download and install the latest Prisma plugin:

```bash
steampipe plugin install prismacloud
```

### Credentials

| Item        | Description                                                                                                                                                                                                                                                                                                                            |
| ----------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Credentials | The Prisma plugin uses a URL and either username/password or a JSON Web Token (JWT) to authenticate to the Prisma APIs.                                                                                                                                                                                                                |
| Permissions | You must create a [Prisma Cloud account](https://docs.paloaltonetworks.com/prismacloud/prismacloud-cloud/prismacloud-cloud-admin/get-started-with-prismacloud-cloud-access/get-started-with-prismacloud-cloud-identity-and-access-management/manage-access-to-prismacloud-cloud.html) with the necessary permissions to query the API. |
| Radius      | The Prisma plugin query scope is generally the same as the Prisma API. You can list resources and details that you have access to within your Prisma Cloud account.                                                                                                                                                                    |
| Resolution  | Credentials in the Steampipe configuration file (`~/.steampipe/config/prismacloud.spc`)                                                                                                                                                                                                                                                |

### Configuration

Installing the latest prismacloud plugin will create a config file (`~/.steampipe/config/prismacloud.spc`) with a single connection named `prismacloud`:

```hcl
connection "prismacloud" {
  plugin = "prismacloud"

  # Required: URL of the Prisma Cloud instance excluding the protocol.
  # https://pan.dev/prismacloud-cloud/api/cspm/api-urls/
  # url = "api.anz.prismacloud.io"

  # Using username, password authentication
  # Username for authentication.
  # username = "87ef938r-e89c-2ff9-9834-8936d88333s8"

  # Password for authentication.
  # password = "JU+HJS8SDMsCk6yjRqd5cHhsj4k="

  # Using JSON Web Token
  # JSON Web Token for authentication.
  # token = "eyJhbGciOiJIUzI1NiJ9.eyJhY2Nlc3NLZXlJZCI6IjA4Y...H6BNc_Xonw"

  # Customer name for the Prisma Cloud account.
  # customer_name = "My Name - 123236897770856499123"

  # Protocol to be used (http or https).
  # protocol = "https"

  # Port to connect to Prisma Cloud API.
  # port = 443

  # Timeout for API requests in seconds.
  # timeout = 30

  # Skip SSL certificate verification (true/false).
  # skip_ssl_cert_verification = false

  # Logging settings.
  # logging = {
  #  # Enable or disable logging for specific components.
  #  "LogAction"  = true
  # }

  # Disable automatic reconnection (true/false).
  # disable_reconnect = false

  # Maximum number of retries for API requests.
  # max_retries = 9

  # Maximum delay between retries in milliseconds.
  # retry_max_delay = 5000

  # Number of retries for API requests.
  # retries = 3
}
```

- `url` - The URL of the Prisma Cloud instance excluding the protocol (e.g., `api.anz.prismacloudcloud.io`).
- `username` - The username for authentication to the Prisma Cloud API.
- `password` - The password for authentication to the Prisma Cloud API.
- `token` - The JSON Web Token (JWT) for authentication to the Prisma Cloud API.
- `customer_name` - The customer name for the Prisma Cloud account.
- `protocol` - The protocol to be used (http or https).
- `port` - The port to connect to Prisma Cloud API.
- `timeout` - The timeout for API requests in seconds.
- `skip_ssl_cert_verification` - Whether to skip SSL certificate verification.
- `logging` - The logging settings.
- `disable_reconnect` - Whether to disable automatic reconnection.
- `max_retries` - The maximum number of retries for API requests.
- `retry_max_delay` - The maximum delay between retries in milliseconds.
- `retries` - The number of retries for API requests.
