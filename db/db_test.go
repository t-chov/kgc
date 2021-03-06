package db

import (
	_ "embed"
	"reflect"
	"testing"
)

//go:embed gc_test.csv
var csvBytes []byte

func TestInitDbFailure(t *testing.T) {
	invalid := []byte{1, 2, 3, 4, 5}
	got, err := InitDb(invalid)
	if err != nil {
		t.Errorf("got %v, expected nil", got)
	}
}
func TestInitDb(t *testing.T) {
	dbp, _ := InitDb(csvBytes)
	db := *dbp
	tests := []struct {
		name     string
		arg      string
		expected bool
	}{
		{
			name:     "アームスタンド exists.",
			arg:      "アームスタンド",
			expected: true,
		},
		{
			name:     "あああああ does not exist",
			arg:      "あああああ",
			expected: false,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			_, got := db[testCase.arg]
			if got != testCase.expected {
				t.Errorf("db[%s] = %v, expected %v", testCase.arg, got, testCase.expected)
			}
		})
	}
}

func TestFind(t *testing.T) {
	dbp, _ := InitDb(csvBytes)
	db := *dbp
	tests := []struct {
		name     string
		arg      string
		expected *Garbage
	}{
		{
			name: "IH調理器 exists.",
			arg:  "IH調理器",
			expected: &Garbage{
				Name:        "IH調理器",
				GarbageType: "小物金属",
				DetailType:  "粗大ごみ",
				Description: "最長辺30cm以上のものは粗大ごみとして出してください。/長辺30cm未満で、30cm×15cmの回収ボックスの投入口に入るものは小型家電としても出すことができます。",
				Url:         "http://www.city.kawasaki.jp/kurashi/category/24-1-23-1-1-6-9-0-0-0.html",
			},
		},
		{
			name:     "あああああ does not exist",
			arg:      "あああああ",
			expected: nil,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			garbage, _ := db.Find(testCase.arg)
			if !reflect.DeepEqual(garbage, testCase.expected) {
				t.Errorf("find %v, expected %v", garbage, testCase.expected)
			}
		})
	}

}
