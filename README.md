# About pdp-extractor

Package `pdp-extractor` is a GO based product details crawler for the most popuplar shops.
You can use it if you want to show a minimalistic preview of a product based on the URL only.

[![Go Reference](https://pkg.go.dev/badge/github.com/sit-mindshift/pdp-extractor/pdp.svg)](https://pkg.go.dev/github.com/sit-mindshift/pdp-extractor/pdp)

## Installing

Install in the usual Go way:

```sh
$ go get -u github.com/sit-mindshift/pdp-extractor/pdp
```

## Example

To use it:

```go
package main

import (
	"fmt"

	pdp "github.com/sit-mindshift/pdp-extractor/pdp"
)

func main() {

	pdpResult, _ := pdp.Run("https://www.amazon.de/dp/B07SPTL3SQ/ref=nosim?tag=masedeveloper-21")
	fmt.Printf("got PDP with title: %s and description:%s\n", pdpResult.MetaTitle, pdpResult.MetaDescription)
	pdp.WriteScreenshotToFile("./out/amazon.de.png", pdpResult)
	fmt.Printf("DONE!")
}
```
