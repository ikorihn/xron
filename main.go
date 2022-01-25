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
        <note name="mynote">
            <author>Author 1</author>
        </note>
        <category name="novel" id="2" />
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

	for _, n := range xmlquery.Find(doc, "/") {
		children(n, "")
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

func children(node *xmlquery.Node, parent string) {
	fc := node.FirstChild
	if fc == nil {
		if len(node.Attr) > 0 {
			pathstr := node.Data
			for _, a := range node.Attr {
				pathstr += fmt.Sprintf(`[@%s="%s"]`, a.Name.Local, a.Value)
			}
			current := fmt.Sprintf("%s/%s", parent, pathstr)
			fmt.Printf("%s\n", current)
		}
		if strings.TrimSpace(node.InnerText()) != "" {
			fmt.Printf("%s/text() = '%s'\n", parent, node.InnerText())
		}
		return
	}

	pathstr := node.Data
	for _, a := range node.Attr {
		pathstr += fmt.Sprintf(`[@%s="%s"]`, a.Name.Local, a.Value)
	}
	// fmt.Printf("child attr: %v, text: %v, type: %s, %d\n", node.Attr, node.Data, nodeTypeString(node.Type), node.Type)
	if parent == "/" {
		parent = ""
	}
	current := fmt.Sprintf("%s/%s", parent, pathstr)
	fmt.Printf("%s\n", current)

	children(fc, current)
	nextSibling(fc, current)
}
func nextSibling(node *xmlquery.Node, parent string) {
	ns := node.NextSibling
	if ns == nil {
		return
	}

	children(ns, parent)
	nextSibling(ns, parent)
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
