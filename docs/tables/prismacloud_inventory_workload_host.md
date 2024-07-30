---
title: "Steampipe Table: prismacloud_inventory_workload_host - Query Prisma Cloud workload hosts using SQL"
description: "Allows users to query Prisma Cloud workload hosts. This table provides detailed information about hosts, including their unique identifiers, names, and vulnerability details. It can be used to monitor and manage hosts within Prisma Cloud."
---

# Table: prismacloud_inventory_workload_host - Query Prisma Cloud workload hosts using SQL

The Prisma Cloud workload host table in Steampipe provides you with comprehensive information about hosts within workloads in Prisma Cloud. This table allows you, as a security engineer or cloud administrator, to query host-specific details, including unique identifiers, names, and vulnerability details. You can utilize this table to gather insights on hosts, such as their unique identifiers, names, and vulnerability funnel details. The schema outlines the various attributes of the Prisma Cloud workload hosts for you.

## Table Usage Guide

The `prismacloud_inventory_workload_host` table in Steampipe provides detailed information about hosts within Prisma Cloud workloads. This table allows you to query details such as the host's unique identifier, name, title, and vulnerability funnel details, enabling you to manage and monitor your hosts effectively.

## Examples

### Basic Info
Retrieve basic information about Prisma Cloud workload hosts, such as their unique identifiers, names, and titles. This query helps you understand the overall configuration and status of your hosts.

```sql+postgres
select
  id,
  name,
  title
from
  prismacloud_inventory_workload_host;
```

```sql+sqlite
select
  id,
  name,
  title
from
  prismacloud_inventory_workload_host;
```

### Vulnerability funnel details
Get the vulnerability funnel details for each host. This is useful for understanding the security posture and vulnerability status of your hosts.

```sql+postgres
select
  id,
  name,
  vuln_funnel
from
  prismacloud_inventory_workload_host;
```

```sql+sqlite
select
  id,
  name,
  vuln_funnel
from
  prismacloud_inventory_workload_host;
```

### Hosts by unique identifier
Retrieve hosts based on their unique identifiers. This helps in identifying and managing specific hosts within your workloads.

```sql+postgres
select
  id,
  name,
  title
from
  prismacloud_inventory_workload_host
where
  id = 'us-binalyze-console.us-central1-a.c.chronicle-coe-351809.internal';
```

```sql+sqlite
select
  id,
  name,
  title
from
  prismacloud_inventory_workload_host
where
  id = 'us-binalyze-console.us-central1-a.c.chronicle-coe-351809.internal';
```

### Get the vulnerability funnel of workload hosts
Retrieve the vulnerability funnel details for each host, including the total number of vulnerabilities, urgent vulnerabilities, patchable vulnerabilities, exploitable vulnerabilities, and vulnerabilities in packages currently in use. This query helps in understanding the security posture and vulnerability status of your hosts.

```sql+postgres
select
  name,
  vuln_funnel ->> 'total' as total,
  vuln_funnel ->> 'urgent' as urgent,
  vuln_funnel ->> 'patchable' as patchable,
  vuln_funnel ->> 'exploitable' as exploitable,
  vuln_funnel ->> 'packageInUse' as package_in_use
from
  prismacloud_inventory_workload_host;
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
  prismacloud_inventory_workload_host;
```