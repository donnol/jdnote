# 数据库表定义

![ER图](./db_cicada.svg)

## t_userdata



| Field | Type | Nullable | Primary | Description |
| :-: | :-: | :-: | :-: | :-: |
| id | integer |  | * | 记录ID |
| name | text | * |  | 用户名 |
| phone | text | * |  | 手机号码 |
| email | text | * |  | 邮箱 |
| password | text | * |  | 密码 |





索引：


## t_roledata



| Field | Type | Nullable | Primary | Description |
| :-: | :-: | :-: | :-: | :-: |
| id | integer |  | * | 记录ID |
| role | text | * |  | 角色 |





索引：


## t_userroledata



| Field | Type | Nullable | Primary | Description |
| :-: | :-: | :-: | :-: | :-: |
| id | integer |  | * | 记录ID |
| user_id | integer | * |  | 用户ID |
| role_id | integer | * |  | 角色ID |





索引：


## t_actiondata



| Field | Type | Nullable | Primary | Description |
| :-: | :-: | :-: | :-: | :-: |
| id | integer |  | * | 记录ID |
| action | text | * |  | 操作 |





索引：


## t_roleactiondata



| Field | Type | Nullable | Primary | Description |
| :-: | :-: | :-: | :-: | :-: |
| id | integer |  | * | 记录ID |
| role_id | integer | * |  | 角色ID |
| action_id | integer | * |  | 动作ID |





索引：


