---
title: "Steampipe Table: prismacloud_resource - Query Prisma Cloud resources using SQL"
description: "Allows users to query resources in Prisma Cloud. This table provides detailed information about each resource, including its name, ID, type, and more. It can be used to manage and audit resources within Prisma Cloud."
---

# Table: prismacloud_resource - Query Prisma Cloud resources using SQL

The `prismacloud_resource` table in Steampipe provides you with information about resources within Prisma Cloud. This table allows security engineers, cloud administrators, and compliance officers to query resource-specific details, including the resource's name, ID, description, and associated metadata. You can utilize this table to gather insights on resources, such as their modification history, associated members, and types. The schema outlines various attributes of the Prisma Cloud resource, including the resource ID, name, and last modification details.

## Table Usage Guide

The `prismacloud_resource` table in Steampipe provides detailed information about resources within Prisma Cloud. This table allows you to query details such as the resource's name, ID, description, and modification history, enabling you to effectively manage and audit your resources.

## Examples

### Basic info
Retrieve basic information about Prisma Cloud resources, such as resource ID, name, type, and description. This query helps you to understand the overall configuration and details of your resources.

```sql+postgres
select
  id,
  name,
  resource_list_type,
  description
from
  prismacloud_resource;
```

```sql+sqlite
select
  id,
  name,
  resource_list_type,
  description
from
  prismacloud_resource;
```

### List of Resources Modified by a Specific User
Get a list of all resources that were last modified by a specific user. This is useful for tracking changes made by administrators or team members.

```sql+postgres
select
  id,
  name,
  last_modified_by,
  last_modified_ts
from
  prismacloud_resource
where
  last_modified_by = 'admin_user';
```

```sql+sqlite
select
  id,
  name,
  last_modified_by,
  last_modified_ts
from
  prismacloud_resource
where
  last_modified_by = 'admin_user';
```

### Resources with Specific Types
Identify resources of a specific type. This helps in organizing and managing different categories of resources within Prisma Cloud.

```sql+postgres
select
  id,
  name,
  resource_list_type
from
  prismacloud_resource
where
  resource_list_type = 'compute';
```

```sql+sqlite
select
  id,
  name,
  resource_list_type
from
  prismacloud_resource
where
  resource_list_type = 'compute';
```

### Resources Without a Description

Identify resources that are missing a description to ensure all assets are well-documented.

```sql+postgres
select
  id,
  name,
  description
from
  prismacloud_resource
where
  description is null or description = '';
```

```sql+sqlite
select
  id,
  name,
  description
from
  prismacloud_resource
where
  description is null or description = '';
```

### Get Members Associated with Each Resource
Retrieve details about members associated with each resource. This can help in understanding who has access to which resources.

```sql+postgres
select
  id,
  name,
  members
from
  prismacloud_resource;
```

```sql+sqlite
select
  id,
  name,
  members
from
  prismacloud_resource;
```

### Count of Resources by Type
Get a count of resources by their type to understand the distribution of different resource types within your environment.

```sql+postgres
select
  resource_list_type,
  count(*) as resource_count
from
  prismacloud_resource
group by
  resource_list_type;
```

```sql+sqlite
select
  resource_list_type,
  count(*) as resource_count
from
  prismacloud_resource
group by
  resource_list_type;
```