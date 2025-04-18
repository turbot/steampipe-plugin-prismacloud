## v1.1.1 [2025-04-18]

_Bug fixes_

- Fixed Linux AMD64 plugin build failures for `Postgres 14 FDW`, `Postgres 15 FDW`, and `SQLite Extension` by upgrading GitHub Actions runners from `ubuntu-20.04` to `ubuntu-22.04`.

## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#17](https://github.com/turbot/steampipe-plugin-prismacloud/pull/17))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#17](https://github.com/turbot/steampipe-plugin-prismacloud/pull/17))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

## v0.1.0 [2024-09-19]

_Bug fixes_

- Fixed the typo in the `prismacloud_vulnerability*` tables. ([#13](https://github.com/turbot/steampipe-plugin-prismacloud/pull/13))

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#11](https://github.com/turbot/steampipe-plugin-prismacloud/pull/11))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#11](https://github.com/turbot/steampipe-plugin-prismacloud/pull/11))

## v0.0.2 [2024-08-09]

_Bug fixes_

- Fixed the plugin's brand color.

## v0.0.1 [2024-08-09]

_What's new?_

- New tables added
  - [prismacloud_account](https://hub.steampipe.io/plugins/turbot/prismacloud/tables/prismacloud_account)
  - [prismacloud_alert](https://hub.steampipe.io/plugins/turbot/prismacloud/tables/prismacloud_alert)
  - [prismacloud_alert_rule](https://hub.steampipe.io/plugins/turbot/prismacloud/tables/prismacloud_alert_rule)
  - [prismacloud_compliance_breakdown_requirement_summary](https://hub.steampipe.io/plugins/turbot/prismacloud/tables/prismacloud_compliance_breakdown_requirement_summary)
  - [prismacloud_compliance_breakdown_statistic](https://hub.steampipe.io/plugins/turbot/prismacloud/tables/prismacloud_compliance_breakdown_statistic)
  - [prismacloud_compliance_breakdown_summary](https://hub.steampipe.io/plugins/turbot/prismacloud/tables/prismacloud_compliance_breakdown_summary)
  - [prismacloud_compliance_requirement](https://hub.steampipe.io/plugins/turbot/prismacloud/tables/prismacloud_compliance_requirement)
  - [prismacloud_compliance_standard](https://hub.steampipe.io/plugins/turbot/prismacloud/tables/prismacloud_compliance_standard)
  - [prismacloud_iam_permission](https://hub.steampipe.io/plugins/turbot/prismacloud/tables/prismacloud_iam_permission)
  - [prismacloud_iam_role](https://hub.steampipe.io/plugins/turbot/prismacloud/tables/prismacloud_iam_role)
  - [prismacloud_iam_user](https://hub.steampipe.io/plugins/turbot/prismacloud/tables/prismacloud_iam_user)
  - [prismacloud_inventory_api_endpoint](https://hub.steampipe.io/plugins/turbot/prismacloud/tables/prismacloud_inventory_api_endpoint)
  - [prismacloud_inventory_asset_explorer](https://hub.steampipe.io/plugins/turbot/prismacloud/tables/prismacloud_inventory_asset_explorer)
  - [prismacloud_inventory_asset_view](https://hub.steampipe.io/plugins/turbot/prismacloud/tables/prismacloud_inventory_asset_view)
  - [prismacloud_inventory_workload](https://hub.steampipe.io/plugins/turbot/prismacloud/tables/prismacloud_inventory_workload)
  - [prismacloud_inventory_workload_container_image](https://hub.steampipe.io/plugins/turbot/prismacloud/tables/prismacloud_inventory_workload_container_image)
  - [prismacloud_inventory_workload_host](https://hub.steampipe.io/plugins/turbot/prismacloud/tables/prismacloud_inventory_workload_host)
  - [prismacloud_permission_group](https://hub.steampipe.io/plugins/turbot/prismacloud/tables/prismacloud_permission_group)
  - [prismacloud_policy](https://hub.steampipe.io/plugins/turbot/prismacloud/tables/prismacloud_policy)
  - [prismacloud_prioritized_vulnerability](https://hub.steampipe.io/plugins/turbot/prismacloud/tables/prismacloud_prioritized_vulnerability)
  - [prismacloud_report](https://hub.steampipe.io/plugins/turbot/prismacloud/tables/prismacloud_report)
  - [prismacloud_resource](https://hub.steampipe.io/plugins/turbot/prismacloud/tables/prismacloud_resource)
  - [prismacloud_trusted_alert_ip](https://hub.steampipe.io/plugins/turbot/prismacloud/tables/prismacloud_trusted_alert_ip)
  - [prismacloud_vulnerability_asset](https://hub.steampipe.io/plugins/turbot/prismacloud/tables/prismacloud_vulnerability_asset)
  - [prismacloud_vulnerability_burndown](https://hub.steampipe.io/plugins/turbot/prismacloud/tables/prismacloud_vulnerability_burndown)
  - [prismacloud_vulnerability_overview](https://hub.steampipe.io/plugins/turbot/prismacloud/tables/prismacloud_vulnerability_overview)
