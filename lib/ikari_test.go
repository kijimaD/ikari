package ikari

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestReplaceNode(t *testing.T) {
	oldStr := `<html>
  <head>
  </head>
  <body>
    <p>hello</p>
  </body>
</html>`
	old, err := html.Parse(strings.NewReader(oldStr))
	assert.NoError(t, err)

	newStr := `
<html>
  <head>
  </head>
  <body>
    <p><a>hello</a></p>
  </body>
</html>`
	newContent, err := html.Parse(strings.NewReader(newStr))
	assert.NoError(t, err)

	replaceNode(old, newContent)

	var b bytes.Buffer
	err = html.Render(&b, newContent)
	assert.NoError(t, err)

	expect := "<html><head>\n  </head>\n  <body>\n    <p><a>hello</a></p>\n  \n</body></html>"
	assert.Equal(t, expect, b.String())
}

func TestWrapText(t *testing.T) {
	doc, err := html.Parse(strings.NewReader("<p>hello</p>"))
	assert.NoError(t, err)

	WrapTextWithAnchorRecursive(doc, "p", "a")

	var b bytes.Buffer
	err = html.Render(&b, doc)
	assert.NoError(t, err)
	assert.Equal(t, `<html><head></head><body><p id="count0"><a href="#count0">hello</a></p></body></html>`, b.String())
}
