package main

import (
	"fmt"
	"strings"

	"github.com/antchfx/xmlquery"
)

func main() {
	s := `<?xml version="1.0" encoding="UTF-8" ?>
<books>
    <book ID="extension1" available="yes">
        <title>Book Title 1</title>
        <price>100 - 200</price>
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

	doc, err := xmlquery.Parse(strings.NewReader(s))
	if err != nil {
		panic(err)
	}
	// channel := xmlquery.FindOne(doc, "//channel")
	// if n := channel.SelectElement("title"); n != nil {
	// 	fmt.Printf("title: %s\n", n.InnerText())
	// }
	// if n := channel.SelectElement("link"); n != nil {
	// 	fmt.Printf("link: %s\n", n.InnerText())
	// }
	// for i, n := range xmlquery.Find(doc, "//item/title") {
	// 	fmt.Printf("#%d %s\n", i, n.InnerText())
	// }

	for i, n := range xmlquery.Find(doc, "/") {
		children(n)
		// child := n.FirstChild
		// ns := child.NextSibling
		// for ns != nil {
		// 	if child.NextSibling != nil {
		// 		fmt.Printf("attr: %v, text: %v, type: %s\n", ns.Attr, ns.Data, nodeTypeString(ns.Type))
		// 	}
		// 	ns = ns.NextSibling
		// }
	}
}

func children(node *xmlquery.Node) {
	fc := node.FirstChild
	if fc == nil {
		// fmt.Printf("deepest attr: %v, text: %v, type: %s\n", node.Attr, node.Data, nodeTypeString(node.Type))
		return
	}

	pathstr := node.Data
	for _, a := range node.Attr {
		pathstr += fmt.Sprintf(`[@%s='%s']`, a.Name.Local, a.Value)

	}
	// fmt.Printf("child attr: %v, text: %v, type: %s, %d\n", node.Attr, node.Data, nodeTypeString(node.Type), node.Type)
	fmt.Printf("%s\n", pathstr)

	nextSibling(fc)
	children(fc)
}
func nextSibling(node *xmlquery.Node) {
	ns := node.NextSibling
	if ns == nil {
		// fmt.Printf("finalSibling attr: %v, text: %v, type: %s\n", node.Attr, node.Data, nodeTypeString(node.Type))
		children(node)
		return
	}
	children(ns)
	nextSibling(ns)
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
