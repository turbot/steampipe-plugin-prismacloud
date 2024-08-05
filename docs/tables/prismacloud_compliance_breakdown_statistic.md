---
title: "Steampipe Table: prismacloud_compliance_breakdown_statistic - Query Prisma Cloud compliance breakdown statistics using SQL"
description: "Allows users to query Prisma Cloud compliance breakdown statistics. This table provides information about the compliance status of resources, including failed resources by severity, compliance requirements, and more. It can be used to monitor and manage compliance within Prisma Cloud."
---

# Table: prismacloud_compliance_breakdown_statistic - Query Prisma Cloud compliance breakdown statistics using SQL

The Prisma Cloud compliance breakdown statistic table in Steampipe provides you with detailed information about the compliance status of resources within Prisma Cloud. This table allows security engineers and cloud administrators to query compliance-specific details, including the number of resources that have passed or failed compliance checks, categorized by severity levels. The schema outlines various attributes, such as account information, cloud type, compliance standard, and the number of resources with different compliance statuses.

## Table Usage Guide

The `prismacloud_compliance_breakdown_statistic` table in Steampipe provides information about the compliance status of resources within Prisma Cloud. This table allows you to query details such as the number of resources that have passed or failed compliance checks, categorized by severity levels. This helps in managing and monitoring the compliance status of your cloud resources effectively.

**Important Notes**
- For improved performance, it is recommended to use the optional qualifiers (quals) to limit the result set.
- Queries with optional qualifiers are optimized to use filters. The following columns support optional qualifiers:
  - `account_name`
  - `cloud_type`
  - `cloud_region`
  - `policy_compliance_standard_name`
  - `policy_compliance_requirement_name`
  - `policy_compliance_section_id`

## Examples
### Basic info
Retrieve a basic summary of compliance breakdown, including the number of failed and passed resources.

```sql+postgres
select
  account_name,
  cloud_type,
  failed_resources,
  passed_resources,
  total_resources
from
  prismacloud_compliance_breakdown_statistic;
```

```sql+sqlite
select
  account_name,
  cloud_type,
  failed_resources,
  passed_resources,
  total_resources
from
  prismacloud_compliance_breakdown_statistic;
```

### List high severity failed resources
Retrieve the breakdown of high severity failed resources and order by the number of high severity failed resources. This helps in identifying the areas with the most critical compliance issues.

```sql+postgres
select
  account_name,
  cloud_type,
  high_severity_failed_resources
from
  prismacloud_compliance_breakdown_statistic
order by
  high_severity_failed_resources desc;
```

```sql+sqlite
select
  account_name,
  cloud_type,
  high_severity_failed_resources
from
  prismacloud_compliance_breakdown_statistic
order by
  high_severity_failed_resources desc;
```

### List breakdown statistics group by account and cloud type
Retrieve a summary of compliance breakdown grouped by account name and cloud type. This query helps you to understand the compliance status of resources across different accounts and cloud environments.

```sql+postgres
select
  account_name,
  cloud_type,
  sum(critical_severity_failed_resources) as critical_failed_resources,
  sum(high_severity_failed_resources) as high_failed_resources,
  sum(medium_severity_failed_resources) as medium_failed_resources,
  sum(low_severity_failed_resources) as low_failed_resources,
  sum(informational_severity_failed_resources) as informational_failed_resources,
  sum(passed_resources) as passed_resources,
  sum(total_resources) as total_resources
from
  prismacloud_compliance_breakdown_statistic
group by
  account_name,
  cloud_type;
```

```sql+sqlite
select
  account_name,
  cloud_type,
  sum(critical_severity_failed_resources) as critical_failed_resources,
  sum(high_severity_failed_resources) as high_failed_resources,
  sum(medium_severity_failed_resources) as medium_failed_resources,
  sum(low_severity_failed_resources) as low_failed_resources,
  sum(informational_severity_failed_resources) as informational_failed_resources,
  sum(passed_resources) as passed_resources,
  sum(total_resources) as total_resources
from
  prismacloud_compliance_breakdown_statistic
group by
  account_name,
  cloud_type;
```

### Get breakdown statistics for compliance standard and requirement
Retrieve detailed compliance breakdown statistics by joining with the compliance standard and requirement tables. This query helps you to get a comprehensive view of the compliance status of resources, including the associated compliance standards and requirements.

```sql+postgres
select
  r.compliance_id as standard_id,
  r.name as requirement_name,
  b.account_name,
  b.cloud_type,
  b.failed_resources,
  b.passed_resources,
  b.total_resources
from
  prismacloud_compliance_breakdown_statistic as b
  join prismacloud_compliance_requirement as r on b.policy_compliance_requirement_name = r.name;
```

```sql+sqlite
select
  r.compliance_id as standard_id,
  r.name as requirement_name,
  b.account_name,
  b.cloud_type,
  b.failed_resources,
  b.passed_resources,
  b.total_resources
from
  prismacloud_compliance_breakdown_statistic as b
  join prismacloud_compliance_requirement as r on b.policy_compliance_requirement_name = r.name;
```

### Recently updated compliance statistics
Retrieve compliance breakdown statistics that were updated within the last 30 days. This query helps in tracking recent changes and understanding the current compliance status.

```sql+postgres
select
  account_name,
  cloud_type,
  timestamp,
  failed_resources,
  passed_resources,
  total_resources
from
  prismacloud_compliance_breakdown_statistic
where
  timestamp > now() - interval '30 days';
```

```sql+sqlite
select
  account_name,
  cloud_type,
  timestamp,
  failed_resources,
  passed_resources,
  total_resources
from
  prismacloud_compliance_breakdown_statistic
where
  timestamp > datetime('now', '-30 days');
```