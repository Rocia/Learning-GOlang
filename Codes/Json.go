package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// We'll use these two structs to demonstrate encoding and
// decoding of custom types below.
type Response1 struct {
	Page   int
	Dishes []string
}
type Response2 struct {
	Page   int      `json:"page"`
	Dishes []string `json:"Dishes"`
}

func main() {

	bolB, _ := json.Marshal(true)	//Atomic JSON values
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("rocia")
	fmt.Println(string(strB))

	// And here are some for slices and maps, which encode
	// to JSON arrays and objects as you'd expect.
	menu := []string{"ravioli","carbonara","spaghetti","farfelle","lasagne"}
	JsonMenu, _ := json.Marshal(menu)
	fmt.Println(string(JsonMenu))

	mapD := map[string]int{"ravioli": 7, "spaghetti": 9}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	Dict1 := &Response1{		//Json automatically will encode ua custom datatypes check output to see how it makes keys
		Page:   1,
		Dishes: []string{"ravioli","carbonara","spaghetti"}}
	result1, _ := json.Marshal(Dict1)
	fmt.Println(string(result1))

	res2D := &Response2{		//If you use tags while declaring the struct field you can customise the json key names.
		Page:   1,
		Dishes: []string{"ravioli","carbonara","spaghetti"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"Numeric":6.13,"Alphabetical":["a","b"]}`)		//decoding json into data structs in GO

	var dat map[string]interface{}		//Declare a map/dictionary to hold the decoded json data

	if err := json.Unmarshal(byt, &dat); err != nil {		//decode string and check for errors if any
		panic(err)
	}
	fmt.Println(dat)

	num := dat["Numeric"].(float64)	//cast value on field to expected data type
	fmt.Println(num)


	strs := dat["Alphabetical"].([]interface{})		// for nested data you need a series of casts
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "Dishes": ["ravioli","carbonara"]}`			//Decode Json into custom data types
	res := Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Dishes[0])

	enc := json.NewEncoder(os.Stdout)								//Streaming JSON to os.writers std.out
	d := map[string]int{"ravioli": 7, "spaghetti": 9}
	enc.Encode(d)
}


/*
true
1
2.34
"rocia"
["ravioli","carbonara","spaghetti","farfelle","lasagne"]
{"ravioli":7,"spaghetti":9}
{"Page":1,"Dishes":["ravioli","carbonara","spaghetti"]}
{"page":1,"Dishes":["ravioli","carbonara","spaghetti"]}
map[Numeric:6.13 Alphabetical:[a b]]
6.13
a
{1 [ravioli carbonara]}
ravioli
{"ravioli":7,"spaghetti":9}
 */