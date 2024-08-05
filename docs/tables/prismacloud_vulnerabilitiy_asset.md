---
title: "Steampipe Table: prismacloud_vulnerabilitiy_asset - Query Prisma Cloud Vulnerability Assets using SQL"
description: "Allows users to query Prisma Cloud vulnerability assets. This table provides detailed information about the types of assets, their life cycle stages, and their associated vulnerabilities."
---

# Table: prismacloud_vulnerabilitiy_asset - Query Prisma Cloud Vulnerability Assets using SQL

The Prisma Cloud vulnerability asset table in Steampipe provides detailed information about vulnerability assets within Prisma Cloud. This table allows security engineers and cloud administrators to query asset-specific details, including asset types, life cycle stages, and associated vulnerabilities. The schema outlines various attributes such as asset type, life cycle, severity, and statistics of the vulnerable assets.

## Table Usage Guide

The `prismacloud_vulnerabilitiy_asset` table in Steampipe provides detailed information about vulnerability assets within Prisma Cloud. This table allows you to query details such as asset types, life cycle stages, and associated vulnerabilities. This helps in managing and monitoring the security status of your cloud assets effectively.

**Important Notes**
- To query this table you need `vulnerabilityDashboard` feature with `View` permission to access this endpoint. Verify if your permission group includes this feature using the Get [Permission Group by ID](https://pan.dev/prisma-cloud/api/cspm/get-1/) endpoint. You can also check this in the Prisma Cloud console by ensuring that **Dashboard > Vulnerability** is enabled.
- For improved performance, it is recommended to use the optional qualifiers (quals) to limit the result set.
- Queries with optional qualifiers are optimized to use filters. The following columns support optional qualifiers:
  - `asset_type`
  - `life_cycle`
  - `severities`

## Examples

### Basic Info
Retrieve basic information about Prisma Cloud vulnerability assets, including asset type, life cycle, and total vulnerabilities.

```sql+postgres
select
  asset_type,
  life_cycle,
  total_vulnerabilities
from
  prismacloud_vulnerabilitiy_asset;
```

```sql+sqlite
select
  asset_type,
  life_cycle,
  total_vulnerabilities
from
  prismacloud_vulnerabilitiy_asset;
```

### List assets by severity
Retrieve the assets categorized by their severity levels. This helps in identifying assets with critical vulnerabilities that need immediate attention.

```sql+postgres
select
  asset_type,
  life_cycle,
  severities,
  total_assets,
  total_vulnerabilities
from
  prismacloud_vulnerabilitiy_asset
where
  severities = 'critical';
```

```sql+sqlite
select
  asset_type,
  life_cycle,
  severities,
  total_assets,
  total_vulnerabilities
from
  prismacloud_vulnerabilitiy_asset
where
  severities = 'critical';
```

### List assets by life cycle stage
Retrieve the assets based on their life cycle stages. This helps in understanding the distribution of vulnerabilities across different stages of the asset lifecycle.

```sql+postgres
select
  asset_type,
  life_cycle,
  total_assets,
  total_vulnerabilities
from
  prismacloud_vulnerabilitiy_asset
where
  life_cycle = 'run';
```

```sql+sqlite
select
  asset_type,
  life_cycle,
  total_assets,
  total_vulnerabilities
from
  prismacloud_vulnerabilitiy_asset
where
  life_cycle = 'run';
```

### Asset statistics
Retrieve detailed statistics of the vulnerable assets, including the total number of assets and vulnerabilities. This helps in gaining insights into the overall security posture of your assets.

```sql+postgres
select
  asset_type,
  stats,
  total_assets,
  total_vulnerabilities
from
  prismacloud_vulnerabilitiy_asset;
```

```sql+sqlite
select
  asset_type,
  stats,
  total_assets,
  total_vulnerabilities
from
  prismacloud_vulnerabilitiy_asset;
```

### Get statistics of vulnerability assets
This query retrieves detailed statistics of vulnerability assets, including the number of users, assets, packages, providers, registries, repositories, and vulnerabilities categorized by severity levels. It helps in understanding the distribution of vulnerabilities across different asset types and providers.

```sql+postgres
select
  asset_type,
  s ->> 'users' as users,
  s ->> 'assets' as assets,
  s ->> 'packages' as packages,
  s ->> 'provider' as provider,
  s ->> 'registries' as registries,
  s ->> 'repositories' as repositories,
  s -> 'vulnerabilities' ->> 'lowCount' as low_count,
  s -> 'vulnerabilities' ->> 'highCount' as highCount,
  s -> 'vulnerabilities' ->> 'mediumCount' as mediumCount,
  s -> 'vulnerabilities' ->> 'criticalCount' as criticalCount
from
  prismacloud_vulnerabilitiy_asset,
  jsonb_array_elements(stats) as s
order by
  provider;
```

```sql+sqlite
select
  asset_type,
  json_extract(s.value, '$.users') as users,
  json_extract(s.value, '$.assets') as assets,
  json_extract(s.value, '$.packages') as packages,
  json_extract(s.value, '$.provider') as provider,
  json_extract(s.value, '$.registries') as registries,
  json_extract(s.value, '$.repositories') as repositories,
  json_extract(s.value, '$.vulnerabilities.lowCount') as low_count,
  json_extract(s.value, '$.vulnerabilities.highCount') as high_count,
  json_extract(s.value, '$.vulnerabilities.mediumCount') as medium_count,
  json_extract(s.value, '$.vulnerabilities.criticalCount') as critical_count
from
  prismacloud_vulnerabilitiy_asset,
  json_each(stats) as s
order by
  provider;
```