# 笔记接口

## 添加

`POST /note`

Param


Return

* code (*int*) 请求返回码，一般0表示正常，非0表示异常
* msg (*string*) 信息，一般是出错时的描述信息
* data (*object*) 
    * id (*int*) 新纪录ID

Error

* `10001` sql: no rows in result set
id: 119

<details>
<summary>Param</summary>

```json
{}
```

</details>

<details>
<summary>Return</summary>

```json
{
    "code": 10001,
    "msg": "sql: no rows in result set\nid: 119",
    "data": null
}
```

</details>

