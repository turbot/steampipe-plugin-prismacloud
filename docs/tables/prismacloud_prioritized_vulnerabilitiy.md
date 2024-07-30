---
title: "Steampipe Table: prismacloud_prioritized_vulnerabilitiy - Query Prisma Cloud Prioritized Vulnerabilities using SQL"
description: "Allows users to query Prisma Cloud prioritized vulnerabilities. This table provides information about vulnerabilities in various asset types, focusing on exploitable, internet-exposed, and patchable vulnerabilities, among others."
---

# Table: prismacloud_prioritized_vulnerabilitiy - Query Prisma Cloud Prioritized Vulnerabilities using SQL

The Prisma Cloud prioritized vulnerabilities table in Steampipe provides detailed information about vulnerabilities in different asset types. This table allows security engineers and cloud administrators to query vulnerabilities based on their priority, including exploitable, internet-exposed, and patchable vulnerabilities. The schema outlines various attributes, such as asset type, life cycle stage, and the counts of different types of vulnerabilities.

## Table Usage Guide

The `prismacloud_prioritized_vulnerabilitiy` table in Steampipe provides information about prioritized vulnerabilities in various asset types within Prisma Cloud. This table allows you to query details such as the number of exploitable, internet-exposed, and patchable vulnerabilities, helping you to manage and monitor the vulnerability status of your cloud resources effectively.

**Important Notes**
- To query this table you need `vulnerabilityDashboard` feature with `View` permission to access this endpoint. Verify if your permission group includes this feature using the Get [Permission Group by ID](https://pan.dev/prisma-cloud/api/cspm/get-1/) endpoint. You can also check this in the Prisma Cloud console by ensuring that **Dashboard > Vulnerability** is enabled.
- You **_must_** specify `asset_type`, and `life_cycle` in `where` clause in order to use this table.

## Examples

### Basic info
Retrieve basic information about the prioritized vulnerabilities, including the total number of vulnerabilities and the number of urgent vulnerabilities.

```sql+postgres
select
  asset_type,
  total_vulnerabilities,
  urgent_vulnerability_count
from
  prismacloud_prioritized_vulnerabilitiy
where
  asset_type = 'host'
  and life_cycle = 'run';
```

```sql+sqlite
select
  asset_type,
  total_vulnerabilities,
  urgent_vulnerability_count
from
  prismacloud_prioritized_vulnerabilitiy
where
  asset_type = 'host'
  and life_cycle = 'run';
```

### Vulnerabilities by asset type
Retrieve the prioritized vulnerabilities grouped by asset type. This helps in understanding the distribution of vulnerabilities across different asset types.

```sql+postgres
select
  asset_type,
  sum(total_vulnerabilities) as total_vulnerabilities,
  sum(urgent_vulnerability_count) as urgent_vulnerabilities,
  sum(exploitable_vulnerability_count) as exploitable_vulnerabilities
from
  prismacloud_prioritized_vulnerabilitiy
where
  asset_type = 'host'
  and life_cycle = 'run'
group by
  asset_type;
```

```sql+sqlite
select
  asset_type,
  sum(total_vulnerabilities) as total_vulnerabilities,
  sum(urgent_vulnerability_count) as urgent_vulnerabilities,
  sum(exploitable_vulnerability_count) as exploitable_vulnerabilities
from
  prismacloud_prioritized_vulnerabilitiy
where
  asset_type = 'host'
  and life_cycle = 'run'
group by
  asset_type;
```

### Recently updated vulnerabilities
Retrieve the prioritized vulnerabilities data that were updated within the last 30 days. This helps in tracking the recent updates in vulnerability data.

```sql+postgres
select
  asset_type,
  last_updated_date_time,
  total_vulnerabilities,
  urgent_vulnerability_count
from
  prismacloud_prioritized_vulnerabilitiy
where
  asset_type = 'host'
  and life_cycle = 'run'
  and last_updated_date_time > now() - interval '30 days';
```

```sql+sqlite
select
  asset_type,
  last_updated_date_time,
  total_vulnerabilities,
  urgent_vulnerability_count
from
  prismacloud_prioritized_vulnerabilitiy
where
  asset_type = 'host'
  and life_cycle = 'run'
  and last_updated_date_time > datetime('now', '-30 days');
```

### Vulnerabilities by life cycle stage
Retrieve the prioritized vulnerabilities data grouped by the life cycle stage of the asset. This helps in understanding the trends in vulnerability management at different stages of the asset's life cycle.

```sql+postgres
select
  asset_type,
  life_cycle,
  sum(total_vulnerabilities) as total_vulnerabilities,
  sum(urgent_vulnerability_count) as urgent_vulnerabilities,
  sum(exploitable_vulnerability_count) as exploitable_vulnerabilities
from
  prismacloud_prioritized_vulnerabilitiy
where
  asset_type = 'host'
  and life_cycle = 'run'
group by
  asset_type,
  life_cycle;
```

```sql+sqlite
select
  asset_type,
  life_cycle,
  sum(total_vulnerabilities) as total_vulnerabilities,
  sum(urgent_vulnerability_count) as urgent_vulnerabilities,
  sum(exploitable_vulnerability_count) as exploitable_vulnerabilities
from
  prismacloud_prioritized_vulnerabilitiy
where
  asset_type = 'host'
  and life_cycle = 'run'
group by
  asset_type,
  life_cycle;
```

### Vulnerabilities by severity and type
Retrieve the prioritized vulnerabilities data grouped by severity and type, such as exploitable and internet-exposed vulnerabilities. This helps in understanding the distribution of vulnerabilities by their severity and type.

```sql+postgres
select
  asset_type,
  sum(exploitable_vulnerability_count) as exploitable_vulnerabilities,
  sum(internet_exposed_vulnerability_count) as internet_exposed_vulnerabilities,
  sum(patchable_vulnerability_count) as patchable_vulnerabilities
from
  prismacloud_prioritized_vulnerabilitiy
where
  asset_type = 'host'
  and life_cycle = 'run'
group by
  asset_type;
```

```sql+sqlite
select
  asset_type,
  sum(exploitable_vulnerability_count) as exploitable_vulnerabilities,
  sum(internet_exposed_vulnerability_count) as internet_exposed_vulnerabilities,
  sum(patchable_vulnerability_count) as patchable_vulnerabilities
from
  prismacloud_prioritized_vulnerabilitiy
where
  asset_type = 'host'
  and life_cycle = 'run'
group by
  asset_type;
```