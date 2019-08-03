# 登陆接口文档

## 登陆

`POST /auth/login`

Param

- id (_int_) 记录 ID
- name (_string_) 用户名
- phone (_string_) 手机号码
- email (_string_) 邮箱
- password (_string_) 密码

Return

- code (_int_) 请求返回码，一般 0 表示正常，非 0 表示异常
- msg (_string_) 信息，一般是出错时的描述信息
- Data (_object_)
  - id (_int_) 记录 ID
  - name (_string_) 用户名
  - phone (_string_) 手机号码
  - email (_string_) 邮箱
  - password (_string_) 密码

<details>
<summary>Param</summary>

```json
{
  "id": 0,
  "name": "jd",
  "phone": "",
  "email": "",
  "password": "jd"
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
