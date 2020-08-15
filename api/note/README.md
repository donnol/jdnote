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
        "id": 16
    },
    "requestID": "35eb9d01-0431-4fd0-a7b2-d67c7c87b699"
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
    "noteID": 16,
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
    "requestID": "adfc0af1-6735-42c1-b5be-b67983ffe199"
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
        "total": 16,
        "list": [
            {
                "noteID": 16,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1597493300
            },
            {
                "noteID": 15,
                "userName": "0",
                "title": "",
                "detail": "",
                "createdAt": 1597493299
            },
            {
                "noteID": 14,
                "userName": "0",
                "title": "",
                "detail": "testDetail",
                "createdAt": 1597493299
            },
            {
                "noteID": 13,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1597490651
            },
            {
                "noteID": 12,
                "userName": "0",
                "title": "",
                "detail": "",
                "createdAt": 1597490650
            },
            {
                "noteID": 11,
                "userName": "0",
                "title": "",
                "detail": "testDetail",
                "createdAt": 1597490650
            },
            {
                "noteID": 10,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1597488159
            },
            {
                "noteID": 9,
                "userName": "0",
                "title": "",
                "detail": "",
                "createdAt": 1597481727
            },
            {
                "noteID": 8,
                "userName": "0",
                "title": "",
                "detail": "testDetail",
                "createdAt": 1597481727
            },
            {
                "noteID": 7,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1597481724
            }
        ]
    },
    "requestID": "e5e4051e-4931-452b-8fe9-ae68c4833306"
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
noteID=16
```

</details>

<details>
<summary>Return</summary>

```json
{
    "code": 0,
    "msg": "",
    "data": {
        "noteID": 16,
        "userName": "1",
        "title": "mod title",
        "detail": "mod detail",
        "createdAt": 1597493300
    },
    "requestID": "2cd4bdba-f45e-4a53-9ab0-d636e71bfff3"
}
```

</details>

