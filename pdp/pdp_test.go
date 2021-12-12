// Creates a Screenshot of the product relevant informations from the most famous online shops
package pdp_extractor

import (
	"testing"
)

func TestRun(t *testing.T) {
	type args struct {
		url  string
		file string
	}
	tests := []struct {
		name       string
		args       args
		wantResult ProductDetailPageExtractorResults
		wantErr    bool
	}{
		{
			name: "Amazon",
			args: args{url: "https://www.amazon.de/dp/B07SPTL3SQ/ref=nosim?tag=masedeveloper-21", file: "./out/amazon.de.png"},
		},
		{
			name: "Lidl",
			args: args{url: "https://www.lidl.de/p/puma-herren-t-shirt-teamgoal-mit-rundhalsausschnitt/p100339220", file: "./out/lidl.de.png"},
		},
		{
			name: "Kaufland",
			args: args{url: "https://www.kaufland.de/product/404444772/", file: "./out/kaufland.de.png"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := Run(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("extracted\ntitle:\t%s \ndescr.:\t%s\n", gotResult.MetaTitle, gotResult.MetaDescription)
			WriteScreenshotToFile(tt.args.file, gotResult)
		})
	}
}
