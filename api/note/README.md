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
        "id": 13
    },
    "requestID": "e1311f98-dffb-4409-a4fd-a00a58d787ee"
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
    "noteID": 13,
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
    "requestID": "fdf2d5dd-116f-4b9a-b2ea-4f5adcfa10b2"
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
        "total": 13,
        "list": [
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
            }
        ]
    },
    "requestID": "2075a4ef-3bfd-4766-a92d-496717c4b83e"
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
noteID=13
```

</details>

<details>
<summary>Return</summary>

```json
{
    "code": 0,
    "msg": "",
    "data": {
        "noteID": 13,
        "userName": "1",
        "title": "mod title",
        "detail": "mod detail",
        "createdAt": 1597490651
    },
    "requestID": "02c8dae0-d721-4085-be30-286b44e977d7"
}
```

</details>

