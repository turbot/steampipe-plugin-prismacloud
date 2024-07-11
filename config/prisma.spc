connection "prisma" {
  plugin = "prisma"

  # URL of the Prisma Cloud instance.
  url = "api.prismacloud.io"

  # Username for authentication.
  username = "your_username"

  # Password for authentication.
  password = "your_password"

  # Customer name for the Prisma Cloud account.
  customer_name = "your_customer_name"

  # Protocol to be used (http or https).
  protocol = "https"

  # Port to connect to Prisma Cloud API.
  port = 443

  # Timeout for API requests in seconds.
  timeout = 30

  # Skip SSL certificate verification (true/false).
  skip_ssl_cert_verification = false

  # Logging settings.
  # logging = {
  #  # Enable or disable logging for specific components.
  #  "request"  = true
  #  "response" = false
  # }

  # Disable automatic reconnection (true/false).
  disable_reconnect = false

  # Maximum number of retries for API requests.
  max_retries = 5

  # Maximum delay between retries in milliseconds.
  retry_max_delay = 5000

  # Number of retries for API requests.
  retries = 3

  # JSON Web Token for authentication.
  json_web_token = "your_jwt"
}
