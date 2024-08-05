---
title: "Steampipe Table: prismacloud_compliance_requirement - Query Prisma Cloud compliance requirements using SQL"
description: "Allows users to query Prisma Cloud compliance requirements. This table provides information about each requirement, including its name, description, associated policies, and more. It can be used to monitor and manage compliance requirements within Prisma Cloud."
---

# Table: prismacloud_compliance_requirement - Query Prisma Cloud compliance requirements using SQL

The Prisma Cloud compliance requirement table in Steampipe provides you with information about compliance requirements within Prisma Cloud. This table allows you, as a security engineer or cloud administrator, to query requirement-specific details, including name, description, associated policies, and more. You can utilize this table to gather insights on compliance requirements, such as their associated sections, policies, and more. The schema outlines the various attributes of the Prisma Cloud compliance requirement for you, including the requirement's ID, name, and description.

## Table Usage Guide

The `prismacloud_compliance_requirement` table in Steampipe provides information about compliance requirements within Prisma Cloud. This table allows you to query details such as the requirement's name, description, associated policies, and more, enabling you to manage and monitor your compliance requirements effectively.

## Examples

### Basic Info
Retrieve basic information about Prisma Cloud compliance requirements, such as name, description, and creation time. This query helps you to understand the overall configuration and details of your compliance requirements.

```sql+postgres
select
  name,
  description,
  created_on
from
  prismacloud_compliance_requirement;
```

```sql+sqlite
select
  name,
  description,
  created_on
from
  prismacloud_compliance_requirement;
```

### List of requirements with their assigned policies count
Get a list of all compliance requirements along with the count of assigned policies. This is useful for identifying the requirements that have policies associated with them.

```sql+postgres
select
  name,
  id,
  policies_assigned_count
from
  prismacloud_compliance_requirement;
```

```sql+sqlite
select
  name,
  id,
  policies_assigned_count
from
  prismacloud_compliance_requirement;
```

### List requirements modified in last 30 days
Retrieve compliance requirements that were modified recently. This helps in tracking changes and understanding recent modifications.

```sql+postgres
select
  name,
  last_modified_by,
  last_modified_on
from
  prismacloud_compliance_requirement
where
  last_modified_on > now() - interval '30 day';
```

```sql+sqlite
select
  name,
  last_modified_by,
  last_modified_on
from
  prismacloud_compliance_requirement
where
  last_modified_on > datetime('now', '-30 days');
```

### List system default requirements
Get a list of compliance requirements that are marked as system default. This helps in understanding which requirements are predefined by the system.

```sql+postgres
select
  name,
  system_default
from
  prismacloud_compliance_requirement
where
  system_default = true;
```

```sql+sqlite
select
  name,
  system_default
from
  prismacloud_compliance_requirement
where
  system_default = 1;
```

### Requirements with associated sections
Retrieve compliance requirements along with their associated sections. This helps in understanding the structure and details of each compliance requirement.

```sql+postgres
select
  name,
  s ->> 'id' as section_id,
  s ->> 'sectionId' as section_name,
  s ->> 'label' as section_label,
  s ->> 'createdBy' as section_created_by,
  s ->> 'createdOn' as section_created_on,
  s ->> 'viewOrder' as section_view_order,
  s ->> 'standardName' as compliance_standard_name,
  s ->> 'viewOrder' as section_view_order,
  s ->> 'systemDefault' as section_system_default,
  s ->> 'lastModifiedBy' as section_last_modified_by,
  s ->> 'lastModifiedOn' as section_last_modified_on,
  s -> 'associatedPolicyIds' as section_associated_policy_ids,
  s ->> 'policiesAssignedCount' as section_policies_assigned_count
from
  prismacloud_compliance_requirement,
  jsonb_array_elements(requirement_sections) as s;
```

```sql+sqlite
select
  name,
  s.value ->> '$.id' as section_id,
  s.value ->> '$.sectionId' as section_name,
  s.value ->> '$.label' as section_label,
  s.value ->> '$.createdBy' as section_created_by,
  s.value ->> '$.createdOn' as section_created_on,
  s.value ->> '$.viewOrder' as section_view_order,
  s.value ->> '$.standardName' as compliance_standard_name,
  s.value ->> '$.viewOrder' as section_view_order,
  s.value ->> '$.systemDefault' as section_system_default,
  s.value ->> '$.lastModifiedBy' as section_last_modified_by,
  s.value ->> '$.lastModifiedOn' as section_last_modified_on,
  json_extract(s.value, '$.associatedPolicyIds') as section_associated_policy_ids,
  s.value ->> '$.policiesAssignedCount' as section_policies_assigned_count
from
  prismacloud_compliance_requirement,
  json_each(requirement_sections) as s;
```

### Get requirement details for specific compliance standard
It allows users to see the structure of the compliance standard, requirements, and sections making it easier to understand how the requirements, and sections are organized and what they entail.

```sql+postgres
select
  s.name as standard_name,
  s.id as standard_id,
  r.name as requirement_name,
  r.description as requirement_description,
  r.requirement_sections
from
  prismacloud_compliance_standard as s
  join prismacloud_compliance_requirement as r on r.compliance_id = s.id
where
 s.name = 'Azure Security Benchmark (v2)';
```

```sql+sqlite
select
  s.name as standard_name,
  s.id as standard_id,
  r.name as requirement_name,
  r.description as requirement_description,
  json(r.requirement_sections) as requirement_sections
from
  prismacloud_compliance_standard as s
  join prismacloud_compliance_requirement as r on r.compliance_id = s.id
where
  s.name = 'Azure Security Benchmark (v2)';
```