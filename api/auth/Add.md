# 添加用户接口

## 添加

`POST /auth/user`

Param

* Name (*string*) 
* Password (*string*) 

Return

* code (*int*) 请求返回码，一般0表示正常，非0表示异常
* msg (*string*) 信息，一般是出错时的描述信息
* Data (*int*) 

<details>
<summary>Param</summary>

```json
{
    "Name": "jd",
    "Password": "13420693396"
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

