---
title: "Steampipe Table: prismacloud_inventory_asset_view - Analyze Prisma Cloud assets using SQL"
description: "Allows users to query and analyze Prisma Cloud assets. This table provides detailed information about cloud assets, including account details, compliance standards, severity of failed resources, and vulnerability information. It can be used to monitor and manage cloud assets within Prisma Cloud."
---

# Table: prismacloud_inventory_asset_view - Analyze Prisma Cloud assets using SQL

The Prisma Cloud asset view table in Steampipe provides you with comprehensive information about cloud assets within Prisma Cloud. This table allows you, as a security engineer or cloud administrator, to query asset-specific details, including account information, compliance standards, failed resources by severity, and vulnerability details. You can utilize this table to gather insights on cloud assets, such as compliance status, failure severity, and unscanned resources. The schema outlines the various attributes of the Prisma Cloud assets for you.

## Table Usage Guide

The `prismacloud_inventory_asset_view` table in Steampipe provides detailed information about cloud assets within Prisma Cloud. This table allows you to query details such as account ID, compliance standards, failed resources, and vulnerability information, enabling you to manage and monitor your cloud assets effectively.

**Important Notes**
- For improved performance, it is recommended to use the optional qualifiers (quals) to limit the result set.
- Queries with optional qualifiers are optimized to use filters. The following columns support optional qualifiers:
  - `account_name`
  - `cloud_type_name`
  - `compliance_requirement_name`
  - `compliance_standard_name`
  - `group_by`
  - `region_name`
  - `scan_status`
  - `service_name`
- The response includes an attribute `groupedAggregates`, whose content depends on the `group_by` query parameter. The following table shows the attributes that `groupedAggregates` will include for the specified `group_by` query parameter:

  | group_by        | column includes          |
  | -------------- | ------------------------------------ |
  | not specified  | service_name, cloud_type_name |
  | cloudType      | cloud_type_name  |
  | cloud.account  | account_name  |
  | cloud.region   | region_name, cloud_type_name   |
  | cloud.service  | service_name, cloud_type_name |
  | resource.type  | resource_type_name, cloud_type_name  |

- By default, the table will return rows with the `group_by` query parameter value `cloud.service`. For more information, please see [Asset Inventory View](https://pan.dev/prisma-cloud/api/cspm/asset-inventory-v-3/).

## Examples

### Basic Info
Retrieve basic information about Prisma Cloud assets, such as account ID, account name, and cloud type.

```sql+postgres
select
  account_id,
  account_name,
  cloud_type_name
from
  prismacloud_inventory_asset_view;
```

```sql+sqlite
select
  account_id,
  account_name,
  cloud_type_name
from
  prismacloud_inventory_asset_view;
```

### Compliance and scan status
Get the compliance standard name and scan status for each asset. This is useful for understanding the compliance posture and scan results of your assets.

```sql+postgres
select
  compliance_standard_name,
  scan_status
from
  prismacloud_inventory_asset_view;
```

```sql+sqlite
select
  compliance_standard_name,
  scan_status
from
  prismacloud_inventory_asset_view;
```

### Failed resources by severity
Identify the number of resources that failed by severity level. This helps in assessing the impact of policy failures on your assets.

```sql+postgres
select
  critical_severity_failed_resources,
  high_severity_failed_resources,
  medium_severity_failed_resources,
  low_severity_failed_resources,
  informational_severity_failed_resources
from
  prismacloud_inventory_asset_view;
```

```sql+sqlite
select
  critical_severity_failed_resources,
  high_severity_failed_resources,
  medium_severity_failed_resources,
  low_severity_failed_resources,
  informational_severity_failed_resources
from
  prismacloud_inventory_asset_view;
```

### Vulnerability details
Retrieve details about vulnerabilities, including the number of failed resources by vulnerability severity.

```sql+postgres
select
  critical_vulnerability_failed_resources,
  high_vulnerability_failed_resources,
  medium_vulnerability_failed_resources,
  low_vulnerability_failed_resources,
  total_vulnerability_failed_resources
from
  prismacloud_inventory_asset_view;
```

```sql+sqlite
select
  critical_vulnerability_failed_resources,
  high_vulnerability_failed_resources,
  medium_vulnerability_failed_resources,
  low_vulnerability_failed_resources,
  total_vulnerability_failed_resources
from
  prismacloud_inventory_asset_view;
```

### Unscanned resources
Get the total number of unscanned resources. This helps in identifying assets that have not been scanned for compliance or vulnerabilities.

```sql+postgres
select
  unscanned_resources
from
  prismacloud_inventory_asset_view;
```

```sql+sqlite
select
  unscanned_resources
from
  prismacloud_inventory_asset_view;
```