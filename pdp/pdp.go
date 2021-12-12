// The Product Detail Page Extractor crawles all usefull informations from a product details page
// which might be relevant for creating a thubnail preview of a article
package pdp_extractor

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	s "strings"
	"time"

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
	ScreenshotBuffer []byte
	MetaTitle        string
	MetaDescription  string
}

func identifyScreenshotPageSelector(url string) (response string, err error) {
	if s.Contains(url, "lidl") {
		return `//*[@id="__layout"]/div/main/section/article`, nil
	} else if s.Contains(url, "amazon") {
		return `//*[@id="ppd"]`, nil
	} else if s.Contains(url, "kaufland") {
		return `//*[@id="__layout"]/div[1]/div[1]`, nil
	}
	return "", errors.New(fmt.Sprintf("No PDP selector found for %s", url))

}

// Starts the crawling of the product details page
func Run(url string) (result ProductDetailPageExtractorResults, err error) {
	// create context
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	ctx, cancel = context.WithTimeout(ctx, 20*time.Second)

	defer cancel()

	if selector, error := identifyScreenshotPageSelector(url); error == nil {
		if err := chromedp.Run(ctx,
			chromedp.EmulateViewport(1920, 3000),
			elementScreenshot(url, selector, &result.ScreenshotBuffer),
		); err != nil {
			log.Println(err.Error())
			return result, err
		} else {
			log.Printf("created a screenshot of %s using the page selector %s\n", url, selector)
		}

	} else {
		// capture entire browser viewport, returning png with quality=90
		if err := chromedp.Run(ctx,
			chromedp.EmulateViewport(1920, 3000),
			fullScreenshot(url, 90, &result.ScreenshotBuffer)); err != nil {
			log.Println(err.Error())
			return result, err
		} else {
			log.Printf("created a screenshot of %s using fullscreen mode\n", url)
		}
	}
	extractMetaInformation(ctx, &result)
	return result, err
}

func WriteScreenshotToFile(filename string, result ProductDetailPageExtractorResults) (err error) {
	log.Printf("WriteScreenshotToFile of  buff len %d \n", len(result.ScreenshotBuffer))
	err = ioutil.WriteFile(filename, result.ScreenshotBuffer, 0o644)
	if err != nil {
		return err
	}
	return nil
}

func extractMetaInformation(ctx context.Context, result *ProductDetailPageExtractorResults) {

	chromedp.Run(ctx, chromedp.InnerHTML(`head > title`, &result.MetaTitle))
	if len(result.MetaTitle) == 0 {
		chromedp.Run(ctx, chromedp.AttributeValue(`meta[name="title"]`, "content", &result.MetaTitle, nil))
	}
	chromedp.Run(ctx, chromedp.AttributeValue(`meta[name="description"]`, "content", &result.MetaDescription, nil))
}

// elementScreenshot takes a screenshot of a specific element.
func elementScreenshot(urlstr, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
	}
}

// fullScreenshot takes a screenshot of the entire browser viewport.
//
// Note: chromedp.FullScreenshot overrides the device's emulation settings. Use
// device.Reset to reset the emulation and viewport settings.
func fullScreenshot(urlstr string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.FullScreenshot(res, quality),
	}
}
