package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// map string slice struct

func sum(m interface{}) int {

	typ := reflect.TypeOf(m).Kind()
	fmt.Println("type", typ)
	switch typ {

	case reflect.Slice:
		// s := reflect.ValueOf(m)
		// for i := 0; i < s.Len(); i++ {
		// 	t := s.Index(i)
		// 	t2 := reflect.TypeOf(t).Kind()
		// 	//fmt.Println(s.Index(i))
		// 	fmt.Println(t, t2)
		// 	sum(s)
		// }

	// case reflect.Map:
	// 	s := reflect.ValueOf(m)
	// 	fmt.Println(s)
	// 	sum(s)

	case reflect.Struct:
		//		s := reflect.ValueOf(m)

		// fmt.Println("s", s)
		// fmt.Println("")
		//sum(s)
		return 1

	case reflect.String: // string and numbers
		// No action
		// t := m.(string)
		// r, err := strconv.Atoi(t)
		// if err == nil {
		// 	return r
		// }
	case reflect.Float64:
		t := m.(float64)
		return int(t)
	}

	return 0
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
