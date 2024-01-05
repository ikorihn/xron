# xron

Transform XML into xpath and make it greppable.

Like [gron](https://github.com/tomnomnom/gron), xron prints every xml element as xpath style line by line.

## Usage

```bash
$ cat books.xml
<?xml version="1.0" encoding="UTF-8"?>
<books>
  <book ID="extension1" available="yes">
    <title>Book Title 1</title>
    <price>1000</price>
    <category name="novel" id="2"/>
  </book>
  <book ID="extension2">
    <title>Book Title 2</title>
    <price>500</price>
  </book>
</books>

$ xron < books.xml
/
/books
/books/book[@ID="extension1"][@available="yes"]
/books/book[@ID="extension1"][@available="yes"]/title
/books/book[@ID="extension1"][@available="yes"]/title/text() = 'Book Title 1'
/books/book[@ID="extension1"][@available="yes"]/price
/books/book[@ID="extension1"][@available="yes"]/price/text() = '1000'
/books/book[@ID="extension1"][@available="yes"]/category[@name="novel"][@id="2"]
/books/book[@ID="extension2"]
/books/book[@ID="extension2"]/title
/books/book[@ID="extension2"]/title/text() = 'Book Title 2'
/books/book[@ID="extension2"]/price
/books/book[@ID="extension2"]/price/text() = '500'
```

## Install

```bash
$ go install github.com/akavel/xron/cmd/xron@latest
```
