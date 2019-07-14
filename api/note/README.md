# 笔记接口

## 添加

`POST /note`

Param

* title (*string*) 标题
* detail (*string*) 详情

Return

* code (*int*) 请求返回码，一般0表示正常，非0表示异常
* msg (*string*) 信息，一般是出错时的描述信息
* data (*int*) 

<details>
<summary>Param</summary>

```json
{
    "title": "test title",
    "detail": "test detail"
}
```

</details>

<details>
<summary>Return</summary>

```json
{
    "code": 0,
    "msg": "",
    "data": 82
}
```

</details>

