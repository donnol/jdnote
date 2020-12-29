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
        "id": 25
    },
    "timestamp": 1609232706,
    "requestID": "6308e48c-520c-48a6-a30c-40cd18305a62"
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
    "noteID": 25,
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
    "timestamp": 1609232715,
    "requestID": "f2bf94a5-a9c5-4815-99a2-016d794f39c3"
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
        "total": 25,
        "list": [
            {
                "noteID": 25,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "status": 1,
                "createdAt": 1609232706
            },
            {
                "noteID": 24,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "status": 1,
                "createdAt": 1609221504
            },
            {
                "noteID": 23,
                "userName": "0",
                "title": "8",
                "detail": "8",
                "status": 1,
                "createdAt": 1608623767
            },
            {
                "noteID": 22,
                "userName": "0",
                "title": "7_hahah",
                "detail": "7",
                "status": 1,
                "createdAt": 1608623736
            },
            {
                "noteID": 21,
                "userName": "0",
                "title": "6_hahah",
                "detail": "6",
                "status": 1,
                "createdAt": 1608622973
            },
            {
                "noteID": 20,
                "userName": "0",
                "title": "5_hahah",
                "detail": "5",
                "status": 1,
                "createdAt": 1608618438
            },
            {
                "noteID": 19,
                "userName": "0",
                "title": "hahah",
                "detail": "4",
                "status": 1,
                "createdAt": 1608616033
            },
            {
                "noteID": 18,
                "userName": "0",
                "title": "3",
                "detail": "3",
                "status": 1,
                "createdAt": 1608099670
            },
            {
                "noteID": 17,
                "userName": "0",
                "title": "",
                "detail": "",
                "status": 1,
                "createdAt": 1608099231
            },
            {
                "noteID": 16,
                "userName": "0",
                "title": "",
                "detail": "",
                "status": 1,
                "createdAt": 1608099004
            }
        ]
    },
    "timestamp": 1609232726,
    "requestID": "c404b1b9-7578-467e-9401-e3af0911d492"
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
noteID=25
```

</details>

<details>
<summary>Return</summary>

```json
{
    "code": 0,
    "msg": "",
    "data": {
        "noteID": 25,
        "userName": "1",
        "title": "mod title",
        "detail": "mod detail",
        "status": 1,
        "createdAt": 1609232706
    },
    "timestamp": 1609232737,
    "requestID": "0695905c-1d74-4ae8-ae56-02fe322cb862"
}
```

</details>

