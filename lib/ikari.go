package ikari

import (
	"strings"

	"golang.org/x/net/html"
)

func wrapTextWithAnchor(htmlContent string) (string, error) {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return "", err
	}

	wrapTextWithAnchorRecursive(doc, "p", "a")

	var buf strings.Builder
	if err := html.Render(&buf, doc); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func wrapTextWithAnchorRecursive(n *html.Node, targetTag string, anchorTag string) {
	if n.Type == html.ElementNode && n.Data == targetTag {
		// 対象のノードを新しいアンカーノードで置き換える
		newAnchorNode := &html.Node{
			Type: html.ElementNode,
			Data: anchorTag,
			Attr: []html.Attribute{},
			FirstChild: &html.Node{
				Type: html.TextNode,
				Data: n.FirstChild.Data, // リンクのテキストを指定
			},
		}

		// 既存のノードを置き換える
		replaceNode(n, newAnchorNode)
		return
	}

	// 再帰的に子ノードに適用
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		wrapTextWithAnchorRecursive(c, targetTag, anchorTag)
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
