---
title: "Steampipe Table: prismacloud_inventory_workload_container_image - Query Prisma Cloud container images using SQL"
description: "Allows users to query Prisma Cloud container images. This table provides detailed information about container images, including their name, related images, running containers, scan status, and vulnerability details. It can be used to monitor and manage container images within Prisma Cloud."
---

# Table: prismacloud_inventory_workload_container_image - Query Prisma Cloud container images using SQL

The Prisma Cloud container image table in Steampipe provides you with comprehensive information about container images within workloads in Prisma Cloud. This table allows you, as a security engineer or cloud administrator, to query container image-specific details, including their name, related images, running containers, scan status, and vulnerability details. You can utilize this table to gather insights on container images, such as their stages, scan results, and vulnerability funnel details. The schema outlines the various attributes of the Prisma Cloud container images for you.

## Table Usage Guide

The `prismacloud_inventory_workload_container_image` table in Steampipe provides detailed information about container images within Prisma Cloud workloads. This table allows you to query details such as the container image's name, related images, running containers, scan status, and vulnerability funnel details, enabling you to manage and monitor your container images effectively.

## Examples

### Basic Info
Retrieve basic information about Prisma Cloud container images, such as their names and the number of related images. This query helps you understand the overall configuration and status of your container images.

```sql+postgres
select
  name,
  related_images
from
  prismacloud_inventory_workload_container_image;
```

```sql+sqlite
select
  name,
  related_images
from
  prismacloud_inventory_workload_container_image;
```

### Running containers of workload container images
Get the number of running containers for each container image. This is useful for understanding the deployment status of your container images.

```sql+postgres
select
  name,
  running_containers
from
  prismacloud_inventory_workload_container_image;
```

```sql+sqlite
select
  name,
  running_containers
from
  prismacloud_inventory_workload_container_image;
```

### Get scan status of workload container images
Identify whether the scan passed for each container image. This helps in assessing the security posture of your container images.

```sql+sqlite
select
  name,
  scan_passed
from
  prismacloud_inventory_workload_container_image;
```

### Running container counts for non-base images
Get the number of running containers for each non-base container image. This query helps in understanding the deployment status of your non-base container images.

```sql+postgres
select
  name,
  running_containers
from
  prismacloud_inventory_workload_container_image
where
  base = false;
```

```sql+sqlite
select
  name,
  running_containers
from
  prismacloud_inventory_workload_container_image
where
  base = false;
```

### Vulnerability funnel details
Get the vulnerability funnel details for each container image, including the total number of vulnerabilities, urgent vulnerabilities, patchable vulnerabilities, exploitable vulnerabilities, and vulnerabilities in packages currently in use.

```sql+postgres
select
  name,
  vuln_funnel ->> 'total' as total,
  vuln_funnel ->> 'urgent' as urgent,
  vuln_funnel ->> 'patchable' as patchable,
  vuln_funnel ->> 'exploitable' as exploitable,
  vuln_funnel ->> 'packageInUse' as package_in_use
from
  prismacloud_inventory_workload_container_image;
```

```sql+sqlite
select
  name,
  json_extract(vuln_funnel, '$.total') as total,
  json_extract(vuln_funnel, '$.urgent') as urgent,
  json_extract(vuln_funnel, '$.patchable') as patchable,
  json_extract(vuln_funnel, '$.exploitable') as exploitable,
  json_extract(vuln_funnel, '$.packageInUse') as package_in_use
from
  prismacloud_inventory_workload_container_image;
```