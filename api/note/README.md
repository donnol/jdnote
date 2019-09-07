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
        "id": 17
    }
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
    "noteID": 17,
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
                "noteID": 17,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1567843929
            },
            {
                "noteID": 16,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1567825606
            },
            {
                "noteID": 15,
                "userName": "0",
                "title": "",
                "detail": "testDetail",
                "createdAt": 1567824210
            },
            {
                "noteID": 14,
                "userName": "0",
                "title": "",
                "detail": "",
                "createdAt": 1567824201
            },
            {
                "noteID": 13,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1567822514
            },
            {
                "noteID": 12,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1567821986
            },
            {
                "noteID": 11,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1567669630
            },
            {
                "noteID": 10,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1567578639
            },
            {
                "noteID": 9,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1567578072
            },
            {
                "noteID": 8,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1567568480
            }
        ]
    }
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
noteID=17
```

</details>

<details>
<summary>Return</summary>

```json
{
    "code": 0,
    "msg": "",
    "data": {
        "noteID": 17,
        "userName": "119",
        "title": "mod title",
        "detail": "mod detail",
        "createdAt": 1567843929
    }
}
```

</details>

