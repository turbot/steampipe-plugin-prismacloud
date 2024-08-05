---
title: "Steampipe Table: prismacloud_iam_permission - Query Prisma Cloud IAM permissions using SQL"
description: "Allows users to query Prisma Cloud IAM permissions. This table provides detailed information about IAM permissions, including the resources accessed, actions taken, and the entities that granted the permissions."
---

# Table: prismacloud_iam_permission - Query Prisma Cloud IAM permissions using SQL

The Prisma Cloud IAM permission table in Steampipe provides comprehensive information about IAM permissions within Prisma Cloud. This table allows security engineers and cloud administrators to query detailed IAM permission data, including accessed resources, granted permissions, and the associated cloud accounts and regions. The schema outlines various attributes such as resource information, effective actions, exceptions, and the entities involved in granting permissions.

## Table Usage Guide

The `prismacloud_iam_permission` table in Steampipe provides detailed information about IAM permissions within Prisma Cloud. This table allows you to query specifics such as accessed resources, actions taken, and the entities that granted the permissions. This helps in managing and monitoring IAM permissions across your cloud environment effectively.

## Examples

### Basic info
Retrieve basic information about IAM permissions, including resource names, cloud types, and effective actions.

```sql+postgres
select
  dest_resource_name,
  cloud_type,
  effective_action_name
from
  prismacloud_iam_permission;
```

```sql+sqlite
select
  dest_resource_name,
  cloud_type,
  effective_action_name
from
  prismacloud_iam_permission;
```

### List permissions by cloud account and region
Retrieve IAM permissions grouped by cloud account and region. This helps in understanding the distribution of permissions across different accounts and regions.

```sql+postgres
select
  account_name,
  cloud_region,
  count(*) as permission_count
from
  prismacloud_iam_permission
group by
  account_name,
  cloud_region;
```

```sql+sqlite
select
  account_name,
  cloud_region,
  count(*) as permission_count
from
  prismacloud_iam_permission
group by
  account_name,
  cloud_region;
```

### Permissions with exceptions
Identify IAM permissions that have exceptions. This helps in understanding any potential issues or constraints related to permissions.

```sql+postgres
select
  dest_resource_name,
  cloud_type,
  effective_action_name,
  exceptions
from
  prismacloud_iam_permission
where
  exceptions is not null;
```

```sql+sqlite
select
  dest_resource_name,
  cloud_type,
  effective_action_name,
  exceptions
from
  prismacloud_iam_permission
where
  exceptions is not null;
```

### Recently accessed resources
Retrieve IAM permissions for resources that were accessed within the last 30 days. This helps in tracking recent activity and understanding the current usage of permissions.

```sql+postgres
select
  dest_resource_name,
  cloud_type,
  last_access_date,
  last_access_status
from
  prismacloud_iam_permission
where
  last_access_date > now() - interval '30 days';
```

```sql+sqlite
select
  dest_resource_name,
  cloud_type,
  last_access_date,
  last_access_status
from
  prismacloud_iam_permission
where
  last_access_date > datetime('now', '-30 days');
```

### Permissions by effective action name
Get a list of IAM permissions grouped by effective action name. This helps in understanding the types of actions that are being performed with the granted permissions.

```sql+postgres
select
  effective_action_name,
  count(*) as permission_count
from
  prismacloud_iam_permission
group by
  effective_action_name;
```

```sql+sqlite
select
  effective_action_name,
  count(*) as permission_count
from
  prismacloud_iam_permission
group by
  effective_action_name;
```

### List permissions by resource type
Retrieve IAM permissions grouped by resource type. This helps in understanding the distribution of permissions across different resource types.

```sql+postgres
select
  dest_resource_type,
  count(*) as permission_count
from
  prismacloud_iam_permission
group by
  dest_resource_type;
```

```sql+sqlite
select
  dest_resource_type,
  count(*) as permission_count
from
  prismacloud_iam_permission
group by
  dest_resource_type;
```

### Permissions granted by specific entity
Get a list of IAM permissions granted by a specific entity. This helps in understanding which entities are granting permissions and to whom.

```sql+postgres
select
  granted_by_cloud_entity_name,
  dest_resource_name,
  effective_action_name
from
  prismacloud_iam_permission
where
  granted_by_cloud_entity_name = 'specific_entity_name';
```

```sql+sqlite
select
  granted_by_cloud_entity_name,
  dest_resource_name,
  effective_action_name
from
  prismacloud_iam_permission
where
  granted_by_cloud_entity_name = 'specific_entity_name';
```

### List permissions by source cloud type
Retrieve IAM permissions grouped by the source cloud type. This helps in understanding the distribution of permissions across different cloud environments.

```sql+postgres
select
  source_cloud_type,
  count(*) as permission_count
from
  prismacloud_iam_permission
group by
  source_cloud_type;
```

```sql+sqlite
select
  source_cloud_type,
  count(*) as permission_count
from
  prismacloud_iam_permission
group by
  source_cloud_type;
```

### Permissions by policy name
Get a list of IAM permissions associated with a specific policy name. This helps in understanding the impact of specific policies on resource access.

```sql+postgres
select
  granted_by_cloud_policy_name,
  dest_resource_name,
  effective_action_name
from
  prismacloud_iam_permission
where
  granted_by_cloud_policy_name = 'specific_policy_name';
```

```sql+sqlite
select
  granted_by_cloud_policy_name,
  dest_resource_name,
  effective_action_name
from
  prismacloud_iam_permission
where
  granted_by_cloud_policy_name = 'specific_policy_name';
```