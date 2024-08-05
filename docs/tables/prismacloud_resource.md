---
title: "Steampipe Table: prismacloud_resource - Query Prisma Cloud resources using SQL"
description: "Allows users to query Prisma Cloud resources. This table provides information about each resource, including its name, type, associated roles, and more. It can be used to monitor and manage resources within Prisma Cloud."
---

# Table: prismacloud_resource - Query Prisma Cloud resources using SQL

The Prisma Cloud resource table in Steampipe provides you with information about resources within Prisma Cloud. This table allows you, as a security engineer or cloud administrator, to query resource-specific details, including resource name, type, associated roles, and more. You can utilize this table to gather insights on resources, such as their configurations, associated features, and more. The schema outlines the various attributes of the Prisma Cloud resource for you, including the resource ID, name, and associated roles.

## Table Usage Guide

The `prismacloud_resource` table in Steampipe provides information about resources within Prisma Cloud. This table allows you to query details such as the resource's name, type, associated roles, and more, enabling you to manage and monitor your resources effectively.

## Examples

### Basic Info

Retrieve basic information about Prisma Cloud resources, such as resource ID, name, type, and description. This query helps you to understand the overall configuration and details of your resources.

```sql+postgres
select
  id,
  name,
  type,
  description,
  custom
from
  prismacloud_resource;
```

```sql+sqlite
select
  id,
  name,
  type,
  description,
  custom
from
  prismacloud_resource;
```

### List of custom resources

Get a list of all custom Prisma Cloud resources. This is useful for identifying which resources are custom-defined.

```sql+postgres
select
  id,
  name,
  description
from
  prismacloud_resource
where
  custom = true;
```

```sql+sqlite
select
  id,
  name,
  description
from
  prismacloud_resource
where
  custom = 1;
```

### Resources modified by a specific user

Identify resources that were last modified by a specific user. This helps in tracking changes made by administrators or other users.

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

### Resources accepting resource lists

Retrieve resources where resource lists are accepted. This helps in understanding the configurations related to resource list acceptance in your resources.

```sql+postgres
select
  id,
  name,
  accept_resource_lists
from
  prismacloud_resource
where
  accept_resource_lists = true;
```

```sql+sqlite
select
  id,
  name,
  accept_resource_lists
from
  prismacloud_resource
where
  accept_resource_lists = 1;
```

### Resources and their associated roles

Get a list of resources along with their associated roles. This can help in understanding the role assignments within your cloud environment.

```sql+postgres
select
  id,
  name,
  associated_roles
from
  prismacloud_resource;
```

```sql+sqlite
select
  id,
  name,
  associated_roles
from
  prismacloud_resource;
```
