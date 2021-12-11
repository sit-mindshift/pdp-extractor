// Creates a Screenshot of the product relevant informations from the most famous online shops
package pdp_extractor

import "testing"

func TestProductDetailPageExtractor_Run(t *testing.T) {
	tests := []struct {
		name string
		s    *ProductDetailPageExtractor
	}{
		{
			name: "amazon",
			s: &ProductDetailPageExtractor{
				Url:            "https://www.amazon.de/dp/B07SPTL3SQ/ref=nosim?tag=masedeveloper-21",
				ScreenshotFile: "./out/amazon.de.png",
			},
		},
		{
			name: "lidl",
			s: &ProductDetailPageExtractor{
				Url:            "https://www.lidl.de/p/puma-herren-t-shirt-teamgoal-mit-rundhalsausschnitt/p100339220",
				ScreenshotFile: "./out/lidl.de.png",
			},
		},
		{
			name: "kaufland",
			s: &ProductDetailPageExtractor{
				Url:            "https://www.kaufland.de/product/404444772/",
				ScreenshotFile: "./out/kaufland.de.png",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("start PDP with url:%s \n", tt.s.Url)
			tt.s.Run()
			t.Logf("extracted\ntitle:\t%s \ndescr.:\t%s\n", tt.s.MetaTitle, tt.s.MetaDescription)
		})
	}
}
