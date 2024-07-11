---
title: "Steampipe Table: prisma_permission_group - Query Prisma Cloud permission groups using SQL"
description: "Allows users to query Prisma Cloud permission groups. This table provides information about each permission group, including its name, type, associated roles, and more. It can be used to monitor and manage permission groups within Prisma Cloud."
---

# Table: prisma_permission_group - Query Prisma Cloud permission groups using SQL

The Prisma Cloud permission group table in Steampipe provides you with information about permission groups within Prisma Cloud. This table allows you, as a security engineer or cloud administrator, to query permission group-specific details, including group name, type, associated roles, and more. You can utilize this table to gather insights on permission groups, such as their configurations, associated features, and more. The schema outlines the various attributes of the Prisma Cloud permission group for you, including the group ID, name, and associated roles.

## Table Usage Guide

The `prisma_permission_group` table in Steampipe provides information about permission groups within Prisma Cloud. This table allows you to query details such as the permission group's name, type, associated roles, and more, enabling you to manage and monitor your permission groups effectively.

## Examples

### Basic Info
Retrieve basic information about Prisma Cloud permission groups, such as group ID, name, type, and description. This query helps you to understand the overall configuration and details of your permission groups.

```sql+postgres
select
  id,
  name,
  type,
  description,
  custom
from
  prisma_permission_group;
```

```sql+sqlite
select
  id,
  name,
  type,
  description,
  custom
from
  prisma_permission_group;
```

### List of custom permission groups
Get a list of all custom Prisma Cloud permission groups. This is useful for identifying which permission groups are custom-defined.

```sql+postgres
select
  id,
  name,
  description
from
  prisma_permission_group
where
  custom = true;
```

```sql+sqlite
select
  id,
  name,
  description
from
  prisma_permission_group
where
  custom = 1;
```

### Permission groups modified by a specific user
Identify permission groups that were last modified by a specific user. This helps in tracking changes made by administrators or other users.

```sql+postgres
select
  id,
  name,
  last_modified_by,
  last_modified_ts
from
  prisma_permission_group
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
  prisma_permission_group
where
  last_modified_by = 'admin_user';
```

### Permission groups accepting resource lists
Retrieve permission groups where resource lists are accepted. This helps in understanding the configurations related to resource list acceptance in your permission groups.

```sql+postgres
select
  id,
  name,
  accept_resource_lists
from
  prisma_permission_group
where
  accept_resource_lists = true;
```

```sql+sqlite
select
  id,
  name,
  accept_resource_lists
from
  prisma_permission_group
where
  accept_resource_lists = 1;
```

### Get associated roles with the permission groups
Get a list of permission groups along with their associated roles. This can help in understanding the role assignments within your cloud environment.

```sql+postgres
select
  id,
  name,
  associated_roles
from
  prisma_permission_group;
```

```sql+sqlite
select
  id,
  name,
  associated_roles
from
  prisma_permission_group;
```