---
title: "Steampipe Table: prismacloud_vulnerabilitiy_overview - Query Prisma Cloud Vulnerability Overview using SQL"
description: "Allows users to query Prisma Cloud vulnerability overview. This table provides an overview of vulnerabilities in runtime assets, including the total number of vulnerabilities, remediated vulnerabilities, and more."
---

# Table: prismacloud_vulnerabilitiy_overview - Query Prisma Cloud Vulnerability Overview using SQL

The Prisma Cloud vulnerability overview table in Steampipe provides a comprehensive view of vulnerabilities in runtime assets. This table allows security engineers and cloud administrators to query the total number of vulnerabilities, the number of remediated vulnerabilities, and other key metrics. The schema outlines various attributes related to the vulnerability overview, helping users to monitor and manage vulnerabilities effectively.

## Table Usage Guide

The `prismacloud_vulnerabilitiy_overview` table in Steampipe provides information about the vulnerability overview in runtime assets within Prisma Cloud. This table allows you to query details such as the total number of vulnerabilities, remediated vulnerabilities, and more, enabling you to manage and monitor your cloud resources effectively.

**Important Notes**
- To query this table you need `vulnerabilityDashboard` feature with `View` permission to access this endpoint. Verify if your permission group includes this feature using the Get [Permission Group by ID](https://pan.dev/prisma-cloud/api/cspm/get-1/) endpoint. You can also check this in the Prisma Cloud console by ensuring that **Dashboard > Vulnerability** is enabled.

## Examples

### Basic info
Retrieve basic information about the vulnerability overview, including the total number of vulnerabilities in runtime assets.

```sql+postgres
select
  total_vulnerable_runtime_assets,
  total_vulnerabilitiesin_runtime,
  total_remediated_in_runtime
from
  prismacloud_vulnerabilitiy_overview;
```

```sql+sqlite
select
  total_vulnerable_runtime_assets,
  total_vulnerabilitiesin_runtime,
  total_remediated_in_runtime
from
  prismacloud_vulnerabilitiy_overview;
```

### Get vulnerability overview
Retrieve detailed information about vulnerabilities, including the breakdown of vulnerabilities by severity levels and remediated vulnerabilities.

```sql+postgres
select
  jsonb_pretty(total_vulnerable_runtime_assets) as total_vulnerable_runtime_assets,
  jsonb_pretty(total_vulnerabilitiesin_runtime) as total_vulnerabilitiesin_runtime,
  jsonb_pretty(total_remediated_in_runtime) as total_remediated_in_runtime,
  jsonb_pretty(values) as values
from
  prismacloud_vulnerabilitiy_overview;
```

```sql+sqlite
select
  json(total_vulnerable_runtime_assets) as total_vulnerable_runtime_assets,
  json(total_vulnerabilitiesin_runtime) as total_vulnerabilitiesin_runtime,
  json(total_remediated_in_runtime) as total_remediated_in_runtime,
  json(values) as values
from
  prismacloud_vulnerabilitiy_overview;
```

### Get runtime vulnerabilities details
Retrieve the total number of vulnerabilities grouped by severity levels, helping in understanding the distribution of vulnerabilities.

```sql+postgres
select
  total_vulnerabilitiesin_runtime ->> 'criticalCount' as critical_count,
  total_vulnerabilitiesin_runtime ->> 'highCount' as high_count,
  total_vulnerabilitiesin_runtime ->> 'mediumCount' as medium_count,
  total_vulnerabilitiesin_runtime ->> 'lowCount' as low_count
from
  prismacloud_vulnerabilitiy_overview;
```

```sql+sqlite
select
  json_extract(total_vulnerabilitiesin_runtime, '$.criticalCount') as critical_count,
  json_extract(total_vulnerabilitiesin_runtime, '$.highCount') as high_count,
  json_extract(total_vulnerabilitiesin_runtime, '$.mediumCount') as medium_count,
  json_extract(total_vulnerabilitiesin_runtime, '$.lowCount') as low_count
from
  prismacloud_vulnerabilitiy_overview;
```

### Get remediated vulnerabilities
Retrieve the total number of remediated vulnerabilities, helping in tracking remediation efforts.

```sql+postgres
select
  total_remediated_in_runtime ->> 'totalCount' as total_remediated
from
  prismacloud_vulnerabilitiy_overview;
```

```sql+sqlite
select
  json_extract(total_remediated_in_runtime, '$.totalCount') as total_remediated
from
  prismacloud_vulnerabilitiy_overview;
```