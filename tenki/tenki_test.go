package tenki

import (
	"github.com/PuerkitoBio/goquery"
	"os"
	"reflect"
	"testing"
)

func TestNewTenki(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    *Tenki
		wantErr bool
	}{
		{
			name: "TestNewTenki",
			args: args{
				filename: "../test/sapporo.html",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := getTenkiHtml(tt.args.filename)
			_, err := NewTenki(r)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTenki() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestTenki_Humidity(t1 *testing.T) {
	type fields struct {
		filename string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name:   "正常系",
			fields: fields{filename: "../test/sapporo.html"},
			want: []string{
				// Today
				"92", "92", "95", "95", "96", "97",
				"93", "89", "85", "81", "80", "86",
				"85", "78", "77", "82", "84", "89",
				"89", "89", "90", "91", "92", "92",
				// Tomorrow
				"91", "89", "89", "88", "88", "87",
				"83", "80", "75", "72", "69", "67",
				"66", "66", "68", "69", "73", "78",
				"83", "86", "88", "89", "89", "89",
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := getTestData(tt.fields.filename)
			if got := t.Humidity(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Humidity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTenki_Precipitation(t1 *testing.T) {
	type fields struct {
		filename string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name:   "正常系",
			fields: fields{filename: "../test/sapporo.html"},
			want: []string{
				// Today
				"0", "0", "0", "0", "0", "0",
				"0", "0", "0", "0", "0", "0",
				"0", "0", "0", "0", "0", "0",
				"0", "0", "0", "0", "0", "0",
				// Tomorrow
				"0", "0", "0", "0", "0", "0",
				"0", "0", "0", "0", "0", "0",
				"0", "0", "0", "0", "0", "0",
				"0", "0", "0", "0", "0", "0",
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := getTestData(tt.fields.filename)
			if got := t.Precipitation(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Precipitation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTenki_ProbPrecip(t1 *testing.T) {
	type fields struct {
		filename string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name:   "正常系",
			fields: fields{filename: "../test/sapporo.html"},
			want: []string{
				// Today
				"---", "---", "---", "---", "---", "---",
				"---", "---", "---", "---", "---", "---",
				"---", "---", "---", "---", "---", "---",
				"---", "---", "20", "20", "20", "20",
				// Tomorrow
				"20", "20", "20", "20", "30", "20",
				"20", "20", "20", "20", "10", "10",
				"0", "0", "0", "0", "0", "0",
				"0", "0", "0", "0", "0", "0",
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := getTestData(tt.fields.filename)
			if got := t.ProbPrecip(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("ProbPrecip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTenki_Temperature(t1 *testing.T) {
	type fields struct {
		filename string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name:   "正常系",
			fields: fields{filename: "../test/sapporo.html"},
			want: []string{
				// Today
				"24.0", "23.5", "23.5", "23.5", "23.5", "23.5",
				"23.5", "24.0", "24.5", "25.0", "26.1", "27.3",
				"26.6", "25.0", "25.0", "24.5", "24.0", "23.5",
				"23.1", "22.9", "22.8", "22.8", "22.5", "22.2",
				// Tomorrow
				"22.1", "22.1", "22.0", "21.9", "21.7", "21.8",
				"22.5", "23.3", "24.1", "25.0", "25.6", "26.6",
				"27.1", "26.7", "25.6", "24.9", "24.0", "23.1",
				"22.4", "22.0", "21.8", "21.5", "21.1", "21.0",
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := getTestData(tt.fields.filename)
			if got := t.Temperature(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Temperature() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTenki_Weather(t1 *testing.T) {
	type fields struct {
		filename string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name:   "正常系",
			fields: fields{filename: "../test/sapporo.html"},
			want: []string{
				// Today
				"曇り", "曇り", "曇り", "曇り", "曇り", "小雨",
				"曇り", "曇り", "曇り", "曇り", "曇り", "曇り",
				"曇り", "曇り", "曇り", "曇り", "曇り", "晴れ",
				"晴れ", "晴れ", "晴れ", "晴れ", "晴れ", "晴れ",
				// Tomorrow
				"曇り", "曇り", "曇り", "曇り", "曇り", "曇り",
				"曇り", "曇り", "曇り", "晴れ", "晴れ", "晴れ",
				"晴れ", "晴れ", "晴れ", "晴れ", "晴れ", "晴れ",
				"晴れ", "晴れ", "晴れ", "晴れ", "晴れ", "晴れ",
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := getTestData(tt.fields.filename)
			if got := t.Weather(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Weather() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTenki_WindBlow(t1 *testing.T) {
	type fields struct {
		filename string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name:   "正常系",
			fields: fields{filename: "../test/sapporo.html"},
			want: []string{
				// Today
				"南南東", "南南東", "静穏", "南南東", "南東", "静穏",
				"静穏", "静穏", "静穏", "北", "北", "北西",
				"北北西", "北北西", "北北西", "北北西", "北北西", "北北西",
				"北西", "北西", "北西", "北西", "北西", "北西",
				// Tomorrow
				"北西", "北西", "北西", "北西", "北西", "北西",
				"北西", "北北西", "北北西", "南", "南", "南",
				"南南西", "南南西", "南南西", "西南西", "西南西", "西南西",
				"北西", "北西", "北西", "西北西", "西", "西",
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := getTestData(tt.fields.filename)
			if got := t.WindBlow(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("WindBlow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTenki_WindSpeed(t1 *testing.T) {
	type fields struct {
		filename string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name:   "正常系",
			fields: fields{filename: "../test/sapporo.html"},
			want: []string{
				// Today
				"1", "1", "0", "1", "1", "0",
				"0", "0", "0", "1", "1", "1",
				"1", "1", "2", "2", "2", "2",
				"3", "3", "2", "2", "2", "2",
				// Tomorrow
				"2", "2", "3", "3", "3", "3",
				"4", "5", "6", "4", "5", "6",
				"4", "5", "6", "4", "5", "6",
				"4", "5", "6", "4", "5", "6",
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := getTestData(tt.fields.filename)
			if got := t.WindSpeed(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("WindSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTenki_CheckFlyingLaundry(t1 *testing.T) {
	type fields struct {
		doc *goquery.Document
	}
	type args struct {
		start  int
		length int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "正常系",
			fields: fields{
				doc: getTestData("../test/sapporo.html").doc,
			},
			args: args{
				start:  1,
				length: 24,
			},
			want: false,
		},
		{
			name: "南風、風速5m/s未満",
			fields: fields{
				doc: getTestData("../test/sapporo.html").doc,
			},
			args: args{
				start:  34,
				length: 1,
			},
			want: false,
		},
		{
			name: "南風、風速5m",
			fields: fields{
				doc: getTestData("../test/sapporo.html").doc,
			},
			args: args{
				start:  35,
				length: 1,
			},
			want: true,
		},
		{
			name: "南風、風速5m強",
			fields: fields{
				doc: getTestData("../test/sapporo.html").doc,
			},
			args: args{
				start:  36,
				length: 1,
			},
			want: true,
		},
		{
			name: "南南西の風、風速5m/s未満",
			fields: fields{
				doc: getTestData("../test/sapporo.html").doc,
			},
			args: args{
				start:  37,
				length: 1,
			},
			want: false,
		},
		{
			name: "南南西の風、風速5m",
			fields: fields{
				doc: getTestData("../test/sapporo.html").doc,
			},
			args: args{
				start:  38,
				length: 1,
			},
			want: true,
		},
		{
			name: "南南西の風、風速5m強",
			fields: fields{
				doc: getTestData("../test/sapporo.html").doc,
			},
			args: args{
				start:  39,
				length: 1,
			},
			want: true,
		},
		{
			name: "西南西の風、風速5m/s未満",
			fields: fields{
				doc: getTestData("../test/sapporo.html").doc,
			},
			args: args{
				start:  40,
				length: 1,
			},
			want: false,
		},
		{
			name: "西南西の風、風速5m",
			fields: fields{
				doc: getTestData("../test/sapporo.html").doc,
			},
			args: args{
				start:  41,
				length: 1,
			},
			want: true,
		},
		{
			name: "西南西の風、風速5m強",
			fields: fields{
				doc: getTestData("../test/sapporo.html").doc,
			},
			args: args{
				start:  42,
				length: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tenki{
				doc: tt.fields.doc,
			}
			if got := t.CheckFlyingLaundry(tt.args.start, tt.args.length); got != tt.want {
				t1.Errorf("CheckFlyingLaundry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getTenkiHtml(filename string) *os.File {
	f, err := os.Open(filename)
	if err != nil {
		return nil
	}

	return f

}
func getTestData(filename string) *Tenki {
	f := getTenkiHtml(filename)
	t, err := NewTenki(f)
	if err != nil {
		return nil
	}
	return t
}
