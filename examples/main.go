// Copyright 2014 Hari haran. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/hariharan-uno/extract"
)

func main() {
	fmt.Println(extract.Links("http://google.com"))
	fmt.Println(extract.Images("http://google.com"))
}
