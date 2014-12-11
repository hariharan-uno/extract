##extract

HTML scraping library.

It provides simple higher level functions based upon [Cascadia](http://code.google.com/p/cascadia) and [html](http://golang.org/x/net/html) packages.

For example, 
To extract all the links from a web page, 
~~~go
package main

import (
	"fmt"

	"github.com/hariharan-uno/extract"
)

func main() {
	fmt.Println(extract.Links("http://google.com"))
}
~~~

To extract all the URLs of the images from a web page,
~~~go
package main

import (
	"fmt"

	"github.com/hariharan-uno/extract"
)

func main() {
	fmt.Println(extract.Images("http://google.com"))
}
~~~

Currently, only the functions `extract.Links()` and `extract.Images()` are supported. If you'd like a specific function to be supported, please file an issue.

###Credits

Authors of [Cascadia](http://code.google.com/p/cascadia) and [html](http://golang.org/x/net/html)