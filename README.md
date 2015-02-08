## extract [![GoDoc](https://img.shields.io/badge/godoc-Documentation-blue.svg?style=flat)](https://godoc.org/github.com/hariharan-uno/extract)

extract is a simple library for extracting elements from a web page. It provides simple higher level functions based upon [Cascadia](http://code.google.com/p/cascadia) and [html](http://golang.org/x/net/html) packages.

For example,

##### extract all the links from a web page 
~~~go
package main

import (
	"fmt"
	"log"

	"github.com/hariharan-uno/extract"
)

func main() {
	l, err := extract.Links("http://google.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(l)
}
~~~

##### extract all the URLs of the images from a web page
~~~go
package main

import (
	"fmt"
	"log"

	"github.com/hariharan-uno/extract"
)

func main() {
	i, err := extract.Images("http://google.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(i)
}
~~~

Currently, only the functions `extract.Links()` and `extract.Images()` are supported. If you'd like a specific function to be supported, please file an issue.

### Credits

Authors of [Cascadia](http://code.google.com/p/cascadia) and [html](http://golang.org/x/net/html)