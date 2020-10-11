# 获取用户接口

## 获取

`GET /auth/user`

Param

* Name (*string*) 

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
name=jd
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

