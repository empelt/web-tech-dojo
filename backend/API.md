---
title: Web Tech Dojo API
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.23"

---

# Web Tech Dojo API

Base URLs: $API_GATEWAY_URL

# Authentication

- HTTP Authentication, scheme: bearer

# web-tech-dojo

## GET question list

GET /api/question

## Overview
問題の一覧を取得

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|tags|query|string| no |絞り込みに用いるタグ|
|isBookmarked|query|boolean| no |ブックマークされているものに絞るかどうか|
|progress|query|integer| no |進行状態を表すEnumStatus|
|limit|query|integer| no |ページネーションのlimit|
|offset|query|integer| no |ページネーションのoffset(question.id)|

> Response Examples

> 200 Response

```json
{}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

## OPTIONS question list

OPTIONS /api/question

> Response Examples

> 204 Response

```json
{}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|204|[No Content](https://tools.ietf.org/html/rfc7231#section-6.3.5)|none|Inline|

### Responses Data Schema

### Response Header

|Status|Header|Type|Format|Description|
|---|---|---|---|---|
|204|Access-Control-Allow-Origin|string||none|
|204|Access-Control-Allow-Methods|string||none|
|204|Access-Control-Allow-Headers|string||none|

## GET question context

GET /api/question/{id}

## Overview
問題のデータを取得

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|id|path|string| yes |none|

> Response Examples

> 200 Response

```json
{}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

## POST question answer

POST /api/question/{id}

## Overview
回答送信

> Body Parameters

```json
{
  "message": "string"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|id|path|string| yes |none|
|body|body|object| no |none|
|» message|body|string| yes |none|

> Response Examples

> 200 Response

```json
{}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

## OPTIONS question context

OPTIONS /api/question/{id}

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|id|path|string| yes |none|

> Response Examples

> 200 Response

```json
{}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

### Response Header

|Status|Header|Type|Format|Description|
|---|---|---|---|---|
|200|Access-Control-Allow-Origin|string||none|
|200|Access-Control-Allow-Methods|array||none|
|200|Access-Control-Allow-Headers|string||none|

# Data Schema

