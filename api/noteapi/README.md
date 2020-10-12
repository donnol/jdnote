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
        "id": 50
    },
    "requestID": "902a8ede-5056-4141-a1a2-9f96474c19bb"
}
```

</details>

## 修改

`PUT /note`

Param

* noteID (*int*) 记录ID
* title (*string*) 标题
* detail (*string*) 详情

Return

* code (*int*) 请求返回码，一般0表示正常，非0表示异常
* msg (*string*) 信息，一般是出错时的描述信息

<details>
<summary>Param</summary>

```json
{
    "noteID": 50,
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
    "data": null,
    "requestID": "1f05f624-ce3b-473f-b24d-c8e8750c1ed9"
}
```

</details>

## 获取分页

`GET /note/page`

Param

* title (*string*) 标题
* detail (*string*) 详情
* beginTime (*int64*) 开始时间
* endTime (*int64*) 结束时间
* pageIndex (*int*) 分页开始
* pageSize (*int*) 分页大小

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
beginTime=0&detail=&endTime=0&pageIndex=0&pageSize=10&title=
```

</details>

<details>
<summary>Return</summary>

```json
{
    "code": 0,
    "msg": "",
    "data": {
        "total": 50,
        "list": [
            {
                "noteID": 50,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1602513692
            },
            {
                "noteID": 49,
                "userName": "0",
                "title": "",
                "detail": "",
                "createdAt": 1602512947
            },
            {
                "noteID": 48,
                "userName": "0",
                "title": "",
                "detail": "",
                "createdAt": 1602510874
            },
            {
                "noteID": 47,
                "userName": "0",
                "title": "",
                "detail": "",
                "createdAt": 1602510703
            },
            {
                "noteID": 46,
                "userName": "0",
                "title": "",
                "detail": "",
                "createdAt": 1602257336
            },
            {
                "noteID": 45,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1602256903
            },
            {
                "noteID": 44,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1602256700
            },
            {
                "noteID": 43,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1602256317
            },
            {
                "noteID": 42,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1602253486
            },
            {
                "noteID": 41,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1602253166
            }
        ]
    },
    "requestID": "5f7b31ac-0ab2-4fdd-b33d-357ff60aff05"
}
```

</details>

## 获取详情

`GET /note`

Param

* noteID (*int*) 

Return

* code (*int*) 请求返回码，一般0表示正常，非0表示异常
* msg (*string*) 信息，一般是出错时的描述信息
* data (*object*) 
    * noteID (*int*) 笔记ID
    * userName (*string*) 用户名
    * title (*string*) 标题
    * detail (*string*) 详情
    * createdAt (*int64*) 创建时间

<details>
<summary>Param</summary>

```json
noteID=50
```

</details>

<details>
<summary>Return</summary>

```json
{
    "code": 0,
    "msg": "",
    "data": {
        "noteID": 50,
        "userName": "1",
        "title": "mod title",
        "detail": "mod detail",
        "createdAt": 1602513692
    },
    "requestID": "85aae7ea-5d24-439f-aabe-42696aa8a515"
}
```

</details>

