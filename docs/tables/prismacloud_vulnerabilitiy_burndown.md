---
title: "Steampipe Table: prismacloud_vulnerabilitiy_burndown - Query Prisma Cloud Vulnerability Burndown using SQL"
description: "Allows users to query Prisma Cloud vulnerability burndown. This table provides a historical view of vulnerabilities and their remediation over time for different asset types."
---

# Table: prismacloud_vulnerabilitiy_burndown - Query Prisma Cloud Vulnerability Burndown using SQL

The Prisma Cloud vulnerability burndown table in Steampipe provides a historical view of vulnerabilities and their remediation over time for various asset types. This table helps security engineers and cloud administrators to track the progress of vulnerability remediation and understand the trends in vulnerability management. The schema outlines various attributes related to the vulnerability burndown, including the type of asset, life cycle stage, severity, and the number of vulnerabilities recorded and remediated each day.

## Table Usage Guide

The `prismacloud_vulnerabilitiy_burndown` table in Steampipe provides information about the historical trends of vulnerabilities in different asset types within Prisma Cloud. This table allows you to query details such as the number of vulnerabilities recorded and remediated each day, helping you to track and manage the progress of vulnerability remediation effectively.

**Important Notes**
- To query this table you need `vulnerabilityDashboard` feature with `View` permission to access this endpoint. Verify if your permission group includes this feature using the Get [Permission Group by ID](https://pan.dev/prisma-cloud/api/cspm/get-1/) endpoint. You can also check this in the Prisma Cloud console by ensuring that **Dashboard > Vulnerability** is enabled.
- You **_must_** specify `asset_type`, `life_cycle`, and `severities` in `where` clause in order to use this table.

## Examples

### Basic info
Retrieve basic information about the vulnerability burndown, including the total number of vulnerabilities recorded each day and the number of vulnerabilities remediated each day.

```sql+postgres
select
  asset_type,
  day_num,
  epoch_timestamp,
  total_count,
  remediated_count
from
  prismacloud_vulnerabilitiy_burndown;
```

```sql+sqlite
select
  asset_type,
  day_num,
  epoch_timestamp,
  total_count,
  remediated_count
from
  prismacloud_vulnerabilitiy_burndown;
```

### Vulnerability burndown by asset type
Retrieve the burndown of vulnerabilities grouped by asset type. This helps in understanding the trends in vulnerability management for different types of assets.

```sql+postgres
select
  asset_type,
  sum(total_count) as total_vulnerabilities,
  sum(remediated_count) as total_remediated
from
  prismacloud_vulnerabilitiy_burndown
where
  asset_type = 'host'
  and life_cycle = 'run'
  and severities = 'critical'
group by
  asset_type;
```

```sql+sqlite
select
  asset_type,
  sum(total_count) as total_vulnerabilities,
  sum(remediated_count) as total_remediated
from
  prismacloud_vulnerabilitiy_burndown
where
  asset_type = 'host'
  and life_cycle = 'run'
  and severities = 'critical'
group by
  asset_type;
```

### Recent vulnerability burndown
Retrieve the vulnerability burndown data for the last 30 days. This helps in tracking the recent progress of vulnerability remediation.

```sql+postgres
select
  asset_type,
  day_num,
  epoch_timestamp,
  total_count,
  remediated_count
from
  prismacloud_vulnerabilitiy_burndown
where
  asset_type = 'host'
  and life_cycle = 'run'
  and severities = 'critical'
  and epoch_timestamp > now() - interval '30 days';
```

```sql+sqlite
select
  asset_type,
  day_num,
  epoch_timestamp,
  total_count,
  remediated_count
from
  prismacloud_vulnerabilitiy_burndown
where
  asset_type = 'host'
  and life_cycle = 'run'
  and severities = 'critical'
  and epoch_timestamp > now() - interval '30 days';
```

### Vulnerability burndown for critical severity
Retrieve the vulnerability burndown data grouped by severity levels. This helps in understanding the distribution of vulnerabilities by severity and tracking their remediation.

```sql+postgres
select
  asset_type,
  severities,
  sum(total_count) as total_vulnerabilities,
  sum(remediated_count) as total_remediated
from
  prismacloud_vulnerabilitiy_burndown
where
  asset_type = 'host'
  and life_cycle = 'run'
  and severities = 'critical'
group by
  asset_type,
  severities;
```

```sql+sqlite
select
  asset_type,
  severities,
  sum(total_count) as total_vulnerabilities,
  sum(remediated_count) as total_remediated
from
  prismacloud_vulnerabilitiy_burndown
where
  asset_type = 'host'
  and life_cycle = 'run'
  and severities = 'critical'
group by
  asset_type,
  severities;
```

### Vulnerability burndown life cycle stage run
Retrieve the vulnerability burndown data grouped by the life cycle stage of the asset. This helps in understanding the trends in vulnerability management at different stages of the asset's life cycle.

```sql+postgres
select
  asset_type,
  life_cycle,
  sum(total_count) as total_vulnerabilities,
  sum(remediated_count) as total_remediated
from
  prismacloud_vulnerabilitiy_burndown
where
  asset_type = 'host'
  and life_cycle = 'run'
  and severities = 'critical'
group by
  asset_type,
  life_cycle;
```

```sql+sqlite
select
  asset_type,
  life_cycle,
  sum(total_count) as total_vulnerabilities,
  sum(remediated_count) as total_remediated
from
  prismacloud_vulnerabilitiy_burndown
where
  asset_type = 'host'
  and life_cycle = 'run'
  and severities = 'critical'
group by
  asset_type,
  life_cycle;
```