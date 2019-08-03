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
        "id": 94
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
    "id": 94,
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
        * noteID (*int*) 笔记ID
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
        "total": 63,
        "list": [
            {
                "noteID": 94,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1564811225
            },
            {
                "noteID": 93,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1564811206
            },
            {
                "noteID": 92,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1564806878
            },
            {
                "noteID": 91,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1564806333
            },
            {
                "noteID": 90,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1564806204
            },
            {
                "noteID": 89,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1563198544
            },
            {
                "noteID": 88,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1563198415
            },
            {
                "noteID": 87,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1563198187
            },
            {
                "noteID": 86,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1563198153
            },
            {
                "noteID": 85,
                "userName": "0",
                "title": "",
                "detail": "",
                "createdAt": 1563198010
            }
        ]
    }
}
```

</details>

