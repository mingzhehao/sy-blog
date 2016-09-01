package g

import (
	"github.com/russross/blackfriday"
)

func RenderMarkdown(mdStr string) string {

	body := blackfriday.MarkdownBasic([]byte(mdStr))

	return string(body)
}
