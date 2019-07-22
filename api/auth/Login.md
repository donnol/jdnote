# 登陆接口文档

## 登陆

`POST /auth/login`

Param

* id (*int*) 记录ID
* name (*string*) 用户名
* phone (*string*) 手机号码
* email (*string*) 邮箱
* password (*string*) 密码

Return

* code (*int*) 请求返回码，一般0表示正常，非0表示异常
* msg (*string*) 信息，一般是出错时的描述信息
* Data (*object*) 
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
    "data": {
        "id": 114,
        "name": "jd",
        "phone": "",
        "email": "",
        "password": ""
    }
}
```

</details>

