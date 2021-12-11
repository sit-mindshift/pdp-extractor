# About chrompdp-extractoredp

Package `pdp-extractor` is go based product details crawler for the most popuplar shops.
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

	p := pdp.ProductDetailPageExtractor{
		Url:            "https://www.amazon.de/dp/B07SPTL3SQ/ref=nosim?tag=masedeveloper-21",
		ScreenshotFile: "./out/amazon.de.png",
	}
	p.Run()
	fmt.Printf("got PDP with url:%s title: %s and description:%s\n", p.Url, p.MetaTitle, p.MetaDescription)

	p = pdp.ProductDetailPageExtractor{
		Url:            "https://www.lidl.de/p/puma-herren-t-shirt-teamgoal-mit-rundhalsausschnitt/p100339220",
		ScreenshotFile: "./out/lidl.de.png",
	}
	p.Run()
	fmt.Printf("got PDP with url:%s title: %s and description:%s\n", p.Url, p.MetaTitle, p.MetaDescription)

	p = pdp.ProductDetailPageExtractor{
		Url:            "https://www.kaufland.de/product/404444772/",
		ScreenshotFile: "./out/kaufland.de.png",
	}
	p.Run()
	fmt.Printf("got PDP with url:%s title: %s and description:%s\n", p.Url, p.MetaTitle, p.MetaDescription)

	fmt.Printf("DONE!")
}
```
