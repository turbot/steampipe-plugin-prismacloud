---
title: "Steampipe Table: prismacloud_policy - Query Prisma Cloud policies using SQL"
description: "Allows users to query Prisma Cloud policies. This table provides information about each policy, including its name, type, severity, and more. It can be used to monitor and manage policies within Prisma Cloud."
---

# Table: prismacloud_policy - Query Prisma Cloud policies using SQL

The Prisma Cloud policy table in Steampipe provides you with information about policies within Prisma Cloud. This table allows you, as a security engineer or cloud administrator, to query policy-specific details, including policy name, type, severity, and more. You can utilize this table to gather insights on policies, such as their configurations, compliance metadata, and more. The schema outlines the various attributes of the Prisma Cloud policy for you, including the policy ID, name, and associated rules.

## Table Usage Guide

The `prismacloud_policy` table in Steampipe provides information about policies within Prisma Cloud. This table allows you to query details such as the policy's name, type, severity, and more, enabling you to manage and monitor your policies effectively.

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
