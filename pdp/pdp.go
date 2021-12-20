// The Product Detail Page Extractor crawles all usefull informations from a product details page
// which might be relevant for creating a thubnail preview of a article
package pdp_extractor

import (
	"context"
	"fmt"
	"log"
	s "strings"

	"github.com/chromedp/chromedp"
)

// The Input field you should set are
//   - Url : the url of the article details page which should be crawled
//   - Screenshot file : the path to the local file where to create the screenshot of the page
// After calling the Run() method the following output fields will be set
//   - MetaTitle : the title of the product
//   - MetaDescription : the description of the product
// @Author: Sebastian Kroll
type ProductDetailPageExtractorResults struct {
	MetaTitle       string
	MetaDescription string
	MetaImage       string
}

// Starts the crawling of the product details page
func Run(url string) (result ProductDetailPageExtractorResults, err error) {
	log.Printf("Run %s", url)

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-web-security", true),
		chromedp.DisableGPU,
		chromedp.NoDefaultBrowserCheck,
		chromedp.NoFirstRun,
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// create context
	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
		//chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	/*
		ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
	*/

	extractMetaTaskInformation(ctx, url, &result)
	return result, err
}

func extractMetaTaskInformation(ctx context.Context, url string, result *ProductDetailPageExtractorResults) {

	/*
		//var val1, val2 string
		var nodes1 []*cdp.Node
		var nodes2 []*cdp.Node

		chromedp.Run(ctx, chromedp.Tasks{
			chromedp.Navigate(url),
			chromedp.Nodes(`meta[property="og:image"]`, &nodes2, chromedp.AtLeast(0)),
		})

		chromedp.Run(ctx, chromedp.Tasks{
			chromedp.Nodes("#landingImage", &nodes1, chromedp.AtLeast(0)),
		})

		//document.querySelector("head > meta:nth-child(43)")

		for _, n := range nodes1 {
			u := n.AttributeValue("src")
			fmt.Printf("nodes1: %s | src = %s\n", n.LocalName, u)
		}

		for _, n := range nodes2 {
			u := n.AttributeValue("content")
			fmt.Printf("nodes2: %s | content = %s\n", n.LocalName, u)
		}
	*/

	chromedp.Run(ctx, chromedp.Navigate(url))

	chromedp.Run(ctx, chromedp.AttributeValue(`meta[name="title"]`, "content", &result.MetaTitle, nil, chromedp.AtLeast(0)))

	if len(result.MetaTitle) == 0 {
		chromedp.Run(ctx, chromedp.Navigate(url), chromedp.InnerHTML(`head > title`, &result.MetaTitle, chromedp.AtLeast(0)))
	}
	chromedp.Run(ctx, chromedp.AttributeValue(`meta[name="description"]`, "content", &result.MetaDescription, nil, chromedp.AtLeast(0)))

	chromedp.Run(ctx, chromedp.AttributeValue(`meta[property="og:image"]`, "content", &result.MetaImage, nil, chromedp.AtLeast(0)))

	// document.querySelector("#landingImage")
	if len(result.MetaImage) == 0 {
		log.Printf("fallback meta image")
		chromedp.Run(ctx, chromedp.AttributeValue(`#landingImage`, "src", &result.MetaImage, nil, chromedp.AtLeast(0)))
	}

}

func identifyScreenshotPageSelector(url string) (response string, err error) {
	if s.Contains(url, "lidl") {
		return `//*[@id="__layout"]/div/main/section/article`, nil
	} else if s.Contains(url, "amazon") {
		return `//*[@id="ppd"]`, nil
	} else if s.Contains(url, "kaufland") {
		return `//*[@id="__layout"]/div[1]/div[1]`, nil
	}

	return "", fmt.Errorf("no pdp selector found for %s", url)

}
