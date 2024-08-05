---
title: "Steampipe Table: prismacloud_compliance_breakdown_requirement_summary - Query Prisma Cloud compliance breakdown requirement summaries using SQL"
description: "Allows users to query Prisma Cloud compliance breakdown requirement summaries. This table provides detailed information about the compliance status of resources with respect to specific requirements across different accounts, regions, and compliance standards."
---

# Table: prismacloud_compliance_breakdown_requirement_summary - Query Prisma Cloud compliance breakdown requirement summaries using SQL

The Prisma Cloud compliance breakdown requirement summary table in Steampipe provides you with detailed information about the compliance status of resources with respect to specific requirements within Prisma Cloud. This table allows security engineers and cloud administrators to query compliance breakdown requirement summaries, including section summaries, categorized by accounts, regions, and compliance standards. The schema outlines various attributes, such as account information, cloud type, compliance standard, and the number of resources with different compliance statuses.

## Table Usage Guide

The `prismacloud_compliance_breakdown_requirement_summary` table in Steampipe provides information about the compliance status of resources with respect to specific requirements within Prisma Cloud. This table allows you to query details such as the summaries of sections, categorized by accounts, regions, and compliance standards. This helps in managing and monitoring the compliance status of your cloud resources effectively.

**Important Notes**
- For improved performance, it is recommended to use the optional qualifiers (quals) to limit the result set.
- Queries with optional qualifiers are optimized to use filters. The following columns support optional qualifiers:
  - `account_name`
  - `cloud_type`
  - `cloud_region`
  - `policy_compliance_standard_name`
  - `policy_compliance_requirement_name`
  - `policy_compliance_section_id`

## Examples

### Basic info
Retrieve a basic summary of compliance breakdown requirements, including account name, cloud type, and the title of the section summary.

```sql+postgres
select
  account_name,
  cloud_type,
  title
from
  prismacloud_compliance_breakdown_requirement_summary;
```

```sql+sqlite
select
  account_name,
  cloud_type,
  title
from
  prismacloud_compliance_breakdown_requirement_summary;
```

### List compliance breakdown requirements by account and cloud type
Retrieve a summary of compliance breakdown requirements grouped by account name and cloud type. This query helps you to understand the compliance status of resources across different accounts and cloud environments.

```sql+postgres
select
  account_name,
  cloud_type,
  sum(section_summaries ->> 'critical_severity_failed_resources') as critical_failed_resources,
  sum(section_summaries ->> 'high_severity_failed_resources') as high_failed_resources,
  sum(section_summaries ->> 'medium_severity_failed_resources') as medium_failed_resources,
  sum(section_summaries ->> 'low_severity_failed_resources') as low_failed_resources,
  sum(section_summaries ->> 'informational_severity_failed_resources') as informational_failed_resources,
  sum(section_summaries ->> 'passed_resources') as passed_resources,
  sum(section_summaries ->> 'total_resources') as total_resources
from
  prismacloud_compliance_breakdown_requirement_summary
group by
  account_name,
  cloud_type;
```

```sql+sqlite
select
  account_name,
  cloud_type,
  sum(section_summaries ->> 'critical_severity_failed_resources') as critical_failed_resources,
  sum(section_summaries ->> 'high_severity_failed_resources') as high_failed_resources,
  sum(section_summaries ->> 'medium_severity_failed_resources') as medium_failed_resources,
  sum(section_summaries ->> 'low_severity_failed_resources') as low_failed_resources,
  sum(section_summaries ->> 'informational_severity_failed_resources') as informational_failed_resources,
  sum(section_summaries ->> 'passed_resources') as passed_resources,
  sum(section_summaries ->> 'total_resources') as total_resources
from
  prismacloud_compliance_breakdown_requirement_summary
group by
  account_name,
  cloud_type;
```

### Get breakdown requirement summaries for compliance standard and requirement
Retrieve detailed compliance breakdown requirement summaries by joining with the compliance standard and requirement tables. This query helps you to get a comprehensive view of the compliance status of resources, including the associated compliance standards and requirements.

```sql+postgres
select
  r.compliance_id as compliance_id,
  r.name as requirement_name,
  b.account_name,
  b.cloud_type,
  b.title,
  b.section_summaries
from
  prismacloud_compliance_breakdown_requirement_summary as b
  join prismacloud_compliance_requirement as r on b.policy_compliance_requirement_name = r.name;
```

```sql+sqlite
select
  r.compliance_id as compliance_id,
  r.name as requirement_name,
  b.account_name,
  b.cloud_type,
  b.title,
  b.section_summaries
from
  prismacloud_compliance_breakdown_requirement_summary as b
  join prismacloud_compliance_requirement as r on b.policy_compliance_requirement_name = r.name;
```

### Recently updated compliance breakdown section summaries
Retrieve compliance breakdown requirement summaries that were updated within the last 30 days. This query helps in tracking recent changes and understanding the current compliance status.

```sql+postgres
select
  account_name,
  cloud_type,
  title,
  section_summaries
from
  prismacloud_compliance_breakdown_requirement_summary
where
  sp_ctx ->> 'timestamp' > now() - interval '30 days';
```

```sql+sqlite
select
  account_name,
  cloud_type,
  title,
  section_summaries
from
  prismacloud_compliance_breakdown_requirement_summary
where
  json_extract(sp_ctx, '$.timestamp') > datetime('now', '-30 days');
```

### Get section summary details of breakdown
This helps in understanding the compliance status of different sections within each requirement, allowing for more granular analysis of compliance performance across different accounts and cloud environments.

```sql+postgres
select
  account_name,
  cloud_type,
  s ->> 'ID' as section_id,
  s ->> 'Name' as section_name,
  s ->> 'PassedResources' as passed_resources,
  s ->> 'FailedResources' as failed_resources
from
  prismacloud_compliance_breakdown_requirement_summary,
  jsonb_array_elements(section_summaries) as s;
```

```sql+sqlite
select
  account_name,
  cloud_type,
  json_extract(s.value, '$.ID') as section_id,
  json_extract(s.value, '$.Name') as section_name,
  json_extract(s.value, '$.PassedResources') as passed_resources,
  json_extract(s.value, '$.FailedResources') as failed_resources
from
  prismacloud_compliance_breakdown_requirement_summary,
  json_each(section_summaries) as s;
```