---
title: "Steampipe Table: prismacloud_alert - Query Prisma Cloud alerts using SQL"
description: "Allows users to query Prisma Cloud alerts. This table provides information about each alert, including its status, time, event details, and associated policies. It can be used to monitor and manage alerts within Prisma Cloud."
---

# Table: prismacloud_alert - Query Prisma Cloud alerts using SQL

The Prisma Cloud alert table in Steampipe provides you with information about alerts within Prisma Cloud. This table allows you, as a security engineer or cloud administrator, to query alert-specific details, including status, timestamps, event details, and associated policies. You can utilize this table to gather insights on alerts, such as their history, risk details, and more. The schema outlines the various attributes of the Prisma Cloud alert for you, including the alert's ID, status, and associated resource.

## Table Usage Guide

The `prismacloud_alert` table in Steampipe provides information about alerts within Prisma Cloud. This table allows you to query details such as the alert's ID, status, timestamps, and more, enabling you to manage and monitor your alerts effectively.

**Important notes:**
- For improved performance, it is advised that you use the optional qual `alert_time` to limit the result set to a specific time period.
- Queries with optional qualifiers are optimized to use filters. The following columns support optional qualifiers:
  - `alert_time`
  - `status`
  - `policy_id`
  - `policy_type`
  - `policy_remediable`
  - `policy_compliance_standard_name`
  - `policy_compliance_requirement_name`
  - `policy_compliance_section_id`

## Examples

### Basic Info
Retrieve basic information about Prisma Cloud alerts, such as alert ID, status, and timestamps. This query helps you understand the overall configuration and details of your alerts.

```sql+postgres
select
  id,
  status,
  first_seen,
  last_seen,
  alert_time,
  event_occurred
from
  prismacloud_alert
where
  alert_time >= current_date - interval '5d';
```

```sql+sqlite
select
  id,
  status,
  first_seen,
  last_seen,
  alert_time,
  event_occurred
from
  prismacloud_alert
where
  alert_time >= current_date - interval '5d';
```

### List open alerts
Get a list of all Prisma Cloud alerts filtered by status. This is useful for identifying which alerts are currently active.

```sql+postgres
select
  id,
  status,
  alert_time,
  last_seen
from
  prismacloud_alert
where
  status = 'open';
```

```sql+sqlite
select
  id,
  status,
  alert_time,
  last_seen
from
  prismacloud_alert
where
  status = 'open';
```

### Alerts and associated policies
Get a list of alerts along with their associated policies by joining with the `prismacloud_policy` table. This helps in understanding the policies that are related to specific alerts.

```sql+postgres
select
  a.id as alert_id,
  a.status,
  a.alert_time,
  p.id as policy_id,
  p.name as policy_name,
  p.policy_type
from
  prismacloud_alert a
join
  prismacloud_policy p on a.policy ->> 'policy_id' = p.id and a.policy ->> 'policy_type' = p.policy_type;
```

```sql+sqlite
select
  a.id as alert_id,
  a.status,
  a.alert_time,
  p.id as policy_id,
  p.name as policy_name,
  p.policy_type
from
  prismacloud_alert a
join
  prismacloud_policy p on json_extract(a.policy, '$.policy_id') = p.id and json_extract(a.policy, '$.policy_type') = p.policy_type;
```

### List top 20 oldest open alerts
Retrieve the 20 oldest open alerts from Prisma Cloud, sorted by their alert time. This query is useful for identifying alerts that have been open for the longest period.

```sql+postgres
select
  id,
  status,
  risk_detail,
  alert_time
from
  prismacloud_alert
where
  status = 'Open'
order by
  alert_time asc
limit 20;
```

```sql+sqlite
select
  id,
  status,
  risk_detail,
  alert_time
from
  prismacloud_alert
where
  status = 'Open'
order by
  alert_time asc
limit 20;
```

### Get the latest alerts
Retrieve the most recent alerts from Prisma Cloud that have occurred within the last day. This query helps you stay updated on new alerts and their details.

```sql+postgres
select
  id,
  status,
  alert_time,
  triggered_by,
  event_occurred
from
  prismacloud_alert
where
  alert_time >= now() - interval '1 day';
```

```sql+sqlite
select
  id,
  status,
  alert_time,
  triggered_by,
  event_occurred
from
  prismacloud_alert
where
  alert_time >= datetime('now', '-1 day');
```

### List critical alerts in the last 7 days
This query filters Prisma Cloud alerts from the last 7 days and joins them with high-severity policies to identify critical alerts that require immediate attention.

```sql+postgres
with high_severity_policies as (
  select
    policy_id,
    policy_type,
    severity
  from
    prismacloud_policy
  where
    severity = 'high'
),
recent_alerts as (
  select
    id,
    status,
    policy_id,
    policy_type,
    alert_time
  from
    prismacloud_alert
  where
    alert_time >= now() - interval '7 day'
)
select
  a.id,
  a.status,
  a.policy_id,
  a.policy_type,
  p.severity as policy_severity,
  a.alert_time
from
  recent_alerts as a
join
  high_severity_policies as p on a.policy_id = p.policy_id and a.policy_type = p.policy_type;
```

```sql+sqlite
with high_severity_policies as (
  select
    policy_id,
    policy_type,
    severity
  from
    prismacloud_policy
  where
    severity = 'high'
),
recent_alerts as (
  select
    id,
    status,
    policy_id,
    policy_type,
    alert_time
  from
    prismacloud_alert
  where
    alert_time >= datetime('now', '-7 day')
)
select
  a.id,
  a.status,
  a.policy_id,
  a.policy_type,
  p.severity as policy_severity,
  a.alert_time
from
  recent_alerts as a
join
  high_severity_policies as p on a.policy_id = p.policy_id and a.policy_type = p.policy_type;
```

### Count alerts by policy type and policy ID
Retrieve the number of alerts grouped by policy ID. This query helps you understand which policies are generating the most alerts.

```sql+postgres
select
  a.policy_id,
  p.name as policy_name,
  count(a.id) as alert_count
from
  prismacloud_alert as a
join
  prismacloud_policy as p on p.policy_id = a.policy_id
group by
  a.policy_id, p.name
order by
  alert_count desc;
```

```sql+sqlite
select
  a.policy_id,
  p.name as policy_name,
  count(a.id) as alert_count
from
  prismacloud_alert as a
join
  prismacloud_policy as p on p.policy_id = a.policy_id
group by
  a.policy_id, p.name
order by
  alert_count desc;
```

### Show alerts for a specific policy
Retrieve all alerts associated with a specific policy.

```sql+postgres
select
  a.id,
  a.status,
  a.alert_time,
  a.triggered_by,
  a.event_occurred
from
  prismacloud_alert as a
where
  a.policy_id = 'bb8a13e4-d3d9-4618-925a-0cff3526430e';
```

```sql+sqlite
select
  a.id,
  a.status,
  a.alert_time,
  a.triggered_by,
  a.event_occurred
from
  prismacloud_alert as a
where
  a.policy_id = 'bb8a13e4-d3d9-4618-925a-0cff3526430e';
```

### Count alerts by cloud account
Retrieve the number of alerts grouped by cloud account. This query helps you identify which cloud accounts are generating the most alerts.

```sql+postgres
select
  a.resource ->> 'account' as cloud_account,
  count(a.id) as alert_count
from
  prismacloud_alert as a
group by
  cloud_account
order by
  alert_count desc;
```

```sql+sqlite
select
  json_extract(a.resource, '$.account') as cloud_account,
  count(a.id) as alert_count
from
  prismacloud_alert as a
group by
  cloud_account
order by
  alert_count desc;
```

### Count alerts by resource
Retrieve the number of alerts grouped by resource. This query helps you identify which resources are associated with the most alerts.

```sql+postgres
select
  a.resource ->> 'name' as resource_name,
  count(a.id) as alert_count
from
  prismacloud_alert as a
group by
  resource_name
order by
  alert_count desc;
```

```sql+sqlite
select
  json_extract(a.resource, '$.name') as resource_name,
  count(a.id) as alert_count
from
  prismacloud_alert as a
group by
  resource_name
order by
  alert_count desc;
```

### Count alerts by cloud type
Retrieve the number of alerts grouped by cloud type. This query helps you understand the distribution of alerts across different cloud providers.

```sql+postgres
select
  a.resource ->> 'cloudType' as cloud_type,
  count(a.id) as alert_count
from
  prismacloud_alert as a
group by
  cloud_type
order by
  alert_count desc;
```

```sql+sqlite
select
  json_extract(a.resource, '$.cloudType') as cloud_type,
  count(a.id) as alert_count
from
  prismacloud_alert as a
group by
  cloud_type
order by
  alert_count desc;
```

### Count alerts by cloud service
Retrieve the number of alerts grouped by cloud service. This query helps you identify which cloud services are generating the most alerts.

```sql+postgres
select
  a.resource ->> 'cloudServiceName' as cloud_service,
  count(a.id) as alert_count
from
  prismacloud_alert as a
group by
  cloud_service
order by
  alert_count desc;
```

```sql+sqlite
select
  json_extract(a.resource, '$.cloudServiceName') as cloud_service,
  count(a.id) as alert_count
from
  prismacloud_alert as a
group by
  cloud_service
order by
  alert_count desc;
```

### Count alerts by status
Retrieve the number of alerts grouped by status. This query helps you understand the current state of your alerts.

```sql+postgres
select
  a.status,
  count(a.id) as alert_count
from
  prismacloud_alert as a
group by
  a.status
order by
  alert_count desc;
```

```sql+sqlite
select
  a.status,
  count(a.id) as alert_count
from
  prismacloud_alert as a
group by
  a.status
order by
  alert_count desc;
```

### Count alerts by policy type
Retrieve the number of alerts grouped by policy type, such as Internet exposure or Misconfiguration.

```sql+postgres
select
  p.policy_category as policy_type,
  count(a.id) as alert_count
from
  prismacloud_alert as a
join
  prismacloud_policy as p on a.policy_id = p.policy_id and a.policy_type = p.policy_type
group by
  p.policy_category
order by
  alert_count desc;
```

```sql+sqlite
select
  p.policy_category as policy_type,
  count(a.id) as alert_count
from
  prismacloud_alert as a
join
  prismacloud_policy as p on a.policy_id = p.policy_id and a.policy_type = p.policy_type
group by
  p.policy_category
order by
  alert_count desc;
```

### Count alerts by compliance standard
Retrieve the number of alerts grouped by compliance standard. This query helps you identify which compliance standards are associated with the most alerts.

```sql+postgres
select
  p.compliance_standard_name as compliance_standard,
  count(a.id) as alert_count
from
  prismacloud_alert as a
join
  prismacloud_policy as p on a.policy_id = p.policy_id and a.policy_type = p.policy_type
group by
  p.compliance_standard_name
order by
  alert_count desc;
```

```sql+sqlite
select
  p.compliance_standard_name as compliance_standard,
  count(a.id) as alert_count
from
  prismacloud_alert as a
join
  prismacloud_policy as p on a.policy_id = p.policy_id and a.policy_type = p.policy_type
group by
  p.compliance_standard_name
order by
  alert_count desc;
```

### Count alerts by for compliance standard requirements in the last 7 days
Retrieve the number of alerts grouped by compliance standard requirement. This query helps you identify which specific compliance requirements are generating the most alerts.

```sql+postgres
select
  r.name as compliance_requirement,
  count(a.id) as alert_count
from
  prismacloud_alert as a,
  prismacloud_compliance_requirement as r
where
  a.policy_compliance_requirement_name = r.name
  and a.alert_time >= now() - interval '7 day'
group by
  compliance_requirement
order by
  alert_count desc;
```

```sql+sqlite
select
  r.name as compliance_requirement,
  count(a.id) as alert_count
from
  prismacloud_alert as a,
  prismacloud_compliance_requirement as r
where
  a.policy_compliance_requirement_name = r.name
  and a.alert_time >= datetime('now', '-7 day')
group by
  r.name
order by
  alert_count desc;
```

### Get alerts for custom policies
This query retrieves alerts that are associated with custom policies, helping you monitor alerts generated by policies specifically created for your environment.

```sql+postgres
select
  a.id,
  a.status,
  a.alert_time,
  a.alert_count
from
  prismacloud_alert as a
join
  prismacloud_policy as p on a.policy_id = p.policy_id and a.policy_type = p.policy_type
where
  p.policy_mode = 'custom';
```

```sql+sqlite
select
  a.id,
  a.status,
  a.alert_time,
  a.alert_count
from
  prismacloud_alert as a
join
  prismacloud_policy as p on a.policy_id = p.policy_id and a.policy_type = p.policy_type
where
  p.policy_mode = 'custom';
```