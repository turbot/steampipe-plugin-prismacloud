connection "prismacloud" {
  plugin = "prismacloud"

  # Required: URL of the Prisma Cloud instance exclusing the protocol.
  # https://pan.dev/prismacloud-cloud/api/cspm/api-urls/
  # url = "api.anz.prismacloud.io"

  # Using username, password authentication
  # Username for authentication.
  # username = "87ef938r-e89c-2ff9-9834-8936d88333s8"

  # Password for authentication.
  # password = "JU+HJS8SDMsCk6yjRqd5cHhsj4k="

  # Using JSON Web Token
  # JSON Web Token for authentication.
  # token = "eyJhbGciOiJIUzI1NiJ9.eyJhY2Nlc3NLZXlJZCI6IjA4YWQzOTNmL...H6BNc_Xonw"

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

  # # Logging settings.
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
