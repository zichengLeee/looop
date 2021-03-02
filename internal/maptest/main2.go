package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Age map[string]string
}

func main() {
	var ssss map[string]interface{}
	ss_map, _ := ssss["ssss"].(map[string]interface{})
	fmt.Printf("ss:%v\n", ss_map["nihao"])
	s := &Student{
		Age: map[string]string{
			"123": "123",
		},
	}
	bytes, err := json.Marshal(s)
	if err != nil {
		fmt.Println("1 error:%s", err.Error())
		return
	}
	ss := &Student{Age: make(map[string]string)}
	err = json.Unmarshal(bytes, ss)
	if err != nil {
		fmt.Println("2 error:%s", err.Error())
		return
	}
	ss.Age["234"] = "234"
	fmt.Println(ss)
}
