package main

import (
	"fmt"

	"github.com/hariharan-uno/extract"
)

func main() {
	fmt.Println(extract.Links("http://google.com"))
	fmt.Println(extract.Images("http://google.com"))
}
