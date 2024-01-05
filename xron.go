package xron

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

func ConvertXmlToXpath(r io.Reader) (rows []string) {
	d := xml.NewDecoder(r)
	stack := []string{}
	// TODO[LATER]: instead, have emitRow passed as arg
	emitRow := func(row string) {
		rows = append(rows, row)
	}
	// FIXME: should only happen if really has any root, probably?
	emitRow("/")
	for {
		t, err := d.Token()
		if err == io.EOF {
			return
		} else if err != nil {
			panic(err) // TODO[LATER]: allow returning errors
		}
		switch t := t.(type) {
		case xml.StartElement:
			attrs := []string{}
			for _, a := range t.Attr {
				attrs = append(attrs, fmt.Sprintf(`[@%s=%q]`, a.Name.Local, a.Value))
			}
			stack = append(stack, t.Name.Local+strings.Join(attrs, ""))
			emitRow("/" + strings.Join(stack, "/"))
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			text := trimSpace(t)
			if len(text) == 0 {
				continue
			}
			// TODO: verify if consecutive CDATA are merged
			text = fmt.Sprintf("%q", text)
			// TODO: modify tests to not mandate single quotes
			text = fmt.Sprintf("/text() = '%s'", text[1:len(text)-1])
			prefix := ""
			if len(stack) > 0 {
				prefix = "/" + strings.Join(stack, "/")
			}
			emitRow(prefix + text)
		case xml.ProcInst, xml.Directive, xml.Comment:
			// ignore
		}
	}
}

func trimSpace(c xml.CharData) string {
	const cutset = " \n"
	s := strings.TrimLeft(string(c), cutset)
	if len(s) != len(c) {
		s = " " + s
	}
	n := len(s)
	s = strings.TrimRight(s, cutset)
	if len(s) != n {
		s = s + " "
	}
	if s == " " {
		return ""
	}
	return s
}

/*
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
		if parent == "/" {
			parent = ""
		}
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
*/
