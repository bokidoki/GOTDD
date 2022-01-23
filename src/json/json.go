package main

import (
	"encoding/json"
	"log"
)

type Foo struct {
	Foo string
}

func main() {
	// 将数据编码成json字符串
	jsonstr, _ := json.Marshal(Foo{Foo: "bar"})
	Foo := Foo{}
	// 将Json字符串解码到相应的数据结构
	json.Unmarshal(jsonstr, &Foo)

	log.Println(jsonstr)
	log.Println(Foo)
}
