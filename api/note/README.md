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
        "id": 19
    },
    "requestID": "bdd5c58f-c3c0-43e0-9d96-8407f324ae05"
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
    "noteID": 19,
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
    "requestID": "5d8928fe-5347-4ab7-a462-37d2f8d85e8a"
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
        "total": 19,
        "list": [
            {
                "noteID": 19,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1599980393
            },
            {
                "noteID": 18,
                "userName": "0",
                "title": "天意难违",
                "detail": "# 天\n\n## 意\n\n## 难\n\n## 违\n\n",
                "createdAt": 1598669851
            },
            {
                "noteID": 17,
                "userName": "0",
                "title": "双双双",
                "detail": "# 双\n\n## 哈哈\n\n学用，读写，进出，买卖。\n\n品味决定。\n\n虚无陷阱。",
                "createdAt": 1598363979
            },
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
            }
        ]
    },
    "requestID": "a41fe183-1160-4fcf-aff4-a5f44a6d5701"
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
noteID=19
```

</details>

<details>
<summary>Return</summary>

```json
{
    "code": 0,
    "msg": "",
    "data": {
        "noteID": 19,
        "userName": "1",
        "title": "mod title",
        "detail": "mod detail",
        "createdAt": 1599980393
    },
    "requestID": "1bb5abb7-969a-4f6d-af57-68e9f6dee2ad"
}
```

</details>

