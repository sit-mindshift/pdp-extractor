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
type ProductDetailPageExtractor struct {
	Url             string
	ScreenshotFile  string
	MetaTitle       string
	MetaDescription string
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
func (s *ProductDetailPageExtractor) Run() (err error) {
	// create context
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	ctx, cancel = context.WithTimeout(ctx, 20*time.Second)

	defer cancel()

	var buf []byte

	if selector, error := identifyScreenshotPageSelector(s.Url); error == nil {
		if err := chromedp.Run(ctx,
			chromedp.EmulateViewport(1920, 3000),
			elementScreenshot(s.Url, selector, &buf),
		); err != nil {
			log.Println(err.Error())
			return err
		} else {
			log.Printf("created a screenshot of %s using the page selector %s\n", s.Url, selector)
		}

	} else {
		// capture entire browser viewport, returning png with quality=90
		if err := chromedp.Run(ctx,
			chromedp.EmulateViewport(1920, 3000),
			fullScreenshot(s.Url, 90, &buf)); err != nil {
			log.Println(err.Error())
			return err
		} else {
			log.Printf("created a screenshot of %s using fullscreen mode\n", s.Url)
		}
	}

	s.extractMetaInformation(ctx)

	if err := ioutil.WriteFile(s.ScreenshotFile, buf, 0o644); err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (s *ProductDetailPageExtractor) extractMetaInformation(ctx context.Context) {

	chromedp.Run(ctx, chromedp.InnerHTML(`head > title`, &s.MetaTitle))
	if len(s.MetaTitle) == 0 {
		chromedp.Run(ctx, chromedp.AttributeValue(`meta[name="title"]`, "content", &s.MetaTitle, nil))
	}
	chromedp.Run(ctx, chromedp.AttributeValue(`meta[name="description"]`, "content", &s.MetaDescription, nil))
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
