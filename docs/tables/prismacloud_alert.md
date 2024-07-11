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
- This table supports optional quals. Queries with optional quals are optimised to use CloudWatch filters. Optional quals are supported for the following columns:
  - `status`
  - `policy_id`
  - `policy_type`
  - `policy_remediable`

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

### Alerts with high risk
Identify alerts that have a high-risk level. This helps in prioritizing alerts based on their risk details.

```sql+postgres
select
  id,
  status,
  risk_detail,
  alert_time
from
  prismacloud_alert
where
  risk_detail ->> 'risk_level' = 'high';
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
  json_extract(risk_detail, '$.risk_level') = 'high';
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