package json

import (
	"encoding/json"
	"testing"
)

type Monster struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Sal      float64 `json:"sal"`
	Skill    string  `json:"skill"`
}

// 结构体序列化
func TestStructJson(t *testing.T) {
	monster := Monster{
		Name:     "阿牛",
		Age:      22,
		Birthday: "2021-10-12",
		Sal:      8000.00,
		Skill:    "打牛",
	}

	data, err := json.Marshal(&monster)
	if err != nil {
		t.Logf("序列化失败%s", err)
		return
	}

	t.Log(string(data))
}

// map序列化
func TestMapJson(t *testing.T) {
	var a map[string]any
	a = make(map[string]any)
	a["name"] = "阿黄"
	a["age"] = 22
	data, err := json.Marshal(a)
	if err != nil {
		t.Logf("序列化失败%s", err)
		return
	}
	t.Log(string(data))
}

// 切片序列化
func TestSliceJson(t *testing.T) {
	var slice []map[string]any
	var m1 map[string]any
	m1 = make(map[string]any)
	m1["name"] = "jack"
	m1["age"] = 7
	slice = append(slice, m1)

	var m2 map[string]any
	m2 = make(map[string]any)
	m2["name"] = "tom"
	m2["age"] = 22

	slice = append(slice, m2)

	data, err := json.Marshal(slice)
	if err != nil {
		t.Logf("序列化失败%s", err)
		return
	}
	t.Log(string(data))
}

func TestJsonStruct(t *testing.T) {
	str := "{\"name\":\"阿牛\",\"age\":22,\"birthday\":\"2021-10-12\",\"sal\":8000,\"skill\":\"打牛\"}"

	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		t.Logf("序列化失败%s", err)
	}

	t.Log(monster)
}
