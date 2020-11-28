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
    "timestamp": 1606570686,
    "requestID": "ad31bfbc-2588-431b-b105-c5f83b67907d"
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
    "timestamp": 1606570687,
    "requestID": "00010a1d-c77f-4060-a813-bc3ab222b2b7"
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
        * status (*notemodel.Status*) 状态
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
        "total": 7,
        "list": [
            {
                "noteID": 7,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "status": 1,
                "createdAt": 1606570686
            },
            {
                "noteID": 6,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "status": 1,
                "createdAt": 1606570564
            },
            {
                "noteID": 5,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "status": 1,
                "createdAt": 1606570135
            },
            {
                "noteID": 4,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "status": 1,
                "createdAt": 1606568988
            },
            {
                "noteID": 3,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "status": 1,
                "createdAt": 1606568781
            },
            {
                "noteID": 2,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "status": 1,
                "createdAt": 1606538720
            },
            {
                "noteID": 1,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "status": 1,
                "createdAt": 1606538652
            }
        ]
    },
    "timestamp": 1606570688,
    "requestID": "3615f540-d8f8-4b21-b7b8-0b841cb50533"
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
    * status (*notemodel.Status*) 状态
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
        "status": 1,
        "createdAt": 1606570686
    },
    "timestamp": 1606570690,
    "requestID": "ed7c034d-28ca-4427-aa4c-2bda57d8bea7"
}
```

</details>

