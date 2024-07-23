---
title: "Steampipe Table: prismacloud_iam_role - Query Prisma Cloud roles using SQL"
description: "Allows users to query Prisma Cloud roles. This table provides information about each role, including their name, description, associated users, and more. It can be used to monitor and manage roles within Prisma Cloud."
---

# Table: prismacloud_iam_role - Query Prisma Cloud roles using SQL

The Prisma Cloud role table in Steampipe provides you with information about roles within Prisma Cloud. This table allows you, as a security engineer or cloud administrator, to query role-specific details, including name, description, associated users, and more. You can utilize this table to gather insights on roles, such as their associated account groups, users, and more. The schema outlines the various attributes of the Prisma Cloud role for you, including the role's ID, name, and description.

## Table Usage Guide

The `prismacloud_iam_role` table in Steampipe provides information about roles within Prisma Cloud. This table allows you to query details such as the role's name, description, associated users, and more, enabling you to manage and monitor your roles effectively.

## Examples

### Basic Info
Retrieve basic information about Prisma Cloud roles, such as name, description, and role type. This query helps you to understand the overall configuration and details of your roles.

```sql+postgres
select
  name,
  description,
  role_type
from
  prismacloud_iam_role;
```

```sql+sqlite
select
  name,
  description,
  role_type
from
  prismacloud_iam_role;
```

### List of roles with their associated users
Get a list of all roles along with their associated users. This is useful for identifying which users are assigned to which roles.

```sql+postgres
select
  name,
  id,
  associated_users
from
  prismacloud_iam_role;
```

```sql+sqlite
select
  name,
  id,
  associated_users
from
  prismacloud_iam_role;
```

### Roles and their account groups
Identify roles along with their associated account groups. This helps in understanding role assignments and group memberships.

```sql+postgres
select
  name,
  account_group_ids
from
  prismacloud_iam_role;
```

```sql+sqlite
select
  name,
  account_group_ids
from
  prismacloud_iam_role;
```

### Recently modified roles
Retrieve roles that were modified recently. This helps in tracking changes and understanding recent modifications.

```sql+postgres
select
  name,
  last_modified_by,
  last_modified_ts
from
  prismacloud_iam_role
where
  last_modified_ts > extract(epoch from now()) - 604800;
```

```sql+sqlite
select
  name,
  last_modified_by,
  last_modified_ts
from
  prismacloud_iam_role
where
  last_modified_ts > strftime('%s','now') - 604800;
```

### Roles with restricted dismissal access
Get a list of roles that restrict dismissal access. This helps in understanding security policies and role permissions.

```sql+postgres
select
  name,
  restrict_dismissal_access
from
  prismacloud_iam_role
where
  restrict_dismissal_access = true;
```

```sql+sqlite
select
  name,
  restrict_dismissal_access
from
  prismacloud_iam_role
where
  restrict_dismissal_access = 1;
```