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
        "id": 35
    },
    "requestID": "6f23fc9c-e498-4bce-a955-7c0d9841bccf"
}
```

</details>

## 修改

`PUT /note`

Param

* noteID (*int*) 笔记ID
* title (*string*) 标题
* detail (*string*) 详情

Return

* code (*int*) 请求返回码，一般0表示正常，非0表示异常
* msg (*string*) 信息，一般是出错时的描述信息

<details>
<summary>Param</summary>

```json
{
    "noteID": 35,
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
    "requestID": "cd93cadd-d9e3-47db-b7b8-d3a81fbd4c95"
}
```

</details>

## 获取分页

`GET /note/page`

Param

* title (*string*) 
* detail (*string*) 
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
        * noteID (*int*) 
        * userName (*string*) 用户名
        * title (*string*) 
        * detail (*string*) 
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
        "total": 35,
        "list": [
            {
                "noteID": 35,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1599996545
            },
            {
                "noteID": 34,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1599996379
            },
            {
                "noteID": 33,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1599996159
            },
            {
                "noteID": 32,
                "userName": "0",
                "title": "",
                "detail": "",
                "createdAt": 1599996127
            },
            {
                "noteID": 31,
                "userName": "0",
                "title": "",
                "detail": "",
                "createdAt": 1599996091
            },
            {
                "noteID": 30,
                "userName": "0",
                "title": "",
                "detail": "",
                "createdAt": 1599995367
            },
            {
                "noteID": 29,
                "userName": "0",
                "title": "",
                "detail": "",
                "createdAt": 1599994007
            },
            {
                "noteID": 28,
                "userName": "0",
                "title": "",
                "detail": "",
                "createdAt": 1599993898
            },
            {
                "noteID": 27,
                "userName": "0",
                "title": "",
                "detail": "",
                "createdAt": 1599993806
            },
            {
                "noteID": 26,
                "userName": "0",
                "title": "",
                "detail": "",
                "createdAt": 1599993777
            }
        ]
    },
    "requestID": "42420104-442e-4d9c-84b6-b92ea381bb5c"
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
    * noteID (*int*) 
    * userName (*string*) 用户名
    * title (*string*) 
    * detail (*string*) 
    * createdAt (*int64*) 创建时间

<details>
<summary>Param</summary>

```json
noteID=35
```

</details>

<details>
<summary>Return</summary>

```json
{
    "code": 0,
    "msg": "",
    "data": {
        "noteID": 35,
        "userName": "1",
        "title": "mod title",
        "detail": "mod detail",
        "createdAt": 1599996545
    },
    "requestID": "e0673b11-de28-4cb2-83cd-46b43f794d65"
}
```

</details>

