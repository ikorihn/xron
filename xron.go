package xron

import (
	"fmt"
	"io"
	"strings"

	"github.com/antchfx/xmlquery"
)

func ConvertXmlToXpath(r io.Reader) {
	doc, err := xmlquery.Parse(r)
	if err != nil {
		panic(err)
	}

	for _, n := range xmlquery.Find(doc, "/") {
		traverseChildElement(n, "")
	}
}

func traverseChildElement(node *xmlquery.Node, parent string) {
	fc := node.FirstChild
	if fc == nil {
		if node.Type == xmlquery.ElementNode {
			fmt.Printf("%s/%s\n", parent, formatAttr(node))
		} else if strings.TrimSpace(node.InnerText()) != "" {
			fmt.Printf("%s/text() = '%s'\n", parent, node.InnerText())
		}
		return
	}

	if parent == "/" {
		parent = ""
	}
	current := fmt.Sprintf("%s/%s", parent, formatAttr(node))
	fmt.Printf("%s\n", current)

	traverseChildElement(fc, current)
	traverseNextSibling(fc, current)
}

func traverseNextSibling(node *xmlquery.Node, parent string) {
	ns := node.NextSibling
	if ns == nil {
		return
	}

	traverseChildElement(ns, parent)
	traverseNextSibling(ns, parent)
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
