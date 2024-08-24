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
  // Send a GET request to the URL
  resp, err := http.Get(url)
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
