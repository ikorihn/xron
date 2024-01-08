package xron

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

func ConvertXmlToXpath(r io.Reader) (rows []string) {
	ConvertXmlToXpathFunc(r, func(row string) {
		rows = append(rows, row)
	})
	return
}

func ConvertXmlToXpathFunc(r io.Reader, emitRow func(string)) {
	d := xml.NewDecoder(r)
	// stack of names of surrounding XML elements
	stack := []string{}
	// inText is non-empty when merging consecutive text/CDATA
	inText := []string{}
	// FIXME: should only happen if really has any root, probably?
	emitRow("/")
	for {
		t, err := d.Token()
		// Note: err handling is further down

		if len(inText) != 0 {
			shouldEmit := true
			switch t.(type) {
			case xml.ProcInst, xml.Directive, xml.Comment,
				xml.CharData:
				shouldEmit = false
			}
			if shouldEmit {
				text := strings.Join(inText, "")
				// TODO: modify tests to not mandate single quotes
				text = fmt.Sprintf("/text() = '%s'", text)
				prefix := ""
				if len(stack) > 0 {
					prefix = "/" + strings.Join(stack, "/")
				}
				emitRow(prefix + text)
				inText = nil
			}
		}

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
				// FIXME[LATER]: handle properly if inText
				continue
			}
			text = fmt.Sprintf("%q", text)
			text = text[1 : len(text)-1]
			inText = append(inText, text)
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
