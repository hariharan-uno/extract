##extract

HTML scraping library.

It provides simple higher level functions based upon [Cascadia](http://code.google.com/p/cascadia) and [go.net](http://code.google.com/p/go.net)

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

Currently, only the functions `extract.Links()` and `extract.Images()` is supported. More functions will be supported soon.

###TODO
- [x] Extract Links
- [x] Extract Images
- [ ] Extract Headings

###Credits

Authors of Cascadia & go.net