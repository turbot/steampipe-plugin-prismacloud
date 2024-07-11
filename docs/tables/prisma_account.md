---
title: "Steampipe Table: prisma_account - Query Prisma Cloud accounts using SQL"
description: "Allows users to query Prisma Cloud accounts. This table provides information about each account, including the cloud type, account status, and more. It can be used to monitor account details, status, and modifications."
---

# Table: prisma_account - Query Prisma Cloud accounts using SQL

The Prisma Cloud account table in Steampipe provides you with information about accounts within Prisma Cloud. This table allows you, as a security engineer or cloud administrator, to query account-specific details, including account type, status, cloud type, and more. You can utilize this table to gather insights on accounts, such as account status, last modification details, and more. The schema outlines the various attributes of the Prisma Cloud account for you, including the account ID, name, and associated groups.

## Table Usage Guide

The `prisma_account` table in Steampipe provides information about accounts within Prisma Cloud. This table allows you to query details such as the account's cloud type, status, and more, enabling you to manage and monitor your cloud accounts effectively.

## Examples

### Basic Info
Retrieve basic information about Prisma Cloud accounts, such as account ID, name, cloud type, and status. This query helps you to understand the overall configuration and status of your accounts.

```sql+postgres
select
  account_id,
  name,
  cloud_type,
  status,
  enabled
from
  prisma_account;
```

```sql+sqlite
select
  account_id,
  name,
  cloud_type,
  status,
  enabled
from
  prisma_account;
```

### List of enabled accounts
Get a list of all enabled Prisma Cloud accounts. This is useful for identifying which accounts are currently active and enabled.

```sql+postgres
select
  account_id,
  name,
  cloud_type,
  status
from
  prisma_account
where
  enabled = true;
```

```sql+sqlite
select
  account_id,
  name,
  cloud_type,
  status
from
  prisma_account
where
  enabled = 1;
```

### Accounts modified by a specific user
Identify accounts that were last modified by a specific user. This helps in tracking changes made by administrators or other users.

```sql+postgres
select
  account_id,
  name,
  last_modified_by,
  last_modified_ts
from
  prisma_account
where
  last_modified_by = 'admin_user';
```

```sql+sqlite
select
  account_id,
  name,
  last_modified_by,
  last_modified_ts
from
  prisma_account
where
  last_modified_by = 'admin_user';
```

### List accounts with storage scan enabled
Retrieve accounts where storage scan is enabled. This is useful for ensuring that storage scanning is properly configured for security purposes.

```sql+postgres
select
  account_id,
  name,
  storage_scan_enabled
from
  prisma_account
where
  storage_scan_enabled = true;
```

```sql+sqlite
select
  account_id,
  name,
  storage_scan_enabled
from
  prisma_account
where
  storage_scan_enabled = 1;
```

### List accounts and their groups
Get a list of accounts along with their associated groups. This can help in understanding the organizational structure and group assignments within your cloud environment.

```sql+postgres
select
  account_id,
  name,
  group_ids,
  groups
from
  prisma_account;
```

```sql+sqlite
select
  account_id,
  name,
  group_ids,
  groups
from
  prisma_account;
```