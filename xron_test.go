package xron

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestConvertXmlToXpath(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{

		{
			name: "convert common xml",
			args: args{
				s: `<?xml version="1.0" encoding="UTF-8" ?>
                  <books>
                      <book ID="extension1" available="yes">
                          <title>Book Title 1</title>
                          <price>100 <b>-</b> 200</price>
                          <note name="mynote">
                              <author id="20">Author 1</author>
                          </note>
                          <category name="novel" id="2" />
                      </book>
                      <book ID="extension2" available="no">
                          <title>Book <![CDATA[Title]]> 2</title>
                          <price>2300</price>
                          <category name="fiction" id="10" />
                      </book>
                  </books>`,
			},
			want: []string{
				"/",
				"/books",
				`/books/book[@ID="extension1"][@available="yes"]`,
				`/books/book[@ID="extension1"][@available="yes"]/title`,
				`/books/book[@ID="extension1"][@available="yes"]/title/text() = 'Book Title 1'`,
				`/books/book[@ID="extension1"][@available="yes"]/price`,
				`/books/book[@ID="extension1"][@available="yes"]/price/text() = '100 '`,
				`/books/book[@ID="extension1"][@available="yes"]/price/b`,
				`/books/book[@ID="extension1"][@available="yes"]/price/b/text() = '-'`,
				`/books/book[@ID="extension1"][@available="yes"]/price/text() = ' 200'`,
				`/books/book[@ID="extension1"][@available="yes"]/note[@name="mynote"]`,
				`/books/book[@ID="extension1"][@available="yes"]/note[@name="mynote"]/author[@id="20"]`,
				`/books/book[@ID="extension1"][@available="yes"]/note[@name="mynote"]/author[@id="20"]/text() = 'Author 1'`,
				`/books/book[@ID="extension1"][@available="yes"]/category[@name="novel"][@id="2"]`,
				`/books/book[@ID="extension2"][@available="no"]`,
				`/books/book[@ID="extension2"][@available="no"]/title`,
				`/books/book[@ID="extension2"][@available="no"]/title/text() = 'Book Title 2'`,
				`/books/book[@ID="extension2"][@available="no"]/price`,
				`/books/book[@ID="extension2"][@available="no"]/price/text() = '2300'`,
				`/books/book[@ID="extension2"][@available="no"]/category[@name="fiction"][@id="10"]`,
			},
		},

		{
			name: "show empty tag",
			args: args{
				s: `<?xml version="1.0" encoding="UTF-8" ?>
                  <empty />`,
			},
			want: []string{
				"/",
				`/empty`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertXmlToXpath(strings.NewReader(tt.args.s)); !cmp.Equal(got, tt.want) {
				t.Errorf("ConvertXmlToXpath() = %v, want %v\ndiff=%v", got, tt.want, cmp.Diff(got, tt.want))
			}
		})
	}
}
