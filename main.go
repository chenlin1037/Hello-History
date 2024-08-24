package main

import (
	"net/http"

	"Hello-History/pkg/articlefetcher"
	"Hello-History/pkg/articletranslate"
	"github.com/TruthHun/html2md"
)

func main() {
	http.HandleFunc("/translate", handleTranslate)
	http.ListenAndServe(":8080", nil)
}

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

	// 使用 articletranslate 翻译并转换为 Markdown
	markdownContent, err := articletranslate.FetchArticle(htmlContent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 返回 Markdown 内容
	mdStr:=html2md.Convert(markdownContent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(mdStr))
}