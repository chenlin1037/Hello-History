2. API 文档

API 名称： 文章翻译 API

描述： 该 API 用于接收一个 URL，获取文章内容，并将其翻译成中文并转换为 Markdown 格式。

请求方法： GET

请求地址： /translate

请求参数：

参数名	数据类型	描述	必填
url	string	文章所在的 URL	是
请求示例：

http://localhost:8080/translate?url=https://example.com
响应状态码：

状态码	描述
200	成功
405	方法不允许
400	请求参数错误
500	服务器内部错误
响应示例：

# 标题
## 副标题
... 翻译后的 Markdown 内容 ...
3. 注意事项

以上 API 代码示例仅供参考，实际开发中你可能需要根据具体需求进行调整。
你需要根据你的项目结构和实际情况修改 articlefetcher 和 articletranslate 模块的导入路径。
你还需要处理一些异常情况，例如：
URL 不合法
文章内容获取失败
翻译失败
你可以使用一些工具来帮助你生成 API 文档，例如 Swagger、OpenAPI 等。