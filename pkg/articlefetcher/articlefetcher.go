package articlefetcher

import (
  "bytes"
  "fmt"
  "io/ioutil"
  "net/http"

  "golang.org/x/net/html"
)

// FetchArticle fetches the HTML content of a given URL and extracts the content within the <article> tag.
func FetchArticle(url string) (string, error) {
  // Create a new GET request with headers
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return "", err
  }

  // Set cookies
  req.Header.Set("Cookie", "_ga=GA1.1.1917273616.1723734054; NEXT_LOCALE=zh; lng=en; _clck=1carypx%7C2%7Cfor%7C0%7C1688; _clsk=1brr9sf%7C1725017798849%7C1%7C1%7Ci.clarity.ms%2Fcollect; _ga_YSBN5EJVBL=GS1.1.1725017794.21.1.1725017864.0.0.0")

  // Set user-agent
  req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36")

  // Send the request
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return "", err
  }
  defer resp.Body.Close()

  // Read the HTML content from the response
  htmlContent, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", err
  }

  // Parse the HTML content
  doc, err := html.Parse(bytes.NewReader(htmlContent))
  if err != nil {
    return "", err
  }

  // Find the <article> tag
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

  // If no <article> tag is found, return an error
  if articleNode == nil {
    return "", fmt.Errorf("no <article> tag found in the HTML content")
  }

  // Render the <article> tag as a string
  var buf bytes.Buffer
  writer := &buf
  html.Render(writer, articleNode)
  return buf.String(), nil
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