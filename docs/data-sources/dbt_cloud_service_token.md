---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dbt_cloud_service_token Data Source - terraform-provider-dbt-cloud"
subcategory: ""
description: |-
  
---

# dbt_cloud_service_token (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `service_token_id` (Number) ID of the service token

### Read-Only

- `id` (String) The ID of this resource.
- `name` (String) Service token name
- `service_token_permissions` (Set of Object) Permissions set for the service token (see [below for nested schema](#nestedatt--service_token_permissions))
- `uid` (String) The UID of the service token (part of the token secret)

<a id="nestedatt--service_token_permissions"></a>
### Nested Schema for `service_token_permissions`

Read-Only:

- `all_projects` (Boolean)
- `permission_set` (String)
- `project_id` (Number)


