---
title: "Steampipe Table: prismacloud_trusted_alert_ip - Query Prisma Cloud trusted alert IPs using SQL"
description: "Allows users to query Prisma Cloud trusted alert IPs. This table provides information about each trusted alert IP, including its name, CIDR blocks, and more. It can be used to monitor and manage trusted alert IPs within Prisma Cloud."
---

# Table: prismacloud_trusted_alert_ip - Query Prisma Cloud trusted alert IPs using SQL

The Prisma Cloud trusted alert IP table in Steampipe provides you with information about trusted alert IPs within Prisma Cloud. This table allows you, as a security engineer or cloud administrator, to query trusted alert IP-specific details, including the name, CIDR blocks, and more. You can utilize this table to gather insights on trusted alert IPs, such as their configurations and associated details. The schema outlines the various attributes of the Prisma Cloud trusted alert IP for you, including the trusted alert IP's name, CIDR blocks, and unique identifier.

## Table Usage Guide

The `prismacloud_trusted_alert_ip` table in Steampipe provides information about trusted alert IPs within Prisma Cloud. This table allows you to query details such as the trusted alert IP's name, CIDR blocks, and more, enabling you to manage and monitor your trusted alert IPs effectively.

## Examples

### Basic Info
Retrieve basic information about Prisma Cloud trusted alert IPs, such as the name, CIDR count, and UUID. This query helps you to understand the overall configuration and details of your trusted alert IPs.

```sql+postgres
select
  name,
  cidr_count,
  uuid
from
  prismacloud_trusted_alert_ip;
```

```sql+sqlite
select
  name,
  cidr_count,
  uuid
from
  prismacloud_trusted_alert_ip;
```

### Detailed CIDR Info
Retrieve detailed information about the CIDR blocks for each trusted alert IP, including the CIDR block, UUID, creation time, and description.

```sql+postgres
select
  name,
  c ->> 'CIDR' as cidr_block,
  c ->> 'UUID' as cidr_uuid,
  c ->> 'CreatedOn' as create_time_in_ms,
  c ->> 'Description' as description
from
  prismacloud_trusted_alert_ip,
  jsonb_array_elements(cidrs) as c;
```

```sql+sqlite
select
  name,
  json_extract(c.value, '$.CIDR') as cidr_block,
  json_extract(c.value, '$.UUID') as cidr_uuid,
  json_extract(c.value, '$.CreatedOn') as create_time_in_ms,
  json_extract(c.value, '$.Description') as description
from
  prismacloud_trusted_alert_ip,
  json_each(cidrs) as c;
```

### List trusted alert IP by Name
Get a list of all trusted alert IPs filtered by a specific name. This is useful for identifying and managing trusted alert IPs based on their names.

```sql+postgres
select
  name,
  uuid,
  cidr_count
from
  prismacloud_trusted_alert_ip
where
  name = 'example_name';
```

```sql+sqlite
select
  name,
  uuid,
  cidr_count
from
  prismacloud_trusted_alert_ip
where
  name = 'example_name';
```

### Count of CIDRs per trusted alert IP
Retrieve the count of CIDR blocks for each trusted alert IP. This helps in understanding the distribution of CIDR blocks across different trusted alert IPs.

```sql+postgres
select
  name,
  cidr_count
from
  prismacloud_trusted_alert_ip;
```

```sql+sqlite
select
  name,
  cidr_count
from
  prismacloud_trusted_alert_ip;
```