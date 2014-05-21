##extract

HTML scraping library.

It provides simple higher level functions based upon [Cascadia](http://code.google.com/p/cascadia) and [go.net](http://code.google.com/p/go.net)

For example, to extract all the links from a web page, 
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

###Credits

Authors of Cascadia & go.net