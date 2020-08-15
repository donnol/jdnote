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
        "id": 7
    },
    "requestID": "670b0038-52db-4b92-a03b-dd2f907f90f5"
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
    "noteID": 7,
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
    "requestID": "6b9a8c2f-bf63-4378-8624-97a24502536f"
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
        "total": 9,
        "list": [
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
            },
            {
                "noteID": 6,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1597480860
            },
            {
                "noteID": 5,
                "userName": "0",
                "title": "mod title",
                "detail": "# mod detail\n\n                  1213213\n\n## 12123\n\n#### 222\n\n####### 66666\n\n\n| a | b |\n| --- | --- |\n| 1 | 2 |\n\n\n\u003e\n\n| Syntax | Description |\n| --- | ----------- |\n| Header | Title |",
                "createdAt": 1596965769
            },
            {
                "noteID": 4,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1596964951
            },
            {
                "noteID": 3,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1596964517
            },
            {
                "noteID": 2,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1596962843
            },
            {
                "noteID": 1,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1596962237
            }
        ]
    },
    "requestID": "51988d51-83e0-4d02-91ad-a3f98e3869fd"
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
noteID=7
```

</details>

<details>
<summary>Return</summary>

```json
{
    "code": 0,
    "msg": "",
    "data": {
        "noteID": 7,
        "userName": "1",
        "title": "mod title",
        "detail": "mod detail",
        "createdAt": 1597481724
    },
    "requestID": "88c79af8-395a-40b1-8ec5-92047a18aeb3"
}
```

</details>

