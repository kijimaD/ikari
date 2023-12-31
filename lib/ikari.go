package ikari

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

var replaceCount int

func wrapTextWithAnchor(htmlContent string) (string, error) {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return "", err
	}

	WrapTextWithAnchorRecursive(doc, "p", "a")

	var buf strings.Builder
	if err := html.Render(&buf, doc); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func WrapTextWithAnchorRecursive(n *html.Node, targetTag string, anchorTag string) {
	if n.Type == html.ElementNode && n.Data == targetTag && n.FirstChild != nil {
		n.Attr = append(n.Attr, html.Attribute{
			Key: "id",
			Val: fmt.Sprintf("count%d", replaceCount),
		})

		// 対象のノードを新しいアンカーノードで置き換える
		newAnchorNode := &html.Node{
			Type: html.ElementNode,
			Data: anchorTag,
			Attr: []html.Attribute{
				{
					Key: "href",
					Val: fmt.Sprintf("#count%d", replaceCount),
				},
			},
			FirstChild: &html.Node{
				Type:       html.TextNode,
				Data:       n.FirstChild.Data, // リンクのテキストを指定
				FirstChild: n.FirstChild,
			},
		}

		replaceCount++

		// 既存のノードを置き換える
		replaceNode(n, newAnchorNode)
		return
	}

	// 再帰的に子ノードに適用
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		WrapTextWithAnchorRecursive(c, targetTag, anchorTag)
	}
}

// ノードを置き換えるヘルパー関数
func replaceNode(oldNode, newNode *html.Node) {
	if oldNode.Parent == nil {
		// 親がいない場合は新しいノード全体で置き換える
		*oldNode = *newNode
	} else {
		// 親がいる場合は親ノード内で置き換える
		oldNode.FirstChild = newNode
	}
}
