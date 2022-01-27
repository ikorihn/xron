package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/antchfx/xmlquery"
)

func main() {
	s := `<?xml version="1.0" encoding="UTF-8" ?>
<books>
    <book ID="extension1" available="yes">
        <title>Book Title 1</title>
        <price>100 - 200</price>
        <note name="mynote">
            <author id="20">Author 1</author>
        </note>
        <category name="novel" id="2" />
        <empty />
    </book>
    <book ID="Book001" available="no">
        <title>Book Title 2</title>
        <price>400 - 500</price>
    </book>
    <book ID="extension2">
        <title>Book Title 3</title>
        <price>abc - 500</price>
    </book>
    <book available="soon">
        <title>Book Title 5</title>
        <price>a100 - c200</price>
    </book>
    <book ID="Book003" available="yes">
        <title>Book Title 6</title>
        <price>c100 - 200</price>
    </book>
</books>
`

	convertXmlToXpath(strings.NewReader(s))
}

func convertXmlToXpath(r io.Reader) {
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
