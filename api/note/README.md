# 笔记接口

## 添加

`POST /note`

Param


Return

* code (*int*) 请求返回码，一般0表示正常，非0表示异常
* msg (*string*) 信息，一般是出错时的描述信息
* data (*object*) 
    * id (*int*) 新纪录ID

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
    "code": 0,
    "msg": "",
    "data": {
        "id": 92
    }
}
```

</details>

## 修改

`PUT /note`

Param

* id (*int*) 记录ID
* title (*string*) 标题
* detail (*string*) 详情

Return

* code (*int*) 请求返回码，一般0表示正常，非0表示异常
* msg (*string*) 信息，一般是出错时的描述信息

<details>
<summary>Param</summary>

```json
{
    "id": 92,
    "title": "mod title",
    "detail": "mod detail"
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

## 获取分页

`GET /note/page`

Param

* start (*int*) 分页开始
* size (*int*) 分页大小

Return

* code (*int*) 请求返回码，一般0表示正常，非0表示异常
* msg (*string*) 信息，一般是出错时的描述信息
* data (*object*) 
    * total (*int*) 总数
    * list (*object list*) 列表
        * userName (*string*) 用户名
        * title (*string*) 标题
        * detail (*string*) 详情
        * createdAt (*int64*) 创建时间

<details>
<summary>Param</summary>

```json
size=10&start=0
```

</details>

<details>
<summary>Return</summary>

```json
{
    "code": 0,
    "msg": "",
    "data": {
        "total": 61,
        "list": [
            {
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1564806878
            },
            {
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1564806333
            },
            {
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1564806204
            },
            {
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1563198544
            },
            {
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1563198415
            },
            {
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1563198187
            },
            {
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1563198153
            },
            {
                "userName": "0",
                "title": "",
                "detail": "",
                "createdAt": 1563198010
            },
            {
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1563196739
            },
            {
                "userName": "0",
                "title": "test title",
                "detail": "test detail",
                "createdAt": 1563111143
            }
        ]
    }
}
```

</details>

