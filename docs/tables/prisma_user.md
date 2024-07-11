---
title: "Steampipe Table: prisma_user - Query Prisma Cloud users using SQL"
description: "Allows users to query Prisma Cloud users. This table provides information about each user, including their username, roles, access keys, and more. It can be used to monitor and manage users within Prisma Cloud."
---

# Table: prisma_user - Query Prisma Cloud users using SQL

The Prisma Cloud user table in Steampipe provides you with information about users within Prisma Cloud. This table allows you, as a security engineer or cloud administrator, to query user-specific details, including username, roles, access keys, and more. You can utilize this table to gather insights on users, such as their access key status, roles, and more. The schema outlines the various attributes of the Prisma Cloud user for you, including the user's ID, name, and email address.

## Table Usage Guide

The `prisma_user` table in Steampipe provides information about users within Prisma Cloud. This table allows you to query details such as the user's username, roles, access keys, and more, enabling you to manage and monitor your users effectively.

## Examples

### Basic Info
Retrieve basic information about Prisma Cloud users, such as username, email, and roles. This query helps you to understand the overall configuration and details of your users.

```sql+postgres
select
  username,
  email,
  roles
from
  prisma_user;
```

```sql+sqlite
select
  username,
  email,
  roles
from
  prisma_user;
```

### List of enabled users
Get a list of all enabled Prisma Cloud users. This is useful for identifying which user accounts are currently active.

```sql+postgres
select
  username,
  email,
  enabled
from
  prisma_user
where
  enabled = true;
```

```sql+sqlite
select
  username,
  email,
  enabled
from
  prisma_user
where
  enabled = 1;
```

### Users with access keys allowed
Identify users that are allowed to have access keys. This helps in tracking access key usage and permissions.

```sql+postgres
select
  username,
  email,
  access_keys_allowed
from
  prisma_user
where
  access_keys_allowed = true;
```

```sql+sqlite
select
  username,
  email,
  access_keys_allowed
from
  prisma_user
where
  access_keys_allowed = 1;
```

### Users with expiring access keys
Retrieve users whose access keys are expiring soon. This helps in managing and rotating access keys proactively.

```sql+postgres
select
  username,
  email,
  access_key_name,
  access_key_expiration
from
  prisma_user
where
  access_key_expiration < extract(epoch from now()) + 2592000;
```

```sql+sqlite
select
  username,
  email,
  access_key_name,
  access_key_expiration
from
  prisma_user
where
  access_key_expiration < strftime('%s','now') + 2592000;
```

### Users and their roles
Get a list of users along with their assigned roles. This can help in understanding the role assignments within your cloud environment.

```sql+postgres
select
  username,
  email,
  role_ids
from
  prisma_user;
```

```sql+sqlite
select
  username,
  email,
  role_ids
from
  prisma_user;
```