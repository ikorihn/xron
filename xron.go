package xron

import (
	"fmt"
	"io"
	"strings"

	"github.com/antchfx/xmlquery"
)

func ConvertXmlToXpath(r io.Reader) []string {
	doc, err := xmlquery.Parse(r)
	if err != nil {
		panic(err)
	}

	xpaths := make([]string, 0)
	for _, n := range xmlquery.Find(doc, "/") {
		traverseChildElement(n, "", &xpaths)
	}
	return xpaths
}

func traverseChildElement(node *xmlquery.Node, parent string, xpaths *[]string) {
	fc := node.FirstChild
	if fc == nil {
		if node.Type == xmlquery.ElementNode {
			current := fmt.Sprintf("%s/%s", parent, formatAttr(node))
			*xpaths = append(*xpaths, current)
		} else if strings.TrimSpace(node.InnerText()) != "" {
			current := fmt.Sprintf("%s/text() = '%s'", parent, node.InnerText())
			*xpaths = append(*xpaths, current)
		}
		return
	}

	if parent == "/" {
		parent = ""
	}
	current := fmt.Sprintf("%s/%s", parent, formatAttr(node))
	*xpaths = append(*xpaths, current)

	traverseChildElement(fc, current, xpaths)
	traverseNextSibling(fc, current, xpaths)
}

func traverseNextSibling(node *xmlquery.Node, parent string, xpaths *[]string) {
	ns := node.NextSibling
	if ns == nil {
		return
	}

	traverseChildElement(ns, parent, xpaths)
	traverseNextSibling(ns, parent, xpaths)
}

func formatAttr(node *xmlquery.Node) string {
	s := node.Data
	for _, a := range node.Attr {
		s += fmt.Sprintf(`[@%s="%s"]`, a.Name.Local, a.Value)
	}
	return s
}

func nodeTypeString(n xmlquery.NodeType) string {
	switch n {
	case xmlquery.DocumentNode:
		return "DocumentNode"
	case xmlquery.DeclarationNode:
		return "DeclarationNode"
	case xmlquery.ElementNode:
		return "ElementNode"
	case xmlquery.TextNode:
		return "TextNode"
	case xmlquery.CharDataNode:
		return "CharDataNode"
	case xmlquery.CommentNode:
		return "CommentNode"
	case xmlquery.AttributeNode:
		return "AttributeNode"
	}
	return ""
}
