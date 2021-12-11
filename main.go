package main

import (
	"fmt"

	pdp "mase.rip/pdp/pdp"
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
