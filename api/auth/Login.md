# 登陆接口文档

## 登陆

`POST /login`

Param

* id (*int*) 记录ID
* name (*string*) 用户名
* phone (*string*) 手机号码
* email (*string*) 邮箱
* password (*string*) 密码

Return

* id (*int*) 记录ID
* name (*string*) 用户名
* phone (*string*) 手机号码
* email (*string*) 邮箱
* password (*string*) 密码

<details>
<summary>Param</summary>

```json
{
    "id": 0,
    "name": "jd",
    "phone": "",
    "email": "",
    "password": "13420693396"
}
```

</details>

<details>
<summary>Return</summary>

```json
{
    "code": 0,
    "msg": "",
    "data": null
}
```

</details>

