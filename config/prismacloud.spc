connection "prismacloud" {
  plugin = "prismacloud"

  # Required: URL of the Prisma Cloud instance exclusing the protocol.
  # https://pan.dev/prismacloud-cloud/api/cspm/api-urls/
  url = "api.anz.prismacloud.io"

  # 1. Using username, password authentication
  # Username for authentication.
  # username = "90ef393f-e59c-3ff3-8473-0836d883ee2d"

  # Password for authentication.
  # password = "JU+AXA3iDMsCk8SjRqd5cHoisYg="

  # 2. Using JSON Web Token
  # JSON Web Token for authentication.
  # token = "eyJhbGciOiJIUzI1NiJ9.eyJhY2Nlc3NLZXlJZCI6IjA4YWQzOTNmL...H6BNc_Xonw"

  # # Customer name for the Prisma Cloud account.
  # customer_name = "Wipro Limited (India) - 6438268967708564903"

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
