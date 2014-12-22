// Copyright 2014 Hari haran. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package extract_test

import (
	"fmt"
	"log"

	"github.com/hariharan-uno/extract"
)

func Example() {
	links, err := extract.Links("https://google.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", links)

	images, err := extract.Images("https://google.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", images)
}
