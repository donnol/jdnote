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
    },
    "requestID": "cb421c22-83f8-490f-817f-807756acaaea"
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
    "data": null,
    "requestID": "aa45e0f0-71cc-4415-97b5-fcc62823e7a0"
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
        "total": 17,
        "list": [
            {
                "noteID": 17,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
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
            }
        ]
    },
    "requestID": "4d705a01-5d05-4629-b94c-a329fa8b91b2"
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
        "userName": "1",
        "title": "mod title",
        "detail": "mod detail",
        "createdAt": 1598363979
    },
    "requestID": "d98b3954-eea9-46fd-a6f0-d1e5a23cf3d3"
}
```

</details>

