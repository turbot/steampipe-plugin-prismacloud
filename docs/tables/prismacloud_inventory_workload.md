---
title: "Steampipe Table: prismacloud_inventory_workload - Query Prisma Cloud workloads using SQL"
description: "Allows users to query Prisma Cloud workloads. This table provides detailed information about workloads, including container images, cloud providers, hosts, and their vulnerability status. It can be used to monitor and manage workloads within Prisma Cloud."
---

# Table: prismacloud_inventory_workload - Query Prisma Cloud workloads using SQL

The Prisma Cloud workload table in Steampipe provides you with comprehensive information about workloads within Prisma Cloud. This table allows you, as a security engineer or cloud administrator, to query workload-specific details, including container images, cloud providers, hosts, and their vulnerability status. You can utilize this table to gather insights on workloads, such as their stages, cloud providers, and vulnerabilities. The schema outlines the various attributes of the Prisma Cloud workloads for you, including the number of container images and hosts, and their respective vulnerability statuses.

## Table Usage Guide

The `prismacloud_inventory_workload` table in Steampipe provides detailed information about workloads within Prisma Cloud. This table allows you to query details such as the number of container images in different stages, cloud providers, hosts, and their vulnerability status, enabling you to manage and monitor your workloads effectively.

## Examples

### Basic Info
Retrieve basic information about Prisma Cloud workloads, such as the number of container images and hosts. This query helps you understand the overall configuration and status of your workloads.

```sql+postgres
select
  container_images_build,
  container_images_deploy,
  container_images_run,
  hosts_total
from
  prismacloud_inventory_workload;
```

```sql+sqlite
select
  container_images_build,
  container_images_deploy,
  container_images_run,
  hosts_total
from
  prismacloud_inventory_workload;
```

### Get host cloud providers
Get a list of cloud providers associated with container images. This is useful for understanding which cloud providers are being used for container images.

```sql+postgres
select
  container_images_cloud_providers,
  hosts_cloud_providers
from
  prismacloud_inventory_workload;
```

```sql+sqlite
select
  container_images_cloud_providers,
  hosts_cloud_providers
from
  prismacloud_inventory_workload;
```

### Vulnerable container images and hosts
Identify the number of vulnerable container images and hosts. This helps in assessing the security posture of your workloads.

```sql+postgres
select
  container_images_vulnerable,
  hosts_vulnerable
from
  prismacloud_inventory_workload;
```

```sql+sqlite
select
  container_images_vulnerable,
  hosts_vulnerable
from
  prismacloud_inventory_workload;
```