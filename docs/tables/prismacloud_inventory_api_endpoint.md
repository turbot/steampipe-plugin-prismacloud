---
title: "Steampipe Table: prismacloud_inventory_api_endpoint - Query Prisma Cloud API endpoints using SQL"
description: "Allows users to query and analyze Prisma Cloud API endpoints. This table provides detailed information about API endpoints, including account details, API paths, methods, and risk factors. It can be used to monitor and manage API endpoints within Prisma Cloud."
---

# Table: prismacloud_inventory_api_endpoint - Query Prisma Cloud API endpoints using SQL

The Prisma Cloud API endpoint table in Steampipe provides you with comprehensive information about API endpoints within Prisma Cloud. This table allows you, as a security engineer or cloud administrator, to query API endpoint-specific details, including account information, API paths, HTTP methods, and risk factors. You can utilize this table to gather insights on API endpoints, such as discovery methods, inspection types, and usage statistics. The schema outlines the various attributes of the Prisma Cloud API endpoints for you.

## Table Usage Guide

The `prismacloud_inventory_api_endpoint` table in Steampipe provides detailed information about API endpoints within Prisma Cloud. This table allows you to query details such as account ID, API paths, HTTP methods, and risk factors, enabling you to manage and monitor your API endpoints effectively.

## Examples

### Basic info
Retrieve basic information about Prisma Cloud API endpoints, such as account ID, account name, and API path.

```sql+postgres
select
  account_id,
  account_name,
  api_path
from
  prismacloud_inventory_api_endpoint;
```

```sql+sqlite
select
  account_id,
  account_name,
  api_path
from
  prismacloud_inventory_api_endpoint;
```

### Get API endpoints HTTP methods and hits
Get the count of hits for each API endpoint by HTTP method. This is useful for understanding the usage patterns of your API endpoints.

```sql+postgres
select
  api_path,
  http_method,
  hits
from
  prismacloud_inventory_api_endpoint;
```

```sql+sqlite
select
  api_path,
  http_method,
  hits
from
  prismacloud_inventory_api_endpoint;
```

### Get API endpoints risk factors
Retrieve the risk factors associated with each API endpoint. This helps in assessing the security posture of your API endpoints.

```sql+postgres
select
  api_path,
  path_risk_factors ->> 'internetExposed' as internet_exposed,
  path_risk_factors -> 'owaspAPIAttacks' as owasp_api_attacks,
  path_risk_factors -> 'requestSensitiveData' as request_sensitive_data,
  path_risk_factors ->> 'requiresAuthentication' as requires_authentication,
  path_risk_factors -> 'responseSensitiveData' as response_sensitive_data
from
  prismacloud_inventory_api_endpoint;
```

```sql+sqlite
select
  api_path,
  json_extract(path_risk_factors, '$.internetExposed') as internet_exposed,
  json_extract(path_risk_factors, '$.owaspAPIAttacks') as owasp_api_attacks,
  json_extract(path_risk_factors, '$.requestSensitiveData') as request_sensitive_data,
  json_extract(path_risk_factors, '$.requiresAuthentication') as requires_authentication,
  json_extract(path_risk_factors, '$.responseSensitiveData') as response_sensitive_data
from
  prismacloud_inventory_api_endpoint;
```

### Discovery methods and inspection types
Identify the discovery methods and inspection types for each API endpoint. This helps in understanding how the endpoints were discovered and inspected.

```sql+postgres
select
  api_path,
  discovery_method,
  inspection_type
from
  prismacloud_inventory_api_endpoint;
```

```sql+sqlite
select
  api_path,
  discovery_method,
  inspection_type
from
  prismacloud_inventory_api_endpoint;
```

### Get API endpoint recent changes and observations
Get the timestamp of when each API endpoint was last changed and last observed. This helps in tracking the activity and updates of your API endpoints.

```sql+postgres
select
  api_path,
  last_changed,
  current_date - last_changed::date as days_since_last_change,
  last_observed,
  current_date - last_observed::date as days_since_last_observed
from
  prismacloud_inventory_api_endpoint;
```

```sql+sqlite
select
  api_path,
  last_changed,
  julianday('now') - julianday(last_changed) as days_since_last_change,
  last_observed,
  julianday('now') - julianday(last_observed) as days_since_last_observed
from
  prismacloud_inventory_api_endpoint;
```