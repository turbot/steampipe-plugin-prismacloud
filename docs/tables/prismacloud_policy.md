---
title: "Steampipe Table: prismacloud_policy - Query Prisma Cloud policies using SQL"
description: "Allows users to query Prisma Cloud policies. This table provides information about each policy, including its name, type, severity, and more. It can be used to monitor and manage policies within Prisma Cloud."
---

# Table: prismacloud_policy - Query Prisma Cloud policies using SQL

The Prisma Cloud policy table in Steampipe provides you with information about policies within Prisma Cloud. This table allows you, as a security engineer or cloud administrator, to query policy-specific details, including policy name, type, severity, and more. You can utilize this table to gather insights on policies, such as their configurations, compliance metadata, and more. The schema outlines the various attributes of the Prisma Cloud policy for you, including the policy ID, name, and associated rules.

## Table Usage Guide

The `prismacloud_policy` table in Steampipe provides information about policies within Prisma Cloud. This table allows you to query details such as the policy's name, type, severity, and more, enabling you to manage and monitor your policies effectively.

**Important Notes**
- For improved performance, it is recommended to use the optional qualifiers (quals) to limit the result set.
- Queries with optional qualifiers are optimized to use filters. The following columns support optional qualifiers:
  - `severity`
  - `cloud_type`
  - `policy_type`
  - `enabled`
  - `policy_mode`
  - `remediable`
  - `name`
  - `policy_compliance_standard_name`
  - `policy_compliance_requirement_name`
  - `policy_compliance_section_id`

## Examples

### Basic Info
Retrieve basic information about Prisma Cloud policies, such as policy ID, name, type, and severity. This query helps you to understand the overall configuration and details of your policies.

```sql+postgres
select
  policy_id,
  name,
  policy_type,
  severity,
  enabled
from
  prismacloud_policy;
```

```sql+sqlite
select
  policy_id,
  name,
  policy_type,
  severity,
  enabled
from
  prismacloud_policy;
```

### List of enabled policies
Get a list of all enabled Prisma Cloud policies. This is useful for identifying which policies are currently active and enabled.

```sql+postgres
select
  policy_id,
  name,
  severity
from
  prismacloud_policy
where
  enabled = true;
```

```sql+sqlite
select
  policy_id,
  name,
  severity
from
  prismacloud_policy
where
  enabled = 1;
```

### Policies created by a specific user
Identify policies that were created by a specific user. This helps in tracking which policies were introduced by which administrators or team members.

```sql+postgres
select
  policy_id,
  name,
  created_by,
  created_on
from
  prismacloud_policy
where
  created_by = 'admin_user';
```

```sql+sqlite
select
  policy_id,
  name,
  created_by,
  created_on
from
  prismacloud_policy
where
  created_by = 'admin_user';
```

### List of policies with high severity
Retrieve policies that have a high severity level. This helps in prioritizing policies that may require more immediate attention or enforcement.

```sql+postgres
select
  policy_id,
  name,
  severity
from
  prismacloud_policy
where
  severity = 'high';
```

```sql+sqlite
select
  policy_id,
  name,
  severity
from
  prismacloud_policy
where
  severity = 'high';
```

### Policies with specific compliance metadata
Identify policies associated with specific compliance metadata. This helps in ensuring that certain compliance requirements are being met by the policies.

```sql+postgres
select
  policy_id,
  name,
  compliance_metadata
from
  prismacloud_policy
where
  compliance_metadata @> '[{"complianceId": "CIS"}]';
```

```sql+sqlite
select
  policy_id,
  name,
  compliance_metadata
from
  prismacloud_policy
where
  json_extract(compliance_metadata, '$[0].complianceId') = 'CIS';
```

### Policies with open alerts
Get a list of policies that have open alerts. This helps in identifying which policies have ongoing issues that need attention.

```sql+postgres
select
  policy_id,
  name,
  open_alerts_count
from
  prismacloud_policy
where
  open_alerts_count > 0;
```

```sql+sqlite
select
  policy_id,
  name,
  open_alerts_count
from
  prismacloud_policy
where
  open_alerts_count > 0;
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

### Get the latest alerts for a specific policy
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
  alert_time >= now() - interval '1 day'
  and policy_id = '2378dbf4-b104-4bda-9b05-7417affbba3f';
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
  alert_time >= datetime('now', '-1 day')
  and policy_id = '2378dbf4-b104-4bda-9b05-7417affbba3f';
```

### Count policies by policy type
This query counts the number of policies grouped by policy type, helping you understand the distribution of policies across different policy categories.

```sql+postgres
select
  policy_type,
  count(policy_id) as policy_count
from
  prismacloud_policy
group by
  policy_type;
```

```sql+sqlite
select
  policy_type,
  count(policy_id) as policy_count
from
  prismacloud_policy
group by
  policy_type;
```

### Count alerts by policy type
This query counts the number of alerts grouped by policy type, helping you understand the distribution of alerts across different policy categories.

```sql+postgres
with alerts as (
  select
    id,
    policy_type
  from
    prismacloud_alert
)
select
  policy_type,
  count(id) as alert_count
from
  alerts
group by
  policy_type
order by
  alert_count desc;
```

```sql+sqlite
with alerts as (
  select
    id,
    policy_type
  from
    prismacloud_alert
)
select
  policy_type,
  count(id) as alert_count
from
  alerts
group by
  policy_type
order by
  alert_count desc;
```

### Count policies by mode
This query counts the number of policies grouped by their mode, helping you understand the distribution of policies across different modes.

```sql+postgres
select
  policy_mode,
  count(policy_id) as policy_count
from
  prismacloud_policy
group by
  policy_mode
order by
  policy_count desc;
```

```sql+sqlite
select
  policy_mode,
  count(policy_id) as policy_count
from
  prismacloud_policy
group by
  policy_mode
order by
  policy_count desc;
```

### List policies assigned to each compliance standard
This query retrieves the list of policies assigned to each compliance standard, including the compliance standard's name, the number of policies assigned, and details about each policy.

```sql+postgres
select
  c.name as compliance_name,
  c.policies_assigned_count as compliance_policies_assigned_count,
  p.policy_id,
  p.policy_type,
  p.name as policy_name
from
  prismacloud_compliance_standard as c
  join prismacloud_policy as p on p.compliance_standard_name = c.name;
```

```sql+sqlite
select
  c.name as compliance_name,
  c.policies_assigned_count as compliance_policies_assigned_count,
  p.policy_id,
  p.policy_type,
  p.name as policy_name
from
  prismacloud_compliance_standard as c
  join prismacloud_policy as p on p.compliance_standard_name = c.name;
```