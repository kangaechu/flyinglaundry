package tenki

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"slices"
	"strconv"
	"strings"
)

type Tenki struct {
	doc *goquery.Document
}

// NewTenki returns a new Tenki struct
func NewTenki(r io.Reader) (*Tenki, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	return &Tenki{doc: doc}, nil
}

func (t *Tenki) getData(cssClass string) []string {
	var data []string
	t.doc.Find(cssClass).Each(func(i int, s *goquery.Selection) {
		data = append(data, strings.TrimSpace(s.Text()))
	})
	return data
}

func (t *Tenki) Weather() []string {
	return slices.Concat(
		t.getData("#forecast-point-1h-today tr.weather td"),
		t.getData("#forecast-point-1h-tomorrow tr.weather td"),
	)
}

func (t *Tenki) Temperature() []string {
	return slices.Concat(
		t.getData("#forecast-point-1h-today tr.temperature td"),
		t.getData("#forecast-point-1h-tomorrow tr.temperature td"),
	)
}

func (t *Tenki) ProbPrecip() []string {
	return slices.Concat(
		t.getData("#forecast-point-1h-today tr.prob-precip td"),
		t.getData("#forecast-point-1h-tomorrow tr.prob-precip td"),
	)
}

func (t *Tenki) Precipitation() []string {
	return slices.Concat(
		t.getData("#forecast-point-1h-today tr.precipitation td"),
		t.getData("#forecast-point-1h-tomorrow tr.precipitation td"),
	)
}

func (t *Tenki) Humidity() []string {
	return slices.Concat(
		t.getData("#forecast-point-1h-today tr.humidity td"),
		t.getData("#forecast-point-1h-tomorrow tr.humidity td"),
	)
}

func (t *Tenki) WindBlow() []string {
	return slices.Concat(
		t.getData("#forecast-point-1h-today tr.wind-blow td"),
		t.getData("#forecast-point-1h-tomorrow tr.wind-blow td"),
	)
}

func (t *Tenki) WindSpeed() []string {
	return slices.Concat(
		t.getData("#forecast-point-1h-today tr.wind-speed td"),
		t.getData("#forecast-point-1h-tomorrow tr.wind-speed td"),
	)
}

// CheckFlyingLaundry は風向に南を含み、風速が5m以上の場合にtrueを返す
func (t *Tenki) CheckFlyingLaundry(start, length int) bool {
	windSpeed24h := t.WindSpeed()[start-1 : start+length-1]
	windBlow24h := t.WindBlow()[start-1 : start+length-1]

	flyingLaundry := false
	for i := 0; i < length; i++ {
		windSpeed, err := strconv.Atoi(windSpeed24h[i])
		if err != nil {
			fmt.Println("Error converting wind speed:", err)
			continue
		}
		windBlow := windBlow24h[i]
		if strings.Contains(windBlow, "南") && windSpeed >= 5 {
			flyingLaundry = true
			break
		}
	}
	return flyingLaundry
}
