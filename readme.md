# 灵感来源

Medium上面有很多不错的文章，可惜需要收费，有一个网站readMedium可以破解收费，网站还提供翻译功能，将英文内容翻译为中文，翻译质量还说得过去。爬取翻译后的内容，得到一个markdown格式的文章内容。注意：由于medium是一个被墙的网站，当直接粘贴文章在某些平台时，图片不能被解析。

## TODoList

-正在做前端页面

# 获取Medium文章并保存为中文markdown文档 API 

## 概述
-本服务提供三个API端点用于翻译不同来源的内容。
-直接从文章url翻译
-使用HTML文本翻译（应对ip被封的情况）
-仅包含<article>标签的html文本的翻译（早期版本）

## API 端点

### 1. URL内容翻译
从指定URL获取内容并进行翻译。

**接口：** `/translate`
**方法：** GET
**查询参数：**
- `url` (必需)：需要翻译的内容的URL地址

**请求示例：**
```curl
curl "http://localhost:8080/translate?url=https://example.com/article"
```

**成功响应：**
```json
{
    "success": true,
    "data": "翻译后的markdown格式内容"
}
```

**错误响应：**
```json
{
    "success": false,
    "error": "错误信息描述"
}
```

### 2. HTML内容翻译
直接翻译提供的HTML内容。

**接口：** `/translateFromHtml`
**方法：** POST
**Content-Type：** text/html

**请求示例：**
```curl
curl -X POST \
  http://localhost:8080/translateFromHtml \
  -H "Content-Type: text/html" \
  -d "<html><body><p>需要翻译的内容</p></body></html>"
```

**响应格式：** 与URL翻译接口相同

### 3. 带文章提取的HTML翻译
提取HTML中的`<article>`标签内容并进行翻译。

**接口：** `/translateFromHtml2`
**方法：** POST
**Content-Type：** text/html

**请求示例：**
```curl
curl -X POST \
  http://localhost:8080/translateFromHtml2 \
  -H "Content-Type: text/html" \
  -d "<html><body><article>需要翻译的文章内容</article></body></html>"
```

**响应格式：** 与其他接口相同

## 错误码说明
- 400：请求错误（缺少参数或输入无效）
- 405：方法不允许（HTTP方法错误）
- 500：服务器内部错误（翻译或处理失败）

## 响应格式
所有接口都返回JSON格式的响应：

**成功响应：**
```json
{
    "success": true,
    "data": "翻译后的markdown格式内容"
}
```

**错误响应：**
```json
{
    "success": false,
    "error": "错误描述信息"
}
```

## 使用建议
1. 对于单篇文章翻译，推荐使用URL翻译接口
2. 对于已有HTML内容的翻译，可直接使用HTML翻译接口
3. 如果HTML内容包含在article标签中，建议使用带文章提取的翻译接口以获得更精确的结果