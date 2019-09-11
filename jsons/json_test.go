package jsons

import (
	"strings"
	"testing"
)

func TestEscape(t *testing.T) {
	input := "\\u0026123\\u003c1\\u003e"
	except := "&123<1>"
	actual := Escape(input)
	if !strings.EqualFold(actual, except) {
		t.Errorf("test fail actual:%s, except:%s", actual, except)
	}
}

type User struct {
	Name    string `json:"name" `
	Age     int    `json:"age" ms:"age"`
	Addr    string `json:"addr"`
	Country string `json:"country" ms:"country"`
}

func TestMarshal(t *testing.T) {

	u := User{
		Name:    "jack",
		Age:     22,
		Addr:    "成都",
		Country: "中国",
	}

	gotA, err := Marshal(u)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(gotA))

	gotB, err := Marshal(u, "ms")
	if err != nil {
		t.Error(err)
	}
	t.Log(string(gotB))
}
