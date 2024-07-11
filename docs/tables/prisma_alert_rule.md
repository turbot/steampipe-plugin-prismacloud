---
title: "Steampipe Table: prisma_alert_rule - Query Prisma Cloud alert rules using SQL"
description: "Allows users to query Prisma Cloud alert rules. This table provides information about each alert rule, including its name, status, notification settings, and more. It can be used to monitor and manage alert rules within Prisma Cloud."
---

# Table: prisma_alert_rule - Query Prisma Cloud alert rules using SQL

The Prisma Cloud alert rule table in Steampipe provides you with information about alert rules within Prisma Cloud. This table allows you, as a security engineer or cloud administrator, to query alert rule-specific details, including rule name, status, notification settings, and more. You can utilize this table to gather insights on alert rules, such as their configurations, status, and more. The schema outlines the various attributes of the Prisma Cloud alert rule for you, including the rule ID, name, notification settings, and associated policies.

## Table Usage Guide

The `prisma_alert_rule` table in Steampipe provides information about alert rules within Prisma Cloud. This table allows you to query details such as the alert rule's name, status, notification settings, and more, enabling you to manage and monitor your alert rules effectively.

## Examples

### Basic Info
Retrieve basic information about Prisma Cloud alert rules, such as rule name, description, status, and whether auto-remediation is allowed. This query helps you to understand the overall configuration and status of your alert rules.

```sql+postgres
select
  name,
  description,
  enabled,
  allow_auto_remediate
from
  prisma_alert_rule;
```

```sql+sqlite
select
  name,
  description,
  enabled,
  allow_auto_remediate
from
  prisma_alert_rule;
```

### List of enabled alert rules
Get a list of all enabled Prisma Cloud alert rules. This is useful for identifying which alert rules are currently active and enabled.

```sql+postgres
select
  name,
  description,
  enabled
from
  prisma_alert_rule
where
  enabled = true;
```

```sql+sqlite
select
  name,
  description,
  enabled
from
  prisma_alert_rule
where
  enabled = 1;
```

### Alert rules with specific notification settings
Identify alert rules that notify on dismissal. This helps in understanding the notification settings configured for your alert rules.

```sql+postgres
select
  name,
  description,
  notify_on_dismissed
from
  prisma_alert_rule
where
  notify_on_dismissed = true;
```

```sql+sqlite
select
  name,
  description,
  notify_on_dismissed
from
  prisma_alert_rule
where
  notify_on_dismissed = 1;
```

### Alert rules by owner
Retrieve alert rules managed by a specific owner. This helps in tracking which alert rules are managed by which users or teams.

```sql+postgres
select
  name,
  description,
  owner
from
  prisma_alert_rule
where
  owner = 'admin_user';
```

```sql+sqlite
select
  name,
  description,
  owner
from
  prisma_alert_rule
where
  owner = 'admin_user';
```

### Alert rules with open alerts
Get a list of alert rules with open alerts. This helps in identifying which alert rules have ongoing issues that need attention.

```sql+postgres
select
  name,
  description,
  open_alerts_count
from
  prisma_alert_rule
where
  open_alerts_count > 0;
```

```sql+sqlite
select
  name,
  description,
  open_alerts_count
from
  prisma_alert_rule
where
  open_alerts_count > 0;
```