// Creates a Screenshot of the product relevant informations from the most famous online shops
package pdp_extractor

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {
	type args struct {
		html string
		url  string
	}
	tests := []struct {
		name       string
		args       args
		wantResult ProductDetailPageExtractorResults
		wantErr    bool
	}{
		/*{
			name: "amazon",
			args: args{
				html: "./in/amazon.html",
			},
			wantResult: ProductDetailPageExtractorResults{
				MetaTitle:       `Kewago Auto Organizer Mittelkonsole. Zubehör Auto Aufbewahrung. Die Aufbewahrungsbox ist einfach zwischen Autositz und Mittelkonsole einsteckbar. Praktische Ergänzung zum Auto Getränkehalter : Amazon.de: Auto & Motorrad`,
				MetaDescription: `Kaufen Sie Kewago Auto Organizer Mittelkonsole. Zubehör Auto Aufbewahrung. Die Aufbewahrungsbox ist einfach zwischen Autositz und Mittelkonsole einsteckbar. Praktische Ergänzung zum Auto Getränkehalter im Auto & Motorrad-Shop auf Amazon.de. Große Auswahl und Gratis Lieferung durch Amazon ab 29€.`,
				MetaImage:       `https://images-eu.ssl-images-amazon.com/images/I/61AusEXMQXL.__AC_SX300_SY300_QL70_ML2_.jpg`},
		},
		{
			name: "lidl",
			args: args{
				html: "./in/lidl.html",
			},
			wantResult: ProductDetailPageExtractorResults{
				MetaTitle:       `Puma Herren T-Shirt teamGoal, mit Rundhalsausschnitt`,
				MetaDescription: `Puma Herren T-Shirt teamGoal, mit Rundhalsausschnitt im LIDL Online-Shop kaufen ✓ 90 Tage Rückgaberecht ✓ Schneller Versand  ✓ Jetzt bestellen!`,
				MetaImage:       `https://www.lidl.de/media/b43b4600b210e0a36c7cdb14cb7d6721.jpeg`},
		},
		{
			name: "kaufland",
			args: args{
				html: "./in/kaufland.html",
			},
			wantResult: ProductDetailPageExtractorResults{
				MetaTitle:       `Sammelalbum | Format A4 | Pokemon | | Kaufland.de`,
				MetaDescription: `Sammelkarte Sammelalbum | Format A4 | Pokemon | Sammelkarten-Spiel | Album für 252 Karten Preis ab 21.95€ (19.12.2021). Jetzt kaufen!`,
				MetaImage:       `https://media.cdn.kaufland.de/product-images/original/9a04b48385bef67bfc105a879f0a6648.jpg`},
		},
		{
			name: "otto",
			args: args{
				html: "./in/otto.html",
			},
			wantResult: ProductDetailPageExtractorResults{
				MetaTitle:       `Tefal Pfannen-Set »Duetto«, Edelstahl (Set, 3-tlg), A704S3, Ø 20/24/28 cm, Antihaft-Beschichtung, Thermo-Signal-Technologie, induktionsgeeignet, für alle Herdarten, Edelstahl/Schwarz online kaufen | OTTO`,
				MetaDescription: `Tefal Pfannen-Set »Duetto«, Edelstahl (Set, 3-tlg), A704S3, Ø 20/24/28 cm, Antihaft-Beschichtung, Thermo-Signal-Technologie, induktionsgeeignet, für alle Herdarten, Edelstahl/Schwarz für 79,99€ bei OTTO`,
				MetaImage:       `https://i.otto.de/i/otto/34282388/tefal-pfannen-set-duetto-edelstahl-set-3-tlg-a704s3-o-20-24-28-cm-antihaft-beschichtung-thermo-signal-technologie-induktionsgeeignet-fuer-alle-herdarten-edelstahl-schwarz.jpg?$formatz$`},
		},
		{
			name: "zalando",
			args: args{
				html: "./in/zalando.html",
			},
			wantResult: ProductDetailPageExtractorResults{
				MetaTitle:       `Tezenis Slip - nero/schwarz - Zalando.de`,
				MetaDescription: `Tezenis Slip - nero/schwarz für 4,99 € (19.12.2021) versandkostenfrei bei Zalando bestellen.`,
				MetaImage:       `https://img01.ztat.net/article/spp-media-p1/727cd53f90ca40fda02b6dce352505cb/8ca71bfbdc5f4ceba42f73258b8616df.jpg?imwidth=103&filter=packshot`},
		},*/
		/*{
			name: "lidl2",
			args: args{
				url: "https://www.lidl.de/p/puma-herren-t-shirt-teamgoal-mit-rundhalsausschnitt/p100339220",
			},
			wantResult: ProductDetailPageExtractorResults{
				MetaTitle:       `Puma Herren T-Shirt teamGoal, mit Rundhalsausschnitt`,
				MetaDescription: `Puma Herren T-Shirt teamGoal, mit Rundhalsausschnitt im LIDL Online-Shop kaufen ✓ 90 Tage Rückgaberecht ✓ Schneller Versand  ✓ Jetzt bestellen!`,
				MetaImage:       `https://www.lidl.de/media/b43b4600b210e0a36c7cdb14cb7d6721.jpeg`},
		},
		*/
		{
			name: "zalando2",
			args: args{
				url: "https://www.zalando.de/tezenis-slip-nero-teg81r05s-q11.html",
			},
			wantResult: ProductDetailPageExtractorResults{
				MetaTitle:       `Tezenis Slip - nero/schwarz - Zalando.de`,
				MetaDescription: ` Tezenis Slip - nero/schwarz für 4,99 € (19.12.2021) versandkostenfrei bei Zalando bestellen.`,
				MetaImage:       `https://img01.ztat.net/article/spp-media-p1/727cd53f90ca40fda02b6dce352505cb/8ca71bfbdc5f4ceba42f73258b8616df.jpg?imwidth=103&filter=packshot`},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := tt.args.url
			if url == "" {
				content, err := ioutil.ReadFile(tt.args.html)
				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

					if err == nil {
						fmt.Fprint(w, string(content))
					} else {
						t.Errorf("error hosting file %s", err.Error())
						fmt.Fprint(w, "error")
					}
				}))
				defer ts.Close()
				url = ts.URL
			}

			gotResult, err := Run(url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("\nGot =\t %v \n Want =\t %v", gotResult, tt.wantResult)
			}

			t.Logf("extracted\ntitle:\n%s\ndescr.:\n%s\nmimage:\n%s\n", gotResult.MetaTitle, gotResult.MetaDescription, gotResult.MetaImage)

		})
	}
}
