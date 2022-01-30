package xron

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestConvertXmlToXpath(t *testing.T) {
	type test struct {
		name string
		arg  string
		want []string
	}
	tests := []test{
		{
			name: "",
			arg: `<?xml version="1.0" encoding="UTF-8" ?>
                  <books>
                      <book ID="extension1" available="yes">
                          <title>Book Title 1</title>
                          <price>100 - 200</price>
                          <note name="mynote">
                              <author id="20">Author 1</author>
                          </note>
                          <category name="novel" id="2" />
                      </book>
                  </books>`,
			want: []string{
				"/",
				"/books",
				`/books/book[@ID="extension1"][@available="yes"]`,
				`/books/book[@ID="extension1"][@available="yes"]/title`,
				`/books/book[@ID="extension1"][@available="yes"]/title/text() = 'Book Title 1'`,
				`/books/book[@ID="extension1"][@available="yes"]/price`,
				`/books/book[@ID="extension1"][@available="yes"]/price/text() = '100 - 200'`,
				`/books/book[@ID="extension1"][@available="yes"]/note[@name="mynote"]`,
				`/books/book[@ID="extension1"][@available="yes"]/note[@name="mynote"]/author[@id="20"]`,
				`/books/book[@ID="extension1"][@available="yes"]/note[@name="mynote"]/author[@id="20"]/text() = 'Author 1'`,
				`/books/book[@ID="extension1"][@available="yes"]/category[@name="novel"][@id="2"]`,
			},
		},
	}

	for _, tt := range tests {
		got := ConvertXmlToXpath(strings.NewReader(tt.arg))
		if diff := cmp.Diff(got, tt.want); diff != "" {
			t.Errorf("Mismatch (-got, +want) %s\n", diff)
		}
	}
}
