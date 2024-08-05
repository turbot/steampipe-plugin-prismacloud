---
title: "Steampipe Table: prismacloud_compliance_standard - Query Prisma Cloud compliance standards using SQL"
description: "Allows users to query Prisma Cloud compliance standards. This table provides information about each standard, including its name, description, associated policies, and more. It can be used to monitor and manage compliance standards within Prisma Cloud."
---

# Table: prismacloud_compliance_standard - Query Prisma Cloud compliance standards using SQL

The Prisma Cloud compliance standard table in Steampipe provides you with information about compliance standards within Prisma Cloud. This table allows you, as a security engineer or cloud administrator, to query compliance standard-specific details, including name, description, associated policies, and more. You can utilize this table to gather insights on standards, such as their cloud types, creation and modification details, and more. The schema outlines the various attributes of the Prisma Cloud compliance standard for you, including the standard's ID, name, and description.

## Table Usage Guide

The `prismacloud_compliance_standard` table in Steampipe provides information about compliance standards within Prisma Cloud. This table allows you to query details such as the standard's name, description, associated policies, and more, enabling you to manage and monitor your compliance standards effectively.

## Examples

### Basic Info
Retrieve basic information about Prisma Cloud compliance standards, such as name, description, and whether it is a system default. This query helps you to understand the overall configuration and details of your standards.

```sql+postgres
select
  name,
  description,
  system_default
from
  prismacloud_compliance_standard;
```

```sql+sqlite
select
  name,
  description,
  system_default
from
  prismacloud_compliance_standard;
```

### Assigned policies count with standard compliance
Get a list of all compliance standards along with the number of policies assigned to them. This is useful for identifying which standards have the most policies assigned.

```sql+postgres
select
  name,
  id,
  policies_assigned_count
from
  prismacloud_compliance_standard;
```

```sql+sqlite
select
  name,
  id,
  policies_assigned_count
from
  prismacloud_compliance_standard;
```

### Get assigned policy details for standard compliance
This information is crucial for security engineers and cloud administrators to ensure that compliance requirements are being met and to monitor the status and impact of various security policies.

```sql+postgres
select
  c.name as compliance_name,
  c.id as compliance_id,
  c.created_on as compliance_create_time,
  p.policy_id,
  p.name as policy_name,
  p.severity as policy_severity,
  p.enabled as is_policy_enabled
from
  prismacloud_compliance_standard as c
  join prismacloud_policy as p on p.compliance_requirement_name = c.name;
```

```sql+sqlite
select
  c.name as compliance_name,
  c.id as compliance_id,
  c.created_on as compliance_create_time,
  p.policy_id,
  p.name as policy_name,
  p.severity as policy_severity,
  p.enabled as is_policy_enabled
from
  prismacloud_compliance_standard as c
  join prismacloud_policy as p on p.compliance_requirement_name = c.name;
```

### Recently modified standard compliances
Retrieve compliance standards that were modified recently. This helps in tracking changes and understanding recent modifications.

```sql+postgres
select
  name,
  last_modified_by,
  last_modified_on
from
  prismacloud_compliance_standard
where
  last_modified_on > now() - interval '30 day';
```

```sql+sqlite
select
  name,
  last_modified_by,
  last_modified_on
from
  prismacloud_compliance_standard
where
  last_modified_on > datetime('now', '-30 days');
```

### Standard compliance created by a specific user
Get a list of compliance standards that were created by a specific user. This helps in understanding who created which standards.

```sql+postgres
select
  name,
  created_by
from
  prismacloud_compliance_standard
where
  created_by = 'user@example.com';
```

```sql+sqlite
select
  name,
  created_by
from
  prismacloud_compliance_standard
where
  created_by = 'user@example.com';
```

### Standard compliances for specific cloud types
Retrieve compliance standards that apply to specific cloud environments. This helps in managing standards based on cloud types.

```sql+postgres
select
  name,
  cloud_type
from
  prismacloud_compliance_standard
where
  cloud_type ?| array['aws', 'azure'];
```

```sql+sqlite
select
  name,
  cloud_type
from
  prismacloud_compliance_standard
where
  json_contains(cloud_type, json_array('aws', 'azure'));
```