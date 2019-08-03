# 添加用户接口

## 添加

`POST /auth/user`

Param

- Name (_string_)
- Password (_string_)

Return

- code (_int_) 请求返回码，一般 0 表示正常，非 0 表示异常
- msg (_string_) 信息，一般是出错时的描述信息
- Data (_int_)

<details>
<summary>Param</summary>

```json
{
  "Name": "jd",
  "Password": "jd"
}
```

</details>

<details>
<summary>Return</summary>

```json
{
  "code": 0,
  "msg": "",
  "data": 115
}
```

</details>
