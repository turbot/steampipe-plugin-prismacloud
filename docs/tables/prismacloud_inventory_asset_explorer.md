---
title: "Steampipe Table: prismacloud_inventory_asset_explorer - Explore Prisma Cloud assets using SQL"
description: "Allows users to query and explore Prisma Cloud assets. This table provides detailed information about cloud assets, including account details, alert statuses, compliance standards, and vulnerability information. It can be used to monitor and manage cloud assets within Prisma Cloud."
---

# Table: prismacloud_inventory_asset_explorer - Explore Prisma Cloud assets using SQL

The Prisma Cloud asset explorer table in Steampipe provides you with comprehensive information about cloud assets within Prisma Cloud. This table allows you, as a security engineer or cloud administrator, to query asset-specific details, including account information, alert statuses, compliance standards, and vulnerability details. You can utilize this table to gather insights on cloud assets, such as compliance status, alert severity, and vulnerability counts. The schema outlines the various attributes of the Prisma Cloud assets for you.

## Table Usage Guide

The `prismacloud_inventory_asset_explorer` table in Steampipe provides detailed information about cloud assets within Prisma Cloud. This table allows you to query details such as account ID, alert statuses, compliance standards, and vulnerability information, enabling you to manage and monitor your cloud assets effectively.

**Important Notes**
- For improved performance, it is recommended to use the optional qualifiers (quals) to limit the result set.
- Queries with optional qualifiers are optimized to use filters. The following columns support optional qualifiers:
  - `account_name`
  - `cloud_type`
  - `compliance_requirement_name`
  - `compliance_standard_name`
  - `region_name`
  - `scan_status`

## Examples

### Basic info
Retrieve basic information about Prisma Cloud assets, such as account ID, account name, and asset type.

```sql+postgres
select
  account_id,
  account_name,
  asset_type
from
  prismacloud_inventory_asset_explorer;
```

```sql+sqlite
select
  account_id,
  account_name,
  asset_type
from
  prismacloud_inventory_asset_explorer;
```

### Get asset alert statuses
Get the count of alert statuses for each asset by severity level. This is useful for understanding the alert distribution across your assets.

```sql+postgres
select
  name,
  alert_status_critical,
  alert_status_high,
  alert_status_medium,
  alert_status_low,
  alert_status_informational
from
  prismacloud_inventory_asset_explorer;
```

```sql+sqlite
select
  name,
  alert_status_critical,
  alert_status_high,
  alert_status_medium,
  alert_status_low,
  alert_status_informational
from
  prismacloud_inventory_asset_explorer;
```

### Get compliance standard details of assets
Retrieve the compliance standard name and scan status for each asset. This helps in understanding the compliance posture and scan results of your assets.

```sql+postgres
select
  e.compliance_standard_name,
  s.policies_assigned_count,
  s.system_default,
  e.scan_status
from
  prismacloud_inventory_asset_explorer as e
  join prismacloud_compliance_standard as s on s.name = e.compliance_standard_name;
```

```sql+sqlite
select
  e.compliance_standard_name,
  s.policies_assigned_count,
  s.system_default,
  e.scan_status
from
  prismacloud_inventory_asset_explorer as e
  join prismacloud_compliance_standard as s on s.name = e.compliance_standard_name;
```

### Get vulnerability statuses of assets
Identify the number of vulnerabilities by severity level for each asset. This helps in assessing the impact of vulnerabilities on your assets.

```sql+postgres
select
  name,
  vulnerability_status_critical,
  vulnerability_status_high,
  vulnerability_status_medium,
  vulnerability_status_low
from
  prismacloud_inventory_asset_explorer;
```

```sql+sqlite
select
  name,
  vulnerability_status_critical,
  vulnerability_status_high,
  vulnerability_status_medium,
  vulnerability_status_low
from
  prismacloud_inventory_asset_explorer;
```

### Get assets overall passed resources
Get the list of resources that passed overall checks. This helps in identifying compliant and secure resources within your cloud environment.

```sql+postgres
select
  name,
  overall_passed
from
  prismacloud_inventory_asset_explorer
where
  overall_passed = true;
```

```sql+sqlite
select
  name,
  overall_passed
from
  prismacloud_inventory_asset_explorer
where
  overall_passed = 1;
```