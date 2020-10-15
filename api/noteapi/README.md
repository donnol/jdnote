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
        "id": 52
    },
    "requestID": "25d03a43-03d6-4dd3-bdf1-0ac76fe428da"
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
    "noteID": 52,
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
    "requestID": "49d122bb-0ce2-4d04-95cd-9d815e1f9326"
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
        "total": 52,
        "list": [
            {
                "noteID": 52,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1602772737
            },
            {
                "noteID": 51,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1602772619
            },
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
            }
        ]
    },
    "requestID": "dd4e89e5-b902-4e0a-98af-018788cad691"
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
noteID=52
```

</details>

<details>
<summary>Return</summary>

```json
{
    "code": 0,
    "msg": "",
    "data": {
        "noteID": 52,
        "userName": "1",
        "title": "mod title",
        "detail": "mod detail",
        "createdAt": 1602772737
    },
    "requestID": "0f619126-a23a-41b1-9df2-9868469fe4cf"
}
```

</details>

