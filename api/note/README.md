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
        "id": 112
    },
    "requestID": "ffef65a1-a4b3-4489-8e94-f710726e130f"
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
    "noteID": 112,
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
    "requestID": "a96be485-9b87-4864-a0b2-ad05d04408d9"
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
        "total": 75,
        "list": [
            {
                "noteID": 112,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1581140867
            },
            {
                "noteID": 111,
                "userName": "0",
                "title": "123123",
                "detail": "1231313",
                "createdAt": 1579752275
            },
            {
                "noteID": 110,
                "userName": "0",
                "title": "积沙成塔",
                "detail": "# 积沙成塔\n\n无沙何来塔。",
                "createdAt": 1566523796
            },
            {
                "noteID": 109,
                "userName": "0",
                "title": "12313",
                "detail": "\u003cp\u003e### 45454545\u003c/p\u003e\u003cp\u003e```js\u003c/p\u003e\u003cp\u003efunc main()\u003c/p\u003e\u003cp\u003e```\u003c/p\u003e\u003cp\u003e## hahahahahahah\u003c/p\u003e\u003cp\u003e```go\u003c/p\u003e\u003cp\u003eadfdf\u003c/p\u003e\u003cp\u003e```\u003c/p\u003e\u003cp\u003eabc\u003c/p\u003e\u003cp\u003e\u0026gt; abc\u003c/p\u003e\u003cp\u003e\u003cbr\u003e\u003c/p\u003e\u003cp\u003e\u003cbr\u003e\u003c/p\u003e",
                "createdAt": 1564908092
            },
            {
                "noteID": 107,
                "userName": "0",
                "title": "11111111111111",
                "detail": "# 123\n\n### 123\n\n\n```go\nvar a int\n\nfunc main() {\n     fmt.Println(\"Hello world\") // 打印\n}\n```\n\n| 1| 2 |\n| --- | --- |\n| 3| 3|\n| 3| 3|\n| 3| 3|\n\n1111111111111\n22222222222222222\n\n\u003e a\n\u003e \n\u003e b\n\u003e c\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n",
                "createdAt": 1564907517
            },
            {
                "noteID": 106,
                "userName": "0",
                "title": "123132",
                "detail": "# 123",
                "createdAt": 1564906388
            },
            {
                "noteID": 105,
                "userName": "0",
                "title": "ceshi",
                "detail": "ces",
                "createdAt": 1564905500
            },
            {
                "noteID": 99,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1564888827
            },
            {
                "noteID": 98,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1564887133
            },
            {
                "noteID": 97,
                "userName": "0",
                "title": "mod title",
                "detail": "mod detail",
                "createdAt": 1564886414
            }
        ]
    },
    "requestID": "72ac0983-2606-419f-8f62-23bc6f5c5597"
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
noteID=112
```

</details>

<details>
<summary>Return</summary>

```json
{
    "code": 0,
    "msg": "",
    "data": {
        "noteID": 112,
        "userName": "119",
        "title": "mod title",
        "detail": "mod detail",
        "createdAt": 1581140867
    },
    "requestID": "495f0cca-a6b1-4061-9581-ee0dcfa83221"
}
```

</details>

