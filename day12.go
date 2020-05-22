package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// map string slice struct
// 10326
func sum(m interface{}) int {
	c := 0
	typ := reflect.TypeOf(m).Kind()
	switch typ {

	case reflect.Slice:
		t := m.([]interface{})
		for i := 0; i < len(t); i++ {
			c += sum(t[i])
		}

	case reflect.Map:
		t := m.(map[string]interface{})
		for _, v := range t {
			c += sum(v)
		}
	// 	fmt.Println(s)
	// 	sum(s)

	case reflect.Struct:
		// No action
		// s := reflect.ValueOf(m)
		// fmt.Println("s", s)
		// fmt.Println("")
		// sum(s)
		// return 1

	case reflect.String: // string and numbers
		// No action
		// t := m.(string)
		// r, err := strconv.Atoi(t)
		// if err == nil {
		// 	return r
		// }
	case reflect.Float64:
		t := m.(float64)
		c += int(t)
	}

	return c
}

// Day12 is here
func Day12() {

	data := input("data/day12.json")

	var result map[string]interface{}
	json.Unmarshal([]byte(data), &result)

	c := 0
	for key := range result {

		d := result[key]
		c += sum(d)
		//fmt.Println(d)
		// Each value is an interface{} type, that is type asserted as a string
		//fmt.Println(key) //, value.(string))
	}

	fmt.Println(c)

}
