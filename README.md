# About chrompdp-extractoredp

Package `pdp-extractor` is go based webpage crawler for product pages on the most popuplar shos
you can use it if you want to show a minimalistic preview of a product based on the URL only.

[![Go Reference](https://pkg.go.dev/badge/github.com/sit-mindshift/pdp-extractor/pdp.svg)](https://pkg.go.dev/github.com/sit-mindshift/pdp-extractor/pdp)

## Installing

Install in the usual Go way:

```sh
$ go get -u github.com/sit-mindshift/pdp-extractor/pdp
```

## Example

To use it:

```go
	p := pdp.ProductDetailPageExtractor{
		Url:            "https://www.lidl.de/p/puma-herren-t-shirt-teamgoal-mit-rundhalsausschnitt/p100339220",
		ScreenshotFile: "./out/lidl-shirtpng",
	}
	p.Run()
	fmt.Printf("got PDP with url:%s title: %s and description:%s\n", p.Url, p.MetaTitle, p.MetaDescription)
}))
```
