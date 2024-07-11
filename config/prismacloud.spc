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
  # token = "eyJhbGciOiJIUzI1NiJ9.eyJhY2Nlc3NLZXlJZCI6IjA4YWQzOTNmLWU1OWMtNGFhNS05NzA5LTA4MzZkNzczZWUyZCIsInN1YiI6InNhY2hpbi5rdW1hcjEzNUB3aXByby5jb20iLCJmaXJzdExvZ2luIjpmYWxzZSwicHJpc21hSWQiOiI4MDY4NDEwNDI2NjYzMTM3MjgiLCJpcEFkZHJlc3MiOiIxMjIuMTY0Ljg0LjE1NiIsImlzcyI6Imh0dHBzOi8vYXBpLmFuei5wcmlzbWFjbG91ZC5pbyIsInJlc3RyaWN0IjowLCJpc0FjY2Vzc0tleUxvZ2luIjp0cnVlLCJ1c2VyUm9sZVR5cGVEZXRhaWxzIjp7Imhhc09ubHlSZWFkQWNjZXNzIjpmYWxzZX0sInVzZXJSb2xlVHlwZU5hbWUiOiJTeXN0ZW0gQWRtaW4iLCJpc1NTT1Nlc3Npb24iOmZhbHNlLCJsYXN0TG9naW5UaW1lIjoxNzE5OTAyOTc3OTQwLCJhdWQiOiJodHRwczovL2FwaS5hbnoucHJpc21hY2xvdWQuaW8iLCJ1c2VyUm9sZVR5cGVJZCI6MSwiYXV0aC1tZXRob2QiOiJQQVNTV09SRCIsInNlbGVjdGVkQ3VzdG9tZXJOYW1lIjoiV2lwcm8gTGltaXRlZCAoSW5kaWEpIC0gNjQzODI2ODk2NzcwODU2NDkwMyIsInNlc3Npb25UaW1lb3V0Ijo2MCwidXNlclJvbGVJZCI6IjQwNjliNTQyLWUwNGMtNDFiMy05NDgyLThiYzE1NzRlZWRiYyIsImhhc0RlZmVuZGVyUGVybWlzc2lvbnMiOnRydWUsImV4cCI6MTcyMDYwNjUzNSwiaWF0IjoxNzIwNjA1OTM1LCJ1c2VybmFtZSI6InNhY2hpbi5rdW1hcjEzNUB3aXByby5jb20iLCJ1c2VyUm9sZU5hbWUiOiJTeXN0ZW0gQWRtaW4ifQ.8BrD4XFkY2zZhSy7Syjx9Y4lWcOyDrMqaH6BNc_Xonw"

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
