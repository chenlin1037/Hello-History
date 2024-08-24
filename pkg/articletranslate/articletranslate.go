package articletranslate

import (
  "bytes"
  "strings"

  "github.com/russross/blackfriday/v2"
  "golang.org/x/net/html"
)

// FetchArticle converts HTML content to Markdown
func FetchArticle(inputHTML string) (string, error) {
  // Parse HTML into a node tree
  doc, err := html.Parse(strings.NewReader(inputHTML))
  if err != nil {
    return "", err
  }

  // Remove unwanted content
  removeEnglishContent(doc)

  // Render the node tree back to HTML
  var buf bytes.Buffer
  err = html.Render(&buf, doc)
  if err != nil {
    return "", err
  }

  // Convert HTML to Markdown
  markdown := blackfriday.Run(buf.Bytes())

  // Convert Markdown byte slice to string
  outputMarkdown := strings.TrimSpace(string(markdown))

  return outputMarkdown, nil
}

// removeEnglishContent recursively removes English content from HTML nodes
func removeEnglishContent(node *html.Node) {
  if node.Type == html.ElementNode && node.Data == "br" {
    return // Preserve <br> tags
  }

  // Find <span> tags with class "readmedium-translated-content" and preserve their content, remove English content from sibling nodes
  if node.Type == html.ElementNode && node.Data == "span" {
    for _, attr := range node.Attr {
      if attr.Key == "class" && attr.Val == "readmedium-translated-content" {
        // Replace parent node's entire child content with <span> tag's content
        node.Parent.FirstChild = node.FirstChild
        node.Parent.LastChild = node.LastChild
        break
      }
    }
  }

  // Recursively process child nodes
  for child := node.FirstChild; child != nil; child = child.NextSibling {
    removeEnglishContent(child)
  }
}