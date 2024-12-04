package main

import (
	"io/ioutil"
	"net/http"
	"strings"
	"bytes"

	"Hello-History/pkg/articlefetcher"
	"Hello-History/pkg/articletranslate"
	"github.com/TruthHun/html2md"
	"golang.org/x/net/html"
)

func main() {
	// 三个 API 处理程序
	http.HandleFunc("/translate", handleTranslate)
	http.HandleFunc("/translateFromHtml", handleTranslateFromHtml)
	http.HandleFunc("/translateFromHtml2", handleTranslateFromHtml2)
	http.ListenAndServe(":8080", nil)
}

// 第一个 API 处理程序：从 URL 获取文章内容并翻译
func handleTranslate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// 获取 URL 参数
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "Missing url parameter", http.StatusBadRequest)
		return
	}

	// 使用 articlefetcher 获取文章内容
	htmlContent, err := articlefetcher.FetchArticle(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 使用 articletranslate 翻译文章内容
	translateContent, err := articletranslate.FetchArticle(htmlContent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 将 HTML 转换为 Markdown
	mdStr := html2md.Convert(translateContent)

	// 替换 Markdown 内容中的 &amp; 为 "
	mdStr = strings.ReplaceAll(mdStr, "&amp;#34;", "\"")

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(mdStr))
}

// 第二个 API 处理程序：直接接收 HTML 内容并翻译
func handleTranslateFromHtml(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// 从请求体中获取 HTML 内容
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	htmlContent := string(body)

	// 使用 articletranslate 翻译文章内容
	translateContent, err := articletranslate.FetchArticle(htmlContent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 将 HTML 转换为 Markdown
	mdStr := html2md.Convert(translateContent)

	// 替换 Markdown 内容中的 &amp; 为 "
	mdStr = strings.ReplaceAll(mdStr, "&amp;#34;", "\"")

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(mdStr))
}
// 第三个 API 处理程序
func handleTranslateFromHtml2(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// 从请求体中获取 HTML 内容
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	htmlContent := string(body)

	// 提取 <article> 内容
	articleContent, err := extractArticle([]byte(htmlContent))
	if err != nil {
		http.Error(w, "Failed to extract article content: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 使用 articletranslate 翻译文章内容
	translateContent, err := articletranslate.FetchArticle(articleContent)
	if err != nil {
		http.Error(w, "Failed to translate article: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 将 HTML 转换为 Markdown
	mdStr := html2md.Convert(translateContent)

	// 替换 Markdown 内容中的 &amp; 为 "
	mdStr = strings.ReplaceAll(mdStr, "&amp;#34;", "\"")

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(mdStr))
}



// extractArticle 从给定的 HTML 内容中提取 <article> 标签的内容
func extractArticle(htmlContent []byte) (string, error) {
	doc, err := html.Parse(bytes.NewReader(htmlContent))
	if err != nil {
		return "", err
	}

	var articleNode *html.Node
	var findArticle func(*html.Node)
	findArticle = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "article" {
			articleNode = node
			return
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			findArticle(child)
		}
	}
	findArticle(doc)

	if articleNode == nil {
		return "", nil // 如果没有找到 <article> 标签，返回空字符串而不是错误
	}

	var buf bytes.Buffer
	html.Render(&buf, articleNode)
	return buf.String(), nil
}